package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
	求一个数组，当前位置上的值是可以跳跃的最大步，这个数组能否被跳完，起跳点为第一个元素的位置，即index=0
	例：[2,3,1,1,4] true [3,2,1,0,4] false
	解释：第二个数组例子中，从index=0开始跳，目前最远能跳3步，也可以跳1步。
	假如是3步，那么到达的index为3的值为0，此时0值是跳不完这个数组，所以为false;
	假如是走1步，到达index=1的值为2，可以走2步，到达的index=3的值为0，同上为false;
	假如走2步，到达的index=2的值为1，最多能走1步，到达的index=3的值为0，同上为fasle;
*/
func main() {

	reader := bufio.NewReader(os.Stdin)

	bts, _, err := reader.ReadLine()
	checkErr(err)
	raw := string(bts)
	raw = strings.TrimRight(raw, "\r\t\n")
	arrStrs := strings.Split(raw, " ")
	fmt.Println(arrStrs)
	arr := make([]int, len(arrStrs))
	for i, item := range arrStrs {
		num, err := strconv.Atoi(item)
		checkErr(err)
		arr[i] = num
	}

	fmt.Println(canJump(arr))
}

/*
贪心算法：
	每一步求最优解，与全局最优解对比，更新全局最优解
*/
func canJump(arr []int) bool {
	max := 0 // 记录跳跃最远位置
	length := len(arr)
	// 遍历每一步记录并且更新最远能跳跃的值，只要最远能跳跃的值能跳过终点length就是true，否则为false
	for i := range arr {
		max = int(math.Max(float64(max), float64(i+arr[i]))) //更新max值
		if i > max {
			// 此时max无法抵达当前位置，直接返回false
			return false
		}
	}
	return max >= length-1
}

/*
	动态规划算法
*/
func canJumpII(arr []int) bool {

	return false
}

func checkErr(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}
