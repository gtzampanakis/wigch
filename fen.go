package main

import "errors"

func decodePlacement(placementFen string) (*placement, error) {
    var placement placement
    var mask int64 = 1
    for _, r := range placementFen {
        switch r {
            case '1': mask <<= 1
            case '2': mask <<= 2
            case '3': mask <<= 3
            case '4': mask <<= 4
            case '5': mask <<= 5
            case '6': mask <<= 6
            case '7': mask <<= 7
            case '8': mask <<= 8
            case 'r':
                {
                    placement.occupiedBits |= mask
                    placement.blackBits |= mask
                    placement.rookBits |= mask
                    mask <<= 1
                }
            case 'R':
                {
                    placement.occupiedBits |= mask
                    placement.whiteBits |= mask
                    placement.rookBits |= mask
                    mask <<= 1
                }
            case 'n':
                {
                    placement.occupiedBits |= mask
                    placement.blackBits |= mask
                    placement.knightBits |= mask
                    mask <<= 1
                }
            case 'N':
                {
                    placement.occupiedBits |= mask
                    placement.whiteBits |= mask
                    placement.knightBits |= mask
                    mask <<= 1
                }
            case 'b':
                {
                    placement.occupiedBits |= mask
                    placement.blackBits |= mask
                    placement.bishopBits |= mask
                    mask <<= 1
                }
            case 'B':
                {
                    placement.occupiedBits |= mask
                    placement.whiteBits |= mask
                    placement.bishopBits |= mask
                    mask <<= 1
                }
            case 'Q':
                {
                    placement.occupiedBits |= mask
                    placement.blackBits |= mask
                    placement.queenBits |= mask
                    mask <<= 1
                }
            case 'q':
                {
                    placement.occupiedBits |= mask
                    placement.whiteBits |= mask
                    placement.queenBits |= mask
                    mask <<= 1
                }
            case 'k':
                {
                    placement.occupiedBits |= mask
                    placement.blackBits |= mask
                    placement.kingBits |= mask
                    mask <<= 1
                }
            case 'K':
                {
                    placement.occupiedBits |= mask
                    placement.whiteBits |= mask
                    placement.kingBits |= mask
                    mask <<= 1
                }
            case ' ', '/':
            default:
                return nil, errors.New("Invalid FEN placement")
        }
    }
    return &placement, nil
}

func decodeActiveColor(activeColorFen string) (*color, error) {
    var result color
    var err error
    var r = []rune(activeColorFen)[0]
    switch r {
        case 'w':
            result = COLOR_WHITE
        case 'b':
            result = COLOR_BLACK
        default:
            err = errors.New("Invalid FEN activeColor")
    }
    return &result, err
}

func decodeCastlingAvailability(castlingAvailabilityFen string) (*castlingAvailability, error) {
    var result castlingAvailability
    var err error
    var foundDash = false
    var setErr = func() {
        err = errors.New("Invalid FEN castling availability")
    }
    for _, r := range castlingAvailabilityFen {
        if foundDash {
            setErr()
            break
        }
        switch r {
            case '-': foundDash = true
            case 'K': result |= CASTLING_WHITE_KING_SIDE
            case 'Q': result |= CASTLING_WHITE_QUEEN_SIDE
            case 'k': result |= CASTLING_BLACK_KING_SIDE
            case 'q': result |= CASTLING_BLACK_QUEEN_SIDE
            default: setErr()
        }
    }
    return &result, err
}

func decodeEnPassantSquare(enPassantSquareFen string) (*square, error) {
    var result square
    var err error
    if enPassantSquareFen != "-" {
        result, err = parseNotationSquare(enPassantSquareFen)
        if err != nil {
            err = errors.New("Invalid FEN en passant square")
        }
    }
    return &result, err
}

func decodeHalfmoveClock(halfmoveClockFen string) (int, error) {
    return decodeInteger(halfmoveClockFen)
}

func decodeFullmoveNumber(fullmoveNumberFen string) (int, error) {
    return decodeInteger(fullmoveNumberFen)
}

func isWhitespace(r rune) bool {
    return r == ' ' || r == '\n' || r == '\t' || r == '\r' || r == '\f'
}

func splitString(s string) []string {
    var result []string
    var startI int
    var rPrev, r rune
    r = ' '
    var rs = []rune(s)
    for rsI := 0; rsI < len(rs); rsI++ {
        rPrev = r
        r = rs[rsI]
        if !isWhitespace(r) && isWhitespace(rPrev) {
            startI = rsI
        }
        if isWhitespace(r) && !isWhitespace(rPrev) {
            result = append(result, string(rs[startI:rsI]))
        }
        if !isWhitespace(r) && rsI == len(rs) - 1 {
            result = append(result, string(rs[startI:rsI+1]))
        }
    }
    return result
}

func decodeFen(fen string) (*gameState, error) {
    /* Example fen:
        1r6/5pp1/R1R4p/1r1pP3/2pkQPP1/7P/1P6/2K5 w - - 0 41 
    */
    var gameState = &gameState{}
    var split = splitString(fen)
    var err error
    var placement *placement
    var activeColor *color
    var castlingAvailability *castlingAvailability
    var enPassantSquare *square
    var halfmoveClock int
    var fullmoveNumber int

    placement, err = decodePlacement(split[0])
    if err != nil {
        return nil, err
    }

    activeColor, err = decodeActiveColor(split[1])
    if err != nil {
        return nil, err
    }

    castlingAvailability, err = decodeCastlingAvailability(split[2])
    if err != nil {
        return nil, err
    }

    castlingAvailability, err = decodeCastlingAvailability(split[2])
    if err != nil {
        return nil, err
    }

    enPassantSquare, err = decodeEnPassantSquare(split[3])
    if err != nil {
        return nil, err
    }

    halfmoveClock, err = decodeHalfmoveClock(split[4])
    if err != nil {
        return nil, err
    }

    fullmoveNumber, err = decodeFullmoveNumber(split[5])
    if err != nil {
        return nil, err
    }

    gameState.placement = *placement
    gameState.activeColor = *activeColor
    gameState.castlingAvailability = *castlingAvailability
    gameState.enPassantSquare = *enPassantSquare
    gameState.halfmoveClock = halfmoveClock
    gameState.fullmoveNumber = fullmoveNumber
    return gameState, nil
}
