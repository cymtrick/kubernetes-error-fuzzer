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

package gce

import compute "google.golang.org/api/compute/v1"

// Legacy HTTP Health Checks

// GetHttpHealthCheck returns the given HttpHealthCheck by name.
func (gce *GCECloud) GetHttpHealthCheck(name string) (*compute.HttpHealthCheck, error) {
	return gce.service.HttpHealthChecks.Get(gce.projectID, name).Do()
}

// UpdateHttpHealthCheck applies the given HttpHealthCheck as an update.
func (gce *GCECloud) UpdateHttpHealthCheck(hc *compute.HttpHealthCheck) error {
	op, err := gce.service.HttpHealthChecks.Update(gce.projectID, hc.Name, hc).Do()
	if err != nil {
		return err
	}
	return gce.waitForGlobalOp(op)
}

// DeleteHttpHealthCheck deletes the given HttpHealthCheck by name.
func (gce *GCECloud) DeleteHttpHealthCheck(name string) error {
	op, err := gce.service.HttpHealthChecks.Delete(gce.projectID, name).Do()
	if err != nil {
		return err
	}
	return gce.waitForGlobalOp(op)
}

// CreateHttpHealthCheck creates the given HttpHealthCheck.
func (gce *GCECloud) CreateHttpHealthCheck(hc *compute.HttpHealthCheck) error {
	op, err := gce.service.HttpHealthChecks.Insert(gce.projectID, hc).Do()
	if err != nil {
		return err
	}
	return gce.waitForGlobalOp(op)
}

// ListHttpHealthChecks lists all HttpHealthChecks in the project.
func (gce *GCECloud) ListHttpHealthChecks() (*compute.HttpHealthCheckList, error) {
	// TODO: use PageToken to list all not just the first 500
	return gce.service.HttpHealthChecks.List(gce.projectID).Do()
}

// Legacy HTTPS Health Checks

// GetHttpsHealthCheck returns the given HttpsHealthCheck by name.
func (gce *GCECloud) GetHttpsHealthCheck(name string) (*compute.HttpsHealthCheck, error) {
	return gce.service.HttpsHealthChecks.Get(gce.projectID, name).Do()
}

// UpdateHttpsHealthCheck applies the given HttpsHealthCheck as an update.
func (gce *GCECloud) UpdateHttpsHealthCheck(hc *compute.HttpsHealthCheck) error {
	op, err := gce.service.HttpsHealthChecks.Update(gce.projectID, hc.Name, hc).Do()
	if err != nil {
		return err
	}
	return gce.waitForGlobalOp(op)
}

// DeleteHttpsHealthCheck deletes the given HttpsHealthCheck by name.
func (gce *GCECloud) DeleteHttpsHealthCheck(name string) error {
	op, err := gce.service.HttpsHealthChecks.Delete(gce.projectID, name).Do()
	if err != nil {
		return err
	}
	return gce.waitForGlobalOp(op)
}

// CreateHttpsHealthCheck creates the given HttpsHealthCheck.
func (gce *GCECloud) CreateHttpsHealthCheck(hc *compute.HttpsHealthCheck) error {
	op, err := gce.service.HttpsHealthChecks.Insert(gce.projectID, hc).Do()
	if err != nil {
		return err
	}
	return gce.waitForGlobalOp(op)
}

// ListHttpsHealthChecks lists all HttpsHealthChecks in the project.
func (gce *GCECloud) ListHttpsHealthChecks() (*compute.HttpsHealthCheckList, error) {
	// TODO: use PageToken to list all not just the first 500
	return gce.service.HttpsHealthChecks.List(gce.projectID).Do()
}

// Generic HealthCheck

// GetHealthCheck returns the given HealthCheck by name.
func (gce *GCECloud) GetHealthCheck(name string) (*compute.HealthCheck, error) {
	return gce.service.HealthChecks.Get(gce.projectID, name).Do()
}

// UpdateHealthCheck applies the given HealthCheck as an update.
func (gce *GCECloud) UpdateHealthCheck(hc *compute.HealthCheck) error {
	op, err := gce.service.HealthChecks.Update(gce.projectID, hc.Name, hc).Do()
	if err != nil {
		return err
	}
	return gce.waitForGlobalOp(op)
}

// DeleteHealthCheck deletes the given HealthCheck by name.
func (gce *GCECloud) DeleteHealthCheck(name string) error {
	op, err := gce.service.HealthChecks.Delete(gce.projectID, name).Do()
	if err != nil {
		return err
	}
	return gce.waitForGlobalOp(op)
}

// CreateHealthCheck creates the given HealthCheck.
func (gce *GCECloud) CreateHealthCheck(hc *compute.HealthCheck) error {
	op, err := gce.service.HealthChecks.Insert(gce.projectID, hc).Do()
	if err != nil {
		return err
	}
	return gce.waitForGlobalOp(op)
}

// ListHealthChecks lists all HealthCheck in the project.
func (gce *GCECloud) ListHealthChecks() (*compute.HealthCheckList, error) {
	// TODO: use PageToken to list all not just the first 500
	return gce.service.HealthChecks.List(gce.projectID).Do()
}
