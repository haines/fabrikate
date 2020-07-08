// adapted from https://github.com/helm/helm/blob/bf9c64f48ee8f799c89de6cfbd44302f798b3531/internal/urlutil/urlutil.go

/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"net/url"
	"path/filepath"
)

// EqualURLs normalizes two URLs and then compares for equality.
func EqualURLs(a, b string) bool {
	au, err := url.Parse(a)
	if err != nil {
		a = filepath.Clean(a)
		b = filepath.Clean(b)
		// If urls are paths, return true only if they are an exact match
		return a == b
	}
	bu, err := url.Parse(b)
	if err != nil {
		return false
	}

	for _, u := range []*url.URL{au, bu} {
		if u.Path == "" {
			u.Path = "/"
		}
		u.Path = filepath.Clean(u.Path)
	}
	return au.String() == bu.String()
}
