package core

import (
	"fmt"
	"xpath"
)

func (basic *Basic) Checker() bool {
	if basic.Path != "" && basic.FileName != "" {
		url := basic.Path + basic.FileName
		fmt.Println(url)
		fmt.Println(xpath.Request(url))
	}
	return true
}
