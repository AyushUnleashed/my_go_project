package main

import (
	"fmt"

	"github.com/ayushunleashed/my_go_project/internal/package_a"
	"github.com/ayushunleashed/my_go_project/internal/package_b"
)

func main() {
	fmt.Println("Go my_go_project")
	package_a.My_function_a_1()
	package_a.My_function_a_2()
	package_b.My_function_b_3()

}
