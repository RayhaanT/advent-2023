package main

import (
    "bufio"
    "fmt"
    "strings"
    "sort"
)

type Hand struct {
    vals []int
    bid int
    kind int
}

func parseCard(c byte) (int) {
    n := c - '0'
    if int(n) >= 0 && int(n) < 10 {
        return int(n)
    }
    switch c {
    case 'T': return 10
    case 'J': return 11
    case 'Q': return 12
    case 'K': return 13
    case 'A': return 14
    default: panic("Invalid card")
    }
}

func parseHand(handstr string) (hand []int) {
    hand = []int{0, 0, 0, 0, 0}
    for i, c := range handstr {
        hand[i] = parseCard(byte(c))
    }
    return
}

func countPairs(hand []int) (pairs int) {
    seen := make([]int, 0)
    pairs = 0
    for _, n := range hand {
        success := false
        for _, s := range seen {
            if n == s {
                pairs++
                success = true
            }
        }
        if !success {
            seen = append(seen, n)
        }
    }
    return
}

func getKind(hand []int) (int) {
    sorted := make([]int, 5)
    copy(sorted, hand)
    sort.Ints(sorted)
    if sorted[0] == sorted[4] {
        return 6
    } else if sorted[0] == sorted[3] || sorted[1] == sorted[4] {
        return 5
    } else if (sorted[0] == sorted[1] && sorted[2] == sorted[4]) ||
    (sorted[0] == sorted[2] && sorted[3] == sorted[4]) {
        return 4
    } else if sorted[0] == sorted[2] || sorted[1] == sorted[3] || sorted[2] == sorted[4] {
        return 3
    } else if pairs := countPairs(hand); pairs == 2 {
        return 2
    } else if pairs == 1 {
        return 1
    } else {
        return 0
    }
}

func day7(scanner *bufio.Scanner) {
    hands := make([]Hand, 0)
    wildhands := make([]Hand, 0)
    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), " ")
        hand := parseHand(parts[0])
        bid := stoi(parts[1])
        kind := getKind(hand)
        hands = append(hands, Hand{hand, bid, kind})

        jhand := make([]int, 5)
        copy(jhand, hand)
        jokers := 0
        for i, n := range jhand {
            if n == 11 {
                jhand[i] = 1
                jokers += 1
            }
        }
        wild := make([]int, 5)
        copy(wild, jhand)
        sort.Ints(wild)

        bestkind := 0
        for i := 2; i < 15; i++ {
            if i == 11 {
                continue
            }
            for j := 0; j < jokers; j++ {
                wild[j] = i
            }
            bestkind = max(bestkind, getKind(wild))
        }
        wildhands = append(wildhands, Hand{jhand, bid, bestkind})
    }

    cmp := func(i, j int, hands []Hand) bool {
        if hands[i].kind != hands[j].kind {
            return hands[i].kind < hands[j].kind
        } else {
            index := 0
            for hands[i].vals[index] == hands[j].vals[index] {
                index++
            }
            return hands[i].vals[index] < hands[j].vals[index]
        }
    }
    hcmp := func(i, j int) bool {
        return cmp(i, j, hands)
    }
    wcmp := func(i, j int) bool {
        return cmp(i, j, wildhands)
    }
    sort.Slice(hands, hcmp)
    sort.Slice(wildhands, wcmp)

    p1 := 0
    for i, h := range hands {
        p1 += (i+1)*h.bid
    }

    p2 := 0
    for i, h := range wildhands {
        p2 += (i+1)*h.bid
    }

    fmt.Printf("Part 1: %d\n", p1)
    fmt.Printf("Part 2: %d\n", p2)
}
