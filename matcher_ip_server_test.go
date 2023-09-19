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
	"context"
	"net"
	"net/http"
	"testing"
)

func TestServerIP(t *testing.T) {
	if m, err := ServerIP(); err != nil {
		t.Error(err)
	} else if m != nil {
		t.Errorf("expect nil, but got matcher '%s'", m.String())
	}

	if m, err := ServerIP("localhost"); err == nil {
		t.Errorf("execpt an error, but got a matcher '%s'", m.String())
	}

	req := new(http.Request)
	addr := &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 80}
	req = req.WithContext(context.WithValue(req.Context(), http.LocalAddrContextKey, addr))

	if m, err := ServerIP("127.0.0.1"); err != nil {
		t.Error(err)
	} else if !m.Match(req) {
		t.Errorf("expect match '%s', but got not", addr.IP.String())
	}

	m, err := ServerIP("127.0.0.0/8", "192.168.0.0/16")
	if err != nil {
		t.Error(err)
	} else if !m.Match(req) {
		t.Errorf("expect match '%s', but got not", addr.IP.String())
	}

	addr = &net.TCPAddr{IP: net.ParseIP("1.2.3.4"), Port: 7890}
	req = req.WithContext(context.WithValue(req.Context(), http.LocalAddrContextKey, addr))
	if m.Match(req) {
		t.Errorf("unexpect match '%s', but got matched", addr.IP.String())
	}
}
