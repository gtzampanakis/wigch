package main

import "fmt"

/* Types */

type square int64

type piece int

type placement struct {
    occupiedBits int64

    whiteBits    int64
    blackBits    int64

    pawnBits     int64
    rookBits     int64
    knightBits   int64
    bishopBits   int64
    queenBits    int64
    kingBits     int64
}

type color int

type castlingAvailability int

type gameState struct {
    placement placement
    activeColor color
    castlingAvailability castlingAvailability
    enPassantSquare square
    halfmoveClock int
    fullmoveNumber int
}

type move struct {
    fromSquare square
    toSquare square
    promotionTo piece
}

/* ~~~~~ */

/* Constants */

const (
    PIECE_PAWN   = 1 << iota
    PIECE_ROOK   = 1 << iota
    PIECE_KNIGHT = 1 << iota
    PIECE_BISHOP = 1 << iota
    PIECE_QUEEN  = 1 << iota
    PIECE_KING   = 1 << iota
)

const (
    COLOR_WHITE = 0
    COLOR_BLACK = 1
)

const (
    CASTLING_WHITE_KING_SIDE =  1 << iota
    CASTLING_WHITE_QUEEN_SIDE = 1 << iota
    CASTLING_BLACK_KING_SIDE =  1 << iota
    CASTLING_BLACK_QUEEN_SIDE = 1 << iota
)
    

/* ~~~~~~~~~ */

func main() {
    fmt.Println("Done")
}
