package main

import (
    "bufio"
    "fmt"
)

type Node struct {
    left string
    right string
}

func isLCM(cand int, nums []int) bool {
    for _, n := range nums {
        if cand % n != 0 {
            return false
        }
    }
    return true
}

func day8(scanner *bufio.Scanner) {
    scanner.Scan()
    seq := scanner.Text()
    scanner.Scan()
    nodes := make(map[string]Node)

    for scanner.Scan() {
        line := scanner.Text()
        start := line[0:3]
        left := line[7:10]
        right := line[12:15]
        nodes[start] = Node{left, right}
    }

    seqidx := 0
    curr := "AAA"
    for curr != "ZZZ" {
        if seq[seqidx % len(seq)] == 'R' {
            curr = nodes[curr].right
        } else {
            curr = nodes[curr].left
        }
        seqidx += 1
    }

    keys := make([]string, len(nodes))

    i := 0
    for k := range nodes {
        keys[i] = k
        i++
    }
    p1 := seqidx

    // find length of each run
    lengths := make([]int, 0)
    for _, k := range keys {
        if k[2] != 'A' {
            continue
        }

        curr = k
        seqidx = 0
        for curr[2] != 'Z' {
            if seq[seqidx % len(seq)] == 'R' {
                curr = nodes[curr].right
            } else {
                curr = nodes[curr].left
            }
            seqidx += 1
        }
        
        dist := 1
        if seq[seqidx % len(seq)] == 'R' {
            curr = nodes[curr].right
        } else {
            curr = nodes[curr].left
        }
        seqidx += 1
        for curr[2] != 'Z' {
            if seq[seqidx % len(seq)] == 'R' {
                curr = nodes[curr].right
            } else {
                curr = nodes[curr].left
            }
            seqidx += 1
            dist += 1
        }

        lengths = append(lengths, dist)
    }

    // find lcm of lengths
    biggest := lengths[0]
    for _, l := range lengths {
        biggest = max(l, biggest)
    }
    mult := biggest
    for !isLCM(mult, lengths) {
        mult += biggest
    }

    p2 := mult
    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
