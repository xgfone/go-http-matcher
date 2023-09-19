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
)

// Method returns a new matcher that checks whether the method is one
// of the specified methods.
//
// If methods is empty, return nil instead of an error.
func Method(methods ...string) Matcher {
	switch _len := len(methods); _len {
	case 0:
		return nil

	case 1:
		method := strings.ToUpper(methods[0])
		desc := fmt.Sprintf("Method(`%s`)", method)
		return New(PriorityMethod, desc, func(r *http.Request) bool {
			return r.Method == method
		})
	}

	_methods := make(exactFullMatches, len(methods))
	for i, method := range methods {
		_methods[i] = strings.ToUpper(method)
	}

	desc := fmt.Sprintf("Method(`%s`)", strings.Join(_methods, "`,`"))
	return New(PriorityMethod, desc, func(r *http.Request) bool {
		return _methods.Match(r.Method)
	})
}
