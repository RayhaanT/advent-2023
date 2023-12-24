package main

import (
    "strconv"
    "math"
    "strings"
)

func stoi(s string) (int) {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return n
}

func tostring(i int) string {
    return strconv.Itoa(i)
}

func digits(n int) (int) {
    if n == 0 {
        return 1
    }
    return int(math.Ceil(math.Log10(float64(n))))
}

func parseIntegers(s string) (res []int) {
    res = make([]int, 0)
    chunks := strings.Split(s, " ")
    for _, c := range chunks {
        if c == "" {
            continue
        }
        res = append(res, stoi(c))
    }
    return
}
