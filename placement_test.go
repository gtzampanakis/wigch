package main

import (
    "testing"
)

func TestParsePlacementFromFen(t *testing.T) {
    var placement Placement
    var err error
    var bytesConsumed int

    bytesConsumed, err =
        placement.initPlacementFromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP ", 0)
    if err == nil { t.Fail() }
    if err.Error() != "Unexpected FEN character: ' '" { t.Fail() }

    bytesConsumed, err =
        placement.initPlacementFromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR ", 0)
    if err != nil { t.Fail() }
    if bytesConsumed != 44 { t.Fail() }
}

// func FuzzParsePlacementFromFen(f *testing.F) {
// }
