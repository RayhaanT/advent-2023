package main

import (
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func pow2(n int) (int) {
    out := 1
    for n > 0 {
        n -= 1
        out *= 2
    }
    return out
}

func day4(scanner *bufio.Scanner) {
    p1sum := 0
    copies := make([]int, 0)
    copies = append(copies, 0)
    index := 0

    for scanner.Scan() {
        var line string = scanner.Text()
        copies[index] += 1
        lists := strings.Split(strings.Split(line, ":")[1], "|")
        winningstrs := strings.Split(strings.TrimSpace(lists[0]), " ")
        havestrs := strings.Split(strings.TrimSpace(lists[1]), " ")
        winning := make([]int, 0)
        have := make([]int, 0)
        for _, s := range winningstrs {
            if s == "" {
                continue
            }
            n, err := strconv.Atoi(s)
            if err != nil {
                panic(err)
            }
            winning = append(winning, n)
        }
        for _, s := range havestrs {
            if s == "" {
                continue
            }
            n, err := strconv.Atoi(s)
            if err != nil {
                panic(err)
            }
            have = append(have, n)
        }

        found := 0
        for _, n := range have {
            for _, m := range winning {
                if n == m {
                    found += 1
                    break
                }
            }
        }
        if found != 0 {
            p1sum += pow2(found - 1)
        }
        for i := 0; i < found; i++ {
            ind := i + index + 1
            if ind >= len(copies) {
                copies = append(copies, copies[index])
            } else {
                copies[ind] += copies[index]
            }
        }
        if found == 0 && index+1 >= len(copies) {
            copies = append(copies, 0)
        }
        index += 1
    }

    p2sum := 0
    for _, v := range copies {
        p2sum += v
    }

    fmt.Printf("Part 1: %d\n", p1sum)
    fmt.Printf("Part 2: %d\n", p2sum)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
