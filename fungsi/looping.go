package fungsi

import "fmt"

var fruits = [3]string{
	"melon", "apel", "jambu"}

func MyFruits() {
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
}
