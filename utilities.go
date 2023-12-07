package main

import (
    "strconv"
    "math"
)

func stoi(s string) (int) {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return n
}

func digits(n int) (int) {
    if n == 0 {
        return 1
    }
    return int(math.Ceil(math.Log10(float64(n))))
}
