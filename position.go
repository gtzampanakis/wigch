package main

import (
    "fmt"
)

type Position struct {
    placement Placement 
    activeColor Color 
    castlingAvailability CastlingAvailability 
    enPassantTargetSquare Square 
    halfMoveClock int 
    fullMoveNumber int 
}

type FenParsingError struct {
    msg string
}

func (self FenParsingError) Error() string {
    return self.msg
}


/* 
    Fen Examples (initial position and the sequence occurring from 1. e4 c5 2.
    Nf3):

    rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
    rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1
    rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2
    rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2
*/

func ParseFenByte(b byte, buffer *[nFiles]SquareContent, bufferIndex int) (int, error) {
    var err error

    if b == 'p' {
        buffer[bufferIndex] = SquareContent{filled: true, color: Black, piece: Pawn}
        bufferIndex++
    } else if b == 'r' {
        buffer[bufferIndex] = SquareContent{filled: true, color: Black, piece: Rook}
        bufferIndex++
    } else if b == 'n' {
        buffer[bufferIndex] = SquareContent{filled: true, color: Black, piece: Knight}
        bufferIndex++
    } else if b == 'b' {
        buffer[bufferIndex] = SquareContent{filled: true, color: Black, piece: Bishop}
        bufferIndex++
    } else if b == 'q' {
        buffer[bufferIndex] = SquareContent{filled: true, color: Black, piece: Queen}
        bufferIndex++
    } else if b == 'k' {
        buffer[bufferIndex] = SquareContent{filled: true, color: Black, piece: King}
        bufferIndex++
    } else if b == 'P' {
        buffer[bufferIndex] = SquareContent{filled: true, color: White, piece: Pawn}
        bufferIndex++
    } else if b == 'R' {
        buffer[bufferIndex] = SquareContent{filled: true, color: White, piece: Rook}
        bufferIndex++
    } else if b == 'B' {
        buffer[bufferIndex] = SquareContent{filled: true, color: White, piece: Bishop}
        bufferIndex++
    } else if b == 'Q' {
        buffer[bufferIndex] = SquareContent{filled: true, color: White, piece: Queen}
        bufferIndex++
    } else if b == 'K' {
        buffer[bufferIndex] = SquareContent{filled: true, color: White, piece: King}
        bufferIndex++
    } else if b == '1' {
        for range(1) {
            if bufferIndex < len(buffer) {
                buffer[bufferIndex] = SquareContent{filled: false}
                bufferIndex++
            } else {
                err = FenParsingError{msg: fmt.Sprintf("Out of space in buffer")}
            }
        }
    } else if b == '2' {
        for range(2) {
            if bufferIndex < len(buffer) {
                buffer[bufferIndex] = SquareContent{filled: false}
                bufferIndex++
            } else {
                err = FenParsingError{msg: fmt.Sprintf("Out of space in buffer")}
            }
        }
    } else if b == '3' {
        for range(3) {
            if bufferIndex < len(buffer) {
                buffer[bufferIndex] = SquareContent{filled: false}
                bufferIndex++
            } else {
                err = FenParsingError{msg: fmt.Sprintf("Out of space in buffer")}
            }
        }
    } else if b == '4' {
        for range(4) {
            if bufferIndex < len(buffer) {
                buffer[bufferIndex] = SquareContent{filled: false}
                bufferIndex++
            } else {
                err = FenParsingError{msg: fmt.Sprintf("Out of space in buffer")}
            }
        }
    } else if b == '5' {
        for range(5) {
            if bufferIndex < len(buffer) {
                buffer[bufferIndex] = SquareContent{filled: false}
                bufferIndex++
            } else {
                err = FenParsingError{msg: fmt.Sprintf("Out of space in buffer")}
            }
        }
    } else if b == '6' {
        for range(6) {
            if bufferIndex < len(buffer) {
                buffer[bufferIndex] = SquareContent{filled: false}
                bufferIndex++
            } else {
                err = FenParsingError{msg: fmt.Sprintf("Out of space in buffer")}
            }
        }
    } else if b == '7' {
        for range(7) {
            if bufferIndex < len(buffer) {
                buffer[bufferIndex] = SquareContent{filled: false}
                bufferIndex++
            } else {
                err = FenParsingError{msg: fmt.Sprintf("Out of space in buffer")}
            }
        }
    } else if b == '8' {
        for range(8) {
            if bufferIndex < len(buffer) {
                buffer[bufferIndex] = SquareContent{filled: false}
                bufferIndex++
            } else {
                err = FenParsingError{msg: fmt.Sprintf("Out of space in buffer")}
            }
        }
    } else {
        err = FenParsingError{msg: fmt.Sprintf("Invalid FEN byte: '%c'", b)}
    }
    return bufferIndex, err
}

func parseFenOneRank(fen string, fenIndex int, buffer *[nFiles]SquareContent) (int, error) {
    var bufferIndex = 0
    var err error

    for ; fenIndex < len(fen); fenIndex++ {
        b := fen[fenIndex]
        if b == '/' {
            break
        }
        bufferIndex, err = ParseFenByte(b, buffer, bufferIndex)
        if err != nil {
            return bufferIndex, err
        }
    }
    if bufferIndex < nFiles {
        err = FenParsingError{msg: fmt.Sprintf("FEN rank too short")}
        return fenIndex, err
    }
    if bufferIndex > nFiles {
        err = FenParsingError{msg: fmt.Sprintf("FEN rank too long")}
        return fenIndex, err
    }
    return fenIndex, err
}

// func parsePlacement(fen string) (Placement, string, error) {
// }
// 
// func parseFen(fen string) Position {
//     fen = parsePlacement(fen)
//     fen = parseActiveColor(fen)
//     fen = parseCastlingAvailability(fen)
//     fen = parseEnPassantTargetSquare(fen)
//     fen = parseHalfMoveClock(fen)
//     fen = parseFullMoveNumber(fen)
// }
