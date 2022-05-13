# filesystem

[![Go](https://github.com/go-packagist/cache/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/go-packagist/cache/actions/workflows/go.yml)

## Installation

```bash
go get github.com/go-packagist/filesystem
```

## Usage

```go
package main

import "github.com/go-packagist/filesystem"

func main() {
    filesystem.Get("/path/filename.txt")
    ……
}
```