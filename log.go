package console

import (
	"fmt"
	"time"
)

var logfunc func(a ...any) = longlog

func Log(a ...any) {
	logfunc(a...)
}

func longlog(a ...any) {
	fmt.Println(append([]any{"[" + time.Now().Format(format) + "]"}, a...)...)
}

func shortlog(a ...any) {
	fmt.Println(a...)
}
