package restaurant

import (
	"../pancake"
	. "gopkg.in/check.v1"
)

type WaiterXRayVisionSuite struct {
	waiterXRayVision WaiterXRayVision
	callback         func(int, pancake.Stack)
}

var _ = Suite(&WaiterXRayVisionSuite{})

func (s *WaiterXRayVisionSuite) SetUpTest(c *C) {
	s.waiterXRayVision = WaiterXRayVision{}
}

func (s *WaiterXRayVisionSuite) assertFlipDatStack(c *C, expectedCount int, orientations ...bool) {
	c.Assert(s.waiterXRayVision.FlipDatStack(
		pancake.NewPancakeSmileyStack(orientations...),
		func(int, pancake.Stack) {},
	), Equals, expectedCount)
}

func (s *WaiterXRayVisionSuite) TestFlipDatStackCase1(c *C) {
	s.assertFlipDatStack(c, 1, false)
}

func (s *WaiterXRayVisionSuite) TestFlipDatStackCase2(c *C) {
	s.assertFlipDatStack(c, 1, false, true)
}

func (s *WaiterXRayVisionSuite) TestFlipDatStackCase3(c *C) {
	s.assertFlipDatStack(c, 2, true, false)
}

func (s *WaiterXRayVisionSuite) TestFlipDatStackCase4(c *C) {
	s.assertFlipDatStack(c, 0, true, true, true)
}

func (s *WaiterXRayVisionSuite) TestFlipDatStackCase5(c *C) {
	s.assertFlipDatStack(c, 3, false, false, true, false)
}
