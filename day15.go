package main

import (
    "bufio"
    "fmt"
    "strings"
)

type Lens struct {
    label string
    focus int
}

func removeLens(slice []Lens, s int) []Lens {
    return append(slice[:s], slice[s+1:]...)
}

func hashAlgorithm(s string) int {
    res := 0
    for _, c := range s {
        res += int(c)
        res *= 17
        res = res % 256
    }
    return res
}

func day15(scanner *bufio.Scanner) {
    p1 := 0
    p2 := 0

    scanner.Scan()
    line := scanner.Text()
    comps := strings.Split(line, ",")
    boxes := make([][]Lens, 256)

    for _, s := range comps {
        p1 += hashAlgorithm(s)
        label := s[0:len(s)-1]
        focus := -1
        if s[len(s)-2] == '=' {
            label = s[0:len(s)-2]
            focus = int(s[len(s)-1] - '0')
        }
        bid := hashAlgorithm(label)
        
        if focus == -1 {
            for i := range boxes[bid] {
                if boxes[bid][i].label == label {
                    boxes[bid] = append(boxes[bid][:i], boxes[bid][i+1:]...)
                    break
                }
            }
        } else {
            found := false
            for i := range boxes[bid] {
                if boxes[bid][i].label == label {
                    found = true
                    boxes[bid][i].focus = focus
                    break
                }
            }
            if !found {
                boxes[bid] = append(boxes[bid], Lens{label, focus})
            }
        }
    }

    for i, q := range boxes {
        for j, l := range q {
            p2 += (i+1)*(j+1)*l.focus
        }
    }

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
