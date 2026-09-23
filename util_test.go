package main

import "testing"

func TestDecodeInteger(t *testing.T) {
    var res int
    var err error

    res, err = decodeInteger("")
    if (err == nil) { t.Fail() }

    res, err = decodeInteger("-0")
    if (err != nil) { t.Fail() }
    if res != 0 { t.Fail() }

    res, err = decodeInteger("0")
    if (err != nil) { t.Fail() }
    if res != 0 { t.Fail() }

    res, err = decodeInteger("1")
    if (err != nil) { t.Fail() }
    if res != 1 { t.Fail() }

    res, err = decodeInteger("12")
    if (err != nil) { t.Fail() }
    if res != 12 { t.Fail() }

    res, err = decodeInteger("-1")
    if (err != nil) { t.Fail() }
    if res != -1 { t.Fail() }

    res, err = decodeInteger("-1234567890")
    if (err != nil) { t.Fail() }
    if res != -1234567890 { t.Fail() }
}
