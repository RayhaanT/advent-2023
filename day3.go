package main

import (
    "bufio"
    "fmt"
    "os"
)

type Coord struct {
    a, b interface{}
}

func isdigit(c byte) (int) {
    n := int(c - '0')
    if n < 10 && n >= 0 {
        return n
    }
    return -1
}

func day3(scanner *bufio.Scanner) {
    file, err := os.Open("inputs/day3.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    lines := make([]string, 0)
    x := 0
    y := 0
    p1sum := 0
    for scanner.Scan() {
        y += 1
        var line string = scanner.Text()
        x = len(line)
        lines = append(lines, line)
    }

    gears := make(map[Coord][]int)

    for i := 0; i < y; i++ {
        for j := 0; j < x; j++ {
            if isdigit(lines[i][j]) == -1 {
                continue
            }
            number := 0
            success := false
            mygears := make([]Coord, 0)
            for j < x && isdigit(lines[i][j]) != -1 {
                number *= 10
                number += isdigit(lines[i][j])

                if !success {
                    for dx := -1; dx < 2; dx++ {
                        for dy := -1; dy < 2; dy++ {
                            yp := i + dy
                            xp := j + dx
                            if yp < 0 || yp >= y || xp < 0 || xp >= x {
                                continue
                            }
                            if isdigit(lines[yp][xp]) == -1 && lines[yp][xp] != '.' {
                                success = true
                                if (lines[yp][xp] == '*') {
                                    mygears = append(mygears, Coord{yp, xp})
                                }
                            }
                        }
                    }
                }

                j++
            }
            if success {
                p1sum += number
            }

            for i := 0; i < len(mygears); i++ {
                gears[mygears[i]] = append(gears[mygears[i]], number)
            }
        }
    }

    p2sum := 0
    for _, v := range gears { 
        if len(v) == 2 {
            p2sum += v[0]*v[1]
        }
    }

    fmt.Printf("Part 1: %d\n", p1sum)
    fmt.Printf("Part 2: %d\n", p2sum)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
