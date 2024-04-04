package main

import "fmt"

func main() {

	fmt.Println("Maping course")


	langwages := make(map[string]string)

	langwages["JS"] = "Javascript"
	langwages["RB"] = "Ruby"
	langwages["PY"] = "Python"
	// langwages["JS"] = "Javascript"

	fmt.Println("List of Langwages", langwages)


	delete(langwages,"RB")

	fmt.Println("new list", langwages)

}