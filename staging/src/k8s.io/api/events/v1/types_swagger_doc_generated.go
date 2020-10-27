/*
Copyright The Kubernetes Authors.

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

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_Event = map[string]string{
	"":                         "Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.",
	"eventTime":                "eventTime is the time when this Event was first observed. It is required.",
	"series":                   "series is data about the Event series this event represents or nil if it's a singleton Event.",
	"reportingController":      "reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.",
	"reportingInstance":        "reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.",
	"action":                   "action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.",
	"reason":                   "reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.",
	"regarding":                "regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.",
	"related":                  "related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.",
	"note":                     "note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.",
	"type":                     "type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.",
	"deprecatedSource":         "deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.",
	"deprecatedFirstTimestamp": "deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.",
	"deprecatedLastTimestamp":  "deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.",
	"deprecatedCount":          "deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.",
}

func (Event) SwaggerDoc() map[string]string {
	return map_Event
}

var map_EventList = map[string]string{
	"":         "EventList is a list of Event objects.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is a list of schema objects.",
}

func (EventList) SwaggerDoc() map[string]string {
	return map_EventList
}

var map_EventSeries = map[string]string{
	"":                 "EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in \"k8s.io/client-go/tools/events/event_broadcaster.go\" shows how this struct is updated on heartbeats and can guide customized reporter implementations.",
	"count":            "count is the number of occurrences in this series up to the last heartbeat time.",
	"lastObservedTime": "lastObservedTime is the time when last Event from the series was seen before last heartbeat.",
}

func (EventSeries) SwaggerDoc() map[string]string {
	return map_EventSeries
}

// AUTO-GENERATED FUNCTIONS END HERE
