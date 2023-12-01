package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func digits(line string) ([]int, []int) {
    var nums = make([]int, len(line))
    var indices = make([]int, len(line))
    var ind = 0
    for i := 0; i < len(line); i++ {
        nums[i] = -1
        indices[i] = -1
    }
    for i := 0; i < len(line); i++ {
        var n = int(line[i] - '0')
        if 0 <= n && n < 10 {
            nums[ind] = n
            indices[ind] = i
            ind++
        }
    }
    return nums, indices
}

func numbers(line string) (int, int, int, int) {
    var names = []string{
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    }
    var firsti = -1
    var firstn = -1
    var lasti = -1
    var lastn = -1

    for i := 0; i < len(names); i++ {
        var fi = strings.Index(line, names[i])
        if fi != -1 && (firsti == -1 || fi < firsti) {
            firsti = fi
            firstn = i+1
        }

        var li = strings.LastIndex(line, names[i])
        if li != -1 && (lasti == -1 || li > lasti) {
            lasti = li
            lastn = i+1
        }
    }

    return firsti, firstn, lasti, lastn
}

func main() {
    file, err := os.Open("inputs/day1.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var p1sum int = 0
    var p2sum int = 0
    for scanner.Scan() {
        var line string = scanner.Text()
        var nums, indices = digits(line)
        var fi, fn, li, ln = numbers(line)
        var last = 0
        var lasti = -1
        for i := 0; i < len(nums); i++ {
            if nums[i] == -1 {
                break
            }
            last = nums[i]
            lasti = indices[i]
        }

        p1sum += nums[0]*10 + last

        var first = nums[0]
        if (fi < indices[0] && fi != -1) || first == -1 {
            first = fn
        }
        if (li > lasti && li != -1) || last == -1 {
            last = ln
        }

        p2sum += first*10 + last
    }

    fmt.Printf("Part 1: %d\n", p1sum)
    fmt.Printf("Part 2: %d\n", p2sum)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
