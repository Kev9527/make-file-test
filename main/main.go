package main

import (
	"fmt"

	"github.com/Kev9627/make_file_test/svc/a"
	"github.com/Kev9627/make_file_test/svc/b"
)

func main() {
	fmt.Println("test")
	a.PrintBB(b.Caculate(1, 2))
}
