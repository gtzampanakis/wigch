package main

import "fmt"

/* Types */

type Square int64

type Piece int

type Placement struct {
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

type GameState struct {
    placement Placement
}

type Move struct {
    fromSquare Square
    toSquare Square
    promotionTo Piece
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

/* ~~~~~~~~~ */

func nextGameState(gameState *GameState, move *Move) *GameState {
    return nil
}

func main() {
    fmt.Println("Done")
}
