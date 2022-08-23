package csaupgrade

import (
	"bytes"
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

// Upgrades the Manager information for fields managed with CSA
// Prepares fields owned by `csaManager` for 'Update' operations for use now
// with the given `ssaManager` for `Apply` operations
//
// csaManager - Name of FieldManager formerly used for `Update` operations
// ssaManager - Name of FieldManager formerly used for `Apply` operations
// subResource - Name of subresource used for api calls or empty string for main resource
func UpgradeManagedFields(
	obj runtime.Object,
	csaManagerName string,
	ssaManagerName string,
	subResource string,
) (runtime.Object, error) {
	accessor, err := meta.Accessor(obj)
	if err != nil {
		return nil, fmt.Errorf("error accessing object metadata: %w", err)
	}

	// Create managed fields clone since we modify the values
	var managedFields []metav1.ManagedFieldsEntry
	managedFields = append(managedFields, accessor.GetManagedFields()...)

	// Locate SSA manager
	ssaManagerIndex, ssaManagerExists := findFirstIndex(managedFields,
		func(entry metav1.ManagedFieldsEntry) bool {
			return entry.Manager == ssaManagerName &&
				entry.Operation == metav1.ManagedFieldsOperationApply &&
				entry.Subresource == ""
		})

	if ssaManagerExists {
		ssaManager := managedFields[ssaManagerIndex]

		// find Update manager of same APIVersion, union ssa fields with it.
		// discard all other Update managers of the same name
		csaManagerIndex, csaManagerExists := findFirstIndex(managedFields,
			func(entry metav1.ManagedFieldsEntry) bool {
				return entry.Manager == csaManagerName &&
					entry.Operation == metav1.ManagedFieldsOperationUpdate &&
					entry.Subresource == "" &&
					entry.APIVersion == ssaManager.APIVersion
			})

		if csaManagerExists {
			csaManager := managedFields[csaManagerIndex]

			// Union the csa manager with the existing SSA manager
			ssaFieldSet, err := fieldsToSet(*ssaManager.FieldsV1)
			if err != nil {
				return nil, fmt.Errorf("failed to convert fields to set: %w", err)
			}

			csaFieldSet, err := fieldsToSet(*csaManager.FieldsV1)
			if err != nil {
				return nil, fmt.Errorf("failed to convert fields to set: %w", err)
			}

			combinedFieldSet := ssaFieldSet.Union(&csaFieldSet)
			combinedFieldSetEncoded, err := setToFields(*combinedFieldSet)
			if err != nil {
				return nil, fmt.Errorf("failed to encode field set: %w", err)
			}

			managedFields[ssaManagerIndex].FieldsV1 = &combinedFieldSetEncoded
		}
	} else {
		// SSA manager does not exist. Find the most recent matching CSA manager,
		// convert it to an SSA manager.
		//
		// (find first index, since managed fields are sorted so that most recent is
		//  first in the list)
		csaManagerIndex, csaManagerExists := findFirstIndex(managedFields, func(entry metav1.ManagedFieldsEntry) bool {
			return entry.Manager == csaManagerName && entry.Operation == metav1.ManagedFieldsOperationUpdate && entry.Subresource == ""
		})

		if !csaManagerExists {
			// There are no CSA managers that need to be converted. Nothing to do
			// Return early
			return obj, nil
		}

		// Convert the entry to apply operation
		managedFields[csaManagerIndex].Operation = metav1.ManagedFieldsOperationApply
		managedFields[csaManagerIndex].Manager = ssaManagerName
	}

	// Create version of managed fields which has no CSA managers with the given name
	filteredManagers := filter(managedFields, func(entry metav1.ManagedFieldsEntry) bool {
		return !(entry.Manager == csaManagerName &&
			entry.Operation == metav1.ManagedFieldsOperationUpdate &&
			entry.Subresource == "")
	})

	copied := obj.DeepCopyObject()
	copiedAccessor, err := meta.Accessor(copied)
	if err != nil {
		return nil, fmt.Errorf("failed to get meta accessor for copied object: %w", err)
	}
	copiedAccessor.SetManagedFields(filteredManagers)
	return copied, nil
}

func findFirstIndex[T any](
	collection []T,
	predicate func(T) bool,
) (int, bool) {
	for idx, entry := range collection {
		if predicate(entry) {
			return idx, true
		}
	}

	return -1, false
}

func filter[T any](
	collection []T,
	predicate func(T) bool,
) []T {
	result := make([]T, 0, len(collection))

	for _, value := range collection {
		if predicate(value) {
			result = append(result, value)
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

// FieldsToSet creates a set paths from an input trie of fields
func fieldsToSet(f metav1.FieldsV1) (s fieldpath.Set, err error) {
	err = s.FromJSON(bytes.NewReader(f.Raw))
	return s, err
}

// SetToFields creates a trie of fields from an input set of paths
func setToFields(s fieldpath.Set) (f metav1.FieldsV1, err error) {
	f.Raw, err = s.ToJSON()
	return f, err
}
