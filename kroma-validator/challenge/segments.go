package challenge

type Hash = [32]byte

type Segments struct {
	Start  uint64
	Size   uint64
	Degree uint64
	Hashes []Hash
}

func NewSegments(start, size uint64, hashes []Hash) *Segments {
	return &Segments{
		Start:  start,
		Size:   size,
		Degree: size / (uint64(len(hashes)) - 1),
		Hashes: hashes,
	}
}

func NewEmptySegments(start, size, sections uint64) *Segments {
	return NewSegments(start, size, make([]Hash, sections))
}

func (s *Segments) NextSegmentsRange(position uint64) (start, size uint64) {
	return s.Start + position*s.Degree, s.Degree
}

func (s *Segments) SetHashValue(index int, hash Hash) {
	s.Hashes[index] = hash
}

func (s *Segments) BlockNumbers() []uint64 {
	arr := make([]uint64, len(s.Hashes))
	for i := range arr {
		arr[i] = s.Start + uint64(i)*s.Degree
	}

	return arr
}
