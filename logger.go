package console

import (
	"bytes"
	"log"
)

var Logger = log.New(&logwriter{}, "", 0)

type logwriter struct{}

func (*logwriter) Write(p []byte) (int, error) {
	errorfunc(string(bytes.TrimRight(p, "\n")))
	return len(p), nil
}
