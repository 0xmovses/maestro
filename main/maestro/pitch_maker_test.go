package maestro

import (
	"testing"

	"github.com/rvmelkonian/maestro/main/shared"
)

func TestPitchMaker(t *testing.T) {
	pitchMaker := NewPitchMaker()

	t.Run("transposeUp should be correct", func(t *testing.T) {
		got := pitchMaker.transposeUp(0, 1)
		want := int(1)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("transposeUp should wrap around 12", func(t *testing.T) {
		got := pitchMaker.transposeUp(11, 1)
		want := int(0)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("transposeUp should wrap around 0", func(t *testing.T) {
		got := pitchMaker.transposeUp(1, 13)
		want := int(2)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("transposeSetUp should be correct", func(t *testing.T) {
		got := pitchMaker.transposeSetUp([]int{0, 1, 2}, 1)
		want := []int{1, 2, 3}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetUp should wrap around 12", func(t *testing.T) {
		got := pitchMaker.transposeSetUp([]int{11, 0, 1}, 1)
		want := []int{0, 1, 2}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetUp should wrap around 0", func(t *testing.T) {
		got := pitchMaker.transposeSetUp([]int{1, 2, 3}, 13)
		want := []int{2, 3, 4}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetDown should be correct", func(t *testing.T) {
		got := pitchMaker.transposeSetDown([]int{0, 1, 2}, 1)
		want := []int{11, 0, 1}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetDown should wrap around 12", func(t *testing.T) {
		got := pitchMaker.transposeSetDown([]int{11, 0, 1}, 1)
		want := []int{10, 11, 0}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetDown should wrap around 0", func(t *testing.T) {
		got := pitchMaker.transposeSetDown([]int{1, 2, 3}, 13)
		want := []int{0, 1, 2}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("findInterval should be correct for single dyad", func(t *testing.T) {
		args := [][]int{
			{0, 1},
			{0, 7},
			{0, 9},
			{1, 4},
			{7, 3},
			{10, 2},
			{11, 6},
			{6, 6},
		}
		want := []int{1, 7, 9, 3, 8, 4, 7, 12}

		for i := 0; i < len(args); i++ {
			for j := 0; j < len(args[i]); j++ {

				got := pitchMaker.findInterval(args[i][0], args[i][1])
				if got != want[i] {
					t.Errorf("got %d want %d", got, want[i])
				}
			}
		}
	})

	t.Run("findIntervals should be correct", func(t *testing.T) {
		got := pitchMaker.findIntervals([]int{0, 2, 4, 5, 7, 9, 11})
		want := []int{2, 2, 1, 2, 2, 2}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("findIntervals should wrap around 12", func(t *testing.T) {
		got := pitchMaker.findIntervals([]int{11, 0, 2, 4, 5, 7, 9})
		want := []int{1, 2, 2, 1, 2, 2}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("findIntervals should wrap around 0", func(t *testing.T) {
		got := pitchMaker.findIntervals([]int{1, 2, 4, 5, 7, 9, 11})
		want := []int{1, 2, 1, 2, 2, 2}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	// t.Run("chromaticInvertSet set should be correct", func(t *testing.T) {
	// 	args := [][]int{
	// 		{0, 2, 4, 5, 7, 9, 11},
	// 		{0, 7, 2, 6, 1, 4, 11},
	// 		{2, 8, 9, 3, 4, 10, 1},
	// 	}

	// 	want := [][]int{
	// 		{0, 10, 8, 7, 5, 3, 1},
	// 		{0, 5, 10, 6, 11, 8, 1},
	// 		{2, 8, 7, 11, 10, 4, 1},
	// 	}

	// 	for i, arg := range args {
	// 		got := pitchMaker.invertSet(arg)
	// 		fmt.Printf("arg %v \n got %v \n want %v \n", arg, got, want[i])

	// 		if !shared.SliceEqual(got, want[i]) {
	// 			t.Errorf("got %v want %v", got, want[i])
	// 		}
	// 	}
	// })
}
