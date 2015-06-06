package gondorcet

import (
	"fmt"
	"strings"
)

// AddCandidates adds a variadic set of candidates and returns a tuple of candidates and results
func AddCandidates(candidates ...string) ([]string, map[string]int) {
	candidateSlice := make([]string, 0)
	resultMap := make(map[string]int)
	for _, candidate := range candidates {
		// TODO: if candidate already in, throw error or skip
		candidateSlice = append(candidateSlice, candidate)
		resultMap[candidate] = 0
	}
	return candidateSlice, resultMap
}

// Contains determines whether a key is in a list
func Contains(key string, list []string) bool {
	for _, element := range list {
		if element == key {
			return true
		}
	}
	return false
}

// CheckVoteLegality determines a inputted vote is legal
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

// AddCandidate adds a candidate to an election
func (e *Election) AddCandidate(candidate string) {
	// TODO: if candidate already in candidates, throw error or skip
	e.candidates = append(e.candidates, candidate)
	e.results[candidate] = 0
}

// AddVote adds a vote to an election
func (e *Election) AddVote(vote string) {
	legal := e.CheckVoteLegality(vote)

	if legal {
		e.votes = append(e.votes, vote)
	} else {
		fmt.Println("throwing out illegal vote")
	}
}

// ClearResults clears the results of an election
func (e *Election) ClearResults() {
	for candidate := range e.results {
		e.results[candidate] = 0
	}
}

// SetWinner sets the winner in an election
func (e *Election) SetWinner() {
	// TODO: detect ties and take 2nd-order approach
	maximum := 0
	for candidate, wins := range e.results {
		if wins > maximum {
			maximum = wins
			e.winner = candidate
		}
	}
}