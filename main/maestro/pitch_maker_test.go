package maestro

import (
	"fmt"
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

	t.Run("findInterval should be correct for each dyad", func(t *testing.T) {
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

	t.Run("invertInterval should correctly invert an single interval", func(t *testing.T) {
		args := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		want := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

		for i := 0; i < len(args); i++ {
			got := pitchMaker.invertInterval(args[i])
			if got != want[i] {
				t.Errorf("got %d want %d", got, want[i])
			}
		}
	})

	t.Run("invertIntervals should correctly invert a slice of intervals", func(t *testing.T) {
		got := pitchMaker.invertIntervals([]int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
		})
		want := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

		if !shared.SliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	args := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		{0, 11, 10, 4, 2, 9, 3, 6, 8, 7, 1, 5},
		{1, 6, 3},
		{4, 2, 7, 5},
		{6, 1, 8, 3, 10, 11},
	}

	t.Run("invertPitchSet should correctly invert a series of pitch sets", func(t *testing.T) {

		want := [][]int{
			{0, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			{0, 1, 2, 8, 10, 3, 9, 6, 4, 5, 11, 7},
			{1, 8, 11},
			{4, 6, 1, 3},
			{6, 11, 4, 9, 2, 1},
		}

		for i := 0; i < len(args); i++ {
			fmt.Printf("args: %v \n", args[i])
			got := pitchMaker.invertPitchSet(args[i])
			if !shared.SliceEqual(got, want[i]) {
				t.Errorf("got %v want %v", got, want[i])
			}
		}
	})

	t.Run("retrogradeSet should correctly retrograte a series of pitch sets", func(t *testing.T) {
		want := [][]int{
			{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			{5, 1, 7, 8, 6, 3, 9, 2, 4, 10, 11, 0},
			{3, 6, 1},
			{5, 7, 2, 4},
			{11, 10, 3, 8, 1, 6},
		}

		for i := 0; i < len(args); i++ {
			got := pitchMaker.retrogradeSet(args[i])
			if !shared.SliceEqual(got, want[i]) {
				t.Errorf("got %v want %v", got, want[i])
			}
		}
	})

	t.Run("retrogradeInvertSet should correctly retrograte and invert a series of pitch sets", func(t *testing.T) {
		want := [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0},
			{7, 11, 5, 4, 6, 9, 3, 10, 8, 2, 1, 0},
			{11, 8, 1},
			{3, 1, 6, 4},
			{1, 2, 9, 4, 11, 6},
		}

		for i := 0; i < len(args); i++ {
			got := pitchMaker.retrogradeInvertSet(args[i])
			if !shared.SliceEqual(got, want[i]) {
				t.Errorf("got %v want %v", got, want[i])
			}
		}
	})
}
