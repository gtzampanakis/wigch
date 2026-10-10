package main

import (
    "fmt"
)

type Placement struct {
    contents [nFiles]Rank
}

/* rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 */

/* initPlacementFromFen initializes the placement by parsing the passed fen.
 * Starts parsing from fenOffset and continues until encountering a space.
 * Returns number of bytes consumed or an error. */
func (placement *Placement) initPlacementFromFen(fen string, fenOffset int) (bytesConsumed int, err error) {
    if fenOffset < 0 {
        err = fmt.Errorf("fenOffset must be nonnegative")
        return 0, err
    }
    /* rankI is the index of the rank being filled. */
    var rankI = 0
    /* Repeatedly call initRankFromFen. */
    for {
        var isLastRank = false
        var endByte byte
        if rankI < nRanks - 1 {
            /* If it is not the last rank that is being filled instruct
             * initRankFromFen to treat '/' as the end byte.*/
            endByte = '/'
        } else if rankI == nRanks - 1 {
            /* If it is the last rank that is being filled instruct
             * initRankFromFen to treat ' ' as the end byte. */
            isLastRank = true
            endByte = ' '
        } else {
            err = fmt.Errorf("placement has too many ranks")
            return 0, err
        }

        /* Keep calling initRankFromFen and incrementing rankI and
         * bytesConsumed.  */
        var rank = placement.contents[rankI]
        var bytesConsumedByThisInitRankCall, errByInitRank =
            rank.initRankFromFen(fen, fenOffset + bytesConsumed, endByte)
        if (errByInitRank != nil) {
            return 0, errByInitRank
        }
        rankI++
        bytesConsumed += bytesConsumedByThisInitRankCall
        if isLastRank {
            break
        }
    }
    /* The loop that calls initRankFromFen runs exactly nRanks times.
     * Therefore there is no need to check that nRanks ranks were parsed. */
     return bytesConsumed, nil
}
