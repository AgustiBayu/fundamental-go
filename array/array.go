package array

import "fmt"

func Array() {
	var fruits [3]string

	fruits[0] = "mangga"
	fruits[1] = "apel"
	fruits[2] = "jambu"

	for _, v := range fruits {
		fmt.Printf("%s", v)
	}
	for i := 1; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
}
