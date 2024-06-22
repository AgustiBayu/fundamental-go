package variabel

import (
	"fmt"
	"testing"
)

const (
	hari = "senin"
	sekarang
)

func TestVariabel(t *testing.T) {
	data := new(string)
	*data = "hello"
	fmt.Println(data)

	fmt.Print(sekarang)
}
