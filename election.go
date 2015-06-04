package gondorcet

// Election consists of a slice of candidates, slice of votes, a delimiter, a winner, and results
type Election struct {
	candidates []string
	votes []string
	delimiter string

	winner string
	results map[string]int
}

// GetCandidates gets the candidates
func (e Election) GetCandidates() []string {
	return e.candidates
}

// GetVotes gets the votes
func (e Election) GetVotes() []string {
	return e.votes
}

// GetWinner gets the winner
func (e Election) GetWinner() string {
	return e.winner
}

// GetResults gets the results
func (e Election) GetResults() map[string]int {
	return e.results
}

// NewElection constructs a new election using the default '>' delimiter
func NewElection(candidates ...string) *Election {
	candidateSlice, resultSlice := AddCandidates(candidates...)
	return &Election{delimiter: ">", winner: "", candidates: candidateSlice, results: resultSlice}
}

// NewElection constructs a new election using the user-specified delimiter
func NewElectionWithDelimiter(delimiter string, candidates ...string) *Election {
	candidateSlice, resultSlice := AddCandidates(candidates...)
	return &Election{delimiter: delimiter, winner: "", candidates: candidateSlice, results: resultSlice}
}