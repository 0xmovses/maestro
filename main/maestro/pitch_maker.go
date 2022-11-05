package maestro

type PitchMaker interface {
	transposeUp(pitch, interval int) int
	transposeDown(pitch, interval int) int
	transposeSetUp(pitchSet []int, interval int) []int
	transposeSetDown(pitchSet []int, interval int) []int
	findInterval(pitch1, pitch2 int) int
	findIntervals(pitchSet []int) []int
	invertSet(pitchSet []int) []int
	retrogradeSet(pitchSet []int) []int
	retrogradeInvertSet(pitchSet []int) []int
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
		result = pitch2 - pitch1
	} else {
		result = 12 - (pitch1 - pitch2)
	}

	return result
}

func (p *pitchMake) retrogradeSet(pitchSet []int) []int {
	retrogradedSet := make([]int, len(pitchSet))
	for i, _ := range pitchSet {
		retrogradedSet[i] = pitchSet[len(pitchSet)-1-i]
	}

	return retrogradedSet
}

func (p *pitchMake) retrogradeInvertSet(pitchSet []int) []int {
	retrogradedInvertedSet := make([]int, len(pitchSet))
	for i, _ := range pitchSet {
		retrogradedInvertedSet[i] = 12 - pitchSet[len(pitchSet)-1-i]

		if retrogradedInvertedSet[i] == 12 {
			retrogradedInvertedSet[i] = 0
		}
		if retrogradedInvertedSet[i] > 12 {
			retrogradedInvertedSet[i] = retrogradedInvertedSet[i] - 12
		}
	}

	return retrogradedInvertedSet
}
