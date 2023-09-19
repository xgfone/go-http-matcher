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
	"fmt"
	"net/http"
	"strings"

	"github.com/xgfone/go-defaults/assists"
)

// GetHost is used to customize the host, which must be lower case.
var GetHost = func(r *http.Request) string {
	if r.TLS != nil && r.TLS.ServerName != "" {
		return strings.ToLower(r.TLS.ServerName)
	}
	return strings.ToLower(assists.TrimPort(r.Host))
}

// Host returns a new matcher that checks whether the host is one
// of the specified hosts, which supports the exact or wildcard domain,
// such as "www.example.com" or "*.example.com".
//
// If hosts is empty, return nil instead of an error.
func Host(hosts ...string) Matcher {
	switch _len := len(hosts); _len {
	case 0:
		return nil

	case 1:
		host := strings.ToLower(hosts[0])
		desc := fmt.Sprintf("Host(`%s`)", host)
		match := _buildHostMatcher(host)
		return New(PriorityHost*len(host), desc, func(r *http.Request) bool {
			return match(GetHost(r))
		})
	}

	var maxlen int
	matches := make([]func(string) bool, len(hosts))
	for i, host := range hosts {
		host = strings.ToLower(host)
		matches[i] = _buildHostMatcher(host)
		if _len := len(host); _len > maxlen {
			maxlen = _len
		}
	}

	desc := fmt.Sprintf("Host(`%s`)", strings.Join(hosts, "`,`"))
	return New(PriorityHost*maxlen, desc, func(r *http.Request) bool {
		host := GetHost(r)
		for _, match := range matches {
			if match(host) {
				return true
			}
		}
		return false
	})
}

func _buildHostMatcher(host string) func(string) bool {
	switch {
	case host == "*":
		return func(string) bool { return true }

	case host[0] == '*':
		host = host[1:]
		return func(s string) bool { return strings.HasSuffix(s, host) }

	default:
		return func(s string) bool { return s == host }
	}
}
