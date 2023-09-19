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
	"net/netip"
	"strings"

	"github.com/xgfone/go-defaults"
)

// GetClientIP is used to customize the client ip.
var GetClientIP = func(r *http.Request) netip.Addr {
	return defaults.GetClientIP(r.Context(), r)
}

// ClientIP returns a new matcher that checks whether the client ip,
// that's remote address ip, is one of the specified ips.
//
// If ips is empty, return (nil, nil) instead of an error.
func ClientIP(ips ...string) (Matcher, error) {
	if len(ips) == 0 {
		return nil, nil
	}

	checker, err := newIPCheckers(ips...)
	if err != nil {
		return nil, err
	}

	desc := fmt.Sprintf("ClientIp(`%s`)", strings.Join(ips, "`,`"))
	return New(PriorityClientIP, desc, func(r *http.Request) bool {
		return checker.ContainsAddr(GetClientIP(r))
	}), nil
}
