package util

import (
	"fmt"
	"math/rand"
	"testing"
)

// 固定随机数生成器以获得可预测的结果
var fixedRand = rand.New(rand.NewSource(1))

func TestMod4(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 1},
	}

	for _, test := range tests {
		got := mod4(test.input)
		if got != test.want {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.want, got)
		}
	}
}

func TestRandInt(t *testing.T) {
	min := 1
	max := 5

	got := randInt(min, max)

	if got < min || got >= max {
		t.Errorf("Expected a number between %v and %v, but got %v", min, max, got)
	}
}

func TestChange(t *testing.T) {
	// 一个简单的测试用例
	curNum := 49
	bianTime := 1

	for i := 0; i < 100; i++ {
		got := Change(curNum, bianTime)
		fmt.Println(got)
		if got != 24 && got != 28 && got != 32 && got != 36 {
			t.Errorf("Expected a positive number, but got %v", got)
		}
	}

}
