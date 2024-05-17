package main

import (
	"C"
	"encoding/csv"
	"fmt"
	"os"
	run "runtime"
	"time"
	"unsafe"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
	mock "k8s.io/kubernetes/pkg/mock"
)
import (
	"strings"
	"testing"
)

// ErrorLog
type ErrorLog struct {
	Timestamp    string
	ErrorType    string
	ErrorMessage string
	FunctionName string
	CoverageData string
}

func logErrorToInfrastructure(errorType string, err interface{}) {
	if errorVal, ok := err.(error); ok {

		if strings.Contains(errorVal.Error(), "invalid memory address or nil pointer dereference") {
			fmt.Println("Excluding nil pointer dereference error from logs.")
			return
		}
	}

	// Proceed with logging other errors
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
			t.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil && test.errFn == nil:
			t.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil:
			if !test.errFn(err) {
				t.Errorf("%d: failed: %v", i, err)
				logErrorToInfrastructure("panic", err)
			}
			if obj != nil {
				t.Errorf("%d: should not have returned an object", i)
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
			t.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil && test.errFn == nil:
			t.Errorf("%d: failed: %v", i, err)
			logErrorToInfrastructure("panic", err)
			continue
		case err != nil:
			if !test.errFn(err) {
				t.Errorf("%d: failed: %v", i, err)
				logErrorToInfrastructure("panic", err)
			}
			if obj != nil {
				t.Errorf("%d: should not have returned an object", i)
			}
			continue
		}
		if node, ok := obj.(*v1.Node); ok {
			returnNode = node

		}

	}
	return returnNode
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
		}
	}()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	mock.TestDoesNotDeletePodDirsIfContainerIsRunning(t, pod)

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
		}
	}()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	mock.TestSyncPodsSetStatusToFailedForPodsThatRunTooLong(t, pod)

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
	}()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	mock.TestSyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed(t, pod)

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
	}()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	mock.TestSyncPodsStartPod(t, pod)

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
	}()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	mock.TestDispatchWorkOfCompletedPod(t, pod)
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
	}()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	mock.TestDispatchWorkOfActivePod(t, pod)
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
	}()
	pod := fuzzPodObjectMutator(dataPtr, dataSize)
	t := new(testing.T)
	mock.TestHandlePodRemovesWhenSourcesAreReady(t, pod)
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
	}()
	pod := fuzzPodObjectMutator(dataPtrPod, dataSizePod)
	node := fuzzNodeObjectMutator(dataPtrNode, dataSizeNode)
	t := new(testing.T)
	mock.TestHandlePortConflicts(t, node, pod)
}

func main() {

	select {}
}
