package common

import (
	"fmt"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`\{\s?\w+\s?\:`)

// Color 会根据给定颜色名称返回已经着色的字符串(可用的话)
type Color func(colorName, s string) string

// FormatMessage 使用给定的color， args格式化消息。
func FormatMessage(color Color, args ...interface{}) string {
	msg := ""

	if len(args) > 0 {
		if src, ok := args[0].(string); ok {
			format := ""
			for {
				loc := reg.FindStringIndex(src)
				if loc != nil && len(loc) == 2 {
					index := strings.Index(src[loc[1]:], "}")
					if index < 0 {
						format += src[0:loc[1]]
						src = src[loc[1]:]
					} else {
						colorName := strings.Trim(src[loc[0]+1:loc[1]], " ")
						colorName = strings.Trim(colorName, ":")
						colorName = strings.Trim(colorName, " ")
						format += src[:loc[0]]
						format += color(colorName, src[loc[1]:loc[1]+index])
						src = src[loc[1]+index+1:]
					}
				} else {
					format += src
					src = ""
					break
				}
			}
			msg += fmt.Sprintf(format, args[1:]...)

		} else if err, ok := args[0].(error); ok {
			msg += err.Error()
		} else {
			msg += fmt.Sprintf("%v", args[0])
		}
	}
	return msg
}

// NoColor 是一个空的着色函数，它始终返回原始的s。
func NoColor(colorName, s string) string {
	return s
}
