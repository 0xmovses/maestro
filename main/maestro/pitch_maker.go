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
	fmt.Printf("transposing pitch %d up by interval %d \n", pitch, interval)
	if pitch == 0 {
		pitch = 12
	}
	if pitch+interval >= 12 {
		return pitch + interval - 12
	}
	fmt.Printf("transposed pitch %d \n", pitch+interval)
	return pitch + interval
}

func (p *pitchMake) transposeDown(pitch, interval int) int {
	fmt.Printf("transposing pitch %d down by interval %d \n", pitch, interval)
	if pitch == 0 {
		pitch = 12
	}
	if pitch-interval < 0 {
		return pitch - interval + 12
	}
	fmt.Printf("transposed pitch %d \n", pitch-interval)
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
		result = pitch2 - pitch1
	} else {
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
	for _, pitch := range pitchSet {
		fmt.Printf("pitch args%d \n", pitch)
	}
	invertedSet := make([]int, len(pitchSet))
	invertedSet[0] = pitchSet[0]

	for i, pitch := range pitchSet {
		if i == 0 {
			continue
		}

		interval := p.findInterval(pitchSet[i-1], pitch)
		invertedInterval := p.invertInterval(interval)
		if pitchSet[i-1] > pitch {
			fmt.Printf("pitch-1 %d is greater than pitch %d \n", pitchSet[i-1], pitch)
			invertedSet[i] = p.transposeUp(invertedSet[i-1], invertedInterval)
		} else {
			fmt.Printf("pitch-1 %d is less than pitch %d \n", pitchSet[i-1], pitch)
			invertedSet[i] = p.transposeDown(invertedSet[i-1], interval)
		}
		fmt.Printf("invertedSet: %v \n", invertedSet)
	}

	return invertedSet
}

func getCircularNextPitch(pitchSet []int, i int) int {
	var next int
	if i == len(pitchSet)-1 {
		next = pitchSet[0]
		return makeZeroTheGreatest(next)
	}
	next = pitchSet[i+1]
	return makeZeroTheGreatest(next)
}

func makeZeroTheGreatest(pitch int) int {
	if pitch == 0 {
		pitch = 12
	}
	return pitch
}

func (p *pitchMake) retrogradeSet(pitchSet []int) []int {
	retrogradedSet := make([]int, len(pitchSet))
	for i, _ := range pitchSet {
		retrogradedSet[i] = pitchSet[len(pitchSet)-1-i]
	}
	return retrogradedSet
}

// work on matrix generation next
