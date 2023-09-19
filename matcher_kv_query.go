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

// Query returns a new matcher that checks whether the query
// has the specified key-value argument.
//
// Both key and value are the exact match.
// If value is empty, it matches all the requests that has the query key
// and ignores the query value.
//
// If key is empty, return nil instead of an error.
func Query(key, value string) Matcher {
	if key == "" {
		return nil
	}

	desc := kvdesc("Query", key, value)
	return New(PriorityQuery, desc, func(r *http.Request) bool {
		values, ok := r.URL.Query()[key]
		if !ok || (value != "" && !contains(values, value)) {
			return false
		}
		return true
	})
}

// Querym returns a new matcher that checks whether the query
// has all the specified key-value arguments.
//
// Both keys and values are the exact match.
// If value is empty, it matches all the requests that has the query key
// and ignores the query value.
//
// If querym is empty, returnnil instead of an error.
func Querym(querym map[string]string) Matcher {
	if len(querym) == 0 {
		return nil
	}

	desc := kvsdesc("Query", querym)
	return New(PriorityQuery*len(querym), desc, func(r *http.Request) bool {
		query := r.URL.Query()
		for key, value := range querym {
			values, ok := query[key]
			if !ok || (value != "" && !contains(values, value)) {
				return false
			}
		}
		return true
	})
}
