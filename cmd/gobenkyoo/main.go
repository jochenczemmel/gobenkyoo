package main

import "fmt"

func main() {

	// TODO: call real code
	expectTest()
}

// expectTest is used to explore the expect test scripts.
func expectTest() {

	// this version is only for setup system test
	fmt.Printf("Q: world\nA: ")
	var answer string
	fmt.Scanf("%s", &answer)
	if answer == "世界" {
		// fmt.Println("xxx")
		fmt.Println("ok")
	} else {
		fmt.Println("wrong")
	}

	fmt.Print("save answer (y/n): ")
	fmt.Scanf("%s", &answer)
	if answer == "y" || answer == "Y" {
		fmt.Println("saved")
	}
}
