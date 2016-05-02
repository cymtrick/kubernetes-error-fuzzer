// Copyright 2015 go-dockerclient authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package docker

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// APIImages represent an image returned in the ListImages call.
type APIImages struct {
	ID          string            `json:"Id" yaml:"Id"`
	RepoTags    []string          `json:"RepoTags,omitempty" yaml:"RepoTags,omitempty"`
	Created     int64             `json:"Created,omitempty" yaml:"Created,omitempty"`
	Size        int64             `json:"Size,omitempty" yaml:"Size,omitempty"`
	VirtualSize int64             `json:"VirtualSize,omitempty" yaml:"VirtualSize,omitempty"`
	ParentID    string            `json:"ParentId,omitempty" yaml:"ParentId,omitempty"`
	RepoDigests []string          `json:"RepoDigests,omitempty" yaml:"RepoDigests,omitempty"`
	Labels      map[string]string `json:"Labels,omitempty" yaml:"Labels,omitempty"`
}

// Image is the type representing a docker image and its various properties
type Image struct {
	ID              string    `json:"Id" yaml:"Id"`
	RepoTags        []string  `json:"RepoTags,omitempty" yaml:"RepoTags,omitempty"`
	Parent          string    `json:"Parent,omitempty" yaml:"Parent,omitempty"`
	Comment         string    `json:"Comment,omitempty" yaml:"Comment,omitempty"`
	Created         time.Time `json:"Created,omitempty" yaml:"Created,omitempty"`
	Container       string    `json:"Container,omitempty" yaml:"Container,omitempty"`
	ContainerConfig Config    `json:"ContainerConfig,omitempty" yaml:"ContainerConfig,omitempty"`
	DockerVersion   string    `json:"DockerVersion,omitempty" yaml:"DockerVersion,omitempty"`
	Author          string    `json:"Author,omitempty" yaml:"Author,omitempty"`
	Config          *Config   `json:"Config,omitempty" yaml:"Config,omitempty"`
	Architecture    string    `json:"Architecture,omitempty" yaml:"Architecture,omitempty"`
	Size            int64     `json:"Size,omitempty" yaml:"Size,omitempty"`
	VirtualSize     int64     `json:"VirtualSize,omitempty" yaml:"VirtualSize,omitempty"`
	RepoDigests     []string  `json:"RepoDigests,omitempty" yaml:"RepoDigests,omitempty"`
}

// ImagePre012 serves the same purpose as the Image type except that it is for
// earlier versions of the Docker API (pre-012 to be specific)
type ImagePre012 struct {
	ID              string    `json:"id"`
	Parent          string    `json:"parent,omitempty"`
	Comment         string    `json:"comment,omitempty"`
	Created         time.Time `json:"created"`
	Container       string    `json:"container,omitempty"`
	ContainerConfig Config    `json:"container_config,omitempty"`
	DockerVersion   string    `json:"docker_version,omitempty"`
	Author          string    `json:"author,omitempty"`
	Config          *Config   `json:"config,omitempty"`
	Architecture    string    `json:"architecture,omitempty"`
	Size            int64     `json:"size,omitempty"`
}

var (
	// ErrNoSuchImage is the error returned when the image does not exist.
	ErrNoSuchImage = errors.New("no such image")

	// ErrMissingRepo is the error returned when the remote repository is
	// missing.
	ErrMissingRepo = errors.New("missing remote repository e.g. 'github.com/user/repo'")

	// ErrMissingOutputStream is the error returned when no output stream
	// is provided to some calls, like BuildImage.
	ErrMissingOutputStream = errors.New("missing output stream")

	// ErrMultipleContexts is the error returned when both a ContextDir and
	// InputStream are provided in BuildImageOptions
	ErrMultipleContexts = errors.New("image build may not be provided BOTH context dir and input stream")

	// ErrMustSpecifyNames is the error rreturned when the Names field on
	// ExportImagesOptions is nil or empty
	ErrMustSpecifyNames = errors.New("must specify at least one name to export")
)

// ListImagesOptions specify parameters to the ListImages function.
//
// See https://goo.gl/xBe1u3 for more details.
type ListImagesOptions struct {
	All     bool
	Filters map[string][]string
	Digests bool
	Filter  string
}

// ListImages returns the list of available images in the server.
//
// See https://goo.gl/xBe1u3 for more details.
func (c *Client) ListImages(opts ListImagesOptions) ([]APIImages, error) {
	path := "/images/json?" + queryString(opts)
	resp, err := c.do("GET", path, doOptions{})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var images []APIImages
	if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
		return nil, err
	}
	return images, nil
}

// ImageHistory represent a layer in an image's history returned by the
// ImageHistory call.
type ImageHistory struct {
	ID        string   `json:"Id" yaml:"Id"`
	Tags      []string `json:"Tags,omitempty" yaml:"Tags,omitempty"`
	Created   int64    `json:"Created,omitempty" yaml:"Created,omitempty"`
	CreatedBy string   `json:"CreatedBy,omitempty" yaml:"CreatedBy,omitempty"`
	Size      int64    `json:"Size,omitempty" yaml:"Size,omitempty"`
}

// ImageHistory returns the history of the image by its name or ID.
//
// See https://goo.gl/8bnTId for more details.
func (c *Client) ImageHistory(name string) ([]ImageHistory, error) {
	resp, err := c.do("GET", "/images/"+name+"/history", doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok && e.Status == http.StatusNotFound {
			return nil, ErrNoSuchImage
		}
		return nil, err
	}
	defer resp.Body.Close()
	var history []ImageHistory
	if err := json.NewDecoder(resp.Body).Decode(&history); err != nil {
		return nil, err
	}
	return history, nil
}

// RemoveImage removes an image by its name or ID.
//
// See https://goo.gl/V3ZWnK for more details.
func (c *Client) RemoveImage(name string) error {
	resp, err := c.do("DELETE", "/images/"+name, doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok && e.Status == http.StatusNotFound {
			return ErrNoSuchImage
		}
		return err
	}
	resp.Body.Close()
	return nil
}

// RemoveImageOptions present the set of options available for removing an image
// from a registry.
//
// See https://goo.gl/V3ZWnK for more details.
type RemoveImageOptions struct {
	Force   bool `qs:"force"`
	NoPrune bool `qs:"noprune"`
}

// RemoveImageExtended removes an image by its name or ID.
// Extra params can be passed, see RemoveImageOptions
//
// See https://goo.gl/V3ZWnK for more details.
func (c *Client) RemoveImageExtended(name string, opts RemoveImageOptions) error {
	uri := fmt.Sprintf("/images/%s?%s", name, queryString(&opts))
	resp, err := c.do("DELETE", uri, doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok && e.Status == http.StatusNotFound {
			return ErrNoSuchImage
		}
		return err
	}
	resp.Body.Close()
	return nil
}

// InspectImage returns an image by its name or ID.
//
// See https://goo.gl/jHPcg6 for more details.
func (c *Client) InspectImage(name string) (*Image, error) {
	resp, err := c.do("GET", "/images/"+name+"/json", doOptions{})
	if err != nil {
		if e, ok := err.(*Error); ok && e.Status == http.StatusNotFound {
			return nil, ErrNoSuchImage
		}
		return nil, err
	}
	defer resp.Body.Close()

	var image Image

	// if the caller elected to skip checking the server's version, assume it's the latest
	if c.SkipServerVersionCheck || c.expectedAPIVersion.GreaterThanOrEqualTo(apiVersion112) {
		if err := json.NewDecoder(resp.Body).Decode(&image); err != nil {
			return nil, err
		}
	} else {
		var imagePre012 ImagePre012
		if err := json.NewDecoder(resp.Body).Decode(&imagePre012); err != nil {
			return nil, err
		}

		image.ID = imagePre012.ID
		image.Parent = imagePre012.Parent
		image.Comment = imagePre012.Comment
		image.Created = imagePre012.Created
		image.Container = imagePre012.Container
		image.ContainerConfig = imagePre012.ContainerConfig
		image.DockerVersion = imagePre012.DockerVersion
		image.Author = imagePre012.Author
		image.Config = imagePre012.Config
		image.Architecture = imagePre012.Architecture
		image.Size = imagePre012.Size
	}

	return &image, nil
}

// PushImageOptions represents options to use in the PushImage method.
//
// See https://goo.gl/zPtZaT for more details.
type PushImageOptions struct {
	// Name of the image
	Name string

	// Tag of the image
	Tag string

	// Registry server to push the image
	Registry string

	OutputStream  io.Writer `qs:"-"`
	RawJSONStream bool      `qs:"-"`
}

// PushImage pushes an image to a remote registry, logging progress to w.
//
// An empty instance of AuthConfiguration may be used for unauthenticated
// pushes.
//
// See https://goo.gl/zPtZaT for more details.
func (c *Client) PushImage(opts PushImageOptions, auth AuthConfiguration) error {
	if opts.Name == "" {
		return ErrNoSuchImage
	}
	headers, err := headersWithAuth(auth)
	if err != nil {
		return err
	}
	name := opts.Name
	opts.Name = ""
	path := "/images/" + name + "/push?" + queryString(&opts)
	return c.stream("POST", path, streamOptions{
		setRawTerminal: true,
		rawJSONStream:  opts.RawJSONStream,
		headers:        headers,
		stdout:         opts.OutputStream,
	})
}

// PullImageOptions present the set of options available for pulling an image
// from a registry.
//
// See https://goo.gl/iJkZjD for more details.
type PullImageOptions struct {
	Repository    string `qs:"fromImage"`
	Registry      string
	Tag           string
	OutputStream  io.Writer `qs:"-"`
	RawJSONStream bool      `qs:"-"`
}

// PullImage pulls an image from a remote registry, logging progress to
// opts.OutputStream.
//
// See https://goo.gl/iJkZjD for more details.
func (c *Client) PullImage(opts PullImageOptions, auth AuthConfiguration) error {
	if opts.Repository == "" {
		return ErrNoSuchImage
	}

	headers, err := headersWithAuth(auth)
	if err != nil {
		return err
	}
	return c.createImage(queryString(&opts), headers, nil, opts.OutputStream, opts.RawJSONStream)
}

func (c *Client) createImage(qs string, headers map[string]string, in io.Reader, w io.Writer, rawJSONStream bool) error {
	path := "/images/create?" + qs
	return c.stream("POST", path, streamOptions{
		setRawTerminal: true,
		rawJSONStream:  rawJSONStream,
		headers:        headers,
		in:             in,
		stdout:         w,
	})
}

// LoadImageOptions represents the options for LoadImage Docker API Call
//
// See https://goo.gl/JyClMX for more details.
type LoadImageOptions struct {
	InputStream io.Reader
}

// LoadImage imports a tarball docker image
//
// See https://goo.gl/JyClMX for more details.
func (c *Client) LoadImage(opts LoadImageOptions) error {
	return c.stream("POST", "/images/load", streamOptions{
		setRawTerminal: true,
		in:             opts.InputStream,
	})
}

// ExportImageOptions represent the options for ExportImage Docker API call.
//
// See https://goo.gl/le7vK8 for more details.
type ExportImageOptions struct {
	Name         string
	OutputStream io.Writer
}

// ExportImage exports an image (as a tar file) into the stream.
//
// See https://goo.gl/le7vK8 for more details.
func (c *Client) ExportImage(opts ExportImageOptions) error {
	return c.stream("GET", fmt.Sprintf("/images/%s/get", opts.Name), streamOptions{
		setRawTerminal: true,
		stdout:         opts.OutputStream,
	})
}

// ExportImagesOptions represent the options for ExportImages Docker API call
//
// See https://goo.gl/huC7HA for more details.
type ExportImagesOptions struct {
	Names        []string
	OutputStream io.Writer `qs:"-"`
}

// ExportImages exports one or more images (as a tar file) into the stream
//
// See https://goo.gl/huC7HA for more details.
func (c *Client) ExportImages(opts ExportImagesOptions) error {
	if opts.Names == nil || len(opts.Names) == 0 {
		return ErrMustSpecifyNames
	}
	return c.stream("GET", "/images/get?"+queryString(&opts), streamOptions{
		setRawTerminal: true,
		stdout:         opts.OutputStream,
	})
}

// ImportImageOptions present the set of informations available for importing
// an image from a source file or the stdin.
//
// See https://goo.gl/iJkZjD for more details.
type ImportImageOptions struct {
	Repository string `qs:"repo"`
	Source     string `qs:"fromSrc"`
	Tag        string `qs:"tag"`

	InputStream   io.Reader `qs:"-"`
	OutputStream  io.Writer `qs:"-"`
	RawJSONStream bool      `qs:"-"`
}

// ImportImage imports an image from a url, a file or stdin
//
// See https://goo.gl/iJkZjD for more details.
func (c *Client) ImportImage(opts ImportImageOptions) error {
	if opts.Repository == "" {
		return ErrNoSuchImage
	}
	if opts.Source != "-" {
		opts.InputStream = nil
	}
	if opts.Source != "-" && !isURL(opts.Source) {
		f, err := os.Open(opts.Source)
		if err != nil {
			return err
		}
		opts.InputStream = f
		opts.Source = "-"
	}
	return c.createImage(queryString(&opts), nil, opts.InputStream, opts.OutputStream, opts.RawJSONStream)
}

// BuildImageOptions present the set of informations available for building an
// image from a tarfile with a Dockerfile in it.
//
// For more details about the Docker building process, see
// http://goo.gl/tlPXPu.
type BuildImageOptions struct {
	Name                string             `qs:"t"`
	Dockerfile          string             `qs:"dockerfile"`
	NoCache             bool               `qs:"nocache"`
	SuppressOutput      bool               `qs:"q"`
	Pull                bool               `qs:"pull"`
	RmTmpContainer      bool               `qs:"rm"`
	ForceRmTmpContainer bool               `qs:"forcerm"`
	Memory              int64              `qs:"memory"`
	Memswap             int64              `qs:"memswap"`
	CPUShares           int64              `qs:"cpushares"`
	CPUQuota            int64              `qs:"cpuquota"`
	CPUPeriod           int64              `qs:"cpuperiod"`
	CPUSetCPUs          string             `qs:"cpusetcpus"`
	InputStream         io.Reader          `qs:"-"`
	OutputStream        io.Writer          `qs:"-"`
	RawJSONStream       bool               `qs:"-"`
	Remote              string             `qs:"remote"`
	Auth                AuthConfiguration  `qs:"-"` // for older docker X-Registry-Auth header
	AuthConfigs         AuthConfigurations `qs:"-"` // for newer docker X-Registry-Config header
	ContextDir          string             `qs:"-"`
	Ulimits             []ULimit           `qs:"-"`
	BuildArgs           []BuildArg         `qs:"-"`
}

// BuildArg represents arguments that can be passed to the image when building
// it from a Dockerfile.
//
// For more details about the Docker building process, see
// http://goo.gl/tlPXPu.
type BuildArg struct {
	Name  string `json:"Name,omitempty" yaml:"Name,omitempty"`
	Value string `json:"Value,omitempty" yaml:"Value,omitempty"`
}

// BuildImage builds an image from a tarball's url or a Dockerfile in the input
// stream.
//
// See https://goo.gl/xySxCe for more details.
func (c *Client) BuildImage(opts BuildImageOptions) error {
	if opts.OutputStream == nil {
		return ErrMissingOutputStream
	}
	headers, err := headersWithAuth(opts.Auth, c.versionedAuthConfigs(opts.AuthConfigs))
	if err != nil {
		return err
	}

	if opts.Remote != "" && opts.Name == "" {
		opts.Name = opts.Remote
	}
	if opts.InputStream != nil || opts.ContextDir != "" {
		headers["Content-Type"] = "application/tar"
	} else if opts.Remote == "" {
		return ErrMissingRepo
	}
	if opts.ContextDir != "" {
		if opts.InputStream != nil {
			return ErrMultipleContexts
		}
		var err error
		if opts.InputStream, err = createTarStream(opts.ContextDir, opts.Dockerfile); err != nil {
			return err
		}
	}

	qs := queryString(&opts)
	if len(opts.Ulimits) > 0 {
		if b, err := json.Marshal(opts.Ulimits); err == nil {
			item := url.Values(map[string][]string{})
			item.Add("ulimits", string(b))
			qs = fmt.Sprintf("%s&%s", qs, item.Encode())
		}
	}

	if len(opts.BuildArgs) > 0 {
		v := make(map[string]string)
		for _, arg := range opts.BuildArgs {
			v[arg.Name] = arg.Value
		}
		if b, err := json.Marshal(v); err == nil {
			item := url.Values(map[string][]string{})
			item.Add("buildargs", string(b))
			qs = fmt.Sprintf("%s&%s", qs, item.Encode())
		}
	}

	return c.stream("POST", fmt.Sprintf("/build?%s", qs), streamOptions{
		setRawTerminal: true,
		rawJSONStream:  opts.RawJSONStream,
		headers:        headers,
		in:             opts.InputStream,
		stdout:         opts.OutputStream,
	})
}

func (c *Client) versionedAuthConfigs(authConfigs AuthConfigurations) interface{} {
	if c.serverAPIVersion == nil {
		c.checkAPIVersion()
	}
	if c.serverAPIVersion != nil && c.serverAPIVersion.GreaterThanOrEqualTo(apiVersion119) {
		return AuthConfigurations119(authConfigs.Configs)
	}
	return authConfigs
}

// TagImageOptions present the set of options to tag an image.
//
// See https://goo.gl/98ZzkU for more details.
type TagImageOptions struct {
	Repo  string
	Tag   string
	Force bool
}

// TagImage adds a tag to the image identified by the given name.
//
// See https://goo.gl/98ZzkU for more details.
func (c *Client) TagImage(name string, opts TagImageOptions) error {
	if name == "" {
		return ErrNoSuchImage
	}
	resp, err := c.do("POST", fmt.Sprintf("/images/"+name+"/tag?%s",
		queryString(&opts)), doOptions{})

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrNoSuchImage
	}

	return err
}

func isURL(u string) bool {
	p, err := url.Parse(u)
	if err != nil {
		return false
	}
	return p.Scheme == "http" || p.Scheme == "https"
}

func headersWithAuth(auths ...interface{}) (map[string]string, error) {
	var headers = make(map[string]string)

	for _, auth := range auths {
		switch auth.(type) {
		case AuthConfiguration:
			var buf bytes.Buffer
			if err := json.NewEncoder(&buf).Encode(auth); err != nil {
				return nil, err
			}
			headers["X-Registry-Auth"] = base64.URLEncoding.EncodeToString(buf.Bytes())
		case AuthConfigurations, AuthConfigurations119:
			var buf bytes.Buffer
			if err := json.NewEncoder(&buf).Encode(auth); err != nil {
				return nil, err
			}
			headers["X-Registry-Config"] = base64.URLEncoding.EncodeToString(buf.Bytes())
		}
	}

	return headers, nil
}

// APIImageSearch reflect the result of a search on the Docker Hub.
//
// See https://goo.gl/AYjyrF for more details.
type APIImageSearch struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	IsOfficial  bool   `json:"is_official,omitempty" yaml:"is_official,omitempty"`
	IsAutomated bool   `json:"is_automated,omitempty" yaml:"is_automated,omitempty"`
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	StarCount   int    `json:"star_count,omitempty" yaml:"star_count,omitempty"`
}

// SearchImages search the docker hub with a specific given term.
//
// See https://goo.gl/AYjyrF for more details.
func (c *Client) SearchImages(term string) ([]APIImageSearch, error) {
	resp, err := c.do("GET", "/images/search?term="+term, doOptions{})
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var searchResult []APIImageSearch
	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return nil, err
	}
	return searchResult, nil
}

// SearchImagesEx search the docker hub with a specific given term and authentication.
//
// See https://goo.gl/AYjyrF for more details.
func (c *Client) SearchImagesEx(term string, auth AuthConfiguration) ([]APIImageSearch, error) {
	headers, err := headersWithAuth(auth)
	if err != nil {
		return nil, err
	}

	resp, err := c.do("GET", "/images/search?term="+term, doOptions{
		headers: headers,
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var searchResult []APIImageSearch
	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return nil, err
	}

	return searchResult, nil
}
