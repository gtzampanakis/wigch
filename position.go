package main

import (
)

type Position struct {
    placement Placement 
    activeColor Color 
    castlingAvailability CastlingAvailability 
    enPassantTargetSquare Square 
    halfMoveClock int 
    fullMoveNumber int 
}
