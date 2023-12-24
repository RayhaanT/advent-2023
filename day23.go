package main

import (
    "bufio"
    "fmt"
)

func coordValid(c Coord, grid []string) bool {
    return c.x >= 0 && c.y >= 0 && c.y < len(grid[0]) && c.x < len(grid[0])
}

func longestPathDAG(grid []string, seen map[Coord]int, pos Coord) int {
    if pos.y == len(grid)-1 {
        return 1
    }

    seen[pos] = -1
    best := -1
    for dx := -1; dx < 2; dx++ {
        for dy := -1; dy < 2; dy++ {
            if dx != 0 && dy != 0 {
                continue
            }
            c := Coord{pos.x + dx, pos.y + dy}
            if !coordValid(c, grid) || grid[c.y][c.x] == '#' {
                continue
            }

            s := grid[c.y][c.x]
            if (s == '>' && dx != 1) ||
            (s == '<' && dx != -1) ||
            (s == '^' && dy != -1) ||
            (s == 'v' && dy != 1) {
                continue
            }

            d := seen[c]
            if d == 0 {
                d = longestPathDAG(grid, seen, c)
            }
            if d != -1 {
                best = max(best, d + 1)
            }
        }
    }
    seen[pos] = best
    return best
}

func longestPath(grid []string, seen map[Coord]bool, pos Coord) int {
    if pos.y == len(grid)-1 {
        return 1
    }

    seen[pos] = true
    best := -1
    for dx := -1; dx < 2; dx++ {
        for dy := -1; dy < 2; dy++ {
            if dx != 0 && dy != 0 {
                continue
            }
            c := Coord{pos.x + dx, pos.y + dy}
            if !coordValid(c, grid) || grid[c.y][c.x] == '#' || seen[c] {
                continue
            }

            d := longestPath(grid, seen, c)
            if d != -1 {
                best = max(best, d + 1)
            }
        }
    }
    seen[pos] = false
    return best
}

func day23(scanner *bufio.Scanner) {
    p1 := 0
    p2 := 0

    grid := make([]string, 0)

    for scanner.Scan() {
        grid = append(grid, scanner.Text())
    }

    p1 = longestPathDAG(grid, make(map[Coord]int, 0), Coord{1, 0}) - 1
    p2 = longestPath(grid, make(map[Coord]bool, 0), Coord{1, 0}) - 1

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
