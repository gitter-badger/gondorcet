package gondorcet

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var candidates = []string{"Ben Carson", "Ted Cruz", "Rand Paul", "Bobby Jindal", "Lincoln Chafee", "Martin O'Malley", "Bernie Sanders", "Hillary Clinton"}

func TestNewElection(t *testing.T) {
	Convey("Given an election", t, func() {
		election := NewElection()

		Convey("Members should be set correctly", func() {
			So(election.candidates, ShouldResemble, []string{})
			So(election.delimiter, ShouldEqual, ">")
			So(election.winner, ShouldEqual, "")
			So(election.votes, ShouldResemble, []string{})
			So(election.results, ShouldResemble, map[string]int{})
		})
	})

	Convey("Given an election initialized with candidates", t, func() {
		election := NewElection(candidates...)

		Convey("Members should be set correctly", func() {
			So(election.candidates, ShouldResemble, candidates)
			So(election.delimiter, ShouldEqual, ">")
			So(election.winner, ShouldEqual, "")
			So(election.votes, ShouldResemble, []string{})
			So(election.results, ShouldResemble,
				map[string]int{"Lincoln Chafee":0,
				"Martin O'Malley":0,
				"Bernie Sanders":0,
				"Hillary Clinton":0,
				"Ben Carson":0,
				"Ted Cruz":0,
				"Rand Paul":0,
				"Bobby Jindal":0})
		})
	})
}

func TestNewElectionWithDelimiter(t *testing.T) {
	Convey("Given an election with delimiter", t, func() {
		election := NewElectionWithDelimiter(",")

		Convey("Members should be set correctly", func() {
			So(election.candidates, ShouldResemble, []string{})
			So(election.delimiter, ShouldEqual, ",")
			So(election.winner, ShouldEqual, "")
			So(election.votes, ShouldResemble, []string{})
			So(election.results, ShouldResemble, map[string]int{})
		})
	})

	Convey("Given an election with delimiter initialized with candidates", t, func() {
		election := NewElectionWithDelimiter(",", candidates...)

		Convey("Members should be set correctly", func() {
			So(election.candidates, ShouldResemble, candidates)
			So(election.delimiter, ShouldEqual, ",")
			So(election.winner, ShouldEqual, "")
			So(election.votes, ShouldResemble, []string{})
			So(election.results, ShouldResemble,
				map[string]int{"Lincoln Chafee":0,
				"Martin O'Malley":0,
				"Bernie Sanders":0,
				"Hillary Clinton":0,
				"Ben Carson":0,
				"Ted Cruz":0,
				"Rand Paul":0,
				"Bobby Jindal":0})
		})
	})
}