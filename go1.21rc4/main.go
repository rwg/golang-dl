// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.21rc4 command runs the go command from Go 1.21rc4.
//
// To install, run:
//
//	$ go install github.com/rwg/golang-dl/go1.21rc4@latest
//	$ go1.21rc4 download
//
// And then use the go1.21rc4 command as if it were your normal go
// command.
//
// See the release notes at https://tip.golang.org/doc/go1.21
//
// File bugs at https://go.dev/issue/new
package main

import "github.com/rwg/golang-dl/internal/version"

func main() {
	version.Run("go1.21rc4")
}