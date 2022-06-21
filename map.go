package main

import "fmt"

/*
input  = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
*/
func main() {

	input := [][]string{
		/*{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"}, */
		{"B", "C"},
		{"D", "B"},
		{"C", "A"},
	}
	dict := make(map[string]bool)

	for _, innerObject := range input {
		if i, ok := dict[innerObject[0]]; !ok || !i {
			dict[innerObject[0]] = true
		}
		if _, ok := dict[innerObject[1]]; !ok {
			dict[innerObject[1]] = false
		}
	}

	result := ""

	for key, value := range dict {
		if value == false {
			result = key
			break
		}
	}

	fmt.Println(result)
}
