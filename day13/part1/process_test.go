package part1

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
)

func TestProcess(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))

	expect := 405
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	value = process(file.GetFile("../input.txt"))
	log.Println("The answer is", value)

}
