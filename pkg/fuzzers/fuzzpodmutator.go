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
	"reflect"
)

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
		flag.Set("log_file", "myfile.log")
		flag.Parse()
		klog.Info("klog initialized")
	})
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
		if str != "" {
			result = append(result, str...)
			result = append(result, '\n')
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intStr := fmt.Sprintf("%d", val.Int())
		result = append(result, intStr...)
		result = append(result, '\n')

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintStr := fmt.Sprintf("%d", val.Uint())
		result = append(result, uintStr...)
		result = append(result, '\n')

	case reflect.Float32, reflect.Float64:
		floatStr := fmt.Sprintf("%g", val.Float())
		result = append(result, floatStr...)
		result = append(result, '\n')
	}

	return result, nil
}

func handleUnexportedField(v reflect.Value) ([]byte, error) {
	var result []byte
	switch v.Kind() {
	case reflect.String:
		result = append(result, v.String()...)
		result = append(result, '\n')
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
func podToFuzzerFormat(pod interface{}) ([]byte, error) {
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

func fuzzPodObjectMutator(dataPtr unsafe.Pointer, dataSize C.size_t) *v1.Pod {
	dataSlice := C.GoBytes(dataPtr, C.int(dataSize))
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

func fuzzNodeObjectMutator(dataPtr unsafe.Pointer, dataSize C.size_t) *v1.Node {
	dataSlice := C.GoBytes(dataPtr, C.int(dataSize))
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

func fuzzContainerStatusObjectMutator(dataPtr unsafe.Pointer, dataSize C.size_t) *v1.Pod {
	dataSlice := C.GoBytes(dataPtr, C.int(dataSize))
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
func DoesNotDeletePodDirsIfContainerIsRunning(dataPtr unsafe.Pointer, dataSize C.size_t) {

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			pod := fuzzPodObjectMutator(dataPtr, dataSize)
			podBytes, err2 := podToFuzzerFormat(pod)
			if err2 != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
				return
			}

			// Write the pod bytes to a file
			file, err2 := os.OpenFile("DoesNotDeletePodDirsIfContainerIsRunning.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err2 != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
				return
			}

			if _, err2 := file.Write(podBytes); err2 != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
				return
			}
			defer file.Close()
		}

	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	kubelet.TestDoesNotDeletePodDirsIfContainerIsRunning(t, pod)

}

//export SyncPodsSetStatusToFailedForPodsThatRunTooLong
func SyncPodsSetStatusToFailedForPodsThatRunTooLong(dataPtr unsafe.Pointer, dataSize C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		} else {
			pod := fuzzPodObjectMutator(dataPtr, dataSize)
			podBytes, err2 := podToFuzzerFormat(pod)
			if err2 != nil {
				logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
				return
			}

			// Write the pod bytes to a file
			file, err2 := os.OpenFile("SyncPodsSetStatusToFailedForPodsThatRunTooLong.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err2 != nil {
				logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
				return
			}

			if _, err2 := file.Write(podBytes); err2 != nil {
				logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
				return
			}
			defer file.Close()
		}

	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	kubelet.TestSyncPodsSetStatusToFailedForPodsThatRunTooLong(t, pod)

}

//export SyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed
func SyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed(dataPtr unsafe.Pointer, dataSize C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtr, dataSize)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	kubelet.TestSyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed(t, pod)

}

//export TestSyncPodsStartPod
func TestSyncPodsStartPod(dataPtr unsafe.Pointer, dataSize C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtr, dataSize)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	kubelet.TestSyncPodsStartPod(t, pod)

}

//export TestDispatchWorkOfCompletedPod
func TestDispatchWorkOfCompletedPod(dataPtr unsafe.Pointer, dataSize C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtr, dataSize)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	kubelet.TestDispatchWorkOfCompletedPod(t, pod)
}

//export TestDispatchWorkOfActivePod
func TestDispatchWorkOfActivePod(dataPtr unsafe.Pointer, dataSize C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtr, dataSize)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	kubelet.TestDispatchWorkOfActivePod(t, pod)
}

//export TestHandlePodRemovesWhenSourcesAreReady
func TestHandlePodRemovesWhenSourcesAreReady(dataPtr unsafe.Pointer, dataSize C.size_t) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			errMsg := ""
			if ok {
				errMsg = err.Error()
			} else {
				errMsg = fmt.Sprint(r)
			}
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtr, dataSize)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

	}()

	initializeKlog()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
			// Ensure that nil pointer dereference errors are excluded
			if strings.Contains(errMsg, "runtime error: invalid memory address or nil pointer dereference") {
				fmt.Println("Excluding nil pointer dereference error from logs.")
				return
			}
			logErrorToInfrastructure("panic", errMsg)
		}
		pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
		podBytes, err2 := podToFuzzerFormat(pod)
		if err2 != nil {
			logErrorToInfrastructure("conversion_error", fmt.Sprintf("Failed to convert pod to fuzzer format: %v", err2))
			return
		}

		// Write the pod bytes to a file
		file, err2 := os.OpenFile("pod_fuzzer_inpu.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err2 != nil {
			logErrorToInfrastructure("file_error", fmt.Sprintf("Failed to open file: %v", err2))
			return
		}

		if _, err2 := file.Write(podBytes); err2 != nil {
			logErrorToInfrastructure("file_write_error", fmt.Sprintf("Failed to write pod to file: %v", err2))
			return
		}
		defer file.Close()

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
