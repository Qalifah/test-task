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
