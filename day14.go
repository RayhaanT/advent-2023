package main

import (
    "bufio"
    "fmt"
    "strings"
)

func rollRocks(grid [][]rune, x int, y int, dir int) {
    nd := 0
    bd := 0
    switch dir {
        case 0: nd = y-1; bd = len(grid)
        case 1: nd = x-1; bd = len(grid[0])
        case 2: nd = y+1; bd = len(grid)
        case 3: nd = x+1; bd = len(grid[0])
    }
    dd := -1
    if dir >= 2 {
        dd = 1
    }

    res := y
    if dir % 2 == 1 {
        res = x
    }
    for nd >= 0 && nd < bd {
        if (dir % 2 == 0 && grid[nd][x] != '.') || (dir % 2 == 1 && grid[y][nd] != '.') {
            break
        }
        res = nd
        nd += dd
    }

    grid[y][x] = '.'
    if dir % 2 == 0 {
        y = res
    } else {
        x = res
    }
    grid[y][x] = 'O'
}

func spinCycle(grid [][]rune) {
    for dir := 0; dir < 4; dir++ {
        for y := range grid {
            my := y
            if dir == 2 {
                my = len(grid) - y - 1
            }
            row := grid[my]
            for x := range row {
                mx := x
                if dir == 3 {
                    mx = len(row) - x - 1
                }
                if row[mx] == 'O' {
                    rollRocks(grid, mx, my, dir)
                }
            }
        }
    }
}

func day14(scanner *bufio.Scanner) {
    p1 := 0
    p2 := 0

    grid := make([][]rune, 0)
    simple := make([][]rune, 0)
    for scanner.Scan() {
        grid = append(grid, []rune(scanner.Text()))
        simple = append(simple, []rune(scanner.Text()))
    }

    // Part one
    for y, row := range simple {
        for x, c := range row {
            if c == 'O' {
                rollRocks(simple, x, y, 0)
            }
        }
    }
    for y, row := range simple {
        for _, cell := range row {
            if cell == 'O' {
                p1 += len(simple) - y
            }
        }
    }

    // Part 2
    set := make(map[string]int)
    cycle := 1
    tot := 1000000000
    loop := 0

    for cycle <= tot {
        spinCycle(grid)

        str := make([]string, len(grid))
        for i := range grid {
            str[i] = string(grid[i])
        }
        flat := strings.Join(str, "")
        if set[flat] > 0 {
            loop = cycle - set[flat]
            break
        }
        set[flat] = cycle
        cycle++
    }

    remaining := tot - cycle
    tocomp := remaining % loop
    for i := 0; i < tocomp; i++ {
        spinCycle(grid)
    }

    for y, row := range grid {
        for _, cell := range row {
            if cell == 'O' {
                p2 += len(grid) - y
            }
        }
    }

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
