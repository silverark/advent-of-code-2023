package part2

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	value := process(file.GetFile("../input_test.txt"))

	expect := 46
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}

	currentTime := time.Now()
	value = process(file.GetFile("../input.txt"))
	completedIn := time.Since(currentTime)
	log.Println("The answer is", value, "found in", completedIn)

	if value != 37384986 {
		t.Fatalf("OptimisedFailed")
	}
}
