package part2

import (
	"log"
	"silverark/aoc-2023/pkg/file"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	value := process(file.GetFile("../input2_test.txt"))

	expect := 6
	if value != expect {
		t.Fatalf("Received %v, but expected %v", value, expect)
	}
	//
	now := time.Now()
	value = process(file.GetFile("../input.txt"))
	taken := time.Since(now)
	log.Println("The answer is", value, "found in", taken)
}
