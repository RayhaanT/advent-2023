package main

import (
    "bufio"
    "fmt"
    "strings"
)

func countSpringArrangements2(conditions []rune, cont []int, memo map[string]int) int {
    if len(conditions) == 0 {
        if len(cont) == 0 {
            return 1
        } else {
            return 0
        }
    }
    if len(cont) == 0 {
        clean := true
        for _, c := range conditions {
            if c == '#' {
                clean = false
                break
            }
        }
        if clean {
            return 1
        } else {
            return 0
        }
    }

    slen := []rune(tostring(len(cont)))
    rid := append(conditions, slen...)
    str := string(rid)
    if memo[str] != 0 {
        if memo[str] == -1 {
            return 0
        }
        return memo[str]
    }

    count := 0
    sz := cont[0]
    space := findContSpace(conditions, sz)
    for _, s := range space {
        end := s+sz
        // Consumed whole string
        if end >= len(conditions) {
            if len(cont) == 1 {
                count++
            }
        } else {
            count += countSpringArrangements2(conditions[s+sz+1:], cont[1:], memo)
        }
    }
    if count != 0 {
        memo[str] = count
    } else {
        memo[str] = -1
    }
    return count
}

func findContSpace(state []rune, size int) []int {
    left := 0
    right := size
    if right > len(state) {
        return make([]int, 0)
    }

    res := make([]int, 0)
    window := make([]rune, size)
    for i := 0; i < size; i++ {
        window[i] = state[i + left]
    }
    for right <= len(state) {
        cont := true
        for _, c := range window {
            if c == '.' {
                cont = false
                break
            }
        }
        if cont {
            if right == len(state) || state[right] != '#' {
                res = append(res, left)
            }
        }
        if window[0] == '#' || right == len(state) {
            break
        }

        left++
        window = append(window[1:], state[right])
        right++
    }

    return res
}

func extendConditions(conditions []rune, cont []int) ([]rune, []int) {
    condcp := make([]rune, len(conditions))
    contcp := make([]int, len(cont))
    copy(condcp, conditions)
    copy(contcp, cont)

    for i := 0; i < 4; i++ {
        shellr := make([]rune, len(condcp))
        copy(shellr, condcp)
        conditions = append(conditions, '?')
        conditions = append(conditions, shellr...)

        shelli := make([]int, len(contcp))
        copy(shelli, contcp)
        cont = append(cont, shelli...)
    }

    return conditions, cont
}

func day12(scanner *bufio.Scanner) {
    p1 := 0
    p2 := 0

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")
        conditions := []rune(line[0])
        cont := make([]int, 0)
        for _, s := range strings.Split(line[1], ",") {
            cont = append(cont, stoi(s))
        }

        memo := make(map[string]int, 0)
        p1 += countSpringArrangements2(conditions, cont, memo)

        memo = make(map[string]int, 0)
        conditions, cont = extendConditions(conditions, cont)
        p2 += countSpringArrangements2(conditions, cont, memo)
    }

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
