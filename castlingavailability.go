package main

import (
    "fmt"
)

/* CastlingAvailability is a two-by-two array with the first index being Color
 * (White or Black) and the second being Side (KingSide or QueenSide). */
type CastlingAvailability [2][2]bool

func (castAv *CastlingAvailability) initCastlingAvailabilityFromFen(
    fen string,
    fenOffset int,
) (bytesConsumed int, err error) {
    var byteI = fenOffset
    if fenOffset < 0 {
        err = fmt.Errorf("fenOffset must be nonnegative")
    } else {
        var done = false
        var sawDash = false
        for ; byteI < len(fen); byteI++ {
            if done {
                break
            }
            var b byte = fen[byteI]
            if b == '-' {
                sawDash = true
            } else if b == ' ' {
                done = true
            } else if !sawDash && b == 'K' {
                castAv[0][0] = true
            } else if !sawDash && b == 'Q' {
                castAv[0][1] = true
            } else if !sawDash && b == 'k' {
                castAv[1][0] = true
            } else if !sawDash && b == 'q' {
                castAv[1][1] = true
            } else {
                err = fmt.Errorf("Unexpected FEN character: '%c'", b)
                done = true
            }
        }
    }
    return byteI - fenOffset, err
}
