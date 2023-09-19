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
	"net/url"
	"testing"
)

func TestQuery(t *testing.T) {
	if m := Query("", ""); m != nil {
		t.Errorf("expect nil, but got matcher '%s'", m.String())
	}

	req := &http.Request{URL: &url.URL{RawQuery: "k1=v1&k2=v2"}}
	if m := Query("k1", "v1"); !m.Match(req) {
		t.Errorf("expect match '%s' with '%s', but got not", req.URL.RawQuery, m.String())
	}

	if m := Query("k1", ""); !m.Match(req) {
		t.Errorf("expect match '%s' with '%s', but got not", req.URL.RawQuery, m.String())
	}

	if Query("k3", "v3").Match(req) {
		t.Errorf("unexpect match '%s', but got matched", req.URL.RawQuery)
	}
}

func TestQuerym(t *testing.T) {
	if m := Querym(nil); m != nil {
		t.Errorf("expect nil, but got matcher '%s'", m.String())
	}

	req := &http.Request{URL: &url.URL{RawQuery: "k1=v1&k2=v2"}}
	if m := Querym(map[string]string{"k1": "v1"}); !m.Match(req) {
		t.Errorf("expect match '%s' with '%s', but got not", req.URL.RawQuery, m.String())
	}

	if m := Querym(map[string]string{"k1": ""}); !m.Match(req) {
		t.Errorf("expect match '%s' with '%s', but got not", req.URL.RawQuery, m.String())
	}

	if m := Querym(map[string]string{"k1": "", "k3": ""}); m.Match(req) {
		t.Errorf("unexpect match '%s' with '%s', but got not", req.URL.RawQuery, m.String())
	}

	if Querym(map[string]string{"k3": "v3"}).Match(req) {
		t.Errorf("unexpect match '%s', but got matched", req.URL.RawQuery)
	}
}
