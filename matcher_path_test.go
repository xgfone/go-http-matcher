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

func TestPath(t *testing.T) {
	if m := Path(); m != nil {
		t.Errorf("expect nil, but got matcher '%s'", m.String())
	}

	req := &http.Request{URL: &url.URL{Path: "/path/to"}}
	if !Path("/path/to").Match(req) {
		t.Errorf("expect match '%s', but got not", req.URL.Path)
	}

	if !Path("/path/to", "/path/other").Match(req) {
		t.Errorf("expect match '%s', but got not", req.URL.Path)
	}

	if Path("/path").Match(req) {
		t.Errorf("unexpect match '%s', but got matched", req.URL.Path)
	}

	if Path("/", "/path").Match(req) {
		t.Errorf("unexpect match '%s', but got matched", req.URL.Path)
	}
}

func TestPathPrefix(t *testing.T) {
	if m := PathPrefix(); m != nil {
		t.Errorf("expect nil, but got matcher '%s'", m.String())
	}

	req := &http.Request{URL: &url.URL{Path: "/"}}
	if !PathPrefix("/").Match(req) {
		t.Errorf("expect match '%s', but got not", req.URL.Path)
	}

	req.URL.Path = "/path/"
	if !PathPrefix("/path").Match(req) {
		t.Errorf("expect match '%s', but got not", req.URL.Path)
	}

	req.URL.Path = "/pathto"
	if PathPrefix("/path").Match(req) {
		t.Errorf("unexpect match '%s', but got true", req.URL.Path)
	} else if PathPrefix("/path/", "/to").Match(req) {
		t.Errorf("unexpect match '%s', but got true", req.URL.Path)
	}

	req.URL.Path = "/path/to"
	if !PathPrefix("/").Match(req) {
		t.Errorf("expect match '%s', but got not", req.URL.Path)
	}

	if !PathPrefix("/path/to").Match(req) {
		t.Errorf("expect match '%s', but got not", req.URL.Path)
	}

	if !PathPrefix("/path", "/other").Match(req) {
		t.Errorf("expect match '%s', but got not", req.URL.Path)
	}

	if PathPrefix("/path/other").Match(req) {
		t.Errorf("unexpect match '%s', but got matched", req.URL.Path)
	}

	if PathPrefix("/path/other", "/other").Match(req) {
		t.Errorf("unexpect match '%s', but got matched", req.URL.Path)
	}
}
