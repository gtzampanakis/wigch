package main

import (
    "errors"
    "fmt"
    "unicode"
)

func utilDecodeInteger(numberString string) (int, error) {
    var sign = 1
    var acc = 0
    var err error = nil
    var wasSignParsed = false
    var emptyInput = true
    var shouldSetErr = false

    for _, r := range(numberString) {
        emptyInput = false
        if unicode.IsSpace(r) {
            continue
        } else if !wasSignParsed && r == '-' {
            sign = -1
            wasSignParsed = true
        } else if r == '0' {
            acc *= 10; acc += 0;
        } else if r == '1' {
            acc *= 10; acc += 1;
        } else if r == '2' {
            acc *= 10; acc += 2;
        } else if r == '3' {
            acc *= 10; acc += 3;
        } else if r == '4' {
            acc *= 10; acc += 4;
        } else if r == '5' {
            acc *= 10; acc += 5;
        } else if r == '6' {
            acc *= 10; acc += 6;
        } else if r == '7' {
            acc *= 10; acc += 7;
        } else if r == '8' {
            acc *= 10; acc += 8;
        } else if r == '9' {
            acc *= 10; acc += 9;
        } else {
            shouldSetErr = true
            break
        }
    }
    if (emptyInput) { shouldSetErr = true; }

    if (shouldSetErr) {
        err = errors.New(fmt.Sprintf("Unable to parse integer from string: \"%s\"", numberString))
    }

    acc *= sign
    return acc, err
}

func utilIsWhitespace(r rune) bool {
    if r == ' ' {
        return true
    }
    if r == '\n' {
        return true
    }
    if r == '\r' {
        return true
    }
    return false
}
