package main

import "fmt"

/*
   题目：如何仅用递归函数和栈操作逆序一个栈
   内容：一个栈依次压入 1、2、3、4、5,那么从栈顶到栈底分别为 5、4、3、2、1。将这个栈转置
        后,从栈顶到栈底为 1、2、3、4、5,也就是实现栈中元素的逆序,但是只能用递归函数来实
        现,不能用其他数据结构。
   要求：无
   难度：尉 ★★☆☆
*/

func getAndRemoveLastElement(stack *[]int) int {
    result := (*stack)[len(*stack)-1]
    *stack = (*stack)[:len(*stack)-1]
    if len(*stack) < 1 {
        return result
    }
    last := getAndRemoveLastElement(stack)
    *stack = append(*stack, result)
    return last
}

func reverseStack(stack *[]int) {
    if len(*stack) < 1 {
        return
    }
    last := getAndRemoveLastElement(stack)
    reverseStack(stack)
    *stack = append(*stack, last)
}

func testReverseStack() {
    stack := []int{1, 2, 3, 4, 5}
    reverseStack(&stack)
    fmt.Println(stack)
}
