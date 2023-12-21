package main

import (
    "bufio"
    "fmt"
)

func galaxyDist(g Coord, h Coord, factor int, rowEmpty []bool, colEmpty []bool) (dist int) {
    dist = 0
    dx := 1
    if h.x < g.x {
        dx = -1
    }
    dy := 1
    if h.y < g.y {
        dy = -1
    }
    for x := g.x; x != h.x; x += dx {
        dist++
        if colEmpty[x] {
            dist += factor - 1
        }
    }
    for y := g.y; y != h.y; y += dy {
        dist++
        if rowEmpty[y] {
            dist += factor - 1
        }
    }
    return
}

func day11(scanner *bufio.Scanner) {
    p1 := 0
    p2 := 0

    galaxies := make([]Coord, 0)
    rowEmpty := make([]bool, 0)
    colEmpty := make([]bool, 0)
    first := true
    y := 0

    for scanner.Scan() {
        line := scanner.Text()
        if first {
            first = false
            for range line {
                colEmpty = append(colEmpty, true)
            }
        }

        empty := true
        for x, c := range line {
            if c == '#' {
                galaxies = append(galaxies, Coord{x, y})
                empty = false
                colEmpty[x] = false
            }
        }
        rowEmpty = append(rowEmpty, empty)
        y++
    }

    for i, g := range galaxies {
        for j := i+1; j < len(galaxies); j++ {
            h := galaxies[j]
            p1 += galaxyDist(g, h, 2, rowEmpty, colEmpty)
            p2 += galaxyDist(g, h, 1000000, rowEmpty, colEmpty)
        }
    }

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
