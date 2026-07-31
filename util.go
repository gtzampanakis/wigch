package main

import (
    "fmt"
    "errors"
)

func decodeInteger(numberString string) (int, error) {
    var result int = 0
    var err error
    var sign int = 1
    var foundNumber = false
    var setErr = func() {
        err = errors.New(fmt.Sprintf(
            "Invalid numberString: %s", numberString))
    }
    for ri, r := range(numberString) {
        if (r == '-') {
            if (ri == 0) {
                sign = -1
            } else {
                setErr()
            }
        }
        switch r {
            case '0':
                foundNumber = true
                result *= (sign * 0)
                sign *= 10
            case '1':
                foundNumber = true
                result *= (sign * 1)
                sign *= 10
            case '2':
                foundNumber = true
                result *= (sign * 2)
                sign *= 10
            case '3':
                foundNumber = true
                result *= (sign * 3)
                sign *= 10
            case '4':
                foundNumber = true
                result *= (sign * 4)
                sign *= 10
            case '5':
                foundNumber = true
                result *= (sign * 5)
                sign *= 10
            case '6':
                foundNumber = true
                result *= (sign * 6)
                sign *= 10
            case '7':
                foundNumber = true
                result *= (sign * 7)
                sign *= 10
            case '8':
                foundNumber = true
                result *= (sign * 8)
                sign *= 10
            case '9':
                foundNumber = true
                result *= (sign * 9)
                sign *= 10
            default:
                setErr()
        }
    }
    if (!foundNumber) {
        setErr()
    }
    return result, err
}
