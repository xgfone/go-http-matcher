// Copyright 2024 xgfone
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
	"log/slog"
	"net/netip"

	"github.com/xgfone/go-toolkit/netx"
)

func extracthost(addr string) string {
	host, _ := netx.SplitHostPort(addr)
	return host
}

func parseip(prefix, ip string) netip.Addr {
	addr, err := netip.ParseAddr(extracthost(ip))
	if err != nil {
		slog.Error(prefix+": fail to parse ip", "ip", ip, "err", err)
	}
	return addr
}
