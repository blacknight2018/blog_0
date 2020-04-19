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
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println(file, line)
	}
}
