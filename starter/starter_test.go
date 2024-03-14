package starter_test

import (
	"io"
	"os"
	"testing"
	"thinking/starter"
)

func TestSolution(t *testing.T) {
	og := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	p := starter.NewFoodLines(3, 4, 3, 2, 5)
	p.Solve()

	w.Close()
	os.Stdout = og

	output, _ := io.ReadAll(r)

	expected := "2\n3\n3\n4\n"
	if string(output) != expected {
		t.Errorf("Output: %q, Expected: %q", output, expected)
	}
}
