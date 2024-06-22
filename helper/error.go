package helper

import "fmt"

func PanifIfError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
