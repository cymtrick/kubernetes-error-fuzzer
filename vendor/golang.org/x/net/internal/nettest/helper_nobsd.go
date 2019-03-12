// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build aix linux solaris

package nettest

func supportsIPv6MulticastDeliveryOnLoopback() bool {
	return true
}
