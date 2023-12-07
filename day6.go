package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "math"
)

func off10(n int) (int) {
    if n == 0 {
        return 1
    }
    return int(math.Pow10(digits(int(n))))
} 

func concatNum(lst []int) (int) {
    c := 0
    for _, t := range lst {
        c *= off10(t)
        c += t
    }
    return c
}

func qroot(a, b, c float64) (int, int) {
    det := b*b - 4*a*c
    if det < 0 {
        return -1, -1
    }
    det = math.Sqrt(det)
    r1 := (-b + det)/(2*a)
    r2 := (-b - det)/(2*a)
    hold := r1
    r1 = min(r1, r2)
    r2 = max(hold, r2)
    if math.Abs(r1 - math.Round(r1)) < 1e-8 {
        r1 = math.Round(r1) + 1
    } else {
        r1 = math.Ceil(r1)
    }
    if math.Abs(r2 - math.Round(r2)) < 1e-8 {
        r2 = math.Round(r2) - 1
    } else {
        r2 = math.Floor(r2)
    }
    return int(r1), int(r2)
}

func parseLine(line string) ([]int) {
    data := make([]int, 0)
    list := strings.Split(strings.Split(line, ":")[1], " ")
    for _, s := range list {
        if s == "" { continue }
        data = append(data, stoi(s))
    }
    return data
}

func day6(scanner *bufio.Scanner) {
    file, err := os.Open("inputs/day6.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner.Scan()
    times := parseLine(scanner.Text())
    scanner.Scan()
    records := parseLine(scanner.Text())

    p1 := 1
    for i := 0; i < len(records); i++ {
        r1, r2 := qroot(1, float64(-times[i]), float64(records[i]))
        ways := (r2 - r1) + 1
        p1 *= int(ways)
    }

    newtime := concatNum(times)
    newrec := concatNum(records)
    r1, r2 := qroot(1, float64(newtime), float64(newrec))
    p2 := (r2 - r1) + 1

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
