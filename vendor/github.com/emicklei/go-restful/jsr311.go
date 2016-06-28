package restful

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
)

// RouterJSR311 implements the flow for matching Requests to Routes (and consequently Resource Functions)
// as specified by the JSR311 http://jsr311.java.net/nonav/releases/1.1/spec/spec.html.
// RouterJSR311 implements the Router interface.
// Concept of locators is not implemented.
type RouterJSR311 struct{}

// SelectRoute is part of the Router interface and returns the best match
// for the WebService and its Route for the given Request.
func (r RouterJSR311) SelectRoute(
	webServices []*WebService,
	httpRequest *http.Request) (selectedService *WebService, selectedRoute *Route, err error) {

	// Identify the root resource class (WebService)
	dispatcher, finalMatch, err := r.detectDispatcher(httpRequest.URL.Path, webServices)
	if err != nil {
		return nil, nil, NewError(http.StatusNotFound, "")
	}
	// Obtain the set of candidate methods (Routes)
	routes := r.selectRoutes(dispatcher, finalMatch)
	if len(routes) == 0 {
		return dispatcher, nil, NewError(http.StatusNotFound, "404: Page Not Found")
	}

	// Identify the method (Route) that will handle the request
	route, ok := r.detectRoute(routes, httpRequest)
	return dispatcher, route, ok
}

// http://jsr311.java.net/nonav/releases/1.1/spec/spec3.html#x3-360003.7.2
func (r RouterJSR311) detectRoute(routes []Route, httpRequest *http.Request) (*Route, error) {
	// http method
	methodOk := []Route{}
	for _, each := range routes {
		if httpRequest.Method == each.Method {
			methodOk = append(methodOk, each)
		}
	}
	if len(methodOk) == 0 {
		if trace {
			traceLogger.Printf("no Route found (in %d routes) that matches HTTP method %s\n", len(routes), httpRequest.Method)
		}
		return nil, NewError(http.StatusMethodNotAllowed, "405: Method Not Allowed")
	}
	inputMediaOk := methodOk

	// content-type
	contentType := httpRequest.Header.Get(HEADER_ContentType)
	inputMediaOk = []Route{}
	for _, each := range methodOk {
		if each.matchesContentType(contentType) {
			inputMediaOk = append(inputMediaOk, each)
		}
	}
	if len(inputMediaOk) == 0 {
		if trace {
			traceLogger.Printf("no Route found (from %d) that matches HTTP Content-Type: %s\n", len(methodOk), contentType)
		}
		return nil, NewError(http.StatusUnsupportedMediaType, "415: Unsupported Media Type")
	}

	// accept
	outputMediaOk := []Route{}
	accept := httpRequest.Header.Get(HEADER_Accept)
	if len(accept) == 0 {
		accept = "*/*"
	}
	for _, each := range inputMediaOk {
		if each.matchesAccept(accept) {
			outputMediaOk = append(outputMediaOk, each)
		}
	}
	if len(outputMediaOk) == 0 {
		if trace {
			traceLogger.Printf("no Route found (from %d) that matches HTTP Accept: %s\n", len(inputMediaOk), accept)
		}
		return nil, NewError(http.StatusNotAcceptable, "406: Not Acceptable")
	}
	// return r.bestMatchByMedia(outputMediaOk, contentType, accept), nil
	return &outputMediaOk[0], nil
}

// http://jsr311.java.net/nonav/releases/1.1/spec/spec3.html#x3-360003.7.2
// n/m > n/* > */*
func (r RouterJSR311) bestMatchByMedia(routes []Route, contentType string, accept string) *Route {
	// TODO
	return &routes[0]
}

// http://jsr311.java.net/nonav/releases/1.1/spec/spec3.html#x3-360003.7.2  (step 2)
func (r RouterJSR311) selectRoutes(dispatcher *WebService, pathRemainder string) []Route {
	filtered := &sortableRouteCandidates{}
	for _, each := range dispatcher.Routes() {
		pathExpr := each.pathExpr
		matches := pathExpr.Matcher.FindStringSubmatch(pathRemainder)
		if matches != nil {
			lastMatch := matches[len(matches)-1]
			if len(lastMatch) == 0 || lastMatch == "/" { // do not include if value is neither empty nor ‘/’.
				filtered.candidates = append(filtered.candidates,
					routeCandidate{each, len(matches) - 1, pathExpr.LiteralCount, pathExpr.VarCount})
			}
		}
	}
	if len(filtered.candidates) == 0 {
		if trace {
			traceLogger.Printf("WebService on path %s has no routes that match URL path remainder:%s\n", dispatcher.rootPath, pathRemainder)
		}
		return []Route{}
	}
	sort.Sort(sort.Reverse(filtered))

	// select other routes from candidates whoes expression matches rmatch
	matchingRoutes := []Route{filtered.candidates[0].route}
	for c := 1; c < len(filtered.candidates); c++ {
		each := filtered.candidates[c]
		if each.route.pathExpr.Matcher.MatchString(pathRemainder) {
			matchingRoutes = append(matchingRoutes, each.route)
		}
	}
	return matchingRoutes
}

// http://jsr311.java.net/nonav/releases/1.1/spec/spec3.html#x3-360003.7.2 (step 1)
func (r RouterJSR311) detectDispatcher(requestPath string, dispatchers []*WebService) (*WebService, string, error) {
	filtered := &sortableDispatcherCandidates{}
	for _, each := range dispatchers {
		matches := each.pathExpr.Matcher.FindStringSubmatch(requestPath)
		if matches != nil {
			filtered.candidates = append(filtered.candidates,
				dispatcherCandidate{each, matches[len(matches)-1], len(matches), each.pathExpr.LiteralCount, each.pathExpr.VarCount})
		}
	}
	if len(filtered.candidates) == 0 {
		if trace {
			traceLogger.Printf("no WebService was found to match URL path:%s\n", requestPath)
		}
		return nil, "", errors.New("not found")
	}
	sort.Sort(sort.Reverse(filtered))
	return filtered.candidates[0].dispatcher, filtered.candidates[0].finalMatch, nil
}

// Types and functions to support the sorting of Routes

type routeCandidate struct {
	route           Route
	matchesCount    int // the number of capturing groups
	literalCount    int // the number of literal characters (means those not resulting from template variable substitution)
	nonDefaultCount int // the number of capturing groups with non-default regular expressions (i.e. not ‘([^  /]+?)’)
}

func (r routeCandidate) expressionToMatch() string {
	return r.route.pathExpr.Source
}

func (r routeCandidate) String() string {
	return fmt.Sprintf("(m=%d,l=%d,n=%d)", r.matchesCount, r.literalCount, r.nonDefaultCount)
}

type sortableRouteCandidates struct {
	candidates []routeCandidate
}

func (rcs *sortableRouteCandidates) Len() int {
	return len(rcs.candidates)
}
func (rcs *sortableRouteCandidates) Swap(i, j int) {
	rcs.candidates[i], rcs.candidates[j] = rcs.candidates[j], rcs.candidates[i]
}
func (rcs *sortableRouteCandidates) Less(i, j int) bool {
	ci := rcs.candidates[i]
	cj := rcs.candidates[j]
	// primary key
	if ci.literalCount < cj.literalCount {
		return true
	}
	if ci.literalCount > cj.literalCount {
		return false
	}
	// secundary key
	if ci.matchesCount < cj.matchesCount {
		return true
	}
	if ci.matchesCount > cj.matchesCount {
		return false
	}
	// tertiary key
	if ci.nonDefaultCount < cj.nonDefaultCount {
		return true
	}
	if ci.nonDefaultCount > cj.nonDefaultCount {
		return false
	}
	// quaternary key ("source" is interpreted as Path)
	return ci.route.Path < cj.route.Path
}

// Types and functions to support the sorting of Dispatchers

type dispatcherCandidate struct {
	dispatcher      *WebService
	finalMatch      string
	matchesCount    int // the number of capturing groups
	literalCount    int // the number of literal characters (means those not resulting from template variable substitution)
	nonDefaultCount int // the number of capturing groups with non-default regular expressions (i.e. not ‘([^  /]+?)’)
}
type sortableDispatcherCandidates struct {
	candidates []dispatcherCandidate
}

func (dc *sortableDispatcherCandidates) Len() int {
	return len(dc.candidates)
}
func (dc *sortableDispatcherCandidates) Swap(i, j int) {
	dc.candidates[i], dc.candidates[j] = dc.candidates[j], dc.candidates[i]
}
func (dc *sortableDispatcherCandidates) Less(i, j int) bool {
	ci := dc.candidates[i]
	cj := dc.candidates[j]
	// primary key
	if ci.matchesCount < cj.matchesCount {
		return true
	}
	if ci.matchesCount > cj.matchesCount {
		return false
	}
	// secundary key
	if ci.literalCount < cj.literalCount {
		return true
	}
	if ci.literalCount > cj.literalCount {
		return false
	}
	// tertiary key
	return ci.nonDefaultCount < cj.nonDefaultCount
}
