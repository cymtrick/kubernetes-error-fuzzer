// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// generated by build/assets.sh; DO NOT EDIT

// Code generated by go-bindata.
// sources:
// pages/assets/html/containers.html
// DO NOT EDIT!

package pages

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pagesAssetsHtmlContainersHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\xcd\x72\xdb\x38\x12\x3e\x4b\x4f\xd1\xc3\xda\xc3\x6c\x55\x48\xc5\x49\x2e\x9b\x95\x55\xa5\x51\x92\x1d\xed\x38\xb6\xcb\xb2\x67\x6a\x8e\x10\xd9\x22\x11\x83\x00\x07\x00\x25\x6b\x5d\x7e\xf7\x2d\x00\xa4\xc4\x5f\xc9\x7f\x95\x8c\x2e\x96\x08\x74\xf7\xd7\x5f\x77\x03\x0d\xc2\xe3\x9f\x7c\x7f\x08\x30\x13\xd9\x56\xd2\x38\xd1\xf0\xee\xed\xc9\x07\xf8\x8f\x10\x31\x43\x98\xf3\x30\x80\x29\x63\x70\x65\x86\x14\x5c\xa1\x42\xb9\xc6\x28\x18\x0e\x01\xce\x68\x88\x5c\x61\x04\x39\x8f\x50\x82\x4e\x10\xa6\x19\x09\x13\x2c\x47\xde\xc0\xef\x28\x15\x15\x1c\xde\x05\x6f\xe1\x67\x33\xc1\x2b\x86\xbc\x7f\xfe\x7b\x08\xb0\x15\x39\xa4\x64\x0b\x5c\x68\xc8\x15\x82\x4e\xa8\x82\x15\x65\x08\x78\x17\x62\xa6\x81\x72\x08\x45\x9a\x31\x4a\x78\x88\xb0\xa1\x3a\xb1\x66\x0a\x25\xc1\x10\xe0\xcf\x42\x85\x58\x6a\x42\x39\x10\x08\x45\xb6\x05\xb1\xaa\xce\x03\xa2\x0d\x5e\xf3\x49\xb4\xce\x3e\x8e\x46\x9b\xcd\x26\x20\x16\x6b\x20\x64\x3c\x62\x6e\x9e\x1a\x9d\xcd\x67\x9f\xcf\x17\x9f\xfd\x77\xc1\x5b\x23\x71\xc3\x19\x2a\x05\x12\xff\xca\xa9\xc4\x08\x96\x5b\x20\x59\xc6\x68\x48\x96\x0c\x81\x91\x0d\x08\x09\x24\x96\x88\x11\x68\x61\xd0\x6e\x24\xd5\x94\xc7\x6f\x40\x89\x95\xde\x10\x89\x43\x80\x88\x2a\x2d\xe9\x32\xd7\x35\xaa\x4a\x6c\x54\xd5\x26\x08\x0e\x84\x83\x37\x5d\xc0\x7c\xe1\xc1\x2f\xd3\xc5\x7c\xf1\x66\x08\xf0\xc7\xfc\xfa\xd7\x8b\x9b\x6b\xf8\x63\x7a\x75\x35\x3d\xbf\x9e\x7f\x5e\xc0\xc5\x15\xcc\x2e\xce\x3f\xcd\xaf\xe7\x17\xe7\x0b\xb8\xf8\x02\xd3\xf3\x3f\xe1\xb7\xf9\xf9\xa7\x37\x80\x54\x27\x28\x01\xef\x32\x69\xf0\x0b\x09\xd4\x90\x68\xe2\x06\xb0\x40\xac\x01\x58\x09\x07\x48\x65\x18\xd2\x15\x0d\x81\x11\x1e\xe7\x24\x46\x88\xc5\x1a\x25\xa7\x3c\x86\x0c\x65\x4a\x95\x09\xa5\x02\xc2\xa3\x21\x00\xa3\x29\xd5\x44\xdb\x27\x2d\xa7\x82\xa1\xef\x4f\x86\xc3\x71\xa2\x53\x36\x19\x02\x8c\x13\x24\xd1\xc4\x86\x60\xac\xa9\x66\x38\x09\xa7\xd1\x9a\x2a\x21\xc1\x87\xfb\xfb\xe0\x13\x55\x19\x23\xdb\x73\x92\xe2\xc3\xc3\x78\xe4\xa6\xb8\xe9\x3f\xf9\x3e\x9c\x11\x8d\x4a\xdb\x4c\xa0\x0c\x23\x83\x00\x52\xca\xe9\x8a\x62\x04\xb3\xc5\x02\x8c\x35\x3b\x9b\x51\x7e\x0b\x12\xd9\xa9\xa7\xf4\x96\xa1\x4a\x10\xb5\x07\x89\xc4\xd5\xa9\x77\x7f\x1f\x5c\x09\xa1\x1f\x1e\x94\xc1\x1d\x8e\x96\x42\x68\xa5\x25\xc9\xfc\xf7\xc1\x49\x70\x12\xa4\x94\x07\xa1\x52\xde\x64\xb8\xb7\x7c\x91\x19\x0f\x09\x33\xce\xa5\xf8\x52\x3b\x56\x49\x8f\xb5\xa7\x68\x0c\x05\x37\xc9\x8e\x52\xb5\x00\x1f\xa4\xea\xbf\x64\x4d\x16\xa1\xa4\x99\xde\x7b\xa2\xdc\x6f\x25\xc3\xb6\x9d\x6f\x7f\xe5\x28\xb7\xfe\x49\x70\xf2\x36\x78\x67\x11\x7f\x53\xde\x64\x3c\x72\x32\x8f\x50\xd0\x45\x71\xbf\x0a\xbd\xcd\xf0\xd4\xd3\x78\xa7\x47\xdf\xc8\x9a\xb8\xa7\x5e\xb7\xe6\xd8\xae\x4f\xfe\x37\x45\x32\xda\x50\xf9\x6c\x9d\x15\x5a\x5f\x09\x64\x98\x10\xa9\xdb\xda\xc6\xa3\xb2\x1e\xc6\x4b\x11\x6d\x0b\x03\x11\x5d\x43\xc8\x88\x52\xa7\xde\x0e\x89\xcb\x3b\x5f\x25\x62\x13\x12\x85\x1e\x4c\x8a\x75\x6c\x4c\x9a\xb9\xe1\xed\x85\x99\xaf\x52\xff\xe4\x9d\x07\x34\x3a\xf5\x98\x88\x85\xb7\x13\x1b\x91\xdd\xd7\x9a\xbd\x52\x64\x32\x1c\x54\x07\x32\x12\xa3\x6f\xc0\xa2\x34\x43\xa6\x92\x4f\x26\xed\x82\x4d\x4e\x8c\xdc\x28\xa2\x6b\xf3\x57\xb0\x52\x7c\x29\x91\x44\xa1\xcc\xd3\xa5\x93\xbe\xbf\x97\x84\xc7\x08\xff\xc8\x88\x44\xae\x67\x3b\x37\x3f\x9e\x42\x70\x59\x7f\xa6\x1e\x1e\xac\x41\x46\x27\x15\x67\x9b\x92\xc1\x19\xe5\xb7\x0f\x0f\xde\xa4\x63\xe8\x1a\xef\xb4\x41\x47\x26\xe3\x11\xa3\x05\x00\xe4\x91\x51\x3c\x1e\x09\xb6\x27\xc5\x02\x77\x3f\xee\xef\xe9\x0a\x82\xb9\x72\xa4\x1e\xe1\x0a\x8a\xcf\x38\xf9\xb0\x07\x19\x04\xa3\x48\x84\xb7\x86\xb1\x4f\xf6\x2f\xec\x7d\x72\x60\x92\x0f\x3d\xa6\x1d\xb8\x2a\x90\x45\xbe\x0c\xab\x8c\xbc\x2c\x76\xef\x27\x35\x7d\xe3\x51\xf2\xbe\x1a\xb8\x8a\x30\xa3\x4a\xfb\xb1\x14\x79\xd6\x88\x9c\xaa\x28\xb0\x61\x6b\x22\x1c\xd4\x92\xb3\x36\xbf\x0c\x56\xdb\x88\x4f\x35\xa6\x36\x88\xb5\xf9\xfb\x08\x36\x82\x57\x61\xad\x9f\x42\xc7\xa0\x8b\xc1\x42\x13\x9d\xbf\x06\x81\x9f\x24\x5d\xa3\x04\xa7\xaf\x49\x60\xce\x8e\xf2\xe7\x52\x43\x59\x71\xcb\x5f\x03\x9f\x4b\x79\xa7\x06\x3a\x28\x1a\xab\x8c\xf0\xd2\x8a\x51\xe3\x33\xb2\x44\x66\xb9\xab\xea\x0e\x7e\xc3\xad\xa1\xce\x4c\x9f\x40\x73\xf0\x77\xc2\x72\x5b\xb9\xcd\xba\xa8\xb3\xe6\x9c\xdd\x63\x1b\x3c\x0f\xda\x42\x0b\x49\x62\x1c\x2f\xe5\xa4\x00\x64\x54\xf5\x91\x35\xd8\x73\x65\xcd\xb7\xb8\xea\x47\xf5\x54\xbe\x2a\xfa\xdb\x7c\x55\x07\xeb\x7c\x0d\x76\x74\x0d\xc6\xa3\x9c\x59\x6f\x4a\x26\x8b\x07\x7d\xd9\xda\x55\xe3\xce\xab\x79\x4a\x62\x3c\x9e\xa1\xb0\xfb\xf4\xa7\x2a\x54\x3e\x26\x67\x9d\x6a\x97\xac\x95\x91\x2a\x2e\xa7\xcd\xec\x17\x2e\x4f\x7c\x6a\x65\xcc\xbe\x55\x9b\x65\x42\xb8\x94\xfb\xdf\xc7\x7c\xbb\x42\x25\x72\x19\xa2\x9a\xae\x09\x65\xa6\x6d\x7e\x85\x1a\x9c\x2b\xc1\x6c\xeb\xd9\xa8\x3f\x67\x72\x96\xe5\x55\x63\xbd\x89\x56\x61\xa2\x37\x7f\x80\x84\x9a\xae\x4d\x93\x5e\x58\xf4\x6d\x6f\x0a\x19\xe1\xc8\xdc\x77\x6f\x32\xbb\xbc\x71\xe1\xdf\x6b\x2c\x16\xef\x0c\x43\x03\x27\x38\x33\xcd\xf2\xce\xf1\xc3\x26\x0f\xd5\x51\x42\xa4\x89\x63\x99\xa3\x99\xa4\x5c\xbb\x87\x6d\x63\x50\x53\x93\x73\xba\x53\xa3\xaa\x6a\xda\xc8\xab\x41\xec\xf0\xe5\x2b\xb9\x7b\x25\x77\xbe\x92\x3b\xb0\xaa\x1a\x1e\xcd\x44\xdd\xa1\xbd\xc5\x7e\x9f\x42\xf1\x22\x97\xd4\xed\xcb\xdd\x99\x32\x26\x36\xe6\x40\x22\xda\x41\x32\x16\x1a\x06\x21\xf8\x4a\xc2\x84\x72\x9c\xf3\x95\x08\xce\xf3\xd4\xca\x95\x6b\x4c\x1b\x7d\xb9\xd4\xec\x7e\x3b\x27\xbe\x62\x2a\xe4\xf6\xfb\x26\xbc\xb3\x79\x20\xe7\xdd\x84\xc0\xbd\x2d\xb0\x6a\x5e\x4e\x6f\x45\x59\xb3\x02\xe8\xff\xf0\x80\xe1\xfe\xa4\x29\xe4\x6f\x38\xd5\x07\xe4\x9f\x93\x55\x85\x9e\x57\x2a\x94\xae\x22\x69\x3b\x7d\xb4\x46\x7a\xdd\x2d\x24\x5f\xe0\xe8\x62\x43\xb2\xd7\x5a\xe4\x36\x24\xeb\x5c\x16\xda\x1e\x57\xac\x3e\xc3\xeb\x8a\xf4\x11\xcf\x9b\xa5\x57\x78\x57\xeb\x42\x9f\xbd\x99\xdd\x28\xd3\x1a\xf5\x77\xe2\xb6\xf2\x8a\xfa\xcb\x24\x4d\x89\xdc\x1e\x68\x03\xcc\x2c\x63\x81\xf2\xb8\xdd\x08\xd4\xa7\x15\xc5\x7c\xb1\x46\xb9\xa6\xb8\x39\xdc\x1e\x54\x3b\x84\xdc\x20\xf6\x63\x92\xc7\xe8\xd5\x55\x9a\xd3\xec\xae\x65\xf8\x21\xde\x5c\x4a\x11\xa2\x52\xc7\xba\x9d\xaa\x3b\x59\x29\xe2\x6b\x91\x3d\xca\xa1\x9e\x3e\xe3\x3b\xba\x69\x5b\x8e\xc7\x38\xd8\xe1\x4d\xc3\xc0\x87\xc9\xb5\xd0\x84\x41\x99\x87\x1f\x6c\x66\x56\xf8\x09\xb3\xdc\xd7\x66\x8a\xef\x02\x6f\x5f\x6a\xec\x49\x81\xf2\xd5\x93\x51\x35\xbb\xbc\x81\x33\x41\x22\x98\xae\x51\x1e\xd0\xc7\x04\x89\xea\x8a\x76\x6f\xa4\xaa\xc8\x2c\x26\xc8\xec\x11\x5a\xf6\x2a\xcb\x50\xfa\x66\xff\xef\xc4\xd7\xad\xf2\x17\x89\xe4\x36\x12\x1b\xde\xa7\xd3\xa9\x5a\x96\xd3\x7a\x95\xb6\x53\xe3\xe8\xee\xfc\x1d\xd3\xa4\xdc\xa8\xbf\x53\xa6\xa4\xd6\xdc\xf1\x30\x2c\xe5\xa8\xf1\xa4\x02\x40\x8a\x0d\x74\x1f\x78\x0e\x86\xb0\x31\xad\xbd\x1c\xff\xcb\x9e\x2d\x6b\xae\x4a\x11\x4b\xb4\x2f\x50\xa1\xf5\xe9\x9a\xe8\x2f\x89\x84\xea\x0f\x3f\x32\x07\x55\xe9\x95\xeb\x88\x1b\x48\x84\xf6\x1d\x15\x9d\x9a\xa1\xbe\x57\x29\xe9\x0b\xce\xb6\xde\xe4\x57\xa1\xa1\x0c\x98\x3b\x24\x77\x48\xb6\xd9\x7c\x0a\x5c\xca\x57\xa2\x01\x36\x14\x2c\x7a\x0e\xda\x99\x60\xd1\x63\xe1\x0e\x06\x9d\xb8\xbb\x1f\xb6\x23\xf7\xde\xab\x66\x97\xc6\xbb\xe6\xea\xf3\xc4\xa2\x3c\x47\xbd\x11\xf2\xf6\x89\x55\x39\x78\x79\x39\x16\x86\x8b\xcd\xfe\x29\x85\x38\x68\x8e\x46\x52\x64\x26\xf9\xdb\x05\xb2\xcc\xb5\x16\xbb\x78\x2d\x35\x87\xa5\xe6\x7e\x84\x2b\x92\x33\x0d\xa5\x9c\xaf\x45\x1c\x33\xf4\x8a\xf7\xd9\x4e\xc8\xf1\xcc\x1d\x4a\x5f\x21\xc3\xd0\x1e\x01\x76\xc6\x20\x22\x9a\x14\xa2\x15\x0c\x40\x24\x25\x7e\x42\x54\x26\xb2\x3c\x3b\xf5\xb4\xcc\xb1\x78\x88\x77\x19\xe1\x11\x46\xa7\xde\x8a\x30\x85\x1d\x29\xe6\xd2\xab\xdb\x70\x19\xeb\xee\xfc\xaa\x25\x66\x48\x24\x56\xe6\x0e\xca\x4c\x70\x9e\xb5\x58\xca\x59\xb7\x49\xaf\x49\xb0\x9f\x22\xcf\x3d\x90\xc2\x78\xec\xbe\x5b\xc7\x6c\x77\xc9\x30\x5a\x6e\x0f\x32\xd6\xce\xf9\xe2\xf5\xd0\x81\xb4\x7d\xca\x82\x9c\x48\x91\xc7\x49\x96\xeb\xf6\x2a\xb8\x5b\x96\x4b\x78\xcb\xad\x46\xd5\xde\xbe\x9f\x61\xf6\xb3\x94\xc2\xbe\x3e\x6e\x6d\x01\xa5\x2d\xb4\x33\xfa\x8d\x35\x9c\x6f\x54\xe8\x17\xf5\xc3\xb6\xcc\x2f\x94\xa1\xda\x2a\x8d\xe9\xe3\x3b\xc8\xd5\x4e\xc6\xed\x7d\x9d\x4d\x64\xbf\xa6\x9e\x65\x6a\x96\x2b\x2d\xd2\xaf\xa8\x25\x0d\x9f\xca\xc7\x91\xc5\x6a\x70\x88\x81\xa9\xbb\xe1\x36\x79\x0c\x85\xf5\xe6\x8a\x75\x28\x57\x1a\xbd\x94\x75\xc2\x4f\x9d\x9e\xa3\xf9\x30\x68\x1e\x36\x3b\x6e\x41\x7e\x58\x6a\x74\xdc\x9d\x1c\xcb\x8e\xc7\x35\x55\x19\x98\xbe\xd9\xb6\x35\x1f\x9b\xeb\x05\xe5\x59\xae\x6b\xad\x6e\xf5\x86\xc4\x8f\xdc\x45\x9c\x1f\x8a\x9c\x6b\xaf\x73\xff\xde\x6d\xdd\x5d\x72\x56\x7d\x8f\xdc\x9a\xb0\x1c\x4f\x4f\xde\x36\x20\xf7\x2f\x34\x9d\x08\x6b\xdd\x60\x43\x53\xf7\x02\xf8\x4c\x0e\x5d\x33\x72\x94\xc6\xa2\x8d\xf8\x7b\x32\x59\x6b\xb5\x9c\x15\x29\x18\xab\x98\x59\x32\x11\xde\x36\x19\x68\xef\x8f\xcd\x9e\xfc\x15\xc3\xd2\xb3\x74\x77\x0c\x56\x87\x2a\x03\x87\xaf\xd2\x4b\x61\xa5\x89\xd4\x97\x24\xc6\x9f\xef\xef\x83\xdd\x0d\xaa\xbb\x71\x7e\x03\xe6\x59\xed\xfc\x6d\x1f\xb5\x8e\x5b\xf6\xa9\xbb\xca\xb5\x5f\xcb\x7b\x5d\xfb\xdf\x47\xe6\x13\x49\xb2\x71\xd7\x23\xc6\x4c\xfd\x26\xa6\x98\x54\xbf\xb9\x77\x17\xf6\xe3\x91\xfb\xd7\x96\xff\x07\x00\x00\xff\xff\x4b\x13\x4b\x6c\x3d\x25\x00\x00")

func pagesAssetsHtmlContainersHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pagesAssetsHtmlContainersHtml,
		"pages/assets/html/containers.html",
	)
}

func pagesAssetsHtmlContainersHtml() (*asset, error) {
	bytes, err := pagesAssetsHtmlContainersHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pages/assets/html/containers.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"pages/assets/html/containers.html": pagesAssetsHtmlContainersHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"pages": {nil, map[string]*bintree{
		"assets": {nil, map[string]*bintree{
			"html": {nil, map[string]*bintree{
				"containers.html": {pagesAssetsHtmlContainersHtml, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
