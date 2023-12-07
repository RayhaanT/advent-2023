package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
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

func off10(n float64) (int) {
    if n == 0 {
        return 1
    }
    return int(math.Pow10(int(math.Ceil(math.Log10(n)))))
} 

func concatNum(lst []int) (int) {
    c := 0
    for _, t := range lst {
        c *= off10(float64(t))
        c += t
    }
    return c
}

func qroot(a float64, b float64, c float64) (int, int) {
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

func main() {
    file, err := os.Open("inputs/day6.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    first := true
    times := make([]int, 0)
    records := make([]int, 0)
    for scanner.Scan() {
        var line string = scanner.Text()
        list := strings.Split(strings.Split(line, ":")[1], " ")
        if first {
            first = false
            for _, s := range list {
                if s == "" {
                    continue
                }
                times = append(times, stoi(s))
            }
        } else {
            for _, s := range list {
                if s == "" {
                    continue
                }
                records = append(records, stoi(s))
            }
        }
    }

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

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
