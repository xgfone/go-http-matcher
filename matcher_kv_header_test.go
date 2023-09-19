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
	"testing"
)

func TestHeader(t *testing.T) {
	if m := Header("", ""); m != nil {
		t.Errorf("expect nil, but got matcher '%s'", m.String())
	}

	req := &http.Request{Header: make(http.Header, 2)}
	req.Header.Set("k1", "v1")
	req.Header.Set("k2", "v2")

	if m := Header("K1", "v1"); !m.Match(req) {
		t.Errorf("expect match '%v' with '%s', but got not", req.Header, m.String())
	}

	if m := Header("K1", ""); !m.Match(req) {
		t.Errorf("expect match '%v' with '%s', but got not", req.Header, m.String())
	}

	if Header("K3", "v3").Match(req) {
		t.Errorf("unexpect match '%v', but got matched", req.Header)
	}
}

func TestHeaderm(t *testing.T) {
	if m := Headerm(nil); m != nil {
		t.Errorf("expect nil, but got matcher '%s'", m.String())
	}

	req := &http.Request{Header: make(http.Header, 2)}
	req.Header.Set("k1", "v1")
	req.Header.Set("k2", "v2")

	if m := Headerm(map[string]string{"K1": "v1"}); !m.Match(req) {
		t.Errorf("expect match '%v' with '%s', but got not", req.Header, m.String())
	}

	if m := Headerm(map[string]string{"K1": ""}); !m.Match(req) {
		t.Errorf("expect match '%v' with '%s', but got not", req.Header, m.String())
	}

	if m := Headerm(map[string]string{"K1": "", "k3": ""}); m.Match(req) {
		t.Errorf("unexpect match '%v' with '%s', but got not", req.Header, m.String())
	}

	if Headerm(map[string]string{"K3": "v3"}).Match(req) {
		t.Errorf("unexpect match '%v', but got matched", req.Header)
	}
}
