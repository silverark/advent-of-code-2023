package part2

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
)

func TestProcess(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))

	expect := 2
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}
	//
	value = process(file.GetFile("../input.txt"))
	log.Println("The answer is", value)

	if value != 1082 {
		t.Fatalf("Received %v, but expected %v", value, 1082)
	}
}
