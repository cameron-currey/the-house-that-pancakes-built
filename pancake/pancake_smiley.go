package pancake

type PancakeSmiley struct {
	faceUp bool
}

func NewPancakeSmiley(faceUp bool) *PancakeSmiley {
	return &PancakeSmiley{faceUp}
}

func (s *PancakeSmiley) Flip() {
	s.faceUp = !s.faceUp
}

func (s *PancakeSmiley) FaceUp() bool {
	return s.faceUp
}

func (s *PancakeSmiley) String() string {
	if s.faceUp {
		return "😍"
	} else {
		return "😶"
	}
}
