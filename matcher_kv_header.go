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

import "net/http"

// Header returns a new matcher that checks whether the headers
// has the specified key-value argument.
//
// The key is the case-insensitive match, but, the value is the exact match.
// If value is empty, it matches all the requests that has the header key
// and ignores the header value.
//
// If key is empty, return nil instead of an error.
func Header(key, value string) Matcher {
	if key == "" {
		return nil
	}

	key = http.CanonicalHeaderKey(key)
	desc := kvdesc("Header", key, value)
	return New(PriorityHeader, desc, func(r *http.Request) bool {
		values, ok := r.Header[key]
		if !ok || (value != "" && !contains(values, value)) {
			return false
		}
		return true
	})
}

// Headerm returns a new matcher that checks whether the headers
// has all the specified key-value arguments.
//
// The keys are the case-insensitive match, but, the values are the exact match.
// If value is empty, it matches all the requests that has the header key
// and ignores the header value.
//
// If headerm is empty, return nil instead of an error.
func Headerm(headerm map[string]string) Matcher {
	if len(headerm) == 0 {
		return nil
	}

	headers := make(map[string]string, len(headerm))
	for key, value := range headerm {
		headers[http.CanonicalHeaderKey(key)] = value
	}

	desc := kvsdesc("Header", headers)
	return New(PriorityHeader*len(headers), desc, func(r *http.Request) bool {
		header := r.Header
		for key, value := range headers {
			values, ok := header[key]
			if !ok || (value != "" && !contains(values, value)) {
				return false
			}
		}
		return true
	})
}
