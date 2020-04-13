// +build !providerless

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

package azure

// create file share
func (az *Cloud) createFileShare(accountName, accountKey, name string, sizeGiB int) error {
	return az.FileClient.CreateFileShare(accountName, accountKey, name, sizeGiB)
}

func (az *Cloud) deleteFileShare(accountName, accountKey, name string) error {
	return az.FileClient.DeleteFileShare(accountName, accountKey, name)
}

func (az *Cloud) resizeFileShare(accountName, accountKey, name string, sizeGiB int) error {
	return az.FileClient.ResizeFileShare(accountName, accountKey, name, sizeGiB)
}
