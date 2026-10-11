package main

import (
    "testing"
)

func TestCastlingAvailability(t *testing.T) {
    var castAv CastlingAvailability
    var bytesConsumed int
    var err error

    bytesConsumed, err = castAv.initCastlingAvailabilityFromFen("- ", 0)
    if (err != nil) { t.Fail() }
    if bytesConsumed != 2 { t.Fail() }
    if castAv[0][0] != false { t.Fail() }
    if castAv[0][1] != false { t.Fail() }
    if castAv[1][0] != false { t.Fail() }
    if castAv[1][1] != false { t.Fail() }

    bytesConsumed, err = castAv.initCastlingAvailabilityFromFen("-k ", 0)
    if (err == nil) { t.Fail() }

    bytesConsumed, err = castAv.initCastlingAvailabilityFromFen("kqKQ ", 0)
    if (err != nil) { t.Fail() }
    if bytesConsumed != 5 { t.Fail() }
    if castAv[0][0] != true { t.Fail() }
    if castAv[0][1] != true { t.Fail() }
    if castAv[1][0] != true { t.Fail() }
    if castAv[1][1] != true { t.Fail() }
}
