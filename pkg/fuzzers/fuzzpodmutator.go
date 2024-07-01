package main

import (
	"C"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	run "runtime"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
	klog "k8s.io/klog/v2"
	kubelet "k8s.io/kubernetes/pkg/kubelet"
)
import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

func generateRandomHash() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(randomBytes)
	return hex.EncodeToString(hash[:]), nil
}

// ErrorLog
type ErrorLog struct {
	Timestamp    string
	ErrorType    string
	ErrorMessage string
	FunctionName string
	CoverageData string
}

var initOnce sync.Once

// Initialize klog only once
func initializeKlog() {
	initOnce.Do(func() {
		klog.InitFlags(nil)
		flag.Set("logtostderr", "false")
		flag.Set("log_file", "klog.log")
		flag.Parse()
		klog.Info("klog initialized")
	})
}

func storePodInDatabasePod(pod *v1.Pod, hash string) error {
	podJSON, err := json.Marshal(pod)
	if err != nil {
		return fmt.Errorf("failed to marshal pod: %v", err)
	}

	db, err := sql.Open("sqlite3", "pods.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pods (
        hash TEXT PRIMARY KEY,
        pod_json TEXT
    )`)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	_, err = db.Exec("INSERT INTO pods (hash, pod_json) VALUES (?, ?)", hash, string(podJSON))
	if err != nil {
		return fmt.Errorf("failed to insert into database: %v", err)
	}

	return nil
}

func storePodInDatabaseNode(node *v1.Node, hash string) error {
	podJSON, err := json.Marshal(node)
	if err != nil {
		return fmt.Errorf("failed to marshal pod: %v", err)
	}

	db, err := sql.Open("sqlite3", "pods.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pods (
        hash TEXT PRIMARY KEY,
        pod_json TEXT
    )`)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	_, err = db.Exec("INSERT INTO pods (hash, pod_json) VALUES (?, ?)", hash, string(podJSON))
	if err != nil {
		return fmt.Errorf("failed to insert into database: %v", err)
	}

	return nil
}

// Recursively extract field values from the map and convert them to a suitable format for libfuzzer
func extractFieldValues(data interface{}) ([]byte, error) {
	var result []byte
	val := reflect.ValueOf(data)

	if !val.IsValid() {
		return result, nil
	}

	// Dereference pointer if needed
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return result, nil
		}
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			mapValue := val.MapIndex(key)
			fieldValue, err := extractFieldValues(mapValue.Interface())
			if err != nil {
				return nil, err
			}
			result = append(result, fieldValue...)
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			elementValue := val.Index(i)
			fieldValue, err := extractFieldValues(elementValue.Interface())
			if err != nil {
				return nil, err
			}
			result = append(result, fieldValue...)
		}

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			fieldValue := val.Field(i)
			if fieldValue.CanInterface() {
				fieldBytes, err := extractFieldValues(fieldValue.Interface())
				if err != nil {
					return nil, err
				}
				result = append(result, fieldBytes...)
			} else {
				// Handle unexported fields
				fieldBytes, err := handleUnexportedField(fieldValue)
				if err != nil {
					return nil, err
				}
				result = append(result, fieldBytes...)
			}
		}

	case reflect.String:
		str := val.String()
		if str != "" && str != "0" && str != "v1" && str != "pod" && str != "Pod" {
			result = append(result, str...)
			result = append(result, '\n')
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Int() != 0 {
			intStr := fmt.Sprintf("%d", val.Int())
			result = append(result, intStr...)
			result = append(result, '\n')
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val.Uint() != 0 {
			uintStr := fmt.Sprintf("%d", val.Uint())
			result = append(result, uintStr...)
			result = append(result, '\n')
		}

	case reflect.Float32, reflect.Float64:
		if val.Float() != 0 {
			floatStr := fmt.Sprintf("%g", val.Float())
			result = append(result, floatStr...)
			result = append(result, '\n')
		}
	}

	return result, nil
}

func handleUnexportedField(v reflect.Value) ([]byte, error) {
	var result []byte
	switch v.Kind() {
	case reflect.String:
		str := v.String()
		if str != "" && str != "0" && str != "v1" && str != "pod" && str != "Pod" {
			result = append(result, str...)
			result = append(result, '\n')
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = append(result, fmt.Sprintf("%d", v.Int())...)
		result = append(result, '\n')
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result = append(result, fmt.Sprintf("%d", v.Uint())...)
		result = append(result, '\n')
	case reflect.Float32, reflect.Float64:
		result = append(result, fmt.Sprintf("%g", v.Float())...)
		result = append(result, '\n')
	}
	return result, nil
}

// Convert pod object to a byte slice in a format suitable for libfuzzer
func runtimeObjectToFuzzerFormat(pod interface{}) ([]byte, error) {
	return extractFieldValues(pod)
}

func logErrorToInfrastructure(errorType string, err interface{}) {

	pc, _, _, ok := run.Caller(1)
	var functionName string
	if ok {
		functionName = run.FuncForPC(pc).Name()
	} else {
		functionName = "unknown"
	}

	logEntry := ErrorLog{
		Timestamp:    time.Now().Format(time.RFC3339),
		ErrorType:    errorType,
		ErrorMessage: fmt.Sprintf("%v", err),
		FunctionName: functionName,
		CoverageData: "ExampleCoverageData", // Placeholder
	}

	appendLogEntryToCSV("error_logs.csv", logEntry)
}

func appendLogEntryToCSV(fileName string, logEntry ErrorLog) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{
		logEntry.Timestamp,
		logEntry.ErrorType,
		logEntry.ErrorMessage,
		logEntry.FunctionName,
		logEntry.CoverageData,
	}); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

func fuzzPodObjectMutator(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) *v1.Pod {
	dataSlice := C.GoBytes(dataPtrPod, C.int(dataSizePod))
	t := new(testing.T)

	//valiadtion test for unwrapping the pod object
	now := metav1.Now()
	startTime := metav1.NewTime(now.Time.Add(-1 * time.Minute))
	exceededActiveDeadlineSeconds := int64(30)
	obj1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cool",
			Namespace: "test",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{Name: "foo"},
			},
			ActiveDeadlineSeconds: &exceededActiveDeadlineSeconds,
		},
		Status: v1.PodStatus{
			StartTime: &startTime,
		},
	}

	//validation wrapper ends here

	//dataslice from the protobuf is wrapped around the runtime unknown object

	wire1, err := (&runtime.Unknown{
		TypeMeta: runtime.TypeMeta{Kind: "Pod", APIVersion: "v1"},
		Raw:      dataSlice,
	}).Marshal()
	if err != nil {
		t.Fatal(err)
	}

	wire1 = append([]byte{0x6b, 0x38, 0x73, 0x00}, wire1...)
	obj1WithKind := obj1.DeepCopyObject()
	obj1WithKind.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Pod"})
	testCases := []struct {
		obj   runtime.Object
		data  []byte
		errFn func(error) bool
	}{
		{
			obj:  obj1WithKind,
			data: wire1,
		},
	}
	scheme := runtime.NewScheme()
	returnPod := &v1.Pod{}
	for i, test := range testCases {
		scheme.AddKnownTypes(schema.GroupVersion{Version: "v1"}, &v1.Pod{})
		s := protobuf.NewSerializer(scheme, scheme)
		obj, err := runtime.Decode(s, test.data)

		switch {
		case err == nil && test.errFn != nil:
			fmt.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil && test.errFn == nil:
			fmt.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil:
			if !test.errFn(err) {
				fmt.Errorf("%d: failed: %v", i, err)
				logErrorToInfrastructure("panic", err)
			}
			if obj != nil {
				fmt.Errorf("%d: should not have returned an object", i)
			}
			continue
		}
		if pod, ok := obj.(*v1.Pod); ok {
			returnPod = pod

		}

	}
	return returnPod
}

func fuzzNodeObjectMutator(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) *v1.Node {
	dataSlice := C.GoBytes(dataPtrPod, C.int(dataSizePod))
	t := new(testing.T)

	//valiadtion test for unwrapping the Node object
	obj1 := &v1.Node{}

	//validation wrapper ends here

	//dataslice from the protobuf is wrapped around the runtime unknown object

	wire1, err := (&runtime.Unknown{
		TypeMeta: runtime.TypeMeta{Kind: "Node", APIVersion: "v1"},
		Raw:      dataSlice,
	}).Marshal()
	if err != nil {
		t.Fatal(err)
	}

	wire1 = append([]byte{0x6b, 0x38, 0x73, 0x00}, wire1...)
	obj1WithKind := obj1.DeepCopyObject()
	obj1WithKind.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Node"})
	testCases := []struct {
		obj   runtime.Object
		data  []byte
		errFn func(error) bool
	}{
		{
			obj:  obj1WithKind,
			data: wire1,
		},
	}
	scheme := runtime.NewScheme()
	returnNode := &v1.Node{}
	for i, test := range testCases {
		scheme.AddKnownTypes(schema.GroupVersion{Version: "v1"}, &v1.Node{})
		s := protobuf.NewSerializer(scheme, scheme)
		obj, err := runtime.Decode(s, test.data)

		switch {
		case err == nil && test.errFn != nil:
			fmt.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil && test.errFn == nil:
			fmt.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil:
			if !test.errFn(err) {
				fmt.Errorf("%d: failed: %v", i, err)
				logErrorToInfrastructure("panic", err)
			}
			if obj != nil {
				fmt.Errorf("%d: should not have returned an object", i)
			}
			continue
		}
		if node, ok := obj.(*v1.Node); ok {
			returnNode = node

		}

	}
	return returnNode
}

func fuzzContainerStatusObjectMutator(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) *v1.Pod {
	dataSlice := C.GoBytes(dataPtrPod, C.int(dataSizePod))
	t := new(testing.T)
	//valiadtion test for unwrapping the pod object
	now := metav1.Now()
	startTime := metav1.NewTime(now.Time.Add(-1 * time.Minute))
	exceededActiveDeadlineSeconds := int64(30)

	// //valiadtion test for unwrapping the pod object
	// now := metav1.Now()
	// startTime := metav1.NewTime(now.Time.Add(-1 * time.Minute))
	obj1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cool",
			Namespace: "test",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{Name: "foo"},
			},
			ActiveDeadlineSeconds: &exceededActiveDeadlineSeconds,
		},
		Status: v1.PodStatus{
			StartTime: &startTime,
		},
	}

	//validation wrapper ends here

	//dataslice from the protobuf is wrapped around the runtime unknown object

	wire1, err := (&runtime.Unknown{
		TypeMeta: runtime.TypeMeta{Kind: "Pod", APIVersion: "v1"},
		Raw:      dataSlice,
	}).Marshal()
	if err != nil {
		t.Fatal(err)
	}
	obj1wire, err := obj1.Marshal()
	fmt.Printf("Go bytes in hex format: %x\n", obj1wire)
	fmt.Printf("Go bytes in hex format: %x\n", dataSlice)
	wire1 = append([]byte{0x6b, 0x38, 0x73, 0x00}, wire1...)
	obj1WithKind := obj1.DeepCopyObject()
	obj1WithKind.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Pod"})
	testCases := []struct {
		obj   runtime.Object
		data  []byte
		errFn func(error) bool
	}{
		{
			obj:  obj1WithKind,
			data: wire1,
		},
	}
	scheme := runtime.NewScheme()
	returnPod := &v1.Pod{}
	for i, test := range testCases {
		scheme.AddKnownTypes(schema.GroupVersion{Version: "v1"}, &v1.Pod{})
		s := protobuf.NewSerializer(scheme, scheme)
		obj, err := runtime.Decode(s, test.data)

		switch {
		case err == nil && test.errFn != nil:
			fmt.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil && test.errFn == nil:
			fmt.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil:
			if !test.errFn(err) {
				fmt.Errorf("%d: failed: %v", i, err)
				logErrorToInfrastructure("panic", err)
			}
			if obj != nil {
				fmt.Errorf("%d: should not have returned an object", i)
			}
			continue
		}
		if pod, ok := obj.(*v1.Pod); ok {
			returnPod = pod

		}

	}
	return returnPod
}

//export DoesNotDeletePodDirsIfContainerIsRunning
func DoesNotDeletePodDirsIfContainerIsRunning(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		// if r := recover(); r != nil {
		// 	err, ok := r.(error)
		// 	errMsg := ""
		// 	if ok {
		// 		errMsg = err.Error()
		// 	} else {
		// 		errMsg = fmt.Sprint(r)
		// 	}
		// 	if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
		// 		fmt.Println("Excluding nil pointer dereference error from logs.")
		// 		return
		// 	}
		// 	logErrorToInfrastructure("panic", errMsg)
		// } else {
		hash, err := generateRandomHash()
		if err != nil {
			fmt.Printf("Error generating hash: %v\n", err)
			return
		}
		klog.Info("object hash: error-%s", hash)
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

		err = storePodInDatabasePod(pod, "error-"+hash)
		if err != nil {
			logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
			return
		}

		podBytes, err := runtimeObjectToFuzzerFormat(pod)
		if err != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
			return
		}

		// Write the pod bytes to a file (keeping this part as it was in the original function)
		file, err := os.OpenFile("DoesNotDeletePodDirsIfContainerIsRunning.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
			return
		}
		defer file.Close()

		if _, err := file.Write(podBytes); err != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
			return
		}
		// }
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestDoesNotDeletePodDirsIfContainerIsRunning(t, pod)
}

//export SyncPodsSetStatusToFailedForPodsThatRunTooLong
func SyncPodsSetStatusToFailedForPodsThatRunTooLong(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		// if r := recover(); r != nil {
		// 	err, ok := r.(error)
		// 	errMsg := ""
		// 	if ok {
		// 		errMsg = err.Error()
		// 	} else {
		// 		errMsg = fmt.Sprint(r)
		// 	}
		// 	if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
		// 		fmt.Println("Excluding nil pointer dereference error from logs.")
		// 		return
		// 	}
		// 	logErrorToInfrastructure("panic", errMsg)
		// } else {
		hash, err := generateRandomHash()
		if err != nil {
			fmt.Printf("Error generating hash: %v\n", err)
			return
		}
		klog.Info("object hash: error-%s", hash)
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

		err = storePodInDatabasePod(pod, "error-"+hash)
		if err != nil {
			logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
			return
		}

		podBytes, err := runtimeObjectToFuzzerFormat(pod)
		if err != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
			return
		}

		// Write the pod bytes to a file (keeping this part as it was in the original function)
		file, err := os.OpenFile("SyncPodsSetStatusToFailedForPodsThatRunTooLong.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
			return
		}
		defer file.Close()

		if _, err := file.Write(podBytes); err != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
			return
		}
		// }
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestSyncPodsSetStatusToFailedForPodsThatRunTooLong(t, pod)

}

//export SyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed
func SyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("SyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestSyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed(t, pod)

}

//export TestSyncPodsStartPod
func TestSyncPodsStartPod(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestSyncPodsStartPod.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestSyncPodsStartPod(t, pod)

}

//export TestDispatchWorkOfCompletedPod
func TestDispatchWorkOfCompletedPod(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestDispatchWorkOfCompletedPod.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestDispatchWorkOfCompletedPod(t, pod)
}

//export TestDispatchWorkOfActivePod
func TestDispatchWorkOfActivePod(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestDispatchWorkOfActivePod.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestDispatchWorkOfActivePod(t, pod)
}

//export TestHandlePodRemovesWhenSourcesAreReady
func TestHandlePodRemovesWhenSourcesAreReady(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestHandlePodRemovesWhenSourcesAreReady.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestHandlePodRemovesWhenSourcesAreReady(t, pod)
}

//export TestHandlePortConflicts
func TestHandlePortConflicts(dataPtrPod unsafe.Pointer, dataSizePod C.size_t, dataPtrNode unsafe.Pointer, dataSizeNode C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: ", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
			node := fuzzNodeObjectMutator(dataPtrNode, dataSizeNode)
			err = storePodInDatabasePod(pod, "error-pod-"+hash)
			err = storePodInDatabaseNode(node, "error-node-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			nodeBytes, err := runtimeObjectToFuzzerFormat(node)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestHandlePortConflicts.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
			if _, err := file.Write(nodeBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()

	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	node := fuzzNodeObjectMutator(dataPtrNode, dataSizeNode)
	t := new(testing.T)

	kubelet.TestHandlePortConflicts(t, node, pod)
}

//export TestHandleHostNameConflicts
func TestHandleHostNameConflicts(dataPtrPod unsafe.Pointer, dataSizePod C.size_t, dataPtrNode unsafe.Pointer, dataSizeNode C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestHandleHostNameConflicts.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	node := fuzzNodeObjectMutator(dataPtrNode, dataSizeNode)
	t := new(testing.T)

	kubelet.TestHandleHostNameConflicts(t, node, pod)
}

//export TestHandleNodeSelectorBasedOnOS
func TestHandleNodeSelectorBasedOnOS(dataPtrPod unsafe.Pointer, dataSizePod C.size_t, dataPtrNode unsafe.Pointer, dataSizeNode C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestHandleNodeSelectorBasedOnOS.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	node := fuzzNodeObjectMutator(dataPtrNode, dataSizeNode)
	t := new(testing.T)

	kubelet.TestHandleNodeSelectorBasedOnOS(t, node, pod)
}

//export TestHandleMemExceeded
func TestHandleMemExceeded(dataPtrPod unsafe.Pointer, dataSizePod C.size_t, dataPtrNode unsafe.Pointer, dataSizeNode C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestHandleMemExceeded.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	node := fuzzNodeObjectMutator(dataPtrNode, dataSizeNode)
	t := new(testing.T)

	kubelet.TestHandleMemExceeded(t, node, pod)
}

//export TestHandlePluginResources
func TestHandlePluginResources(dataPtrPod unsafe.Pointer, dataSizePod C.size_t, dataPtrNode unsafe.Pointer, dataSizeNode C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestHandlePluginResources.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()
	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	node := fuzzNodeObjectMutator(dataPtrNode, dataSizeNode)
	t := new(testing.T)

	kubelet.TestHandlePluginResources(t, node, pod)
}

//export TestPurgingObsoleteStatusMapEntries
func TestPurgingObsoleteStatusMapEntries(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestPurgingObsoleteStatusMapEntries.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)

	kubelet.TestPurgingObsoleteStatusMapEntries(t, pod)
}

//export TestValidateContainerLogStatus
func TestValidateContainerLogStatus(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestValidateContainerLogStatus.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	status := fuzzContainerStatusObjectMutator(dataPtrPod, dataSizePod)
	// t := new(testing.T)
	fmt.Println(status)
	// kubelet.TestValidateContainerLogStatus(t, status)
}

//export TestCreateMirrorPod
func TestCreateMirrorPod(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestCreateMirrorPod.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	fmt.Println(pod)
	kubelet.TestCreateMirrorPod(t, pod)
}

//export TestDeleteOutdatedMirrorPod
func TestDeleteOutdatedMirrorPod(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestDeleteOutdatedMirrorPod.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	fmt.Println(pod)
	kubelet.TestDeleteOutdatedMirrorPod(t, pod)
}

//export TestDeleteOrphanedMirrorPods
func TestDeleteOrphanedMirrorPods(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestDeleteOrphanedMirrorPods.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	fmt.Println(pod)
	kubelet.TestDeleteOrphanedMirrorPods(t, pod)
}

//export TestNetworkErrorsWithoutHostNetwork
func TestNetworkErrorsWithoutHostNetwork(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestNetworkErrorsWithoutHostNetwork.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	fmt.Println(pod)
	kubelet.TestNetworkErrorsWithoutHostNetwork(t, pod)
}

//export TestFilterOutInactivePods
func TestFilterOutInactivePods(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestFilterOutInactivePods.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	fmt.Println(pod)
	kubelet.TestFilterOutInactivePods(t, pod)
}

//export TestSyncPodsSetStatusToFailedForPodsThatRunTooLong
func TestSyncPodsSetStatusToFailedForPodsThatRunTooLong(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestSyncPodsSetStatusToFailedForPodsThatRunTooLong.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestSyncPodsSetStatusToFailedForPodsThatRunTooLong(t, pod)
}

//export TestDeletePodDirsForDeletedPods
func TestDeletePodDirsForDeletedPods(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestDeletePodDirsForDeletedPods.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestDeletePodDirsForDeletedPods(t, pod)
}

//export TestDoesNotDeletePodDirsForTerminatedPods
func TestDoesNotDeletePodDirsForTerminatedPods(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestDoesNotDeletePodDirsForTerminatedPods.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()
	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestDoesNotDeletePodDirsForTerminatedPods(t, pod)
}

//export TestDoesNotDeletePodDirsIfContainerIsRunning
func TestDoesNotDeletePodDirsIfContainerIsRunning(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestDoesNotDeletePodDirsIfContainerIsRunning.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestDoesNotDeletePodDirsIfContainerIsRunning(t, pod)
}

//export TestGetPodsToSync
func TestGetPodsToSync(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestGetPodsToSync.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	klog.InfoS("etest")
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestGetPodsToSync(t, pod)
}

//export TestGenerateAPIPodStatusWithSortedContainers
func TestGenerateAPIPodStatusWithSortedContainers(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestGenerateAPIPodStatusWithSortedContainers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestGenerateAPIPodStatusWithSortedContainers(t, pod)
}

//export TestGenerateAPIPodStatusWithReasonCache
func TestGenerateAPIPodStatusWithReasonCache(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestGenerateAPIPodStatusWithReasonCache.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestGenerateAPIPodStatusWithReasonCache(t, pod)
}

//export TestGenerateAPIPodStatusWithDifferentRestartPolicies
func TestGenerateAPIPodStatusWithDifferentRestartPolicies(dataPtrPod unsafe.Pointer, dataSizePod C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestGenerateAPIPodStatusWithDifferentRestartPolicies.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	t := new(testing.T)
	kubelet.TestGenerateAPIPodStatusWithDifferentRestartPolicies(t, pod)
}

//export TestHandlePodAdditionsInvokesPodAdmitHandlers
func TestHandlePodAdditionsInvokesPodAdmitHandlers(dataPtrPod unsafe.Pointer, dataSizePod C.size_t, dataPtrNode unsafe.Pointer, dataSizeNode C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestHandlePodAdditionsInvokesPodAdmitHandlers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	node := fuzzNodeObjectMutator(dataPtrNode, dataSizeNode)
	t := new(testing.T)

	kubelet.TestHandlePodAdditionsInvokesPodAdmitHandlers(t, node, pod)
}

//export TestPodResourceAllocationReset
func TestPodResourceAllocationReset(dataPtrPod unsafe.Pointer, dataSizePod C.size_t, dataPtrNode unsafe.Pointer, dataSizeNode C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			hash, err := generateRandomHash()
			if err != nil {
				fmt.Printf("Error generating hash: %v\n", err)
				return
			}
			klog.Info("object hash: error-%s", hash)
			pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)

			err = storePodInDatabasePod(pod, "error-"+hash)
			if err != nil {
				logErrorToInfrastructure("database_error", fmt.Sprintf("Failed to store pod in database: %v", err))
				return
			}

			podBytes, err := runtimeObjectToFuzzerFormat(pod)
			if err != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err))
				return
			}

			// Write the pod bytes to a file (keeping this part as it was in the original function)
			file, err := os.OpenFile("TestPodResourceAllocationReset.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err))
				return
			}
			defer file.Close()

			if _, err := file.Write(podBytes); err != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err))
				return
			}
		}
	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	node := fuzzNodeObjectMutator(dataPtrNode, dataSizeNode)
	t := new(testing.T)

	kubelet.TestPodResourceAllocationReset(t, node, pod)
}

func main() {

	select {}
}
