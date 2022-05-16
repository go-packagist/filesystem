# filesystem

[![Go](https://github.com/go-packagist/filesystem/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/filesystem/actions/workflows/go.yml)

> 未完结版

## Installation

```bash
go get github.com/go-packagist/filesystem
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/go-packagist/filesystem"
)

func main() {
	fs := filesystem.NewManager(&filesystem.Config{
		Default: "local",
		Disk: map[string]interface{}{
			"local": &filesystem.LocalDriveConfig{
				Root: "temp",
			},
		},
	})

	fmt.Println(fs.Disk("local").Exists("test.txt"))
}

```
