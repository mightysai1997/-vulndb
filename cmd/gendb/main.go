// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command gendb provides a tool for converting YAML reports into JSON
// database.
package main

import (
	"context"
	"flag"
	"log"

	"golang.org/x/vulndb/internal/database"
)

var (
	repoDir = flag.String("repo", ".", "Directory containing vulndb repo")
	jsonDir = flag.String("out", "out", "Directory to write JSON database to")
	indent  = flag.Bool("indent", false, "Indent JSON for debugging")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	if err := database.Generate(ctx, *repoDir, *jsonDir, *indent); err != nil {
		log.Fatal(err)
	}
}
