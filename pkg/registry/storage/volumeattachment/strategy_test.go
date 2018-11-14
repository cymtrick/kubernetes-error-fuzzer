/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package volumeattachment

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/diff"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/kubernetes/pkg/apis/storage"
)

func getValidVolumeAttachment(name string) *storage.VolumeAttachment {
	return &storage.VolumeAttachment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: storage.VolumeAttachmentSpec{
			Attacher: "valid-attacher",
			Source: storage.VolumeAttachmentSource{
				PersistentVolumeName: &name,
			},
			NodeName: "valid-node",
		},
	}
}

func TestVolumeAttachmentStrategy(t *testing.T) {
	ctx := genericapirequest.WithRequestInfo(genericapirequest.NewContext(), &genericapirequest.RequestInfo{
		APIGroup:   "storage.k8s.io",
		APIVersion: "v1",
		Resource:   "volumeattachments",
	})
	if Strategy.NamespaceScoped() {
		t.Errorf("VolumeAttachment must not be namespace scoped")
	}
	if Strategy.AllowCreateOnUpdate() {
		t.Errorf("VolumeAttachment should not allow create on update")
	}

	volumeAttachment := getValidVolumeAttachment("valid-attachment")

	Strategy.PrepareForCreate(ctx, volumeAttachment)

	errs := Strategy.Validate(ctx, volumeAttachment)
	if len(errs) != 0 {
		t.Errorf("unexpected error validating %v", errs)
	}

	// Create with status should drop status
	statusVolumeAttachment := volumeAttachment.DeepCopy()
	statusVolumeAttachment.Status = storage.VolumeAttachmentStatus{Attached: true}
	Strategy.PrepareForCreate(ctx, statusVolumeAttachment)
	if !apiequality.Semantic.DeepEqual(statusVolumeAttachment, volumeAttachment) {
		t.Errorf("unexpected objects difference after creating with status: %v", diff.ObjectDiff(statusVolumeAttachment, volumeAttachment))
	}

	// Update of spec is disallowed
	newVolumeAttachment := volumeAttachment.DeepCopy()
	newVolumeAttachment.Spec.NodeName = "valid-node-2"

	Strategy.PrepareForUpdate(ctx, newVolumeAttachment, volumeAttachment)

	errs = Strategy.ValidateUpdate(ctx, newVolumeAttachment, volumeAttachment)
	if len(errs) == 0 {
		t.Errorf("Expected a validation error")
	}

	// modifying status should be dropped
	statusVolumeAttachment = volumeAttachment.DeepCopy()
	statusVolumeAttachment.Status = storage.VolumeAttachmentStatus{Attached: true}

	Strategy.PrepareForUpdate(ctx, statusVolumeAttachment, volumeAttachment)

	if !apiequality.Semantic.DeepEqual(statusVolumeAttachment, volumeAttachment) {
		t.Errorf("unexpected objects difference after modfying status: %v", diff.ObjectDiff(statusVolumeAttachment, volumeAttachment))
	}
}

func TestVolumeAttachmentStatusStrategy(t *testing.T) {
	ctx := genericapirequest.WithRequestInfo(genericapirequest.NewContext(), &genericapirequest.RequestInfo{
		APIGroup:   "storage.k8s.io",
		APIVersion: "v1",
		Resource:   "volumeattachments",
	})

	volumeAttachment := getValidVolumeAttachment("valid-attachment")

	// modifying status should be allowed
	statusVolumeAttachment := volumeAttachment.DeepCopy()
	statusVolumeAttachment.Status = storage.VolumeAttachmentStatus{Attached: true}

	expectedVolumeAttachment := statusVolumeAttachment.DeepCopy()
	StatusStrategy.PrepareForUpdate(ctx, statusVolumeAttachment, volumeAttachment)
	if !apiequality.Semantic.DeepEqual(statusVolumeAttachment, expectedVolumeAttachment) {
		t.Errorf("unexpected objects differerence after modifying status: %v", diff.ObjectDiff(statusVolumeAttachment, expectedVolumeAttachment))
	}

	// modifying spec should be dropped
	newVolumeAttachment := volumeAttachment.DeepCopy()
	newVolumeAttachment.Spec.NodeName = "valid-node-2"

	StatusStrategy.PrepareForUpdate(ctx, newVolumeAttachment, volumeAttachment)
	if !apiequality.Semantic.DeepEqual(newVolumeAttachment, volumeAttachment) {
		t.Errorf("unexpected objects differerence after modifying spec: %v", diff.ObjectDiff(newVolumeAttachment, volumeAttachment))
	}
}

func TestBetaAndV1StatusUpdate(t *testing.T) {
	tests := []struct {
		requestInfo    genericapirequest.RequestInfo
		newStatus      bool
		expectedStatus bool
	}{
		{
			genericapirequest.RequestInfo{
				APIGroup:   "storage.k8s.io",
				APIVersion: "v1",
				Resource:   "volumeattachments",
			},
			true,
			false,
		},
		{
			genericapirequest.RequestInfo{
				APIGroup:   "storage.k8s.io",
				APIVersion: "v1beta1",
				Resource:   "volumeattachments",
			},
			true,
			true,
		},
	}
	for _, test := range tests {
		va := getValidVolumeAttachment("valid-attachment")
		newAttachment := va.DeepCopy()
		newAttachment.Status.Attached = test.newStatus
		context := genericapirequest.WithRequestInfo(genericapirequest.NewContext(), &test.requestInfo)
		Strategy.PrepareForUpdate(context, newAttachment, va)
		if newAttachment.Status.Attached != test.expectedStatus {
			t.Errorf("expected status to be %v got %v", test.expectedStatus, newAttachment.Status.Attached)
		}
	}

}

func TestBetaAndV1StatusCreate(t *testing.T) {
	tests := []struct {
		requestInfo    genericapirequest.RequestInfo
		newStatus      bool
		expectedStatus bool
	}{
		{
			genericapirequest.RequestInfo{
				APIGroup:   "storage.k8s.io",
				APIVersion: "v1",
				Resource:   "volumeattachments",
			},
			true,
			false,
		},
		{
			genericapirequest.RequestInfo{
				APIGroup:   "storage.k8s.io",
				APIVersion: "v1beta1",
				Resource:   "volumeattachments",
			},
			true,
			true,
		},
	}
	for _, test := range tests {
		va := getValidVolumeAttachment("valid-attachment")
		va.Status.Attached = test.newStatus
		context := genericapirequest.WithRequestInfo(genericapirequest.NewContext(), &test.requestInfo)
		Strategy.PrepareForCreate(context, va)
		if va.Status.Attached != test.expectedStatus {
			t.Errorf("expected status to be %v got %v", test.expectedStatus, va.Status.Attached)
		}
	}
}
