package common

import "io"

type Provider interface {
	Color(colorName, s string) string
	Output() io.Writer
	Log(entry *Entry)
}
