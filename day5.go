package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Interval struct {
    dest int
    source int
    length int
}

func applyMap(m []Interval, n int) (int) {
    for _, i := range m {
        if i.source <= n && i.source + i.length > n {
            return i.dest + (n - i.source)
        }
    }
    return n
}

func mapInterval(m []Interval, live []Interval) ([]Interval) {
    mapped := make([]Interval, 0)
    for _, tomap := range m {
        remaining := make([]Interval, 0)
        for _, l := range live {
            if l.dest + l.length > tomap.source && l.dest < tomap.source + tomap.length {
                start := max(l.dest, tomap.source)
                end := min(l.dest + l.length, tomap.source + tomap.length)
                off := max(l.dest - tomap.source, 0)
                mapped = append(mapped, Interval{tomap.dest + off, tomap.dest + off, end - start})
                if l.dest < tomap.source {
                    remaining = append(remaining, Interval{l.dest, l.dest, tomap.source - l.dest})
                }
                if l.dest + l.length > tomap.source + tomap.length {
                    st := tomap.source + tomap.length
                    remaining = append(remaining, Interval{st, st, l.dest + l.length - st})
                }
            } else {
                remaining = append(remaining, l)
            }
        }
        live = remaining
    }
    for _, r := range live {
        mapped = append(mapped, r)
    }
    return mapped
}

func day5(scanner *bufio.Scanner) {
    file, err := os.Open("inputs/day5.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    seeds := make([]int, 0)
    maps := make([][]Interval, 0)
    index := 0
    for scanner.Scan() {
        var line string = scanner.Text()
        if line == "" {
            index += 1
            maps = append(maps, make([]Interval, 0))
            continue
        }
        if index == 0 {
            seedlst := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")
            for _, s := range seedlst {
                seeds = append(seeds, stoi(s))
            }
            continue
        }
        if strings.Contains(line, ":") {
            continue
        }
        nums := strings.Split(line, " ")
        a := stoi(nums[0])
        b := stoi(nums[1])
        length := stoi(nums[2])
        maps[index - 1] = append(maps[index - 1], Interval{a, b, length})
    }

    p1 := -1
    for _, s := range seeds {
        end := s
        for _, m := range maps {
            end = applyMap(m, end)
        }
        if p1 == -1 || end < p1 {
            p1 = end
        }
    }

    live := make([]Interval, 0)
    for i := 0; i < len(seeds); i += 2 {
        live = append(live, Interval{seeds[i], seeds[i], seeds[i+1]})
    }

    for _, m := range maps {
        live = mapInterval(m, live)
    }

    p2 := -1
    for _, i := range live {
        if i.dest < p2 || p2 == -1 {
            p2 = i.dest
        }
    }

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
