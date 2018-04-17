/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/sample-apiserver/pkg/apis/wardle"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	// Add non-generated conversion functions to handle the *int32 -> int32
	// conversion. A pointer is useful in the versioned type so we can default
	// it, but a plain int32 is more convenient in the internal type. These
	// functions are the same as the autogenerated ones in every other way.
	err := scheme.AddConversionFuncs(
		Convert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec,
		Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec,
	)
	if err != nil {
		return err
	}

	return nil
}

// Convert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec is an autogenerated conversion function.
func Convert_v1alpha1_FlunderSpec_To_wardle_FlunderSpec(in *FlunderSpec, out *wardle.FlunderSpec, s conversion.Scope) error {
	if in.ReferenceType != nil {
		// assume that ReferenceType is defaulted
		switch *in.ReferenceType {
		case FlunderReferenceType:
			out.ReferenceType = wardle.FlunderReferenceType
			out.FlunderReference = in.Reference
		case FischerReferenceType:
			out.ReferenceType = wardle.FischerReferenceType
			out.FischerReference = in.Reference
		}
	}

	return nil
}

// Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec is an autogenerated conversion function.
func Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(in *wardle.FlunderSpec, out *FlunderSpec, s conversion.Scope) error {
	switch in.ReferenceType {
	case wardle.FlunderReferenceType:
		t := FlunderReferenceType
		out.ReferenceType = &t
		out.Reference = in.FlunderReference
	case wardle.FischerReferenceType:
		t := FischerReferenceType
		out.ReferenceType = &t
		out.Reference = in.FischerReference
	}

	return nil
}
