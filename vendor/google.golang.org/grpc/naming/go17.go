// +build go1.7, !go1.8

/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package naming

import (
	"net"

	"golang.org/x/net/context"
)

var (
	lookupHost = func(ctx context.Context, host string) ([]string, error) { return net.LookupHost(host) }
	lookupSRV  = func(ctx context.Context, service, proto, name string) (string, []*net.SRV, error) {
		return net.LookupSRV(service, proto, name)
	}
)
