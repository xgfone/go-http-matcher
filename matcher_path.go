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

// GetPath is used to customize the path.
var GetPath = func(r *http.Request) (path string) {
	if path = strings.TrimRight(r.URL.Path, "/"); path == "" {
		path = "/"
	}
	return
}

// Path returns a new matcher that checks whether the path is one
// of the specified paths.
//
// If paths is empty, return nil instead of an error.
func Path(paths ...string) Matcher {
	switch _len := len(paths); _len {
	case 0:
		return nil

	case 1:
		path := fixPath(paths[0])
		desc := fmt.Sprintf("Path(`%s`)", path)
		return New(PriorityPath*len(path), desc, func(r *http.Request) bool {
			return GetPath(r) == path
		})
	}

	var maxlen int
	mpaths := make(exactFullMatches, len(paths))
	for i, path := range paths {
		mpaths[i] = fixPath(path)
		if _len := len(mpaths[i]); _len > maxlen {
			maxlen = _len
		}
	}

	desc := fmt.Sprintf("Path(`%s`)", strings.Join(mpaths, "`,`"))
	return New(PriorityPath*maxlen, desc, func(r *http.Request) bool {
		return mpaths.Match(GetPath(r))
	})
}

// PathPrefix returns a new matcher that checks whether the path has the prefix
// that is in the specified path prefixes.
//
// If pathPrefixes is empty, return (nil, nil) instead of an error.
func PathPrefix(pathPrefixes ...string) Matcher {
	switch _len := len(pathPrefixes); _len {
	case 0:
		return nil

	case 1:
		prefix := fixPath(pathPrefixes[0])
		desc := fmt.Sprintf("PathPrefix(`%s`)", prefix)
		if prefix == "/" {
			return New(PriorityPathPrefix, desc, AlwaysTrue)
		}

		return New(PriorityPathPrefix*len(prefix), desc, func(r *http.Request) bool {
			return matchpathprefix(GetPath(r), prefix)
		})
	}

	var maxlen int
	prefixs := make(exactPrefixMatches, len(pathPrefixes))
	for i, prefix := range pathPrefixes {
		prefixs[i] = fixPath(prefix)
		if _len := len(prefixs[i]); _len > maxlen {
			maxlen = _len
		}
	}

	desc := fmt.Sprintf("PathPrefix(`%s`)", strings.Join(prefixs, "`,`"))
	return New(PriorityPathPrefix*maxlen, desc, func(r *http.Request) bool {
		return prefixs.Match(GetPath(r), matchpathprefix)
	})
}

func matchpathprefix(path, prefix string) bool {
	if !strings.HasPrefix(path, prefix) {
		return false
	}

	mlen := len(prefix)
	return mlen == len(path) || path[mlen] == '/'
}

func fixPath(path string) string {
	if path == "/" {
		return path
	} else if path = strings.TrimRight(path, "/"); path != "" {
		return path
	}
	return "/"
}
