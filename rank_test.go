package main

import (
    "math/rand"
    "testing"
)

func TestParseRankFromFen(t *testing.T) {
    var rank Rank
    var unconsumedFenOffset int
    var err error

    _, err = rank.initRankFromFen("", 0, '/')
    if err.Error() != "FEN rank end byte not found" { t.Fail() }

    _, err = rank.initRankFromFen("r/", 0, '/')
    if err.Error() != "FEN rank too short" { t.Fail() }

    _, err = rank.initRankFromFen("rrrrrrrrr/", 0, '/')
    if err.Error() != "FEN rank too long" { t.Fail() }

    unconsumedFenOffset, err = rank.initRankFromFen("rrrrrrrr/", 0, '/')
    if unconsumedFenOffset != nFiles + 1 { t.Fail() }
    if err != nil { t.Fail() }
    if rank.contents[0].filled != true { t.Fail() }
    if rank.contents[0].piece != Rook { t.Fail() }
    if rank.contents[0].color != Black { t.Fail() }

    var pieceBytes = []byte{'p', 'r', 'n', 'b', 'q', 'k', 'P', 'R', 'N', 'B', 'Q', 'K'}

    /* Test valid piece combinations. */
    for i := 0; i < 100; i++ {
        var fenBytes [nFiles+1]byte
        for j := 0; j < nFiles; j++ {
            fenBytes[j] = pieceBytes[rand.Intn(len(pieceBytes))]
        }
        fenBytes[nFiles] = '/'
        var fen = string(fenBytes[:])
        _, err := rank.initRankFromFen(fen, 0, '/')
        if err != nil { t.Fail() }
    }
}

func FuzzParseRankFromFen(f *testing.F) {
    var rank Rank
    f.Add("rrrrrrrrr/", 0)
    f.Fuzz(func(t *testing.T, fen string, fenOffset int) {
        rank.initRankFromFen(fen, fenOffset, '/')
    })
}
