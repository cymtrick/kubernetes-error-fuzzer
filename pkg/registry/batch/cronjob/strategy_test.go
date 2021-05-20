/*
Copyright 2016 The Kubernetes Authors.

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

package cronjob

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/kubernetes/pkg/apis/batch"
	api "k8s.io/kubernetes/pkg/apis/core"
)

func TestCronJobStrategy(t *testing.T) {
	ctx := genericapirequest.NewDefaultContext()
	if !Strategy.NamespaceScoped() {
		t.Errorf("CronJob must be namespace scoped")
	}
	if Strategy.AllowCreateOnUpdate() {
		t.Errorf("CronJob should not allow create on update")
	}

	validPodTemplateSpec := api.PodTemplateSpec{
		Spec: api.PodSpec{
			RestartPolicy: api.RestartPolicyOnFailure,
			DNSPolicy:     api.DNSClusterFirst,
			Containers:    []api.Container{{Name: "abc", Image: "image", ImagePullPolicy: "IfNotPresent", TerminationMessagePolicy: api.TerminationMessageReadFile}},
		},
	}
	cronJob := &batch.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "mycronjob",
			Namespace:  metav1.NamespaceDefault,
			Generation: 999,
		},
		Spec: batch.CronJobSpec{
			Schedule:          "* * * * ?",
			ConcurrencyPolicy: batch.AllowConcurrent,
			JobTemplate: batch.JobTemplateSpec{
				Spec: batch.JobSpec{
					Template: validPodTemplateSpec,
				},
			},
		},
	}

	Strategy.PrepareForCreate(ctx, cronJob)
	if len(cronJob.Status.Active) != 0 {
		t.Errorf("CronJob does not allow setting status on create")
	}
	if cronJob.Generation != 1 {
		t.Errorf("expected Generation=1, got %d", cronJob.Generation)
	}
	errs := Strategy.Validate(ctx, cronJob)
	if len(errs) != 0 {
		t.Errorf("Unexpected error validating %v", errs)
	}
	now := metav1.Now()

	// ensure we do not change generation for non-spec updates
	updatedLabelCronJob := cronJob.DeepCopy()
	updatedLabelCronJob.Labels = map[string]string{"a": "true"}
	Strategy.PrepareForUpdate(ctx, updatedLabelCronJob, cronJob)
	if updatedLabelCronJob.Generation != 1 {
		t.Errorf("expected Generation=1, got %d", updatedLabelCronJob.Generation)
	}

	updatedCronJob := &batch.CronJob{
		ObjectMeta: metav1.ObjectMeta{Name: "bar", ResourceVersion: "4"},
		Spec: batch.CronJobSpec{
			Schedule: "5 5 5 * ?",
		},
		Status: batch.CronJobStatus{
			LastScheduleTime: &now,
		},
	}

	// ensure we do not change status
	Strategy.PrepareForUpdate(ctx, updatedCronJob, cronJob)
	if updatedCronJob.Status.Active != nil {
		t.Errorf("PrepareForUpdate should have preserved prior version status")
	}
	if updatedCronJob.Generation != 2 {
		t.Errorf("expected Generation=2, got %d", updatedCronJob.Generation)
	}
	errs = Strategy.ValidateUpdate(ctx, updatedCronJob, cronJob)
	if len(errs) == 0 {
		t.Errorf("Expected a validation error")
	}

	// Make sure we correctly implement the interface.
	// Otherwise a typo could silently change the default.
	var gcds rest.GarbageCollectionDeleteStrategy = Strategy
	if got, want := gcds.DefaultGarbageCollectionPolicy(genericapirequest.NewContext()), rest.DeleteDependents; got != want {
		t.Errorf("DefaultGarbageCollectionPolicy() = %#v, want %#v", got, want)
	}

	var (
		v1beta1Ctx      = genericapirequest.WithRequestInfo(genericapirequest.NewContext(), &genericapirequest.RequestInfo{APIGroup: "batch", APIVersion: "v1beta1", Resource: "cronjobs"})
		otherVersionCtx = genericapirequest.WithRequestInfo(genericapirequest.NewContext(), &genericapirequest.RequestInfo{APIGroup: "batch", APIVersion: "v100", Resource: "cronjobs"})
	)
	if got, want := gcds.DefaultGarbageCollectionPolicy(v1beta1Ctx), rest.OrphanDependents; got != want {
		t.Errorf("DefaultGarbageCollectionPolicy() = %#v, want %#v", got, want)
	}
	if got, want := gcds.DefaultGarbageCollectionPolicy(otherVersionCtx), rest.DeleteDependents; got != want {
		t.Errorf("DefaultGarbageCollectionPolicy() = %#v, want %#v", got, want)
	}
}

func TestCronJobStatusStrategy(t *testing.T) {
	ctx := genericapirequest.NewDefaultContext()
	if !StatusStrategy.NamespaceScoped() {
		t.Errorf("CronJob must be namespace scoped")
	}
	if StatusStrategy.AllowCreateOnUpdate() {
		t.Errorf("CronJob should not allow create on update")
	}
	validPodTemplateSpec := api.PodTemplateSpec{
		Spec: api.PodSpec{
			RestartPolicy: api.RestartPolicyOnFailure,
			DNSPolicy:     api.DNSClusterFirst,
			Containers:    []api.Container{{Name: "abc", Image: "image", ImagePullPolicy: "IfNotPresent", TerminationMessagePolicy: api.TerminationMessageReadFile}},
		},
	}
	oldSchedule := "* * * * ?"
	oldCronJob := &batch.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:            "mycronjob",
			Namespace:       metav1.NamespaceDefault,
			ResourceVersion: "10",
		},
		Spec: batch.CronJobSpec{
			Schedule:          oldSchedule,
			ConcurrencyPolicy: batch.AllowConcurrent,
			JobTemplate: batch.JobTemplateSpec{
				Spec: batch.JobSpec{
					Template: validPodTemplateSpec,
				},
			},
		},
	}
	now := metav1.Now()
	newCronJob := &batch.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:            "mycronjob",
			Namespace:       metav1.NamespaceDefault,
			ResourceVersion: "9",
		},
		Spec: batch.CronJobSpec{
			Schedule:          "5 5 * * ?",
			ConcurrencyPolicy: batch.AllowConcurrent,
			JobTemplate: batch.JobTemplateSpec{
				Spec: batch.JobSpec{
					Template: validPodTemplateSpec,
				},
			},
		},
		Status: batch.CronJobStatus{
			LastScheduleTime: &now,
		},
	}

	StatusStrategy.PrepareForUpdate(ctx, newCronJob, oldCronJob)
	if newCronJob.Status.LastScheduleTime == nil {
		t.Errorf("CronJob status updates must allow changes to cronJob status")
	}
	if newCronJob.Spec.Schedule != oldSchedule {
		t.Errorf("CronJob status updates must now allow changes to cronJob spec")
	}
	errs := StatusStrategy.ValidateUpdate(ctx, newCronJob, oldCronJob)
	if len(errs) != 0 {
		t.Errorf("Unexpected error %v", errs)
	}
	if newCronJob.ResourceVersion != "9" {
		t.Errorf("Incoming resource version on update should not be mutated")
	}
}
