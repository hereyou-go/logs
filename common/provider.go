package common

import "io"

type Provider interface {
	io.Writer
	Color(colorName, s string) string
	Log(entry *Entry)
}
