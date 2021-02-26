package debug

import "fmt"

func Dbug(i string) (o string) {
	fmt.Println(i)
	fmt.Println("v2.0.0")
	return i
}
