package main

import (
    "fmt"
)

/* rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 */
type Rank struct {
    contents [nFiles]SquareContent
}

func fenByteToFilledSquareContent(b byte) (matched bool, result SquareContent) {
    matched = true
    if b == 'p' {
        result = SquareContent{ filled: true, piece: Pawn, color: Black }
    } else if b == 'r' {
        result = SquareContent{ filled: true, piece: Rook, color: Black }
    } else if b == 'n' {
        result = SquareContent{ filled: true, piece: Knight, color: Black }
    } else if b == 'b' {
        result = SquareContent{ filled: true, piece: Bishop, color: Black }
    } else if b == 'q' {
        result = SquareContent{ filled: true, piece: Queen, color: Black }
    } else if b == 'k' {
        result = SquareContent{ filled: true, piece: King, color: Black }
    } else if b == 'P' {
        result = SquareContent{ filled: true, piece: Pawn, color: White }
    } else if b == 'R' {
        result = SquareContent{ filled: true, piece: Rook, color: White }
    } else if b == 'N' {
        result = SquareContent{ filled: true, piece: Knight, color: White }
    } else if b == 'B' {
        result = SquareContent{ filled: true, piece: Bishop, color: White }
    } else if b == 'Q' {
        result = SquareContent{ filled: true, piece: Queen, color: White }
    } else if b == 'K' {
        result = SquareContent{ filled: true, piece: King, color: White }
    } else {
        matched = false
    }
    return
}

func fenByteToInt(b byte) (matched bool, result int) {
    matched = true
    if b == '1' {
        result = 1
    } else if b == '2' {
        result = 2
    } else if b == '3' {
        result = 3
    } else if b == '4' {
        result = 4
    } else if b == '5' {
        result = 5
    } else if b == '6' {
        result = 6
    } else if b == '7' {
        result = 7
    } else if b == '8' {
        result = 8
    } else {
        matched = false
    }
    return
}

/* parseRankFromFen parses one rank from the passed fen. Starts parsing
 * from fenOffset and continues until encountering endByte. Writes the
 * rank into the passed rank.  Returns number of bytes consumed or an
 * error. */
func parseRankFromFen(fen string, fenOffset int, endByte byte, rank *Rank) (bytesConsumed int, err error) {
    if fenOffset < 0 {
        err = fmt.Errorf("fenOffset must be nonnegative")
        return
    }
    var byteI = fenOffset
    var fileI = 0
    var foundEndByte = false
    for ; byteI < len(fen); byteI++ {
        var b byte = fen[byteI]
        if matched, squareContent := fenByteToFilledSquareContent(b); matched {
            if fileI < nFiles {
                rank.contents[fileI] = squareContent
                fileI++
            } else {
                err = fmt.Errorf("FEN rank too long")
                break
            }
        } else if matched, parsedInt := fenByteToInt(b); matched {
            for i := 0; i < parsedInt; i++ {
                if fileI < nFiles {
                    rank.contents[fileI] = SquareContent{
                        filled: false,
                        piece: Pawn,
                        color: White,
                    }
                    fileI++
                } else {
                    err = fmt.Errorf("FEN rank too long")
                    break
                }
            }
        } else if b == endByte {
            foundEndByte = true
            byteI++
            break
        } else {
            err = fmt.Errorf("Unexpected FEN character: '%c'", b)
            break
        }
    }
    if err == nil {
        if !foundEndByte {
            err = fmt.Errorf("FEN rank end byte not found")
        } else if fileI < nFiles {
            err = fmt.Errorf("FEN rank too short")
        }
    }
    return byteI - fenOffset, err
}
