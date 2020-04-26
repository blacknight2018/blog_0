package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func init() {
	logrus.SetOutput(os.Stdout)
}
func SimpleLog() {
	fmt.Println(1)
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(file, line)
	}
	fmt.Println(2)
	_, file2, line2, ok2 := runtime.Caller(2)
	if ok2 {
		fmt.Println(file2, line2)
	}
	fmt.Println(3)
	_, file3, line3, ok3 := runtime.Caller(2)
	if ok3 {
		fmt.Println(file3, line3)
	}
}
