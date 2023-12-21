package main

import (
    "bufio"
    "fmt"
)

func findHorizontalAxis(grid [][]rune, mask int) int {
    for i := 1; i < len(grid); i++ {
        if i == mask {
            continue
        }
        success := true
        for delta := 1; i - delta >= 0 && i + delta - 1 < len(grid); delta++ {
            if string(grid[i - delta]) != string(grid[i + delta - 1]) {
                success = false
                break
            }
        }
        if success {
            return i
        }
    }
    return -1
}

func findVerticalAxis(grid [][]rune, mask int) int {
    n := len(grid[0])
    for i := 1; i < n; i++ {
        if i == mask {
            continue
        }
        success := true
        for delta := 1; i - delta >= 0 && i + delta - 1 < n; delta++ {
            for j := 0; j < len(grid); j++ {
                if grid[j][i - delta] != grid[j][i + delta - 1] {
                    success = false
                    break
                }
            }
            if !success {
                break
            }
        }
        if success {
            return i
        }
    }
    return -1
}

func findSymmetry(grid [][]rune, mask int) int {
    h := findHorizontalAxis(grid, mask/100)
    if h != -1 {
        return h*100
    } else {
        v := findVerticalAxis(grid, mask)
        if v != -1 {
            return v
        }
    }
    return -1
}

func processMirrors(grid [][]rune) (int, int) {
    r1 := findSymmetry(grid, -1)
    for i, row := range grid {
        for j := range row {
            if grid[i][j] == '.' {
                grid[i][j] = '#'
            } else {
                grid[i][j] = '.'
            }
            r2 := findSymmetry(grid, r1)
            if r2 != -1 {
                return r1, r2
            }
            if grid[i][j] == '.' {
                grid[i][j] = '#'
            } else {
                grid[i][j] = '.'
            }
        }
    }
    panic("No smudge found")
}

func day13(scanner *bufio.Scanner) {
    p1 := 0
    p2 := 0

    grid := make([][]rune, 0)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            d1, d2 := processMirrors(grid)
            p1 += d1
            p2 += d2
            grid = make([][]rune, 0)
        } else {
            grid = append(grid, []rune(line))
        }
    }
    a, b := processMirrors(grid)
    p1 += a
    p2 += b

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
