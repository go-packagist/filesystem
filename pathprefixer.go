package filesystem

import "strings"

type PathPrefixer struct {
	prefix    string
	separator string
}

const separator = "/"

func NewPathPrefixer(prefix string) *PathPrefixer {
	prefix = strings.TrimRight(prefix, "\\/")

	return &PathPrefixer{
		prefix:    prefix + separator,
		separator: separator,
	}
}

func (p *PathPrefixer) PrefixPath(path string) string {
	return p.prefix + strings.TrimLeft(path, "\\/")
}
