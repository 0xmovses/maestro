package maestro

import (
	"testing"

	"github.com/rvmelkonian/maestro/main/shared"
)

func TestPitchMaker(t *testing.T) {
	pitchMaker := NewPitchMaker()

	t.Run("transposeUp should be correct", func(t *testing.T) {
		got := pitchMaker.transposeUp(0, 1)
		want := uint(1)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("transposeUp should wrap around 12", func(t *testing.T) {
		got := pitchMaker.transposeUp(11, 1)
		want := uint(0)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("transposeUp shoul wrap around 0", func(t *testing.T) {
		got := pitchMaker.transposeUp(1, 13)
		want := uint(2)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("transposeSetUp should be correct", func(t *testing.T) {
		got := pitchMaker.transposeSetUp([]uint{0, 1, 2}, 1)
		want := []uint{1, 2, 3}

		if !shared.UintSliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetUp should wrap around 12", func(t *testing.T) {
		got := pitchMaker.transposeSetUp([]uint{11, 0, 1}, 1)
		want := []uint{0, 1, 2}

		if !shared.UintSliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetUp should wrap around 0", func(t *testing.T) {
		got := pitchMaker.transposeSetUp([]uint{1, 2, 3}, 13)
		want := []uint{2, 3, 4}

		if !shared.UintSliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetDown should be correct", func(t *testing.T) {
		got := pitchMaker.transposeSetDown([]uint{0, 1, 2}, 1)
		want := []uint{11, 0, 1}

		if !shared.UintSliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetDown should wrap around 12", func(t *testing.T) {
		got := pitchMaker.transposeSetDown([]uint{11, 0, 1}, 1)
		want := []uint{10, 11, 0}

		if !shared.UintSliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("transposeSetDown should wrap around 0", func(t *testing.T) {
		got := pitchMaker.transposeSetDown([]uint{1, 2, 3}, 13)
		want := []uint{0, 1, 2}

		if !shared.UintSliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("findIntervals should be correct", func(t *testing.T) {
		got := pitchMaker.findIntervals([]uint{0, 2, 4, 5, 7, 9, 11})
		want := []uint{2, 2, 1, 2, 2, 2}

		if !shared.UintSliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("findIntervals should wrap around 12", func(t *testing.T) {
		got := pitchMaker.findIntervals([]uint{11, 0, 2, 4, 5, 7, 9})
		want := []uint{1, 2, 2, 1, 2, 2}

		if !shared.UintSliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("findIntervals should wrap around 0", func(t *testing.T) {
		got := pitchMaker.findIntervals([]uint{1, 2, 4, 5, 7, 9, 11})
		want := []uint{1, 2, 1, 2, 2, 2}

		if !shared.UintSliceEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
