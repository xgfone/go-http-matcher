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
	"log/slog"
	"net"
	"net/http"
	"net/netip"
	"strings"
)

// GetServerIP is used to customize the server ip.
var GetServerIP = func(r *http.Request) netip.Addr {
	addr := r.Context().Value(http.LocalAddrContextKey).(net.Addr)
	switch v := addr.(type) {
	case *net.TCPAddr:
		return ip2addr(v.IP)

	case *net.UDPAddr:
		return ip2addr(v.IP)

	default:
		host := extracthost(v.String())
		addr, err := netip.ParseAddr(host)
		if err != nil {
			slog.Error("matcher.GetServerIP: fail to parse ip", "ip", host, "err", err)
		}
		return addr
	}
}

// ServerIP returns a new matcher that checks whether the server ip,
// that's local address ip, is one of the specified ips.
//
// If ips is empty, return (nil, nil) instead of an error.
func ServerIP(ips ...string) (Matcher, error) {
	if len(ips) == 0 {
		return nil, nil
	}

	checker, err := newIPCheckers(ips...)
	if err != nil {
		return nil, err
	}

	desc := fmt.Sprintf("ServerIp(`%s`)", strings.Join(ips, "`,`"))
	return New(PriorityServerIP, desc, func(r *http.Request) bool {
		return checker.ContainsAddr(GetServerIP(r))
	}), nil
}

func ip2addr(ip net.IP) (addr netip.Addr) {
	switch len(ip) {
	case net.IPv4len:
		addr = netip.AddrFrom4([4]byte(ip))
	case net.IPv6len:
		if ipv4 := ip.To4(); ipv4 != nil {
			addr = netip.AddrFrom4([4]byte(ipv4))
		} else {
			addr = netip.AddrFrom16([16]byte(ip))
		}
	default:
		slog.Warn("ip is not an ipv4 or ipv6", "ip", ip.String())
	}
	return
}
