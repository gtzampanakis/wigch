package main

import (
    "testing"
)

func TestActiveColorParsing(t *testing.T) {
    var color Color
    var bytesConsumed int
    var err error

    bytesConsumed, err = color.initColorFromFen("w ", 0)
    if err != nil { t.Fail() }
    if bytesConsumed != 2 { t.Fail() }
    if color != White { t.Fail() }

    bytesConsumed, err = color.initColorFromFen("b ", 0)
    if err != nil { t.Fail() }
    if bytesConsumed != 2 { t.Fail() }
    if color != Black { t.Fail() }

    bytesConsumed, err = color.initColorFromFen("x ", 0)
    if err == nil { t.Fail() }

    bytesConsumed, err = color.initColorFromFen("w", 0)
    if err == nil { t.Fail() }

    bytesConsumed, err = color.initColorFromFen("", 0)
    if err == nil { t.Fail() }

    bytesConsumed, err = color.initColorFromFen("ww", 0)
    if err == nil { t.Fail() }
}
