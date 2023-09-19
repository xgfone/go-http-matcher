// Copyright 2023 xgfone
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

// Package matcher provides a http request matcher and some implementations.
package matcher

import (
	"fmt"
	"net/http"
)

var (
	// AlwaysTrue is a match function that always reutrns true.
	AlwaysTrue = MatchFunc(func(*http.Request) bool { return true })

	// AlwaysFalse is a match function that always reutrns false.
	AlwaysFalse = MatchFunc(func(*http.Request) bool { return false })
)

// Matcher is used to match a http request.
type Matcher interface {
	Match(*http.Request) bool
	Priority() int
	fmt.Stringer
}

// MatchFunc is a match function.
type MatchFunc func(r *http.Request) bool

// Match implements the interface Matcher#Match.
func (f MatchFunc) Match(r *http.Request) bool { return f(r) }

type matcher struct {
	MatchFunc
	desc string
	prio int
}

func (m matcher) String() string { return m.desc }
func (m matcher) Priority() int  { return m.prio }

// New returns a request matcher.
func New(prio int, desc string, match MatchFunc) Matcher {
	return matcher{prio: prio, desc: desc, MatchFunc: match}
}
