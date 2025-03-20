package main

import (
	"errors"
	"fmt"
	"math"
)

type error interface {
	Error() string
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("negative value")
	}
	return math.Sqrt(x), nil
}

func main() {
	val, err := sqrt(0)
	if err != nil {
		// 处理错误
		fmt.Println(err) // negative value
		return
	}
	// 没有错误，打印结果
	fmt.Println(val)
}
