// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	conversion "k8s.io/kubernetes/pkg/conversion"
	time "time"
)

func DeepCopy_v1_APIGroup(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIGroup)
		out := out.(*APIGroup)
		out.TypeMeta = in.TypeMeta
		out.Name = in.Name
		if in.Versions != nil {
			in, out := &in.Versions, &out.Versions
			*out = make([]GroupVersionForDiscovery, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Versions = nil
		}
		out.PreferredVersion = in.PreferredVersion
		if in.ServerAddressByClientCIDRs != nil {
			in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
			*out = make([]ServerAddressByClientCIDR, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.ServerAddressByClientCIDRs = nil
		}
		return nil
	}
}

func DeepCopy_v1_APIGroupList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIGroupList)
		out := out.(*APIGroupList)
		out.TypeMeta = in.TypeMeta
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]APIGroup, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_APIGroup(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Groups = nil
		}
		return nil
	}
}

func DeepCopy_v1_APIResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIResource)
		out := out.(*APIResource)
		out.Name = in.Name
		out.Namespaced = in.Namespaced
		out.Kind = in.Kind
		if in.Verbs != nil {
			in, out := &in.Verbs, &out.Verbs
			*out = make(Verbs, len(*in))
			copy(*out, *in)
		} else {
			out.Verbs = nil
		}
		return nil
	}
}

func DeepCopy_v1_APIResourceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIResourceList)
		out := out.(*APIResourceList)
		out.TypeMeta = in.TypeMeta
		out.GroupVersion = in.GroupVersion
		if in.APIResources != nil {
			in, out := &in.APIResources, &out.APIResources
			*out = make([]APIResource, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_APIResource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.APIResources = nil
		}
		return nil
	}
}

func DeepCopy_v1_APIVersions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*APIVersions)
		out := out.(*APIVersions)
		out.TypeMeta = in.TypeMeta
		if in.Versions != nil {
			in, out := &in.Versions, &out.Versions
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Versions = nil
		}
		if in.ServerAddressByClientCIDRs != nil {
			in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
			*out = make([]ServerAddressByClientCIDR, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.ServerAddressByClientCIDRs = nil
		}
		return nil
	}
}

func DeepCopy_v1_Duration(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Duration)
		out := out.(*Duration)
		out.Duration = in.Duration
		return nil
	}
}

func DeepCopy_v1_ExportOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ExportOptions)
		out := out.(*ExportOptions)
		out.TypeMeta = in.TypeMeta
		out.Export = in.Export
		out.Exact = in.Exact
		return nil
	}
}

func DeepCopy_v1_GetOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GetOptions)
		out := out.(*GetOptions)
		out.TypeMeta = in.TypeMeta
		out.ResourceVersion = in.ResourceVersion
		return nil
	}
}

func DeepCopy_v1_GroupKind(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupKind)
		out := out.(*GroupKind)
		out.Group = in.Group
		out.Kind = in.Kind
		return nil
	}
}

func DeepCopy_v1_GroupResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupResource)
		out := out.(*GroupResource)
		out.Group = in.Group
		out.Resource = in.Resource
		return nil
	}
}

func DeepCopy_v1_GroupVersion(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupVersion)
		out := out.(*GroupVersion)
		out.Group = in.Group
		out.Version = in.Version
		return nil
	}
}

func DeepCopy_v1_GroupVersionForDiscovery(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupVersionForDiscovery)
		out := out.(*GroupVersionForDiscovery)
		out.GroupVersion = in.GroupVersion
		out.Version = in.Version
		return nil
	}
}

func DeepCopy_v1_GroupVersionKind(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupVersionKind)
		out := out.(*GroupVersionKind)
		out.Group = in.Group
		out.Version = in.Version
		out.Kind = in.Kind
		return nil
	}
}

func DeepCopy_v1_GroupVersionResource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupVersionResource)
		out := out.(*GroupVersionResource)
		out.Group = in.Group
		out.Version = in.Version
		out.Resource = in.Resource
		return nil
	}
}

func DeepCopy_v1_LabelSelector(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LabelSelector)
		out := out.(*LabelSelector)
		if in.MatchLabels != nil {
			in, out := &in.MatchLabels, &out.MatchLabels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.MatchLabels = nil
		}
		if in.MatchExpressions != nil {
			in, out := &in.MatchExpressions, &out.MatchExpressions
			*out = make([]LabelSelectorRequirement, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_LabelSelectorRequirement(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.MatchExpressions = nil
		}
		return nil
	}
}

func DeepCopy_v1_LabelSelectorRequirement(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LabelSelectorRequirement)
		out := out.(*LabelSelectorRequirement)
		out.Key = in.Key
		out.Operator = in.Operator
		if in.Values != nil {
			in, out := &in.Values, &out.Values
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Values = nil
		}
		return nil
	}
}

func DeepCopy_v1_ListMeta(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ListMeta)
		out := out.(*ListMeta)
		out.SelfLink = in.SelfLink
		out.ResourceVersion = in.ResourceVersion
		return nil
	}
}

func DeepCopy_v1_Patch(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Patch)
		out := out.(*Patch)
		_ = in
		_ = out
		return nil
	}
}

func DeepCopy_v1_RootPaths(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RootPaths)
		out := out.(*RootPaths)
		if in.Paths != nil {
			in, out := &in.Paths, &out.Paths
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Paths = nil
		}
		return nil
	}
}

func DeepCopy_v1_ServerAddressByClientCIDR(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServerAddressByClientCIDR)
		out := out.(*ServerAddressByClientCIDR)
		out.ClientCIDR = in.ClientCIDR
		out.ServerAddress = in.ServerAddress
		return nil
	}
}

func DeepCopy_v1_Status(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Status)
		out := out.(*Status)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		out.Status = in.Status
		out.Message = in.Message
		out.Reason = in.Reason
		if in.Details != nil {
			in, out := &in.Details, &out.Details
			*out = new(StatusDetails)
			if err := DeepCopy_v1_StatusDetails(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Details = nil
		}
		out.Code = in.Code
		return nil
	}
}

func DeepCopy_v1_StatusCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatusCause)
		out := out.(*StatusCause)
		out.Type = in.Type
		out.Message = in.Message
		out.Field = in.Field
		return nil
	}
}

func DeepCopy_v1_StatusDetails(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatusDetails)
		out := out.(*StatusDetails)
		out.Name = in.Name
		out.Group = in.Group
		out.Kind = in.Kind
		if in.Causes != nil {
			in, out := &in.Causes, &out.Causes
			*out = make([]StatusCause, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Causes = nil
		}
		out.RetryAfterSeconds = in.RetryAfterSeconds
		return nil
	}
}

func DeepCopy_v1_Time(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Time)
		out := out.(*Time)
		if newVal, err := c.DeepCopy(&in.Time); err != nil {
			return err
		} else {
			out.Time = *newVal.(*time.Time)
		}
		return nil
	}
}

func DeepCopy_v1_Timestamp(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Timestamp)
		out := out.(*Timestamp)
		out.Seconds = in.Seconds
		out.Nanos = in.Nanos
		return nil
	}
}

func DeepCopy_v1_TypeMeta(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TypeMeta)
		out := out.(*TypeMeta)
		out.Kind = in.Kind
		out.APIVersion = in.APIVersion
		return nil
	}
}
