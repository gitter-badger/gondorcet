package gondorcet

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTallyCopeland(t *testing.T) {
	Convey("Given an election", t, func() {
		election := NewElection("A", "B", "C")

		Convey("Determine the obvious winner", func() {
			for i := 0; i < 100; i++ {
				election.AddVote("A>B>C")
			}

			election.TallyCopeland()

			So(election.GetWinner(), ShouldEqual, "A")
		})
	})
}