package main

import (
    "fmt"
)

type Placement struct {
    contents [nFiles]Rank
}

/* rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 */

/* parsePlacementFromFen parses a placement from the passed fen. Starts
 * parsing from fenOffset and continues until encountering a space. Writes
 * the placement into the passed placement. Returns number of bytes
 * consumed or an error. */
func parsePlacementFromFen(fen string, fenOffset int, placement *Placement) (bytesConsumed int, err error) {
    if fenOffset < 0 {
        err = fmt.Errorf("fenOffset must be nonnegative")
        return 0, err
    }
    /* rankI is the index of the rank being filled. */
    var rankI = 0
    /* Repeatedly call parseRankFromFen. */
    for {
        var isLastRank = false
        var endByte byte
        if rankI < nRanks - 1 {
            /* If it is not the last rank that is being filled instruct
             * parseRankFromFen to treat '/' as the end byte.*/
            endByte = '/'
        } else if rankI == nRanks - 1 {
            /* If it is the last rank that is being filled instruct
             * parseRankFromFen to treat ' ' as the end byte. */
            isLastRank = true
            endByte = ' '
        } else {
            err = fmt.Errorf("placement has too many ranks")
            return 0, err
        }

        /* Keep calling parseRankFromFen and incrementing rankI and
         * bytesConsumed.  */
        var bytesConsumedByThisParseRankCall, errByParseRank =
            parseRankFromFen(fen, fenOffset + bytesConsumed, endByte, &placement.contents[rankI])
        if (errByParseRank != nil) {
            return 0, errByParseRank
        }
        rankI++
        bytesConsumed += bytesConsumedByThisParseRankCall
        if isLastRank {
            break
        }
    }
    /* The loop that calls parseRankFromFen runs exactly nRanks times.
     * Therefore there is no need to check that nRanks ranks were parsed. */
     return bytesConsumed, nil
}
