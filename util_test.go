package gondorcet

import (
	"strings"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAddCandidates(t *testing.T) {
	for i := 0; i < len(candidates); i++ {
		Convey("Given a candidate", t, func() {
			candidateSubset := candidates[0:i]

			Convey("When the candidate is added using AddCandidates", func() {
				cs, rs := AddCandidates(candidateSubset...)

				Convey("The candidates and results should be", func() {
					So(cs, ShouldResemble, candidateSubset)

					expected := make(map[string]int)
					for _, candidate := range candidateSubset {
						expected[candidate] = 0
					}
					So(rs, ShouldResemble, expected)
				})
			})
		})
	}
}

func TestContains(t *testing.T) {
	Convey("Given a list of candidates", t, func() {
		// true
		So(Contains("Ben Carson", candidates), ShouldBeTrue)
		So(Contains("Hillary Clinton", candidates), ShouldBeTrue)
		So(Contains("Rand Paul", candidates), ShouldBeTrue)

		// false
		So(Contains("Barack Obama", candidates), ShouldBeFalse)

		// type check
		So(Contains("Bobby Jindal", candidates), ShouldNotHaveSameTypeAs, "string")
		So(Contains("Bobby Jindal", candidates), ShouldNotHaveSameTypeAs, 0)
	})
}

func TestCheckVoteLegality(t *testing.T) {
	Convey("Given elections of differing delimiters", t, func() {
		election := NewElection(candidates...)
		electionComma := NewElectionWithDelimiter(",", candidates...)
		electionPipe := NewElectionWithDelimiter("|", candidates...)

		Convey("When a vote is legal", func() {
			for i := 1; i < len(candidates); i++ {
				voteSlice := candidates[0:i]
				vote := strings.Join(voteSlice, ">")
				voteComma := strings.Join(voteSlice, ",")
				votePipe := strings.Join(voteSlice, "|")

				So(election.CheckVoteLegality(vote), ShouldBeTrue)
				So(electionComma.CheckVoteLegality(voteComma), ShouldBeTrue)
				So(electionPipe.CheckVoteLegality(votePipe), ShouldBeTrue)
			}
		})

		Convey("When a vote is illegal", func() {
			voteExtra := strings.Join(append(candidates, "Barack Obama"), ">")
			voteDuplicate := strings.Join(append(candidates, candidates...), ">")

			So(election.CheckVoteLegality(voteExtra), ShouldBeFalse)
			So(election.CheckVoteLegality(voteDuplicate), ShouldBeFalse)
			
			voteCommaExtra := strings.Join(append(candidates, "Barack Obama"), ",")
			voteCommaDuplicate := strings.Join(append(candidates, candidates...), ",")

			So(electionComma.CheckVoteLegality(voteCommaExtra), ShouldBeFalse)
			So(electionComma.CheckVoteLegality(voteCommaDuplicate), ShouldBeFalse)
			
			votePipeExtra := strings.Join(append(candidates, "Barack Obama"), "|")
			votePipeDuplicate := strings.Join(append(candidates, candidates...), "|")

			So(electionPipe.CheckVoteLegality(votePipeExtra), ShouldBeFalse)
			So(electionPipe.CheckVoteLegality(votePipeDuplicate), ShouldBeFalse)
		})
	})
}

func TestAddCandidate(t *testing.T) {
	Convey("Given elections of differing delimiters", t, func() {
		election := NewElection()
		electionComma := NewElectionWithDelimiter(",")
		electionPipe := NewElectionWithDelimiter("|")

		Convey("AddCandidates should correctly add candidates", func() {
			for _, candidate := range candidates {
				election.AddCandidate(candidate)
				electionComma.AddCandidate(candidate)
				electionPipe.AddCandidate(candidate)
				store := election.GetCandidates()
				storeComma := electionComma.GetCandidates()
				storePipe := electionPipe.GetCandidates()
				So(store[len(store)-1], ShouldResemble, candidate)
				So(storeComma[len(store)-1], ShouldResemble, candidate)
				So(storePipe[len(store)-1], ShouldResemble, candidate)
			}
		})
	})
}

func TestAddVote(t *testing.T) {
	Convey("Given elections of differing delimiters", t, func() {
		election := NewElection(candidates...)
		electionComma := NewElectionWithDelimiter(",", candidates...)
		electionPipe := NewElectionWithDelimiter("|", candidates...)

		Convey("AddVote should correctly add votes to an election", func() {
			voteSlice := make([]string, 0)
			for _, candidate := range candidates {
				voteSlice = append(voteSlice, candidate)
				vote := strings.Join(voteSlice, ">")
				voteComma := strings.Join(voteSlice, ",")
				votePipe := strings.Join(voteSlice, "|")
				election.AddVote(vote)
				electionComma.AddVote(voteComma)
				electionPipe.AddVote(votePipe)
				store := election.GetVotes()
				storeComma := electionComma.GetVotes()
				storePipe := electionPipe.GetVotes()
				So(store[len(store)-1], ShouldResemble, vote)
				So(storeComma[len(store)-1], ShouldResemble, voteComma)
				So(storePipe[len(store)-1], ShouldResemble, votePipe)
			}
		})
	})
}

func TestClearResults(t *testing.T) {
	Convey("Given elections of differing delimiters", t, func() {
		election := NewElection(candidates...)
		electionComma := NewElectionWithDelimiter(",", candidates...)
		electionPipe := NewElectionWithDelimiter("|", candidates...)

		for candidate := range election.results {
			election.results[candidate] = 1
			electionComma.results[candidate] = 1
			electionPipe.results[candidate] = 1
		}

		Convey("ClearResults should correctly clear results of an election", func() {
			election.ClearResults()
			electionComma.ClearResults()
			electionPipe.ClearResults()

			results := []map[string]int{election.results, electionComma.results, electionPipe.results}

			for _, result := range results {
				for _, wins := range result {
					So(wins, ShouldEqual, 0)
				}
			}
		})
	})
}

func TestSetWinner(t *testing.T) {
	Convey("Given an election", t, func() {
		election := NewElection(candidates...)

		i := 1
		for candidate := range election.results {
			election.results[candidate] = i
			i++
		}

		Convey("The winner should be properly set", func() {
			election.SetWinner()

			// TODO: not a good test
			maximum := 0
			var winner string
			for candidate, wins := range election.results {
				if wins > maximum {
					maximum = wins
					winner = candidate
				}
			}

			So(election.GetWinner(), ShouldEqual, winner)
		})
	})
}