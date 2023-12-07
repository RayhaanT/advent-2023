package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    args := os.Args[1:]
    if len(args) < 1 {
        os.Exit(1)
    }
    day := stoi(args[0])

    filename := fmt.Sprintf("inputs/day%d.txt", day)
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    switch day {
        case 1: day1(scanner)
        case 2: day2(scanner)
        case 3: day3(scanner)
        case 4: day4(scanner)
        case 5: day5(scanner)
        case 6: day6(scanner)
    }
}
