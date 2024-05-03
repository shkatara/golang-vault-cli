package utils

import (
	"fmt"
)

func Checkerr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
