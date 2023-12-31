package part2

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
)

func TestProcess(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))
	expect := 952408144115
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}
}
func TestProcessActual(t *testing.T) {
	value := process(file.GetFile("../input.txt"))
	log.Println("The answer is", value)
}
