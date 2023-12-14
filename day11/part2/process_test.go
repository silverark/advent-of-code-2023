package part2

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
)

func TestProcess(t *testing.T) {

	// Double
	value := process(file.GetFile("../input_test.txt"), 2)
	expect := 374
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	// 10x
	value = process(file.GetFile("../input_test.txt"), 10)
	expect = 1030
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	// 100x
	value = process(file.GetFile("../input_test.txt"), 100)
	expect = 8410
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	value = process(file.GetFile("../input.txt"), 1000000)
	log.Println("The answer is", value)
}
