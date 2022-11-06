package maestro

import "fmt"

type PitchMaker interface {
	transposeUp(pitch, interval int) int
	transposeDown(pitch, interval int) int
	transposeSetUp(pitchSet []int, interval int) []int
	transposeSetDown(pitchSet []int, interval int) []int
	findInterval(pitch1, pitch2 int) int
	findIntervals(pitchSet []int) []int
	invertInterval(interval int) int
	invertIntervals(intervals []int) []int
	invertPitchSet(pitchSet []int) []int
	retrogradeSet(pitchSet []int) []int
	retrogradeInvertSet(pitchSet []int) []int
	createToneRowMatrix(pitchSet []int) [][]int
}

var inversionMap = map[int]int{
	0:  0,
	1:  11,
	2:  10,
	3:  9,
	4:  8,
	5:  7,
	6:  6,
	7:  5,
	8:  4,
	9:  3,
	10: 2,
	11: 1,
}

type pitchMake struct {
	PitchMaker
}

func NewPitchMaker() PitchMaker {
	return &pitchMake{}
}

func (p *pitchMake) transposeUp(pitch, interval int) int {
	if pitch+interval >= 12 {
		return pitch + interval - 12
	}
	return pitch + interval
}

func (p *pitchMake) transposeDown(pitch, interval int) int {
	if pitch-interval < 0 {
		return pitch - interval + 12
	}
	return pitch - interval
}

func (p *pitchMake) transposeSetUp(pitchSet []int, interval int) []int {
	transposedSet := make([]int, len(pitchSet))
	for i, pitch := range pitchSet {

		transposedSet[i] = p.transposeUp(pitch, interval)
	}
	return transposedSet
}

func (p *pitchMake) transposeSetDown(pitchSet []int, interval int) []int {
	transposedSet := make([]int, len(pitchSet))
	for i, pitch := range pitchSet {
		transposedSet[i] = p.transposeDown(pitch, interval)
	}
	return transposedSet
}

func (p *pitchMake) findIntervals(pitchSet []int) []int {
	intervals := make([]int, (len(pitchSet) - 1))

	for i := 0; i < len(pitchSet)-1; i++ {
		intervals[i] = p.findInterval(pitchSet[i], pitchSet[i+1])
	}

	return intervals
}

func (p *pitchMake) findInterval(pitch1, pitch2 int) int {
	var result int

	if pitch2 > pitch1 {
		fmt.Println("pitch2 > pitch1")
		fmt.Printf("pitch2: %d, pitch1: %d : ", pitch2, pitch1)
		result = pitch2 - pitch1
	} else {
		fmt.Println("pitch2 < pitch1")
		fmt.Printf("pitch2: %d, pitch1: %d : ", pitch2, pitch1)
		result = 12 - (pitch1 - pitch2)
	}

	return result
}

func (p *pitchMake) invertInterval(interval int) int {
	var i = interval

	if i > 12 {
		i = i - 12
	}
	return inversionMap[i]
}

func (p *pitchMake) invertIntervals(intervals []int) []int {
	invertedIntervals := make([]int, len(intervals))

	for i, interval := range intervals {
		invertedIntervals[i] = p.invertInterval(interval)
	}

	return invertedIntervals
}

func (p *pitchMake) invertPitchSet(pitchSet []int) []int {
	invertedSet := make([]int, len(pitchSet))
	invertedSet[0] = pitchSet[0]

	for i, pitch := range pitchSet {
		if i == 0 {
			continue
		}

		interval := p.findInterval(pitchSet[i-1], pitch)
		invertedInterval := p.invertInterval(interval)
		if pitchSet[i-1] > pitch {
			invertedSet[i] = p.transposeUp(invertedSet[i-1], invertedInterval)
		} else {
			invertedSet[i] = p.transposeDown(invertedSet[i-1], interval)
		}
	}

	return invertedSet
}

func (p *pitchMake) retrogradeSet(pitchSet []int) []int {
	retrogradedSet := make([]int, len(pitchSet))
	for i := range pitchSet {
		retrogradedSet[i] = pitchSet[len(pitchSet)-1-i]
	}
	return retrogradedSet
}

func (p *pitchMake) retrogradeInvertSet(pitchSet []int) []int {
	invertedSet := p.invertPitchSet(pitchSet)
	retrogradeInvertedSet := p.retrogradeSet(invertedSet)
	return retrogradeInvertedSet
}

func (p *pitchMake) createToneRowMatrix(pitchSet []int) [][]int {
	matrix := initializeEmptyMatrix(len(pitchSet))
	invertedSet := p.invertPitchSet(pitchSet)
	fmt.Printf("invertedSet: %v \n", invertedSet)

	for i := range matrix {
		if i == 0 {
			matrix[i] = pitchSet
			continue
		}
		transposeBy := p.findInterval(pitchSet[0], invertedSet[i])
		fmt.Printf("transposeBy: %d \n", transposeBy)
		matrix[i] = p.transposeSetUp(pitchSet, transposeBy)
		// if pitchSet[i] > invertedSet[i] {
		// 	matrix[i] = p.transposeSetDown(pitchSet, invertedSet[i])
		// } else {
		// 	matrix[i] = p.transposeSetUp(pitchSet, invertedSet[i])
		// }
		fmt.Printf("matrix[%d]: %v \n", i, matrix[i])
	}

	return matrix
}

func initializeEmptyMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}
