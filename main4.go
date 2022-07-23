//dictionary
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {

		fmt.Println("[english word] -> [turkish word]")

	}

	query := args[0]

	dict := map[string]string{

		"good":    "kötu",
		"great":   "harika",
		"perfect": "mukemmel",
		"awesome": "mukemmel",
	}

	delete(dict, "awesome")
	//dict = nil
	for k := range dict {
		delete(dict, k)

	}

	turkish := make(map[string]string, len(dict))
	for k, v := range dict {
		turkish[v] = k
	}

	// turkish["good"] = "guzel"
	// dict["great"] = "kusursuz"
	fmt.Printf("english: %q\nturkish: %q\n", dict, turkish)

	fmt.Printf("%#v\n", dict)

	if value, ok := turkish[query]; ok {

		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	if value, ok := dict[query]; ok {

		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)
	fmt.Printf("# of Keys: %d\n", len(dict))

	//without using a map
	// english := []string{"good", "great", "perfect"}
	// turkish := []string{"iyi", "harika", "mukemmel"}

	// for i, w := range english {
	// 	if query == w {

	// 		fmt.Printf("%q means %q\n", w, turkish[i])
	// 		return

	// 	}

	// }

	// fmt.Printf("%q not found\n", query)
}
