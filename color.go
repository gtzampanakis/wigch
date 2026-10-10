package main

import (
    "fmt"
)

type Color int

const (
    White = 0
    Black = 1
)

/* initColorFromFen initializes a color from the passed fen. Starts parsing
 * from fenOffset and attempts to read two bytes. Returns bytesConsumed or an
 * error. */
func (color *Color) initColorFromFen(fen string, fenOffset int) (bytesConsumed int, err error) {
    if fenOffset < 0 {
        err = fmt.Errorf("fenOffset must be nonnegative")
        return
    }
    if fenOffset < len(fen) {
        /* The first byte must be 'w' or 'b'. */
        var b = fen[fenOffset]
        fenOffset++
        bytesConsumed++
        if b == 'w' {
            *color = White
        } else if b == 'b' {
            *color = Black
        } else {
            err = fmt.Errorf("Unexpected FEN character: '%c'", b)
            return
        }
        if fenOffset < len(fen) {
            /* The second byte must be a space. */
            var b = fen[fenOffset]
            fenOffset++
            bytesConsumed++
            if b == ' ' {
                return
            } else {
                err = fmt.Errorf("Unexpected FEN character: '%c'", b)
                return
            }
        } else {
            err = fmt.Errorf("FEN active color end byte not found")
            return
        }
    } else {
        err = fmt.Errorf("Empty FEN")
        return
    }
}
