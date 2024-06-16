package fungsi

import "fmt"

var dataku = [3]int{
	1, 2, 3}

func Output() {
	for i := 0; i < len(dataku); i++ {
		fmt.Println(dataku[i])
	}
}
