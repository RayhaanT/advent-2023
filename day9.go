package main

import (
    "bufio"
    "fmt"
    "strings"
)

func extrapolate(array []int) (int, int) {
    zero := true
    for i := 0; i < len(array); i++ {
        if array[i] != 0 {
            zero = false
        }
    }
    if zero {
        return 0, 0
    } else {
        diff := make([]int, len(array)-1)
        for i := 1; i < len(array); i++ {
            diff[i-1] = array[i] - array[i-1]
        }
        l, r := extrapolate(diff)
        return array[0] - l, array[len(array)-1] + r
    }
}

func day9(scanner *bufio.Scanner) {
    p1 := 0
    p2 := 0
    for scanner.Scan() {
        numsstr := strings.Split(scanner.Text(), " ")
        nums := make([]int, len(numsstr))
        for i, s := range numsstr {
            nums[i] = stoi(s)
        }
        l, r := extrapolate(nums)
        p1 += r
        p2 += l
    }

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
