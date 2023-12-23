package main

import (
    "bufio"
    "fmt"
)

type Beam struct {
    pos Coord
    dir int
}

func beamValid(b Beam, grid []string) bool {
    if b.pos.x < 0 || b.pos.x >= len(grid[0]) ||
    b.pos.y < 0 || b.pos.y >= len(grid) {
        return false
    } else {
        return true
    }
}

func stepBeam(b Beam, grid []string) (Beam, Beam) {
    c := grid[b.pos.y][b.pos.x]
    switch c {
        case '\\':
            switch b.dir {
                case 0: b.dir = 3
                case 1: b.dir = 2
                case 2: b.dir = 1
                case 3: b.dir = 0
            }
        case '/':
            switch b.dir {
                case 0: b.dir = 1
                case 1: b.dir = 0
                case 2: b.dir = 3
                case 3: b.dir = 2
            }
        case '-':
            if b.dir % 2 == 0 {
                return Beam{Coord{b.pos.x-1, b.pos.y}, 3},
                    Beam{Coord{b.pos.x+1, b.pos.y}, 1}
            }
        case '|':
            if b.dir % 2 == 1 {
                return Beam{Coord{b.pos.x, b.pos.y-1}, 0},
                    Beam{Coord{b.pos.x, b.pos.y+1}, 2}
            }
    }

    switch b.dir {
        case 0: b.pos.y--
        case 1: b.pos.x++
        case 2: b.pos.y++
        case 3: b.pos.x--
    }
    return b, Beam{Coord{-1,-1}, 0}
}

func countEnergized(start Beam, grid []string) int {
    energies := make([][4]bool, len(grid)*len(grid[0]))
    beams := make([]Beam, 0)
    beams = append(beams, start)
    energies[start.pos.y*len(grid) + start.pos.x][start.dir] = true
    invalid := Beam{Coord{-1,-1}, 0}

    for len(beams) > 0 {
        todelete := make([]int, 0)
        for i, b := range beams {
            newb := invalid
            b, newb = stepBeam(b, grid)
            beams[i] = b

            ind := newb.pos.y*len(grid) + newb.pos.x
            if beamValid(newb, grid) && !energies[ind][newb.dir] {
                beams = append(beams, newb)
                energies[ind][newb.dir] = true
            }

            ind = b.pos.y*len(grid) + b.pos.x
            if !beamValid(b, grid) || energies[ind][b.dir] {
                todelete = append(todelete, i)
            } else {
                energies[ind][b.dir] = true
            }
        }

        offset := 0
        for _, d := range todelete {
            di := d - offset
            beams = append(beams[:di], beams[di+1:]...)
            offset++
        }
    }

    res := 0
    for _, e := range energies {
        suc := false
        for _, b := range e {
            if b {
                suc = true
                break
            }
        }
        if suc {
            res += 1
        }
    }
    return res
}

func day16(scanner *bufio.Scanner) {
    p1 := 0
    p2 := 0

    grid := make([]string, 0)
    for scanner.Scan() {
        grid = append(grid, scanner.Text())
    }

    p1 = countEnergized(Beam{Coord{0, 0}, 1}, grid)
    p2 = p1
    for y := 0; y < len(grid); y++ {
        p2 = max(countEnergized(Beam{Coord{0, y}, 1}, grid), p2)
        p2 = max(countEnergized(Beam{Coord{len(grid[0]) - 1, y}, 3}, grid), p2)
    }
    for x := 0; x < len(grid[0]); x++ {
        p2 = max(countEnergized(Beam{Coord{x, 0}, 2}, grid), p2)
        p2 = max(countEnergized(Beam{Coord{x, len(grid) - 1}, 0}, grid), p2)
    }

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
