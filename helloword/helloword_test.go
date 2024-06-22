package helloword

import (
	"testing"
)

func TestHelloWord(t *testing.T) {
	expect := "Hallo Agusti"
	result := helloWord()

	if expect != result {
		t.Errorf("data %s dan data expect %s tidak sesuai ", result, expect)
	} else {
		t.Logf("data sama %s", result)
	}
}
