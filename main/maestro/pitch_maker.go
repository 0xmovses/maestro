package maestro

type PitchMaker interface {
	transposeUp(pitch, interval uint) uint
	transposeDown(pitch, interval uint) uint
	transposeSetUp(pitchSet []uint, interval uint) []uint
	transposeSetDown(pitchSet []uint, interval uint) []uint
	findIntervals(pitchSet []uint) []uint
	invertSet(pitchSet []uint) []uint
}

type pitchMake struct {
	PitchMaker
}

func NewPitchMaker() PitchMaker {
	return &pitchMake{}
}

func (p *pitchMake) transposeUp(pitch, interval uint) uint {
	if pitch+interval >= 12 {
		return pitch + interval - 12
	}
	return pitch + interval
}

func (p *pitchMake) transposeDown(pitch, interval uint) uint {
	if pitch-interval < 0 {
		return pitch - interval + 12
	}
	return pitch - interval
}

func (p *pitchMake) transposeSetUp(pitchSet []uint, interval uint) []uint {
	transposedSet := make([]uint, len(pitchSet))
	for i, pitch := range pitchSet {
		transposedSet[i] = p.transposeUp(pitch, interval)
	}
	return transposedSet
}

func (p *pitchMake) transposeSetDown(pitchSet []uint, interval uint) []uint {
	transposedSet := make([]uint, len(pitchSet))
	for i, pitch := range pitchSet {
		transposedSet[i] = p.transposeDown(pitch, interval)
	}
	return transposedSet
}

func (p *pitchMake) findIntervals(pitchSet []uint) []uint {
	intervals := make([]uint, len(pitchSet))
	for i, pitch := range pitchSet {
		if i == 0 {
			intervals[i] = pitch
		} else {
			intervals[i] = pitch - pitchSet[i-1]
		}
	}
	return intervals
}
