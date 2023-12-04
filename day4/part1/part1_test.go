package part1

import (
	"bitbucket.org/silverark/aoc-2023/pkg/file"
	"log"
	"testing"
)

func TestProcess(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))
	expect := 13
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}
	value = process(file.GetFile("../input.txt"))
	log.Println("The answer is", value)
}
