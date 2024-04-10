package main

import (
	"C"
	"fmt"
	"unsafe"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
	mock "k8s.io/kubernetes/pkg/mock"
)
import (
	"testing"
	"time"
)

// FuzzUnknownObjectMutator is the exported function for fuzzing unknown object mutator.
//
//export FuzzUnknownObjectMutator
func FuzzUnknownObjectMutator(dataPtr unsafe.Pointer, dataSize C.size_t) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic in FuzzUnknownObjectMutator:", r)
	// 	}
	// }()
	dataSlice := C.GoBytes(dataPtr, C.int(dataSize))
	t := new(testing.T)
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
	obj1wire, err := obj1.Marshal()

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Go bytes in hex format: %x\n", obj1wire)
	wire1, err := (&runtime.Unknown{
		TypeMeta: runtime.TypeMeta{Kind: "Pod", APIVersion: "v1"},
		Raw:      dataSlice,
	}).Marshal()
	if err != nil {
		t.Fatal(err)
	}
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
	for i, test := range testCases {
		scheme.AddKnownTypes(schema.GroupVersion{Version: "v1"}, &v1.Pod{})
		fmt.Printf("Go bytes in hex format: %x\n", test.data)
		s := protobuf.NewSerializer(scheme, scheme)
		obj, err := runtime.Decode(s, test.data)

		switch {
		case err == nil && test.errFn != nil:
			t.Errorf("%d: failed: %v", i, err)
			continue
		case err != nil && test.errFn == nil:
			t.Errorf("%d: failed: %v", i, err)
			continue
		case err != nil:
			if !test.errFn(err) {
				t.Errorf("%d: failed: %v", i, err)
			}
			if obj != nil {
				t.Errorf("%d: should not have returned an object", i)
			}
			continue
		}
		fmt.Printf("Received data in Go: %v\n", obj)
		if pod, ok := obj.(*v1.Pod); ok {
			// mock.TestSyncPodsSetStatusToFailedForPodsThatRunTooLong(t, pod)

			// err := mock.TestSyncPodsDoesNotSetPodsThatDidNotRunTooLongToFailed(t, pod)()
			mock.TestDoesNotDeletePodDirsIfContainerIsRunning(t, pod)

		} else {
			fmt.Println("Decoded object is not a *v1.Pod")
		}

	}

}

func main() {
	// Keep this running or use a mechanism to keep it alive for fuzzing
	select {}
}
