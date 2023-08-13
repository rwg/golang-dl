// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "testing"

func TestVersionNoPatch(t *testing.T) {
	data := []struct {
		in  string
		out string
	}{
		{"go1.11.4", "devel/release#go1.11.minor"},
		{"go1.12", "go1.12"},
		{"go1.12beta1", "go1.12"},
		{"go1.12rc2", "go1.12"},
	}
	for _, item := range data {
		if out := versionNoPatch(item.in); out != item.out {
			t.Errorf("versionNoPatch(%q) = %q; want %q", item.in, out, item.out)
		}
	}
}
