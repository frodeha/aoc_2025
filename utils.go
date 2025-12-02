package aoc2025

import "fmt"

func Assert(b bool, s string, args ...interface{}) {
	if !b {
		Fatal(fmt.Errorf(s, args...))
	}
}

func Debug(s string, args ...interface{}) {
	if false {
		fmt.Printf(s, args...)
	}
}

func Fatal(err error) {
	if err != nil {
		panic(err)
	}
}
