/*
Copyright 2016 The Kubernetes Authors.
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
package kubediscovery

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type Routes []Route

var routes = Routes{
	Route{
		"ClusterInfoIndex",
		"GET",
		"/cluster-info/v1/",
		NewClusterInfoHandler(),
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
