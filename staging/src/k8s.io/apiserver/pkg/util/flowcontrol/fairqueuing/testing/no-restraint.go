/*
Copyright 2019 The Kubernetes Authors.

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

package testing

import (
	"context"

	fq "k8s.io/apiserver/pkg/util/flowcontrol/fairqueuing"
)

// NewNoRestraintFactory makes a QueueSetFactory that produces
// QueueSets that exert no restraint --- every request is dispatched
// for execution as soon as it arrives.
func NewNoRestraintFactory() fq.QueueSetFactory {
	return noRestraintFactory{}
}

type noRestraintFactory struct{}

type noRestraintCompeter struct{}

type noRestraint struct{}

func (noRestraintFactory) QualifyQueuingConfig(qCfg fq.QueuingConfig) (fq.QueueSetCompleter, error) {
	return noRestraintCompeter{}, nil
}

func (noRestraintCompeter) GetQueueSet(dCfg fq.DispatchingConfig) fq.QueueSet {
	return noRestraint{}
}

func (noRestraint) QualifyQueuingConfig(qCfg fq.QueuingConfig) (fq.QueueSetCompleter, error) {
	return noRestraintCompeter{}, nil
}

func (noRestraint) Quiesce(fq.EmptyHandler) {
}

func (noRestraint) Wait(ctx context.Context, hashValue uint64, descr1, descr2 interface{}) (quiescent, execute bool, afterExecution func()) {
	return false, true, func() {}
}
