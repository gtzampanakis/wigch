package main

import (
    "errors"
    "fmt"
)

func parseNotationSquare(notationSquare string) (square, error) {
    var result square
    var err error
    var setErr = func() {
        err = errors.New(
            fmt.Sprintf("Invalid Notation Square: %s", notationSquare))
    }
    for ri, r := range notationSquare {
        if ri == 0 {
            switch r {
                case 'a': result = 1 << 0
                case 'b': result = 1 << 1
                case 'c': result = 1 << 2
                case 'd': result = 1 << 3
                case 'e': result = 1 << 4
                case 'f': result = 1 << 5
                case 'g': result = 1 << 6
                case 'h': result = 1 << 7
                default:
                    setErr()
                    break
            }
        } else if ri == 1 {
            switch r {
                case '1': result <<= (8 * 0)
                case '2': result <<= (8 * 1)
                case '3': result <<= (8 * 2)
                case '4': result <<= (8 * 3)
                case '5': result <<= (8 * 4)
                case '6': result <<= (8 * 5)
                case '7': result <<= (8 * 6)
                case '8': result <<= (8 * 7)
                default:
                    setErr()
                    break
            }
        } else {
            setErr()
            break
        }
    }
    return result, err
}
