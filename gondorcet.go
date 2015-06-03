package gondorcet

import (
	"fmt"
	"strings"
)

type Election struct {
	candidates []string
	votes []string
	delimiter string

	winner string
	results map[string]int
}

func NewElection(candidates ...string) *Election {
	candidateSlice, resultSlice := AddCandidates(candidates...)
	return &Election{delimiter: ">", winner: "", candidates: candidateSlice, results: resultSlice}
}

func NewElectionWithDelimiter(delimiter string, candidates ...string) *Election {
	candidateSlice, resultSlice := AddCandidates(candidates...)
	return &Election{delimiter: delimiter, winner: "", candidates: candidateSlice, results: resultSlice}
}

func Contains(key string, slice []string) bool {
	for _, element := range slice {
		if element == key {
			return true
		}
	}
	return false
}

func AddCandidates(candidates ...string) ([]string, map[string]int) {
	candidateSlice := make([]string, 0)
	resultSlice := make(map[string]int)
	for _, candidate := range candidates {
		// TODO: if candidate already in, throw error or skip
		candidateSlice = append(candidateSlice, candidate)
		resultSlice[candidate] = 0
	}
	return candidateSlice, resultSlice
}

func (e Election) GetCandidates() []string {
	return e.candidates
}

func (e Election) GetVotes() []string {
	return e.votes
}

func (e Election) GetWinner() string {
	return e.winner
}

func (e Election) GetResults() map[string]int {
	return e.results
}

func (e Election) CheckVoteLegality(vote string) bool {
	voteSlice := strings.Split(vote, e.delimiter)

	var previous string
	if len(voteSlice) <= len(e.candidates) {
		for _, candidate := range voteSlice {
			// if candidate not in list
			if Contains(candidate, e.candidates) == false {
				return false
			}

			// if vote contains duplicates
			if candidate == previous {
				return false
			}

			// set previous
			previous = candidate
		}
	} else {
		return false		
	}

	return true
}

func (e *Election) AddCandidate(candidate string) {
	// TODO: if candidate already in candidates, throw error or skip
	e.candidates = append(e.candidates, candidate)
	e.results[candidate] = 0
}

func (e *Election) AddVote(vote string) {
	legal := e.CheckVoteLegality(vote)

	if legal {
		e.votes = append(e.votes, vote)
	} else {
		fmt.Println("throwing out illegal vote")
	}
}

// TODO: detect ties and take 2nd-order approach
func (e *Election) setWinner() {
	maximum := 0
	for candidate, wins := range e.results {
		if wins > maximum {
			maximum = wins
			e.winner = candidate
		}
	}
}

// TODO:
// func TallyKemenyYoung()
// func TallyMinimax()
// func TallyNanson()
// func TallyDodgson()
// func TallySchulze()
func (e *Election) TallyCopeland() { 
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

	e.setWinner()
}