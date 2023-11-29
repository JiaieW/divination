package util

import (
	"math/rand"
	"strconv"
	"time"
)

// 一变二变三变
func Change(curNum int, bianTime int) int {
	if bianTime > 3 {
		return curNum
	}
	//一分为二
	tian := randInt(1, curNum-2)
	//三才
	ren := 1
	di := curNum - (tian + ren)
	modTian := mod4(tian)
	modDi := mod4(di)
	//一次变之后的余数
	curYu := curNum - modTian - modDi - ren
	return Change(curYu, bianTime+1)
}

// 除4取余数象征四季
func mod4(num int) int {
	yu := num % 4
	if yu == 0 {
		yu = 4
	}
	return yu
}

// 取一个随机数
func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max <= min || min == 0 || max == 0 {
		return max
	}
	return r.Intn(max-min) + min
}

func ArrayToString(arr []int) string {
	var str string
	for _, v := range arr {
		str += strconv.Itoa(v)
	}
	return str
}

func IssetInSlice(v int, bianyao []int) int {
	for _, num := range bianyao {
		if v == num {
			return 0
		}
	}
	return v
}
