package main

import (
	"fmt"

	"github.com/Kev9527/make-file-test/svc/b"

	"github.com/Kev9527/make-file-test/svc/a"
)

func main() {
	fmt.Println("test")
	a.PrintBB(b.Caculate(1, 2))
}
