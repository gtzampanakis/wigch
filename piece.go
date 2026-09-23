package main

type Piece int

const (
    Pawn    = 1 << iota
    Knight  = 1 << iota
    Bishop  = 1 << iota
    Rook    = 1 << iota
    Queen   = 1 << iota
    King    = 1 << iota
)
