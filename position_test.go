package main

import (
    "errors"
    "testing"
)

func TestParseFenByte(t *testing.T) {
    var buffer [nFiles]SquareContent
    var bufferIndex int
    var err error

    bufferIndex, err = ParseFenByte('p', &buffer, 0)
    if (err != nil) {
        t.Fail()
    }
    if (bufferIndex != 1) {
        t.Fail()
    }
    if (buffer[0].filled != true) {
        t.Fail()
    }
    if (buffer[0].color != Black) {
        t.Fail()
    }
    if (buffer[0].piece != Pawn) {
        t.Fail()
    }

    bufferIndex, err = ParseFenByte('7', &buffer, 0)
    if (err != nil) {
        t.Fail()
    }
    if (bufferIndex != 7) {
        t.Fail()
    }
    if (buffer[6].filled != false) {
        t.Fail()
    }

    bufferIndex, err = ParseFenByte('a', &buffer, 0)
    if !errors.Is(err, FenParsingError{msg: "Invalid FEN byte: 'a'"}) {
        t.Fail()
    }

    bufferIndex, err = ParseFenByte('1', &buffer, 8)
    if !errors.Is(err, FenParsingError{msg: "Out of space in buffer"}) {
        t.Fail()
    }

    bufferIndex, err = ParseFenByte('2', &buffer, 7)
    if !errors.Is(err, FenParsingError{msg: "Out of space in buffer"}) {
        t.Fail()
    }

}

func TestParseFenOneRank(t *testing.T) {
    // var fen string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
    var buffer [nFiles]SquareContent
    var err error
    var fenIndex int

    _, err = parseFenOneRank("", 0, &buffer)
    if err == nil {
        t.Fail()
    }
    if err.Error() != "FEN rank too short" {
        t.Fail()
    }

    _, err = parseFenOneRank("r", 0, &buffer)
    if err == nil {
        t.Fail()
    }
    if err.Error() != "FEN rank too short" {
        t.Fail()
    }

    fenIndex, err = parseFenOneRank("rnbqkbnr", 0, &buffer)
    if err != nil {
        t.Fail()
    }
    if fenIndex != 8 {
        t.Fail()
    }
    if buffer[0].filled != true {
        t.Fail()
    }
    if buffer[0].color != Black {
        t.Fail()
    }
    if buffer[0].piece != Rook {
        t.Fail()
    }
    if buffer[7].filled != true {
        t.Fail()
    }
    if buffer[7].color != Black {
        t.Fail()
    }
    if buffer[7].piece != Rook {
        t.Fail()
    }

    fenIndex, err = parseFenOneRank("r7", 0, &buffer)
    if err != nil {
        t.Fail()
    }
    if fenIndex != 2 {
        t.Fail()
    }
    if buffer[0].filled != true {
        t.Fail()
    }
    if buffer[0].color != Black {
        t.Fail()
    }
    if buffer[0].piece != Rook {
        t.Fail()
    }
    if buffer[7].filled != false {
        t.Fail()
    }
}
