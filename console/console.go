package console

import (
	"fmt"
	"io"

	"github.com/hereyou-go/logs/common"
	"github.com/mattn/go-colorable"
)

type Console struct {
	output io.Writer
}

func NewConsole() *Console {
	return &Console{
		output: colorable.NewColorableStdout(),
	}
}

func (c *Console) Color(colorName, s string) string {
	setter := GetColorSetter(colorName)
	if setter == nil {
		return s
	}
	return setter(true) + s + ResetSetter(true)
}

func (c *Console) Log(entry *common.Entry) {
	msg := FormatLog(c, entry)
	_, err := fmt.Fprint(c.output, msg)
	if err != nil {
		panic(fmt.Errorf("[mano.logs.Console.Log] %v", err))
	}
}

func (c *Console) Write(p []byte) (n int, err error) {
	return c.output.Write(p)
}
