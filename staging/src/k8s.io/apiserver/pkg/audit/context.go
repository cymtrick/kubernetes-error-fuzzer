/*
Copyright 2020 The Kubernetes Authors.

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

package audit

import (
	"context"
	"sync"

	"k8s.io/apimachinery/pkg/types"
	auditinternal "k8s.io/apiserver/pkg/apis/audit"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/klog/v2"
)

// The key type is unexported to prevent collisions
type key int

// auditKey is the context key for storing the audit context that is being
// captured and the evaluated policy that applies to the given request.
const auditKey key = iota

// AuditContext holds the information for constructing the audit events for the current request.
type AuditContext struct {
	// RequestAuditConfig is the audit configuration that applies to the request
	RequestAuditConfig RequestAuditConfig

	// Event is the audit Event object that is being captured to be written in
	// the API audit log. It is set to nil when the request is not being audited.
	Event *auditinternal.Event

	// annotations holds audit annotations that are recorded before the event has been initialized.
	// This is represented as a slice rather than a map to preserve order.
	annotations []annotation
	// annotationMutex guards annotations AND event.Annotations
	annotationMutex sync.Mutex

	// auditID is the Audit ID associated with this request.
	auditID types.UID
}

type annotation struct {
	key, value string
}

// AddAuditAnnotation sets the audit annotation for the given key, value pair.
// It is safe to call at most parts of request flow that come after WithAuditAnnotations.
// The notable exception being that this function must not be called via a
// defer statement (i.e. after ServeHTTP) in a handler that runs before WithAudit
// as at that point the audit event has already been sent to the audit sink.
// Handlers that are unaware of their position in the overall request flow should
// prefer AddAuditAnnotation over LogAnnotation to avoid dropping annotations.
func AddAuditAnnotation(ctx context.Context, key, value string) {
	ac := AuditContextFrom(ctx)
	if ac == nil {
		// auditing is not enabled
		return
	}

	ac.annotationMutex.Lock()
	defer ac.annotationMutex.Unlock()

	addAuditAnnotationLocked(ac, key, value)
}

// AddAuditAnnotations is a bulk version of AddAuditAnnotation. Refer to AddAuditAnnotation for
// restrictions on when this can be called.
// keysAndValues are the key-value pairs to add, and must have an even number of items.
func AddAuditAnnotations(ctx context.Context, keysAndValues ...string) {
	ac := AuditContextFrom(ctx)
	if ac == nil {
		// auditing is not enabled
		return
	}

	ac.annotationMutex.Lock()
	defer ac.annotationMutex.Unlock()

	if len(keysAndValues)%2 != 0 {
		klog.Errorf("Dropping mismatched audit annotation %q", keysAndValues[len(keysAndValues)-1])
	}
	for i := 0; i < len(keysAndValues); i += 2 {
		addAuditAnnotationLocked(ac, keysAndValues[i], keysAndValues[i+1])
	}
}

// AddAuditAnnotationsMap is a bulk version of AddAuditAnnotation. Refer to AddAuditAnnotation for
// restrictions on when this can be called.
func AddAuditAnnotationsMap(ctx context.Context, annotations map[string]string) {
	ac := AuditContextFrom(ctx)
	if ac == nil {
		// auditing is not enabled
		return
	}

	ac.annotationMutex.Lock()
	defer ac.annotationMutex.Unlock()

	for k, v := range annotations {
		addAuditAnnotationLocked(ac, k, v)
	}
}

// addAuditAnnotationLocked is the shared code for recording an audit annotation. This method should
// only be called while the auditAnnotationsMutex is locked.
func addAuditAnnotationLocked(ac *AuditContext, key, value string) {
	if ac.Event != nil {
		logAnnotation(ac.Event, key, value)
	} else {
		ac.annotations = append(ac.annotations, annotation{key: key, value: value})
	}
}

// This is private to prevent reads/write to the slice from outside of this package.
// The audit event should be directly read to get access to the annotations.
func addAuditAnnotationsFrom(ctx context.Context, ev *auditinternal.Event) {
	ac := AuditContextFrom(ctx)
	if ac == nil {
		// auditing is not enabled
		return
	}

	ac.annotationMutex.Lock()
	defer ac.annotationMutex.Unlock()

	for _, kv := range ac.annotations {
		logAnnotation(ev, kv.key, kv.value)
	}
}

// LogAnnotation fills in the Annotations according to the key value pair.
func logAnnotation(ae *auditinternal.Event, key, value string) {
	if ae == nil || ae.Level.Less(auditinternal.LevelMetadata) {
		return
	}
	if ae.Annotations == nil {
		ae.Annotations = make(map[string]string)
	}
	if v, ok := ae.Annotations[key]; ok && v != value {
		klog.Warningf("Failed to set annotations[%q] to %q for audit:%q, it has already been set to %q", key, value, ae.AuditID, ae.Annotations[key])
		return
	}
	ae.Annotations[key] = value
}

// WithAuditContext returns a new context that stores the AuditContext.
func WithAuditContext(parent context.Context) context.Context {
	if AuditContextFrom(parent) != nil {
		return parent // Avoid double registering.
	}

	return genericapirequest.WithValue(parent, auditKey, &AuditContext{})
}

// AuditEventFrom returns the audit event struct on the ctx
func AuditEventFrom(ctx context.Context) *auditinternal.Event {
	if o := AuditContextFrom(ctx); o != nil {
		return o.Event
	}
	return nil
}

// AuditContextFrom returns the pair of the audit configuration object
// that applies to the given request and the audit event that is going to
// be written to the API audit log.
func AuditContextFrom(ctx context.Context) *AuditContext {
	ev, _ := ctx.Value(auditKey).(*AuditContext)
	return ev
}

// WithAuditID sets the AuditID on the AuditContext. The AuditContext must already be present in the
// request context. If the specified auditID is empty, no value is set.
func WithAuditID(ctx context.Context, auditID types.UID) {
	if auditID == "" {
		return
	}
	ac := AuditContextFrom(ctx)
	if ac == nil {
		return
	}
	ac.auditID = auditID
	if ac.Event != nil {
		ac.Event.AuditID = auditID
	}
}

// AuditIDFrom returns the value of the audit ID from the request context.
func AuditIDFrom(ctx context.Context) (types.UID, bool) {
	if ac := AuditContextFrom(ctx); ac != nil {
		return ac.auditID, ac.auditID != ""
	}
	return "", false
}

// GetAuditIDTruncated returns the audit ID (truncated) from the request context.
// If the length of the Audit-ID value exceeds the limit, we truncate it to keep
// the first N (maxAuditIDLength) characters.
// This is intended to be used in logging only.
func GetAuditIDTruncated(ctx context.Context) string {
	auditID, ok := AuditIDFrom(ctx)
	if !ok {
		return ""
	}

	// if the user has specified a very long audit ID then we will use the first N characters
	// Note: assuming Audit-ID header is in ASCII
	const maxAuditIDLength = 64
	if len(auditID) > maxAuditIDLength {
		auditID = auditID[:maxAuditIDLength]
	}

	return string(auditID)
}
