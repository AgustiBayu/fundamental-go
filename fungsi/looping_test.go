package fungsi

import (
	"testing"
)

func TestLooping(t *testing.T) {
	var databaru = [3]string{"melon", "apel", "jambu"}

	for i, fruit := range fruits {
		if fruit != databaru[i] {
			t.Errorf("Fruit mismatch at index %d. Expected: %s, Got: %s", i, databaru[i], fruit)
		}
	}
}
