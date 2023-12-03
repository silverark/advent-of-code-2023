package part2

import (
	"bitbucket.org/silverark/aoc-2023/pkg/file"
	"log"
	"testing"
)

func TestProcess(t *testing.T) {

	value := Process(file.GetFile("../input_test.txt"))

	expect := 2286
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	value = Process(file.GetFile("../input.txt"))
	log.Println("The answer is", value)
}
