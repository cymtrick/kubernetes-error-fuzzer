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

package batch

import (
	api "k8s.io/kubernetes/pkg/api"
	v1 "k8s.io/kubernetes/pkg/apis/meta/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_CronJob, InType: reflect.TypeOf(&CronJob{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_CronJobList, InType: reflect.TypeOf(&CronJobList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_CronJobSpec, InType: reflect.TypeOf(&CronJobSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_CronJobStatus, InType: reflect.TypeOf(&CronJobStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_Job, InType: reflect.TypeOf(&Job{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobCondition, InType: reflect.TypeOf(&JobCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobList, InType: reflect.TypeOf(&JobList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobSpec, InType: reflect.TypeOf(&JobSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobStatus, InType: reflect.TypeOf(&JobStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobTemplate, InType: reflect.TypeOf(&JobTemplate{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_batch_JobTemplateSpec, InType: reflect.TypeOf(&JobTemplateSpec{})},
	)
}

func DeepCopy_batch_CronJob(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJob)
		out := out.(*CronJob)
		out.TypeMeta = in.TypeMeta
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_batch_CronJobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_batch_CronJobStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_CronJobList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJobList)
		out := out.(*CronJobList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]CronJob, len(*in))
			for i := range *in {
				if err := DeepCopy_batch_CronJob(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_batch_CronJobSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJobSpec)
		out := out.(*CronJobSpec)
		out.Schedule = in.Schedule
		if in.StartingDeadlineSeconds != nil {
			in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.StartingDeadlineSeconds = nil
		}
		out.ConcurrencyPolicy = in.ConcurrencyPolicy
		if in.Suspend != nil {
			in, out := &in.Suspend, &out.Suspend
			*out = new(bool)
			**out = **in
		} else {
			out.Suspend = nil
		}
		if err := DeepCopy_batch_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_CronJobStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CronJobStatus)
		out := out.(*CronJobStatus)
		if in.Active != nil {
			in, out := &in.Active, &out.Active
			*out = make([]api.ObjectReference, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Active = nil
		}
		if in.LastScheduleTime != nil {
			in, out := &in.LastScheduleTime, &out.LastScheduleTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		} else {
			out.LastScheduleTime = nil
		}
		return nil
	}
}

func DeepCopy_batch_Job(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Job)
		out := out.(*Job)
		out.TypeMeta = in.TypeMeta
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_batch_JobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_batch_JobStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_JobCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobCondition)
		out := out.(*JobCondition)
		out.Type = in.Type
		out.Status = in.Status
		out.LastProbeTime = in.LastProbeTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		out.Reason = in.Reason
		out.Message = in.Message
		return nil
	}
}

func DeepCopy_batch_JobList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobList)
		out := out.(*JobList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Job, len(*in))
			for i := range *in {
				if err := DeepCopy_batch_Job(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_batch_JobSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobSpec)
		out := out.(*JobSpec)
		if in.Parallelism != nil {
			in, out := &in.Parallelism, &out.Parallelism
			*out = new(int32)
			**out = **in
		} else {
			out.Parallelism = nil
		}
		if in.Completions != nil {
			in, out := &in.Completions, &out.Completions
			*out = new(int32)
			**out = **in
		} else {
			out.Completions = nil
		}
		if in.ActiveDeadlineSeconds != nil {
			in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.ActiveDeadlineSeconds = nil
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = new(v1.LabelSelector)
			if err := v1.DeepCopy_v1_LabelSelector(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Selector = nil
		}
		if in.ManualSelector != nil {
			in, out := &in.ManualSelector, &out.ManualSelector
			*out = new(bool)
			**out = **in
		} else {
			out.ManualSelector = nil
		}
		if err := api.DeepCopy_api_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_JobStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobStatus)
		out := out.(*JobStatus)
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]JobCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_batch_JobCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		if in.StartTime != nil {
			in, out := &in.StartTime, &out.StartTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		} else {
			out.StartTime = nil
		}
		if in.CompletionTime != nil {
			in, out := &in.CompletionTime, &out.CompletionTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		} else {
			out.CompletionTime = nil
		}
		out.Active = in.Active
		out.Succeeded = in.Succeeded
		out.Failed = in.Failed
		return nil
	}
}

func DeepCopy_batch_JobTemplate(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobTemplate)
		out := out.(*JobTemplate)
		out.TypeMeta = in.TypeMeta
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_batch_JobTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_batch_JobTemplateSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JobTemplateSpec)
		out := out.(*JobTemplateSpec)
		if err := api.DeepCopy_api_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_batch_JobSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}
