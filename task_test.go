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
