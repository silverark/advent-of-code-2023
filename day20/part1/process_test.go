package part1

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
)

func TestProcessTest(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))
	expect := 32000000
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}
}
func TestProcessTest2(t *testing.T) {
	value := process(file.GetFile("../input2_test.txt"))
	expect := 11687500
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}
}
func TestProcessActual(t *testing.T) {
	value := process(file.GetFile("../input.txt"))
	log.Println("The answer is", value)
}
