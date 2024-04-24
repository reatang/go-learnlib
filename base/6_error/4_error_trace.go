package main

import (
	"errors"
	"fmt"
	"runtime"
)

// 简单的输出栈，并使用了 errors.Warp 机制

var (
	ErrTrace = errors.New("ErrTrace")
)

func WrapError(wrapMsg string, err error) error {
	pc, file, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if !ok {
		return errors.New("WrapError 方法获取堆栈失败")
	}
	if err == nil {
		return nil
	} else {
		return fmt.Errorf("%s \n\tat %s:%d (Method %s)\nCause by: %w", wrapMsg, file, line, f.Name(), err)
	}
}

func a() error {
	return WrapError("a some error", b())
}

func b() error {
	return WrapError("b some error", c())
}

func c() error {
	return WrapError("c some error", ErrTrace)
}

func main() {
	err := a()

	fmt.Println(err)
}
