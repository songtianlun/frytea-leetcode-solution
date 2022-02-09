package frytea_leetcode

import (
    "sort"
)

func firstMissingPositive_with_no_map(nums []int) int {
    sort.Ints(nums)
    res := 1
    for _,v := range nums {
        if v <= 0 {
            continue
        } else if v>0 {
            if res == v {
                res ++
            } else if res-1 == v {
                continue
            } else {
                return res
            }
        } 
    }
    return res
}

func firstMissingPositive_with_map(nums []int) int {
    var numMap map[int]int = make(map[int]int, len(nums))
    for _, v := range nums{
        numMap[v] = v
    }
    for i:=1;i<=len(nums);i++ {
        _, ok := numMap[i]
        if !ok {
            return i
        }
    }
    return len(nums)+1
}

func firstMissingPositive(nums []int) int {
    return firstMissingPositive_with_map(nums)
}