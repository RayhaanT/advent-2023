package main

import (
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

func parsepull(pull string) (int, int, int) {
    pull = strings.TrimSpace(pull)
    parts := strings.Split(pull, ",")
    red := 0
    green := 0
    blue := 0
    for i := 0; i < len(parts); i++ {
        parts[i] = strings.TrimSpace(parts[i])
        // fmt.Printf("%s\n", parts[i])
        spl := strings.Split(parts[i], " ")
        n, err := strconv.Atoi(spl[0])
        if err != nil {
            panic(err)
        }
        if (spl[1] == "red") {
            red = n
        }
        if (spl[1] == "green") {
            green = n
        }
        if (spl[1] == "blue") {
            blue = n
        }
    }
    return red, green, blue
}

func day2(scanner *bufio.Scanner) {
    p1sum := 0
    p2sum := 0
    index := 0
    for scanner.Scan() {
        index += 1
        var line string = scanner.Text()

        lsts := strings.Split(line, ":")[1][1:]
        pulls := strings.Split(lsts, ";")
        success := true
        maxr := 0
        maxg := 0
        maxb := 0
        for i := 0; i < len(pulls); i++ {
            red, green, blue := parsepull(pulls[i])
            if red > maxr {
                maxr = red
            }
            if blue > maxb {
                maxb = blue
            }
            if green > maxg {
                maxg = green
            }
            if red > 12 || green > 13 || blue > 14 {
                success = false
            }
        }
        power := maxr*maxg*maxb
        p2sum += power
        if success {
            p1sum += index
        }
    }

    fmt.Printf("Part 1: %d\n", p1sum)
    fmt.Printf("Part 2: %d\n", p2sum)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
