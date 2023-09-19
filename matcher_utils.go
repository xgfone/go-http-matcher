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
	"bytes"
	"fmt"
	"net/netip"
	"strings"
)

const (
	PriorityQuery      = 1
	PriorityHeader     = 4
	PriorityClientIP   = 20
	PriorityServerIP   = 20
	PriorityMethod     = 40
	PriorityPathPrefix = 50
	PriorityPath       = 500
	PriorityHost       = 5000
)

func contains(vs []string, s string) bool {
	for i := range vs {
		if s == vs[i] {
			return true
		}
	}
	return false
}

type (
	exactFullMatches   []string
	exactPrefixMatches []string
)

func (vs exactFullMatches) Match(value string) bool {
	for _, s := range vs {
		if value == s {
			return true
		}
	}
	return false
}

func (vs exactPrefixMatches) Match(value string, match func(path, prefix string) bool) bool {
	if match == nil {
		match = strings.HasPrefix
	}

	for _, s := range vs {
		if match(value, s) {
			return true
		}
	}
	return false
}

type ipcheckers []netip.Prefix

// ContainsAddr reports whether the checkers contains the ip addr.
func (cs ipcheckers) ContainsAddr(ip netip.Addr) bool {
	for _, c := range cs {
		if c.Contains(ip) {
			return true
		}
	}
	return false
}

func newIPCheckers(ips ...string) (cs ipcheckers, err error) {
	cs = make([]netip.Prefix, len(ips))
	for i, ip := range ips {
		if strings.IndexByte(ip, '/') == -1 {
			if strings.IndexByte(ip, '.') == -1 {
				ip += "/128"
			} else {
				ip += "/32"
			}
		}

		cs[i], err = netip.ParsePrefix(ip)
		if err != nil {
			return
		}
	}
	return
}

func kvsdesc(name string, kvs map[string]string) string {
	buf := bytes.NewBuffer(make([]byte, 0, 18*len(kvs)+2))
	buf.WriteByte('(')
	var i int
	for k, v := range kvs {
		if i > 0 {
			buf.WriteString(" && ")
		}

		if v == "" {
			fmt.Fprintf(buf, "%s(`%s`)", name, k)
		} else {
			fmt.Fprintf(buf, "%s(`%s`,`%s`)", name, k, v)
		}

		i++
	}
	buf.WriteByte(')')
	return buf.String()
}

func kvdesc(name, key, value string) string {
	if value == "" {
		return fmt.Sprintf("%s(`%s`)", name, key)
	}
	return fmt.Sprintf("%s(`%s`,`%s`)", name, key, value)
}
