package gondorcet

import (
	"strings"
)

// TallyCopeland uses Copeland's method to tally votes
func (e *Election) TallyCopeland() {
	e.ClearResults()

	for _, vote := range e.votes {
		voteSlice := strings.Split(vote, e.delimiter)
		for i := 0; i < len(voteSlice); i++ {
			for j := i + 1; j < len(voteSlice); j++ {
				// i always wins, j always loses
				e.results[voteSlice[i]]++
				e.results[voteSlice[j]]--
			}
		}
	}

	e.SetWinner()
}

/* TODO:
// TallyKemenyYoung uses Kemeny-Young method to tally votes
func (e *Election) TallyKemenyYoung() {
	
}

// TallyMinimax uses Minimax to tally votes
func (e *Election) TallyMinimax() {
	
}

// TallyNanson uses Nanson's method to tally votes
func (e *Election) TallyNanson() {
	
}

// TallyDodgson uses Dodgson's method to tally votes
func (e *Election) TallyDodgson() {

}

// TallyRankedPairs uses Ranked Pairs to tally votes
func (e *Election) TallyRankedPairs() {
	
}

// TallySchulze uses Schulze method to tally votes
func (e *Election) TallySchulze() {
	
}
*/