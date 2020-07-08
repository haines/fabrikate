// adapted from https://github.com/helm/helm/blob/bf9c64f48ee8f799c89de6cfbd44302f798b3531/internal/urlutil/urlutil_test.go

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

import "testing"

func TestEqualURLs(t *testing.T) {
	for _, tt := range []struct {
		a, b  string
		match bool
	}{
		{"http://example.com", "http://example.com", true},
		{"http://example.com", "http://another.example.com", false},
		{"https://example.com", "https://example.com", true},
		{"http://example.com/", "http://example.com", true},
		{"https://example.com", "http://example.com", false},
		{"http://example.com/foo", "http://example.com/foo/", true},
		{"http://example.com/foo//", "http://example.com/foo/", true},
		{"http://example.com/./foo/", "http://example.com/foo/", true},
		{"http://example.com/bar/../foo/", "http://example.com/foo/", true},
		{"/foo", "/foo", true},
		{"/foo", "/foo/", true},
		{"/foo/.", "/foo/", true},
		{"%/1234", "%/1234", true},
		{"%/1234", "%/123", false},
		{"/1234", "%/1234", false},
	} {
		if tt.match != EqualURLs(tt.a, tt.b) {
			t.Errorf("Expected %q==%q to be %t", tt.a, tt.b, tt.match)
		}
	}
}
