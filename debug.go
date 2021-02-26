package debug

import "fmt"

func Dbug(i string) (o string) {
	fmt.Println(i)
	fmt.Println("v1.0.0")
	return i
}
