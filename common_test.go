package logs

import (
	"testing"

	"github.com/hereyou-go/logs/level"
)

func TestLog(t *testing.T) {
	SetLevel(level.DEBUG)
	Debug("hello, %s", "one go!")
}
