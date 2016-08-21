//+build prod

//
// Copyright (c) 2014 The godbc Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package godbc

type InvariantSimpleTester interface {
	Invariant() bool
}

type InvariantTester interface {
	InvariantSimpleTester
	String() string
}

func Require(b bool, message ...interface{}) {
}

func Ensure(b bool, message ...interface{}) {
}

func Check(b bool, message ...interface{}) {
}

func InvariantSimple(obj InvariantSimpleTester, message ...interface{}) {
}

func Invariant(obj InvariantTester, message ...interface{}) {
}
