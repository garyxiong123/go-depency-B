package B

import _ "github.com/garyxiong123/go-depency-A/A"

func Go_dependency_B_V1() {
	A.Go_dependency_A_V1()
	println("go_dependency_B_V1")
}
