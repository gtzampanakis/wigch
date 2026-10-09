package main

import (
    "math/rand"
    "testing"
)

func TestParseRankFromFen(t *testing.T) {
    var rank Rank
    var unconsumedFenOffset int
    var err error

    _, err = parseRankFromFen("", 0, '/', &rank)
    if err.Error() != "FEN rank end byte not found" { t.Fail() }

    _, err = parseRankFromFen("r/", 0, '/', &rank)
    if err.Error() != "FEN rank too short" { t.Fail() }

    _, err = parseRankFromFen("rrrrrrrrr/", 0, '/', &rank)
    if err.Error() != "FEN rank too long" { t.Fail() }

    unconsumedFenOffset, err = parseRankFromFen("rrrrrrrr/", 0, '/', &rank)
    if unconsumedFenOffset != nFiles + 1 { t.Fail() }
    if err != nil { t.Fail() }

    var pieceBytes = []byte{'p', 'r', 'n', 'b', 'q', 'k', 'P', 'R', 'N', 'B', 'Q', 'K'}

    /* Test valid piece combinations. */
    for i := 0; i < 100; i++ {
        var fenBytes [nFiles+1]byte
        for j := 0; j < nFiles; j++ {
            fenBytes[j] = pieceBytes[rand.Intn(len(pieceBytes))]
        }
        fenBytes[nFiles] = '/'
        var fen = string(fenBytes[:])
        _, err := parseRankFromFen(fen, 0, '/', &rank)
        if err != nil { t.Fail() }
    }
}

func FuzzParseRankFromFen(f *testing.F) {
    var rank Rank
    f.Add("rrrrrrrrr/", 0)
    f.Fuzz(func(t *testing.T, fen string, fenOffset int) {
        parseRankFromFen(fen, fenOffset, '/', &rank)
    })
}
