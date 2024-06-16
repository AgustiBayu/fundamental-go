package kondisi

import "fmt"

func Kondisiku(a int) {
	if a > 80 {
		fmt.Println("Nilai Anda A")
	} else if a < 80 {
		fmt.Println("Nilai Anda B")
	} else if a < 65 {
		fmt.Println("Nilai Anda C")
	} else {
		fmt.Println("Nilai Anda D")
	}
}
