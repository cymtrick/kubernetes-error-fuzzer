package csaupgrade

import (
	"bytes"
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

const csaAnnotationName = "kubectl.kubernetes.io/last-applied-configuration"

var csaAnnotationFieldSet = fieldpath.NewSet(fieldpath.MakePathOrDie("metadata", "annotations", csaAnnotationName))

// Upgrades the Manager information for fields managed with client-side-apply (CSA)
// Prepares fields owned by `csaManager` for 'Update' operations for use now
// with the given `ssaManager` for `Apply` operations.
//
// This transformation should be performed on an object if it has been previously
// managed using client-side-apply to prepare it for future use with
// server-side-apply.
//
// Caveats:
//  1. This operation is not reversible. Information about which fields the client
//     owned will be lost in this operation.
//  2. Supports being performed either before or after initial server-side apply.
//  3. Client-side apply tends to own more fields (including fields that are defaulted),
//     this will possibly remove this defaults, they will be re-defaulted, that's fine.
//  4. Care must be taken to not overwrite the managed fields on the server if they
//     have changed before sending a patch.
//
// obj - Target of the operation which has been managed with CSA in the past
// csaManagerName - Name of FieldManager formerly used for `Update` operations
// ssaManagerName - Name of FieldManager formerly used for `Apply` operations
// Returns: a copy of the `obj` paramter with its managed fields modified so that
// it may be used with server-side apply in the future.
func UpgradeManagedFields(
	obj runtime.Object,
	csaManagerName string,
	ssaManagerName string,
) (runtime.Object, error) {
	accessor, err := meta.Accessor(obj)
	if err != nil {
		return nil, fmt.Errorf("error accessing object metadata: %w", err)
	}

	// Create managed fields clone since we modify the values
	var managedFields []metav1.ManagedFieldsEntry
	managedFields = append(managedFields, accessor.GetManagedFields()...)

	// Locate SSA manager
	replaceIndex, managerExists := findFirstIndex(managedFields,
		func(entry metav1.ManagedFieldsEntry) bool {
			return entry.Manager == ssaManagerName &&
				entry.Operation == metav1.ManagedFieldsOperationApply &&
				entry.Subresource == ""
		})

	if !managerExists {
		// SSA manager does not exist. Find the most recent matching CSA manager,
		// convert it to an SSA manager.
		//
		// (find first index, since managed fields are sorted so that most recent is
		//  first in the list)
		replaceIndex, managerExists = findFirstIndex(managedFields,
			func(entry metav1.ManagedFieldsEntry) bool {
				return entry.Manager == csaManagerName &&
					entry.Operation == metav1.ManagedFieldsOperationUpdate &&
					entry.Subresource == ""
			})

		if !managerExists {
			// There are no CSA managers that need to be converted. Nothing to do
			// Return early
			return obj, nil
		}

		// Convert CSA manager into SSA manager
		managedFields[replaceIndex].Operation = metav1.ManagedFieldsOperationApply
		managedFields[replaceIndex].Manager = ssaManagerName
	}
	err = unionManagerIntoIndex(managedFields, replaceIndex, csaManagerName)
	if err != nil {
		return nil, err
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

	// Wipe out CSA annotation if it exists
	annotations := copiedAccessor.GetAnnotations()
	delete(annotations, csaAnnotationName)
	copiedAccessor.SetAnnotations(annotations)

	return copied, nil
}

// Locates an Update manager entry named `csaManagerName` with the same APIVersion
// as the manager at the targetIndex. Unions both manager's fields together
// into the manager specified by `targetIndex`. No other managers are modified.
func unionManagerIntoIndex(entries []metav1.ManagedFieldsEntry, targetIndex int, csaManagerName string) error {
	ssaManager := entries[targetIndex]

	// find Update manager of same APIVersion, union ssa fields with it.
	// discard all other Update managers of the same name
	csaManagerIndex, csaManagerExists := findFirstIndex(entries,
		func(entry metav1.ManagedFieldsEntry) bool {
			return entry.Manager == csaManagerName &&
				entry.Operation == metav1.ManagedFieldsOperationUpdate &&
				entry.Subresource == "" &&
				entry.APIVersion == ssaManager.APIVersion
		})

	targetFieldSet, err := fieldsToSet(*ssaManager.FieldsV1)
	if err != nil {
		return fmt.Errorf("failed to convert fields to set: %w", err)
	}

	combinedFieldSet := &targetFieldSet

	// Union the csa manager with the existing SSA manager. Do nothing if
	// there was no good candidate found
	if csaManagerExists {
		csaManager := entries[csaManagerIndex]

		csaFieldSet, err := fieldsToSet(*csaManager.FieldsV1)
		if err != nil {
			return fmt.Errorf("failed to convert fields to set: %w", err)
		}

		combinedFieldSet = combinedFieldSet.Union(&csaFieldSet)
	}

	// Ensure that the resultant fieldset does not include the
	// last applied annotation
	combinedFieldSet = combinedFieldSet.Difference(csaAnnotationFieldSet)

	// Encode the fields back to the serialized format
	combinedFieldSetEncoded, err := setToFields(*combinedFieldSet)
	if err != nil {
		return fmt.Errorf("failed to encode field set: %w", err)
	}

	entries[targetIndex].FieldsV1 = &combinedFieldSetEncoded
	return nil
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

// Included from fieldmanager.internal to avoid dependency cycle
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
