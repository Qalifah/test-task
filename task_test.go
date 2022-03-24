package test_task

import "testing"

func Test_testValidity(t *testing.T) {
	for i := 0; i < 4; i++ {
		status := testValidity(generate(true))
		if !status {
			t.Errorf("expected %v, got %v", true, status)
		}
	}
}

func Test_averageNumber(t *testing.T) {
	avg := averageNumber("23-ab-48-caba-55-haha")
	if avg != 42 {
		t.Errorf("expected %v, got %v", 42, avg)
	}
}

func Test_wholeStory(t *testing.T) {
	text := wholeStory("23-ab-48-caba-55-haha")
	if text != "ab caba haha" {
		t.Errorf("expected %v, got %v", "ab caba haha", text)
	}
}

func Test_storyStats(t *testing.T) {
	shortest, longest, avgLen, avgLenWords := storyStats("23-ab-48-caba-55-haha")
	if shortest != "ab" {
		t.Errorf("expected %v, got %v", "ab", shortest)
	}
	if longest != "haha" {
		t.Errorf("expected %v, got %v", "haha", longest)
	}
	if avgLen != 4 {
		t.Errorf("expected %v, got %v", 4, longest)
	}
	if len(avgLenWords) != 1 {
		t.Errorf("expected len %v, got %v", 1, len(avgLenWords))
	}
}
