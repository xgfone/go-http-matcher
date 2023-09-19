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

package matcher

import (
	"net/http"
	"sort"
	"strings"
)

// RemoveParentheses tries to remove the leading '(' and trailling ')'.
//
// If failing, return the original s.
func RemoveParentheses(s string) string {
	if _len := len(s) - 1; _len > 0 && s[0] == '(' && s[_len] == ')' {
		return s[1:_len]
	}
	return s
}

// Sort sorts the matchers by the priority from high to low.
func Sort(ms []Matcher) {
	if len(ms) > 0 {
		sort.Slice(ms, func(i, j int) bool {
			return ms[j].Priority() < ms[i].Priority()
		})
	}
}

type matchers struct {
	ms []Matcher

	match func(*http.Request, []Matcher) bool
	desc  string
	prio  int
}

func (m matchers) String() string             { return m.desc }
func (m matchers) Priority() int              { return m.prio }
func (m matchers) Match(r *http.Request) bool { return m.match(r, m.ms) }
func (m matchers) Matchers() []Matcher        { return m.ms }

// And returns a new matcher based on AND from a set of matchers,
// which checks whether all the matchers match the request.
//
// If no matchers, the returned mathcer does not match any request.
func And(ms ...Matcher) Matcher {
	switch _len := len(ms); _len {
	case 0:
		return New(0, "", AlwaysFalse)

	case 1:
		return ms[0]
	}

	Sort(ms)
	prio := andprio(ms)
	desc := formatMatchers(" && ", ms)
	return matchers{ms: ms, match: andmatch, desc: desc, prio: prio}
}

func andprio(ms []Matcher) (priority int) {
	for i, _len := 0, len(ms); i < _len; i++ {
		priority += ms[i].Priority()
	}
	return
}

func andmatch(r *http.Request, ms []Matcher) bool {
	_len := len(ms)
	if _len == 0 {
		return false
	}

	for i := 0; i < _len; i++ {
		if !ms[i].Match(r) {
			return false
		}
	}
	return true
}

// Or returns a new matcher based on OR from a set of matchers,
// which checks whether any matcher matches the request.
//
// If no matchers, the returned mathcer will match any request.
func Or(ms ...Matcher) Matcher {
	switch _len := len(ms); _len {
	case 0:
		return New(0, "", AlwaysTrue)

	case 1:
		return ms[0]
	}

	Sort(ms)
	prio := orprio(ms)
	desc := formatMatchers(" || ", ms)
	return matchers{ms: ms, match: ormatch, desc: desc, prio: prio}
}

func orprio(ms []Matcher) (priority int) {
	for _, m := range ms {
		if p := m.Priority(); p > priority {
			priority = p
		}
	}
	return
}

func ormatch(r *http.Request, ms []Matcher) bool {
	_len := len(ms)
	if _len == 0 {
		return true
	}

	for i := 0; i < _len; i++ {
		if ms[i].Match(r) {
			return true
		}
	}
	return false
}

func formatMatchers(sep string, matchers []Matcher) string {
	switch len(matchers) {
	case 0:
		return ""
	case 1:
		return matchers[0].String()
	}

	var b strings.Builder
	b.Grow(64)

	b.WriteByte('(')
	for i, matcher := range matchers {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(matcher.String())
	}
	b.WriteByte(')')

	return b.String()
}
