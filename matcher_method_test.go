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

func TestMethod(t *testing.T) {
	if m := Method(); m != nil {
		t.Errorf("expect nil, but got matcher '%s'", m.String())
	}

	req := &http.Request{Method: http.MethodGet}
	if !Method("GET").Match(req) {
		t.Errorf("expect match '%s', but got not", req.Method)
	}

	if !Method("get", "post").Match(req) {
		t.Errorf("expect match '%s', but got not", req.Method)
	}

	if Method("POST").Match(req) {
		t.Errorf("unexpect match '%s', but got matched", req.Method)
	}

	if Method("put", "post").Match(req) {
		t.Errorf("unexpect match '%s', but got matched", req.Method)
	}
}
