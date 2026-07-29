package main

import "testing"

func TestSplitString(t *testing.T) {
    var result []string

    result = SplitString("")
    if len(result) != 0 {
        t.Fail()
    }

    result = SplitString("a")
    if len(result) != 1 {
        t.Fail()
    }
    if (result[0] != "a") {
        t.Fail()
    }

    result = SplitString("ab")
    if len(result) != 1 {
        t.Fail()
    }
    if (result[0] != "ab") {
        t.Fail()
    }

    result = SplitString("   a b")
    if len(result) != 2 {
        t.Fail()
    }
    if (result[0] != "a") {
        t.Fail()
    }
    if (result[1] != "b") {
        t.Fail()
    }

    result = SplitString("a  bc  d e  f ghi ")
    if len(result) != 6 {
        t.Fail()
    }
    if (result[0] != "a") {
        t.Fail()
    }
    if (result[1] != "bc") {
        t.Fail()
    }
    if (result[2] != "d") {
        t.Fail()
    }
    if (result[3] != "e") {
        t.Fail()
    }
    if (result[4] != "f") {
        t.Fail()
    }
    if (result[5] != "ghi") {
        t.Fail()
    }
}
