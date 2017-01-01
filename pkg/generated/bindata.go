// Code generated by go-bindata.
// sources:
// ../../../translations/kubectl/default/LC_MESSAGES/k8s.mo
// ../../../translations/kubectl/default/LC_MESSAGES/k8s.po
// ../../../translations/kubectl/en_US/LC_MESSAGES/k8s.mo
// ../../../translations/kubectl/en_US/LC_MESSAGES/k8s.po
// ../../../translations/test/default/LC_MESSAGES/k8s.mo
// ../../../translations/test/default/LC_MESSAGES/k8s.po
// ../../../translations/test/en_US/LC_MESSAGES/k8s.mo
// ../../../translations/test/en_US/LC_MESSAGES/k8s.po
// DO NOT EDIT!

package generated

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

var _translationsKubectlDefaultLc_messagesK8sMo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x93\xcf\x6e\xd3\x4e\x10\xc7\x27\xbf\xf6\x77\xf1\x91\x33\x87\x41\x02\xa9\x08\x8d\x59\xbb\xa5\xaa\x5c\x82\x50\x4b\x8b\x22\x25\x6a\x54\x5c\xd4\x23\x1b\xef\xd4\x59\xb4\xd9\xb5\x76\xd7\x4d\x7b\xe4\x21\x78\x05\x2e\xf0\x16\xbc\x02\x67\x9e\x05\x39\x86\x14\x6e\xbd\xb4\x7b\xb0\xfc\x1d\x7d\xfc\x9d\x3f\x9e\xfd\xf9\x60\xf3\x33\x00\xc0\x06\x00\x3c\x04\x80\x1d\x00\xf8\x1f\x00\xc6\xd0\x9f\x0f\x00\x30\x02\x00\x09\x00\x5b\x03\x80\xaf\x00\xf0\x6d\x00\xf0\x63\x00\xf0\x18\x00\x3e\x6d\x00\x7c\x07\x80\x2f\x1b\x00\x83\xdf\x3e\x7f\xce\x7f\xdd\xe3\xac\x51\x32\x32\xc6\x39\xa3\xb4\xd6\x45\x19\xb5\xb3\x01\x9d\x45\x89\x9e\x83\x6b\x7d\xc5\x9b\xb7\x81\x60\x29\x63\x35\x47\xdd\x85\xcd\x35\x86\xb6\x69\x9c\x8f\xac\x3a\x4a\x5b\xa5\x2f\xb5\x6a\xa5\x59\xe3\x01\xa5\x55\x6b\x85\x95\x33\x86\xab\xde\x96\xf0\x89\xfa\x8b\x5b\xb2\x67\xbc\x70\xad\x55\x9b\x77\x9f\xe2\x1e\xba\x80\xa9\x77\x1f\xb9\x8a\x34\x52\xf4\x9e\x7d\xd0\xce\x16\x58\x73\x8c\x7c\x15\xa9\x76\xc4\x57\x72\xd1\x18\x0e\x34\x67\x63\x5c\x72\xca\x5d\x05\x34\x09\xb5\x56\x74\xd0\xd6\x81\x4a\x57\x60\x32\x3d\x29\xe9\xd0\xf3\xea\x4f\xd0\x1b\x19\xb9\xc0\x5c\x64\xdb\x94\xe5\x94\xe5\x98\x8b\x42\x6c\x3f\x13\x42\x88\x64\x7a\x42\xa7\x7c\xa9\xc3\x3f\xdc\xee\x8a\xdb\xc6\x3c\x2b\x5e\xec\x90\xd8\x13\x22\x19\xcb\x10\xa9\xf4\xd2\x06\x23\xa3\xf3\x05\x1e\x78\xb6\x4a\x5a\x3c\x68\xbd\x0d\xf8\x72\xd6\xcb\x54\xa5\xb3\x2e\xf0\xba\x5e\x48\x6d\xd2\xca\x2d\x5e\x25\x93\xd1\xe4\xe8\xa6\x95\x2c\x15\xc9\xa1\xb3\x91\x6d\xa4\xf2\xba\xe1\x02\xbb\xce\x9e\x37\x46\x6a\xbb\x8f\xd5\x5c\xfa\xc0\x71\x78\x56\x1e\xd3\xde\x0d\xd7\xe5\xbd\x60\x4f\x47\xb6\x72\x4a\xdb\xba\xc0\xbd\x99\x8e\xc9\x39\xbd\x65\xcb\xbe\x2f\x68\xea\x58\xe9\x88\x59\xba\x9b\x66\x22\x39\xa7\x5e\xd3\xbb\xd5\x84\x0f\x7b\xdf\x02\x7b\xe3\xb1\xb4\x75\x2b\x6b\xa6\x92\xe5\xa2\x1b\x97\x69\xbd\x34\x74\xec\xfc\x22\x14\x68\x9b\x95\x0c\xc3\x7c\x1f\xfb\xd7\xe1\x96\xc5\x47\x43\xcc\x9e\xee\xaf\x3f\x2d\x90\x6d\x72\xab\xfb\x71\x47\x4b\x83\x4b\x19\xee\x6f\x2b\x7f\x05\x00\x00\xff\xff\xaa\xa3\x15\xc3\x6a\x04\x00\x00")

func translationsKubectlDefaultLc_messagesK8sMoBytes() ([]byte, error) {
	return bindataRead(
		_translationsKubectlDefaultLc_messagesK8sMo,
		"translations/kubectl/default/LC_MESSAGES/k8s.mo",
	)
}

func translationsKubectlDefaultLc_messagesK8sMo() (*asset, error) {
	bytes, err := translationsKubectlDefaultLc_messagesK8sMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "translations/kubectl/default/LC_MESSAGES/k8s.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _translationsKubectlDefaultLc_messagesK8sPo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x53\xdd\x6e\xd3\x4c\x10\xbd\xcf\x53\xcc\xe7\xe8\x93\x5a\xc1\x1a\x3b\x85\xaa\x72\x09\xa2\x0d\x6d\xa9\x68\xd5\x28\x75\x11\x12\x20\xb4\xf1\x4e\x9c\x85\xf5\xac\xb5\x33\xee\xcf\xdb\x23\xdb\x69\x43\x90\x90\x7a\xd3\xde\xcd\xce\x9e\x99\x33\x67\xf6\xec\x10\x72\x64\x01\x09\x9a\xd8\x69\xb1\x9e\x18\x16\x3e\x40\x43\x56\x40\x90\x85\xe3\xc1\x10\x26\xbe\xbe\x0b\xb6\x5c\x0a\x6c\x4d\xb6\x61\x94\xa4\xbb\x83\x21\xe4\x4b\xcb\xb0\xb0\x0e\xc1\x32\x18\xcb\x12\xec\xbc\x11\x34\xd0\x90\xc1\x00\xb2\x44\x60\x5d\x21\x38\x5b\x20\x31\x82\xe6\x2e\x37\x3d\x98\x7c\x3a\x38\x39\x82\x5a\x17\xbf\x74\x89\x6d\xfb\xe3\xd3\xd9\x65\x0e\x07\x57\xf9\xc7\x8b\x19\xcc\x03\x92\xd1\x14\x9b\x78\xde\x04\xe2\xf7\x65\xa5\xad\x8b\x0b\x5f\xbd\xec\x88\xe3\xc1\x70\x50\x71\x69\x0d\x44\x51\x1b\xb0\x84\x36\x8a\xa6\xc1\xff\xc4\x42\xd4\xa9\x51\x9f\x31\xb0\xf5\x94\x41\x89\x22\x78\x2b\xaa\xf4\x0a\x6f\x75\x55\x3b\x64\xb5\x44\xe7\xfc\x37\x8a\x06\xd1\x0c\x6b\x1f\x44\x9d\xb7\xcd\xd4\x61\x53\xb2\xca\x7d\x06\xdd\xd5\xf4\x22\x57\x93\x80\xdd\x3e\xd4\x07\x2d\x98\xb5\xdc\x3b\x2a\x1d\xa9\x74\x04\xa3\x24\x4b\x76\x5e\x24\x49\x92\xac\xc0\x6a\x86\xd7\x96\x37\xb0\xbb\x1d\x76\x07\x46\x69\xf6\xe6\xb5\x4a\xf6\x56\xd8\x33\xcd\xa2\xf2\xd5\xb2\x7d\xc8\xe0\xb0\x57\x0b\x87\xad\x56\x78\xfb\x4f\xf1\xef\xba\xf2\xf3\xd3\xf3\xa3\xb5\xbc\x34\xee\x9b\x4e\x3c\x09\x92\xa8\xfc\xae\xc6\x0c\x5a\xc5\xaf\x6a\xa7\x2d\xed\x43\xb1\xd4\x81\x51\xc6\x57\xf9\xb1\xda\xdb\xc4\xb6\x33\x2c\x30\xa8\x23\x2a\xbc\xb1\x54\x66\xb0\x37\xb7\xd2\x61\xbe\xa8\x13\x24\x0c\xfd\x80\x53\x8f\xc6\x0a\xa4\xf1\x6e\x9c\x26\xab\xeb\x3e\xa7\x2e\x7d\x13\x0a\x9c\xf4\x1c\x19\xac\x49\xce\x34\x95\x8d\x2e\x51\xe5\xa8\xab\xfb\x95\xba\x26\x68\xa7\x8e\x7d\xa8\x38\x03\xaa\xbb\x23\x8f\x47\xfb\xd0\x87\xe3\x2d\x82\xff\xc6\x90\x6e\xef\x6f\xb4\xc8\x00\xa9\x4d\xb4\x4f\x5d\xc8\xad\x40\x74\x55\x1b\x2d\xd8\x59\x49\x13\x79\x59\x99\xd6\x13\x68\x08\xc8\xdd\x4c\xd1\xbd\x45\x1e\x0d\xee\x6c\xf4\x38\xf4\x7a\x94\x68\x10\xdd\x68\x29\x96\xed\x07\xf0\xe4\xee\x80\x9b\xba\xf5\x14\x9a\xb6\xc6\x92\xb1\xd7\xd6\x34\xda\x3d\x14\x33\x68\x32\x0f\x27\x28\xbc\x73\x58\xf4\x24\x0a\xa2\x41\xf4\xbf\xf9\x03\x7a\x83\x01\x61\xe1\x1b\x32\xd1\xda\xf1\xcf\x4a\xf8\xa3\x7f\x9a\x67\xe6\x65\x09\x5f\x93\xef\x4f\x4b\x0a\x37\x9a\xff\xa2\x4c\x9f\x98\x72\x53\xe7\xef\x00\x00\x00\xff\xff\x96\x33\xff\xe8\x78\x05\x00\x00")

func translationsKubectlDefaultLc_messagesK8sPoBytes() ([]byte, error) {
	return bindataRead(
		_translationsKubectlDefaultLc_messagesK8sPo,
		"translations/kubectl/default/LC_MESSAGES/k8s.po",
	)
}

func translationsKubectlDefaultLc_messagesK8sPo() (*asset, error) {
	bytes, err := translationsKubectlDefaultLc_messagesK8sPoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "translations/kubectl/default/LC_MESSAGES/k8s.po", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _translationsKubectlEn_usLc_messagesK8sMo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x93\xcf\x6e\xd3\x40\x10\xc6\x27\xb4\x5c\x7c\xe4\xcc\x61\x90\x40\x2a\x42\x63\xd6\x0e\x54\x95\x4b\x10\x6a\x69\x51\xa5\x44\x8d\x8a\x8b\x7a\x64\xe3\x9d\x3a\x8b\x36\xbb\xd6\xee\xba\x69\x8f\x3c\x04\xaf\xc0\x05\xde\x82\x57\xe0\xcc\xb3\x20\xc7\x90\xc2\xad\x97\x76\x0f\x96\xbf\xd1\xcf\xdf\xfc\xf1\xec\xaf\x07\x9b\x5f\x00\x00\x36\x00\xe0\x21\x00\xbc\x00\x80\xfb\x00\x30\x86\xfe\x7c\x04\x80\x23\x00\x90\x00\xb0\x35\x00\xf8\x06\x00\xdf\x07\x00\x3f\x07\x00\x8f\x01\xe0\xf3\x06\xc0\x0f\x00\xf8\xba\x01\x30\xf8\xe3\xf3\xf7\xdc\xeb\x1e\xa7\x8d\x92\x91\x31\xce\x19\xa5\xb5\x2e\xca\xa8\x9d\x0d\xe8\x2c\x4a\xf4\x1c\x5c\xeb\x2b\xde\xbc\x09\x04\x4b\x19\xab\x39\xea\x2e\x6c\xae\x30\xb4\x4d\xe3\x7c\x64\xd5\x51\xda\x2a\x7d\xa1\x55\x2b\xcd\x1a\x0f\x28\xad\x5a\x2b\xac\x9c\x31\x5c\xf5\xb6\x84\x4f\xd4\x3f\xdc\x92\x3d\xe3\xb9\x6b\xad\xda\xbc\xfd\x14\x77\xd0\x05\x4c\xbd\xfb\xc4\x55\xa4\x23\x45\x1f\xd8\x07\xed\x6c\x81\x35\xc7\xc8\x97\x91\x6a\x47\x7c\x29\x17\x8d\xe1\x40\x73\x36\xc6\x25\x27\xdc\x55\x40\x93\x50\x6b\x45\x7b\x6d\x1d\xa8\x74\x05\x26\xd3\xe3\x92\xf6\x3d\xaf\xfe\x04\xbd\x95\x91\x0b\xcc\x45\x36\xa4\x2c\xa7\x2c\xc7\x5c\x14\x62\xf8\x4c\x08\x21\x92\xe9\x31\x9d\xf0\x85\x0e\xff\x71\xdb\x2b\x6e\x88\x79\x5e\x0c\x5f\x92\xd8\x11\x22\x19\xcb\x10\xa9\xf4\xd2\x06\x23\xa3\xf3\x05\xee\x79\xb6\x4a\x5a\xdc\x6b\xbd\x0d\xf8\x6a\xd6\xcb\x54\xa5\xb3\x2e\xf0\xa6\x5e\x48\x6d\xd2\xca\x2d\x5e\x27\x93\xa3\xc9\xc1\x75\x2b\x59\x2a\x92\x7d\x67\x23\xdb\x48\xe5\x55\xc3\x05\x76\x9d\x3d\x6f\x8c\xd4\x76\x17\xab\xb9\xf4\x81\xe3\xe8\xb4\x3c\xa4\x9d\x6b\xae\xcb\x7b\xce\x9e\x0e\x6c\xe5\x94\xb6\x75\x81\x3b\x33\x1d\x93\x33\x7a\xc7\x96\x7d\x5f\xd0\xd4\xb1\xd2\x11\xb3\x74\x3b\xcd\x44\x72\x46\xbd\xa6\xf7\xab\x09\xef\xf7\xbe\x05\xf6\xc6\x63\x69\xeb\x56\xd6\x4c\x25\xcb\x45\x37\x2e\xd3\x7a\x69\xe8\xd0\xf9\x45\x28\xd0\x36\x2b\x19\x46\xf9\x2e\xf6\xaf\xa3\x2d\x8b\x8f\x46\x98\x3d\xdd\x5d\x7f\x5a\x20\xdb\xe4\x46\xf7\xe3\x96\x96\x06\x97\x32\xdc\xdd\x56\xfe\x0e\x00\x00\xff\xff\x52\x15\x0d\x38\x6a\x04\x00\x00")

func translationsKubectlEn_usLc_messagesK8sMoBytes() ([]byte, error) {
	return bindataRead(
		_translationsKubectlEn_usLc_messagesK8sMo,
		"translations/kubectl/en_US/LC_MESSAGES/k8s.mo",
	)
}

func translationsKubectlEn_usLc_messagesK8sMo() (*asset, error) {
	bytes, err := translationsKubectlEn_usLc_messagesK8sMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "translations/kubectl/en_US/LC_MESSAGES/k8s.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _translationsKubectlEn_usLc_messagesK8sPo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x53\x5d\x4f\xdb\x4a\x10\x7d\xcf\xaf\x98\xeb\xe8\x4a\xa0\x7b\xd7\xb5\x13\x15\x21\xd3\x54\x85\x14\x28\x2a\x88\x28\x98\xaa\x52\x5b\x55\x1b\xef\xc4\xd9\x76\x3d\x6b\xed\x8c\xf9\xf8\xf7\x95\xed\x40\x9a\x4a\x95\x78\x81\xb7\xd9\xd9\x33\x73\xe6\xcc\x9e\x1d\x42\x8e\x2c\x20\x41\x13\x3b\x2d\xd6\x13\xc3\xd2\x07\x68\xc8\x0a\x08\xb2\x70\x3c\x18\xc2\xd4\xd7\xf7\xc1\x96\x2b\x81\x9d\xe9\x2e\x8c\x92\x74\x6f\x30\x84\x7c\x65\x19\x96\xd6\x21\x58\x06\x63\x59\x82\x5d\x34\x82\x06\x1a\x32\x18\x40\x56\x08\xac\x2b\x04\x67\x0b\x24\x46\xd0\xdc\xe5\x66\x87\xd3\x8f\x87\xa7\xc7\x50\xeb\xe2\xa7\x2e\xb1\x6d\x7f\x72\x36\xbf\xca\xe1\xf0\x3a\xff\x70\x39\x87\x45\x40\x32\x9a\x62\x13\x2f\x9a\x40\xfc\xae\xac\xb4\x75\x71\xe1\xab\xff\x3b\xe2\x78\x30\x1c\x54\x5c\x5a\x03\x51\xd4\x06\x2c\xa1\x8d\xa2\x59\xf0\x3f\xb0\x10\x75\x66\xd4\x27\x0c\x6c\x3d\x65\x50\xa2\x08\xde\x89\x2a\xbd\xc2\x3b\x5d\xd5\x0e\x59\xad\xd0\x39\xff\x95\xa2\x41\x34\xc7\xda\x07\x51\x17\x6d\x33\x75\xd4\x94\xac\x72\x9f\x41\x77\x35\xbb\xcc\xd5\x34\x60\xb7\x0f\xf5\x5e\x0b\x66\x2d\xf7\x58\xa5\x23\x95\x8e\x60\x94\x64\xc9\xf8\xbf\x24\x49\x92\x35\x58\xcd\xf1\xc6\xf2\x16\x76\xaf\xc3\x8e\x61\x34\xca\xc6\xaf\x55\xb2\xbf\xc6\x9e\x6b\x16\x95\xaf\x97\xed\x43\x06\x47\xbd\x5a\x38\x6a\xb5\xc2\x9b\xbf\x8a\x7f\xdb\x95\x5f\x9c\x5d\x1c\x6f\xe4\xa5\x71\xdf\x74\xea\x49\x90\x44\xe5\xf7\x35\x66\xd0\x2a\x7e\x55\x3b\x6d\xe9\x00\x8a\x95\x0e\x8c\x32\xb9\xce\x4f\xd4\xfe\x36\xb6\x9d\x61\x89\x41\x1d\x53\xe1\x8d\xa5\x32\x83\xfd\x85\x95\x0e\xf3\x59\x9d\x22\x61\xe8\x07\x9c\x79\x34\x56\x20\x8d\xf7\xe2\x34\x59\x5f\xf7\x39\x75\xe5\x9b\x50\xe0\xb4\xe7\xc8\x60\x43\x72\xae\xa9\x6c\x74\x89\x2a\x47\x5d\x3d\xac\xd4\x35\x41\x3b\x75\xe2\x43\xc5\x19\x50\xdd\x1d\x79\x32\x3a\x80\x3e\x9c\xec\x10\xfc\x33\x81\x74\xf7\x60\xab\x45\x06\x48\x6d\xa2\x7d\xea\x42\xee\x04\xa2\xeb\xda\x68\xc1\xce\x4a\x9a\xc8\xcb\xda\xb4\x9e\x40\x43\x40\xee\x66\x8a\x1e\x2c\xf2\x64\x70\x67\xa3\xa7\xa1\x37\xa3\x44\x83\xe8\x56\x4b\xb1\x6a\x3f\x80\x27\x77\x0f\xdc\xd4\xad\xa7\xd0\xb4\x35\x96\x8c\xbd\xb1\xa6\xd1\xee\xb1\x98\x41\x93\x79\x3c\x41\xe1\x9d\xc3\xa2\x27\x51\x10\x0d\xa2\x7f\xcd\x6f\xd0\x5b\x0c\x08\x4b\xdf\x90\x89\x36\x8e\x7f\x51\xc2\xef\xfd\xd3\xbc\x30\x2f\x4b\xf8\x92\x7c\x7b\x5e\x52\xb8\xd5\xfc\x07\x65\xfa\xcc\x94\xdb\x3a\x7f\x05\x00\x00\xff\xff\x61\x66\xb9\x11\x78\x05\x00\x00")

func translationsKubectlEn_usLc_messagesK8sPoBytes() ([]byte, error) {
	return bindataRead(
		_translationsKubectlEn_usLc_messagesK8sPo,
		"translations/kubectl/en_US/LC_MESSAGES/k8s.po",
	)
}

func translationsKubectlEn_usLc_messagesK8sPo() (*asset, error) {
	bytes, err := translationsKubectlEn_usLc_messagesK8sPoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "translations/kubectl/en_US/LC_MESSAGES/k8s.po", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _translationsTestDefaultLc_messagesK8sMo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x51\xc1\x6e\x13\x31\x10\x7d\x0d\xe1\xb2\x47\x2e\x5c\x38\x18\xa1\x4a\x20\x34\x8b\x77\x03\x55\xe4\x10\x84\x12\x5a\x54\x94\xa8\x51\x59\x50\x6f\xe0\x64\xa7\x1b\xa3\x5d\x7b\x65\x3b\x50\x3e\x80\x4f\xe0\xc8\x1f\xf0\x4d\x7c\x0b\x4a\x36\xd0\x22\xe6\x60\xbd\x79\x7a\xef\xd9\xe3\xf9\x75\xa7\xff\x1d\x00\x6e\x01\xb8\x07\xe0\x29\x80\xdb\x00\x66\xe8\xea\x23\x80\x07\x00\x34\x80\xbb\x00\xbe\x01\xf8\x79\x00\xfc\x00\x70\x08\xe0\x4d\xaf\xf3\xb6\x3d\xe0\x60\xef\xe9\xed\xf3\x76\x15\x39\xc4\x0f\x6d\xbd\xf1\xba\xee\xdf\xc0\xf8\x0f\x87\xe8\x8d\xad\xfa\x37\x30\x16\xde\x7d\xe2\x55\xa4\xd3\x92\xde\xb3\x0f\xc6\x59\x25\x2a\x8e\x91\xaf\x22\x55\x8e\xf8\x4a\x37\x6d\xcd\x81\xd6\x5c\xd7\x2e\x39\xe7\xd6\xf9\x48\xf3\x50\x99\x92\x26\x9b\x2a\x50\xe1\x94\x48\x16\x67\x05\x4d\x3d\xeb\x68\x9c\xa5\x57\x3a\xb2\x12\xb9\xcc\x06\x94\xe5\x94\xe5\x22\x97\x4a\x0e\x1e\x4b\x29\x65\xb2\x38\xa3\x73\xfe\x6c\xc2\x3f\xba\xa3\x9d\x6e\x20\xf2\x4c\x0d\x9e\x91\x1c\x4a\x99\xcc\x74\x88\x54\x78\x6d\x43\xad\xa3\xf3\x4a\x4c\x3c\xdb\x52\x5b\x31\xd9\x78\x1b\xc4\xf3\x65\xd7\xa6\x65\xba\xdc\x12\x2f\xab\x46\x9b\x3a\x5d\xb9\xe6\x45\x32\x3f\x9d\x1f\x5f\x8f\x92\xa5\x32\x99\x3a\x1b\xd9\x46\x2a\xbe\xb6\xac\xc4\x76\xb2\x27\x6d\xad\x8d\x1d\x89\xd5\x5a\xfb\xc0\x71\xfc\xae\x38\xa1\xe1\xb5\x6e\x7b\xef\x25\x7b\x3a\xb6\x2b\x57\x1a\x5b\x29\x31\x5c\x9a\x98\x5c\xd0\x6b\xb6\xec\xbb\x07\x2d\x1c\x97\x26\x8a\x2c\x3d\x4a\x33\x99\x5c\x50\xd7\xd3\x5b\xb7\xf1\x2b\x9e\x76\xb9\x4a\x74\xc1\x33\x6d\xab\x8d\xae\x98\x0a\xd6\xcd\xf6\xbb\x76\x2b\xa1\x13\xe7\x9b\xa0\x84\xed\x36\x14\xc6\xf9\x48\x74\x70\xfc\xd0\x8a\xfb\x63\x91\x3d\x1a\xfd\xb5\x2a\xc1\x36\x41\x5c\xb3\x67\xf1\x45\x07\x71\x58\x0a\x13\xb9\xf9\xc3\x6c\x8f\x3d\x15\x70\xe9\x1c\x7e\x07\x00\x00\xff\xff\x59\x11\x2e\xef\x74\x02\x00\x00")

func translationsTestDefaultLc_messagesK8sMoBytes() ([]byte, error) {
	return bindataRead(
		_translationsTestDefaultLc_messagesK8sMo,
		"translations/test/default/LC_MESSAGES/k8s.mo",
	)
}

func translationsTestDefaultLc_messagesK8sMo() (*asset, error) {
	bytes, err := translationsTestDefaultLc_messagesK8sMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "translations/test/default/LC_MESSAGES/k8s.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _translationsTestDefaultLc_messagesK8sPo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\x61\x6f\xd3\x30\x10\xfd\x9e\x5f\x71\xa4\x42\xda\x04\x0e\x49\x2a\xa6\x29\xa3\x88\x2d\xb4\xa3\x62\xd5\xaa\x2e\x43\x48\x80\x90\x9b\x5c\x13\x43\x62\x47\xbe\x0b\x74\xff\x1e\x39\xc9\xa8\xca\xc4\x97\xc8\x7e\xef\xdd\x3d\xbf\xcb\x4d\x20\x43\x62\x60\x2b\x35\xd5\x92\x95\xd1\x04\x3b\x63\xa1\xd3\x8a\x81\x91\x98\x02\x6f\x02\xa9\x69\x1f\xac\x2a\x2b\x86\x93\xf4\x14\xe2\x30\x3a\xf3\x26\x90\x55\x8a\x60\xa7\x6a\x04\x45\x50\x28\x62\xab\xb6\x1d\x63\x01\x9d\x2e\xd0\x02\x57\x08\x24\x1b\x84\x5a\xe5\xa8\x09\x41\x52\x8f\xad\x2f\xd3\x8f\x97\xd7\x73\x68\x65\xfe\x53\x96\xe8\xda\x2f\x96\x9b\xbb\x0c\x2e\xef\xb3\x0f\xb7\x1b\xd8\x5a\xd4\x85\xd4\x41\x11\x6c\x3b\xab\xe9\x5d\xd9\x48\x55\x07\xb9\x69\x5e\xf6\xc6\x81\x37\xf1\x1a\x2a\x55\x01\xbe\xef\x0e\xc4\xd6\x9d\xfc\xb5\x35\x3f\x30\x67\xb1\x2c\xc4\x27\xb4\xa4\x8c\x4e\xa0\x44\x66\xdc\xb3\x28\x8d\xc0\xbd\x6c\xda\x1a\x49\x54\x58\xd7\xe6\xab\xf6\x3d\x7f\x83\xad\xb1\x2c\x56\xae\x99\xb8\xea\x4a\x12\x99\x49\xa0\xa7\xd6\xb7\x99\x48\x2d\xf6\xf3\x10\xef\x25\x63\xe2\xbc\xa7\x22\x8a\x45\x14\x43\x1c\x26\xe1\xf4\x45\x18\x86\xe1\x28\x16\x1b\xfc\xa5\xe8\x48\x7b\xd6\x6b\xa7\x10\x47\xc9\xf4\xb5\x08\xcf\x47\xed\x8d\x24\x16\xd9\x38\x6c\x63\x13\xb8\x1a\xd2\xc2\x95\xcb\x0a\x6f\xfe\x1b\xfe\x6d\x5f\xbe\x5a\xae\xe6\x87\x78\x51\x30\x34\x4d\x8d\x66\xd4\x2c\xb2\x87\x16\x13\x70\x89\x5f\xb5\xb5\x54\xfa\x02\xf2\x4a\x5a\x42\x9e\xdd\x67\x0b\x71\x7e\xac\x75\x6f\xd8\xa1\x15\x73\x9d\x9b\x42\xe9\x32\x81\xf3\xad\xe2\x5e\xf3\x59\x5c\xa3\x46\x3b\x3c\x70\x6d\xb0\x50\x0c\x51\x70\x16\x44\xe1\x48\x0f\x98\xb8\x33\x9d\xcd\x31\x1d\x3c\x12\x38\x98\xdc\x48\x5d\x76\xb2\x44\x91\xa1\x6c\x1e\x47\x5a\x77\x56\xd6\x62\x61\x6c\x43\x09\xe8\xb6\xbf\xd2\x2c\xbe\x80\xe1\x38\x3b\xd1\xf0\x6c\x06\xd1\xe9\xc5\x51\x8b\x04\x50\x3b\xc0\xfd\xea\x9c\xf7\x0c\xbe\x5b\xca\xef\x6e\xd9\x74\xe9\x3f\x6e\xc2\xbf\x58\xbf\x14\x3b\x63\x9e\xd4\x0d\x5e\xc7\x75\x47\xd8\x78\x7b\x4a\x11\xdb\x2f\xe1\x37\xf0\xb9\x42\x8b\xf0\x5b\x12\x3c\x2f\x40\x31\x36\x7f\xd9\xe8\xc0\xba\xcf\x48\x93\xef\xfd\x09\x00\x00\xff\xff\xbf\xc0\xcb\xd2\x64\x03\x00\x00")

func translationsTestDefaultLc_messagesK8sPoBytes() ([]byte, error) {
	return bindataRead(
		_translationsTestDefaultLc_messagesK8sPo,
		"translations/test/default/LC_MESSAGES/k8s.po",
	)
}

func translationsTestDefaultLc_messagesK8sPo() (*asset, error) {
	bytes, err := translationsTestDefaultLc_messagesK8sPoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "translations/test/default/LC_MESSAGES/k8s.po", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _translationsTestEn_usLc_messagesK8sMo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x51\xc1\x6e\x13\x31\x10\x7d\x0d\xe1\xb2\x47\x2e\x5c\x38\x18\xa1\x4a\x20\x34\x8b\x77\x83\xaa\xc8\x21\x08\x25\xb4\xa8\x28\x51\xa3\xb2\xa0\xde\xc0\xc9\x0e\x1b\xa3\x5d\x7b\x65\x3b\x50\xb8\xf3\x09\x1c\xf9\x03\xbe\x89\x6f\x41\xc9\x06\x5a\xd4\x39\x58\x6f\x9e\xde\x7b\xf6\x78\x7e\xdf\xe9\xff\x00\x80\x5b\x00\xee\x01\x78\x0a\xe0\x36\x80\x19\xba\xfa\x00\xe0\x01\x00\x0d\xe0\x2e\x80\xef\x00\x7e\x1d\x00\x3f\x01\x1c\x02\x78\xdd\xeb\xbc\x6d\x0f\x38\xd8\x7b\x7a\xfb\xbc\x5d\x45\x0e\xf1\x7d\x5b\x6f\xbc\xae\xfb\xd7\x30\x6e\xe0\x10\xbd\xb1\x55\xff\x1a\xc6\xc2\xbb\x4f\xbc\x8a\x74\x5a\xd2\x3b\xf6\xc1\x38\xab\x44\xc5\x31\xf2\x65\xa4\xca\x11\x5f\xea\xa6\xad\x39\xd0\x9a\xeb\xda\x25\xe7\xdc\x3a\x1f\x69\x1e\x2a\x53\xd2\x64\x53\x05\x2a\x9c\x12\xc9\xe2\xac\xa0\xa9\x67\x1d\x8d\xb3\xf4\x52\x47\x56\x22\x97\xd9\x80\xb2\x9c\xb2\x5c\xe4\x52\xc9\xc1\x63\x29\xa5\x4c\x16\x67\x74\xce\x9f\x4d\xf8\x4f\x77\xb4\xd3\x0d\x44\x9e\xab\x2c\x27\x39\x94\x32\x99\xe9\x10\xa9\xf0\xda\x86\x5a\x47\xe7\x95\x98\x78\xb6\xa5\xb6\x62\xb2\xf1\x36\x88\x67\xcb\xae\x4d\xcb\x74\xb9\x25\x5e\x54\x8d\x36\x75\xba\x72\xcd\xf3\x64\x7e\x3a\x3f\xbe\x1a\x25\x4b\x65\x32\x75\x36\xb2\x8d\x54\x7c\x6d\x59\x89\xed\x64\x4f\xda\x5a\x1b\x3b\x12\xab\xb5\xf6\x81\xe3\xf8\x6d\x71\x42\xc3\x2b\xdd\xf6\xde\x8f\xec\xe9\xd8\xae\x5c\x69\x6c\xa5\xc4\x70\x69\x62\x72\x41\xaf\xd8\xb2\xef\x1e\xb4\x70\x5c\x9a\x28\xb2\xf4\x28\xcd\x64\x72\x41\x5d\x4f\x6f\xdc\xc6\xaf\x78\xda\xe5\x2a\xd1\x05\xcf\xb4\xad\x36\xba\x62\x2a\x58\x37\xdb\xef\xda\xad\x84\x4e\x9c\x6f\x82\x12\xb6\xdb\x50\x18\xe7\x23\xd1\xc1\xf1\x43\x2b\xee\x8f\x45\xf6\x68\xf4\xcf\xaa\x04\xdb\x04\x71\xcd\x9e\xc5\x17\x1d\xc4\x61\x29\x4c\xe4\xe6\x2f\xb3\x3d\xf6\x54\xc0\x52\x7f\xc3\x9f\x00\x00\x00\xff\xff\xa9\x73\xeb\x8c\x74\x02\x00\x00")

func translationsTestEn_usLc_messagesK8sMoBytes() ([]byte, error) {
	return bindataRead(
		_translationsTestEn_usLc_messagesK8sMo,
		"translations/test/en_US/LC_MESSAGES/k8s.mo",
	)
}

func translationsTestEn_usLc_messagesK8sMo() (*asset, error) {
	bytes, err := translationsTestEn_usLc_messagesK8sMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "translations/test/en_US/LC_MESSAGES/k8s.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _translationsTestEn_usLc_messagesK8sPo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\x61\x6f\xd3\x30\x10\xfd\x9e\x5f\x71\xa4\x42\xda\x04\x0e\x49\x26\x4d\x53\x46\x11\x6b\x68\x47\xc5\xaa\x55\x5d\x86\x90\x00\x21\x37\xb9\x26\x86\xc4\x8e\x7c\x17\xe8\xf8\xf5\xc8\x49\x46\x55\x26\xbe\x44\xf6\x7b\xef\xee\xf9\x5d\x6e\x02\x19\x12\x03\x5b\xa9\xa9\x96\xac\x8c\x26\xd8\x19\x0b\x9d\x56\x0c\x8c\xc4\x14\x78\x13\x48\x4d\xfb\x60\x55\x59\x31\x9c\xa4\xa7\x10\x87\xd1\xb9\x37\x81\xac\x52\x04\x3b\x55\x23\x28\x82\x42\x11\x5b\xb5\xed\x18\x0b\xe8\x74\x81\x16\xb8\x42\x20\xd9\x20\xd4\x2a\x47\x4d\x08\x92\x7a\x6c\x7d\x95\x7e\xb8\xba\x9e\x43\x2b\xf3\x1f\xb2\x44\xd7\x7e\xb1\xdc\xdc\x65\x70\x75\x9f\xbd\xbf\xdd\xc0\xd6\xa2\x2e\xa4\x0e\x8a\x60\xdb\x59\x4d\x6f\xcb\x46\xaa\x3a\xc8\x4d\xf3\xb2\x37\x0e\xbc\x89\xd7\x50\xa9\x0a\xf0\x7d\x77\x20\xb6\xee\xe4\xaf\xad\xf9\x8e\x39\x8b\x65\x21\x3e\xa2\x25\x65\x74\x02\x25\x32\xe3\x9e\x45\x69\x04\xee\x65\xd3\xd6\x48\xa2\xc2\xba\x36\x5f\xb4\xef\xf9\x1b\x6c\x8d\x65\xb1\x72\xcd\xc4\xac\x2b\x49\x64\x26\x81\x9e\x5a\xdf\x66\x22\xb5\xd8\xcf\x43\xbc\x93\x8c\x89\xf3\x3e\x13\x51\x2c\xa2\x18\xe2\x30\x09\xcf\x5e\x84\x61\x18\x8e\x62\xb1\xc1\x9f\x8a\x8e\xb4\xe7\xbd\xf6\x0c\xe2\x38\x89\x62\x11\x5e\x8c\xda\x1b\x49\x2c\xb2\x71\xd8\xc6\x26\x30\x1b\xd2\xc2\xcc\x65\x85\xd7\xff\x0d\xff\xa6\x2f\x5f\x2d\x57\xf3\x43\xbc\x28\x18\x9a\xa6\x46\x33\x6a\x16\xd9\x43\x8b\x09\xb8\xc4\xaf\xda\x5a\x2a\x7d\x09\x79\x25\x2d\x21\x4f\xef\xb3\x85\xb8\x38\xd6\xba\x37\xec\xd0\x8a\xb9\xce\x4d\xa1\x74\x99\xc0\xc5\x56\x71\xaf\xf9\x24\xae\x51\xa3\x1d\x1e\xb8\x36\x58\x28\x86\x28\x38\x0f\xa2\x70\xa4\x07\x4c\xdc\x99\xce\xe6\x98\x0e\x1e\x09\x1c\x4c\x6e\xa4\x2e\x3b\x59\xa2\xc8\x50\x36\x8f\x23\xad\x3b\x2b\x6b\xb1\x30\xb6\xa1\x04\x74\xdb\x5f\x69\x1a\x5f\xc2\x70\x9c\x9e\x68\x78\x36\x85\xe8\xf4\xf2\xa8\x45\x02\xa8\x1d\xe0\x7e\x75\xce\x7b\x06\xdf\x2d\xe5\x37\xb7\x6c\xba\xf4\x1f\x37\xe1\x5f\xac\x5f\x8a\xad\xfc\xfd\xa4\x6e\xf0\x3a\xae\x3b\xc2\xc6\xdb\x53\x8a\xd8\x7e\x0e\xbf\x82\xcf\x15\x5a\x84\x5f\x92\xe0\x79\x01\x8a\xb1\xf9\xcb\x46\x07\xd6\x7d\x46\x9a\x7c\xef\x4f\x00\x00\x00\xff\xff\x03\xfe\xb1\xf0\x64\x03\x00\x00")

func translationsTestEn_usLc_messagesK8sPoBytes() ([]byte, error) {
	return bindataRead(
		_translationsTestEn_usLc_messagesK8sPo,
		"translations/test/en_US/LC_MESSAGES/k8s.po",
	)
}

func translationsTestEn_usLc_messagesK8sPo() (*asset, error) {
	bytes, err := translationsTestEn_usLc_messagesK8sPoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "translations/test/en_US/LC_MESSAGES/k8s.po", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"translations/kubectl/default/LC_MESSAGES/k8s.mo": translationsKubectlDefaultLc_messagesK8sMo,
	"translations/kubectl/default/LC_MESSAGES/k8s.po": translationsKubectlDefaultLc_messagesK8sPo,
	"translations/kubectl/en_US/LC_MESSAGES/k8s.mo":   translationsKubectlEn_usLc_messagesK8sMo,
	"translations/kubectl/en_US/LC_MESSAGES/k8s.po":   translationsKubectlEn_usLc_messagesK8sPo,
	"translations/test/default/LC_MESSAGES/k8s.mo":    translationsTestDefaultLc_messagesK8sMo,
	"translations/test/default/LC_MESSAGES/k8s.po":    translationsTestDefaultLc_messagesK8sPo,
	"translations/test/en_US/LC_MESSAGES/k8s.mo":      translationsTestEn_usLc_messagesK8sMo,
	"translations/test/en_US/LC_MESSAGES/k8s.po":      translationsTestEn_usLc_messagesK8sPo,
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
	"translations": {nil, map[string]*bintree{
		"kubectl": {nil, map[string]*bintree{
			"default": {nil, map[string]*bintree{
				"LC_MESSAGES": {nil, map[string]*bintree{
					"k8s.mo": {translationsKubectlDefaultLc_messagesK8sMo, map[string]*bintree{}},
					"k8s.po": {translationsKubectlDefaultLc_messagesK8sPo, map[string]*bintree{}},
				}},
			}},
			"en_US": {nil, map[string]*bintree{
				"LC_MESSAGES": {nil, map[string]*bintree{
					"k8s.mo": {translationsKubectlEn_usLc_messagesK8sMo, map[string]*bintree{}},
					"k8s.po": {translationsKubectlEn_usLc_messagesK8sPo, map[string]*bintree{}},
				}},
			}},
		}},
		"test": {nil, map[string]*bintree{
			"default": {nil, map[string]*bintree{
				"LC_MESSAGES": {nil, map[string]*bintree{
					"k8s.mo": {translationsTestDefaultLc_messagesK8sMo, map[string]*bintree{}},
					"k8s.po": {translationsTestDefaultLc_messagesK8sPo, map[string]*bintree{}},
				}},
			}},
			"en_US": {nil, map[string]*bintree{
				"LC_MESSAGES": {nil, map[string]*bintree{
					"k8s.mo": {translationsTestEn_usLc_messagesK8sMo, map[string]*bintree{}},
					"k8s.po": {translationsTestEn_usLc_messagesK8sPo, map[string]*bintree{}},
				}},
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
