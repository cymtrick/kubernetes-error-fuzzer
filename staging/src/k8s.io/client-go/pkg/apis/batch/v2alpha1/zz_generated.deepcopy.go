// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v2alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/api/core/v1"
	batch_v1 "k8s.io/api/batch/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v2alpha1_CronJob, InType: reflect.TypeOf(&CronJob{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v2alpha1_CronJobList, InType: reflect.TypeOf(&CronJobList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v2alpha1_CronJobSpec, InType: reflect.TypeOf(&CronJobSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v2alpha1_CronJobStatus, InType: reflect.TypeOf(&CronJobStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v2alpha1_JobTemplate, InType: reflect.TypeOf(&JobTemplate{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v2alpha1_JobTemplateSpec, InType: reflect.TypeOf(&JobTemplateSpec{})},
	)
}

// DeepCopy_v2alpha1_CronJob is an autogenerated deepcopy function.
func DeepCopy_v2alpha1_CronJob(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJob)
		out := out.(*CronJob)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v2alpha1_CronJobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v2alpha1_CronJobStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v2alpha1_CronJobList is an autogenerated deepcopy function.
func DeepCopy_v2alpha1_CronJobList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJobList)
		out := out.(*CronJobList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CronJob, len(*in))
			for i := range *in {
				if err := DeepCopy_v2alpha1_CronJob(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v2alpha1_CronJobSpec is an autogenerated deepcopy function.
func DeepCopy_v2alpha1_CronJobSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJobSpec)
		out := out.(*CronJobSpec)
		*out = *in
		if in.StartingDeadlineSeconds != nil {
			in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
			*out = new(int64)
			**out = **in
		}
		if in.Suspend != nil {
			in, out := &in.Suspend, &out.Suspend
			*out = new(bool)
			**out = **in
		}
		if err := DeepCopy_v2alpha1_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, c); err != nil {
			return err
		}
		if in.SuccessfulJobsHistoryLimit != nil {
			in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
			*out = new(int32)
			**out = **in
		}
		if in.FailedJobsHistoryLimit != nil {
			in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v2alpha1_CronJobStatus is an autogenerated deepcopy function.
func DeepCopy_v2alpha1_CronJobStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJobStatus)
		out := out.(*CronJobStatus)
		*out = *in
		if in.Active != nil {
			in, out := &in.Active, &out.Active
			*out = make([]api_v1.ObjectReference, len(*in))
			copy(*out, *in)
		}
		if in.LastScheduleTime != nil {
			in, out := &in.LastScheduleTime, &out.LastScheduleTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		return nil
	}
}

// DeepCopy_v2alpha1_JobTemplate is an autogenerated deepcopy function.
func DeepCopy_v2alpha1_JobTemplate(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobTemplate)
		out := out.(*JobTemplate)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v2alpha1_JobTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v2alpha1_JobTemplateSpec is an autogenerated deepcopy function.
func DeepCopy_v2alpha1_JobTemplateSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobTemplateSpec)
		out := out.(*JobTemplateSpec)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := batch_v1.DeepCopy_v1_JobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}
