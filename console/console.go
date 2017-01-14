package console

import (
	"fmt"
	"io"

	"github.com/mattn/go-colorable"
	"github.com/one-go/logs/common"
)

type Console struct {
	output io.Writer
}

func NewConsole() *Console {
	return &Console{
		output: colorable.NewColorableStdout(),
	}
}

func (c *Console) Output() io.Writer {
	return c.output
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
