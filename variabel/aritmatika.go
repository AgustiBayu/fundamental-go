package variabel

import "fmt"

var a int = 10
var b int = 5

func Aritmatika() {
	fmt.Println(tambah())
	fmt.Println(kurang())
	fmt.Println(kali())
	fmt.Println(bagi())
	fmt.Println(mod())
	sayHey()
}
func sayHey() {
	fmt.Println("Apakabar Fatiha")
}
func tambah() int {
	return a + b
}
func kurang() int {
	return a - b
}
func kali() int {
	return a * b
}
func bagi() int {
	return a / b
}
func mod() int {
	return a % b
}
