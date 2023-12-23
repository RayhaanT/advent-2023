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
        case 7: day7(scanner)
        case 8: day8(scanner)
        case 9: day9(scanner)
        case 10: day10(scanner)
        case 11: day11(scanner)
        // case 12: day12(scanner)
        case 13: day13(scanner)
        case 14: day14(scanner)
        // case 15: day15(scanner)
        // case 16: day16(scanner)
        // case 17: day17(scanner)
        // case 18: day18(scanner)
        // case 19: day19(scanner)
        // case 20: day20(scanner)
        // case 21: day21(scanner)
        // case 22: day22(scanner)
        // case 23: day23(scanner)
        // case 24: day24(scanner)
        // case 25: day25(scanner)
        default: panic("Not implemented")
    }
}
