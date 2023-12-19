package main

import (
    "bufio"
    "fmt"
)

func containsCoord(lst []Coord, c Coord) bool {
    for _, n := range lst {
        if c == n {
            return true
        }
    }
    return false
}

func rightCoord(c Coord, dir int) Coord {
    switch dir {
        case 0: return Coord{c.x+1, c.y}
        case 1: return Coord{c.x, c.y+1}
        case 2: return Coord{c.x-1, c.y}
        case 3: return Coord{c.x, c.y-1}
        default: panic("Bad direction")
    }
}

func dxdir(dir int) int {
    if dir == 1 {
        return 1
    } else if dir == 3 {
        return -1
    } else {
        return 0
    }
}

func dydir(dir int) int {
    if dir == 0 {
        return -1
    } else if dir == 2 {
        return 1
    } else {
        return 0
    }
}

func day10(scanner *bufio.Scanner) {
    p1 := 0
    p2 := 0

    grid := make([][]rune, 0)

    for scanner.Scan() {
        grid = append(grid, []rune(scanner.Text()))
    }

    sx := 0
    sy := 0
    for y, row := range grid {
        for x, cell := range row {
            if cell == 'S' {
                sx = x
                sy = y
            }
        }
    }

    coords := make([]Coord, 0)
    coords = append(coords, Coord{sx, sy})
    x := sx-1
    y := sy
    coords = append(coords, Coord{x, y})
    dir := 3
    length := 0
    right := make([]Coord, 0)
    for x != sx || y != sy {
        right = append(right, rightCoord(Coord{x, y}, dir))
        switch grid[y][x] {
            case 'L':
                if dir == 2 {
                    dir = 1
                } else {
                    dir = 0
                }
            case 'J':
                if dir == 2 {
                    dir = 3
                } else {
                    dir = 0
                }
            case '7':
                if dir == 0 {
                    dir = 3
                } else {
                    dir = 2
                }
            case 'F':
                if dir == 0 {
                    dir = 1
                } else {
                    dir = 2
                }
        }
        right = append(right, rightCoord(Coord{x, y}, dir))

        x += dxdir(dir)
        y += dydir(dir)

        c := Coord{x, y}
        coords = append(coords, c)
        length++
    }

    for y, row := range grid {
        for x := range row {
            if !containsCoord(coords, Coord{x, y}) {
                if containsCoord(right, Coord{x, y}) {
                    grid[y][x] = '#'
                } else {
                    grid[y][x] = '.'
                }
            }
        }
    }

    for y, row := range grid  {
        for x, c := range row {
            if c == '#' {
                for dir := 0; dir < 4; dir++ {
                    ex := x + dxdir(dir)
                    ey := y + dydir(dir)
                    for ey >= 0 && ey < len(grid) && ex >= 0 && ex < len(grid[ey]) &&
                        grid[ey][ex] == '.' {
                        grid[ey][ex] = '#'
                        ex += dxdir(dir)
                        ey += dydir(dir)
                    }
                }
            }
        }
    }

    // for _, row := range grid {
    //     for _, cell := range row {
    //         fmt.Printf("%c", cell)
    //     }
    //     fmt.Println()
    // }

    for _, row := range grid {
        for _, cell := range row {
            if cell == '#' {
                p2++
            }
        }
    }

    p1 = (length+1)/2
    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
