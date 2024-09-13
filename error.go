package console

import (
	"fmt"
	"os"
	"time"
)

var error func(a ...any) = longerror

func Error(a ...any) {
	error(a...)
}

func longerror(a ...any) {
	fmt.Fprintln(os.Stderr, append([]any{"[" + time.Now().Format(format) + "] [ERROR]"}, a...)...)
}

func shorterror(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}
