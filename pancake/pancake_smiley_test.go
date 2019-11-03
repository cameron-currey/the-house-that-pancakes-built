package pancake

import (
	. "gopkg.in/check.v1"
)

type PancakeSmileySuite struct {
	pancakeSmiley PancakeSmiley
}

var _ = Suite(&PancakeSmileySuite{})

func (s *PancakeSmileySuite) SetUpTest(c *C) {
	s.pancakeSmiley = PancakeSmiley{true}
}

func (s *PancakeSmileySuite) TestFlipTogglesFaceUp(c *C) {
	s.pancakeSmiley.faceUp = false
	s.pancakeSmiley.Flip()
	c.Assert(s.pancakeSmiley.faceUp, Equals, true)
}

func (s *PancakeSmileySuite) TestFaceUpReturnsFaceUp(c *C) {
	c.Assert(s.pancakeSmiley.FaceUp(), Equals, true)
}

func (s *PancakeSmileySuite) TestStringReturnsSmileyWhenUpside(c *C) {
	s.pancakeSmiley.faceUp = true
	c.Assert(s.pancakeSmiley.String(), Equals, "😍")
}

func (s *PancakeSmileySuite) TestStringReturnsNoFaceWhenNotUpside(c *C) {
	s.pancakeSmiley.faceUp = false
	c.Assert(s.pancakeSmiley.String(), Equals, "😶")
}
