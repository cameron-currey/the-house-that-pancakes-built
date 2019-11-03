package pancake

import (
	. "gopkg.in/check.v1"
)

type StackSuite struct {
	stack Stack
}

var _ = Suite(&StackSuite{})

func (s *StackSuite) SetUpTest(c *C) {
	s.stack = Stack{
		NewPancakeSmiley(true),
		NewPancakeSmiley(false),
		NewPancakeSmiley(true),
	}
}

func (s *StackSuite) TestNewPancakeSmileyStackReturnsOrientedList(c *C) {
	result := NewPancakeSmileyStack(true, false, true)
	c.Assert(result, DeepEquals, s.stack)
}

func (s *StackSuite) TestFlipRangeFlipsPancakesBeforeIndex(c *C) {
	s.stack.FlipRange(2)
	c.Assert(s.stack, DeepEquals, Stack{
		NewPancakeSmiley(false),
		NewPancakeSmiley(true),
		NewPancakeSmiley(true),
	})
}

func (s *StackSuite) TestStringBuildsStringLeftToRight(c *C) {
	c.Assert(s.stack.String(), Equals, s.stack[0].String()+s.stack[1].String()+s.stack[2].String())
}
