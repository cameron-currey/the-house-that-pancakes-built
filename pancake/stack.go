package pancake

import (
	"bytes"
)

type Stack []Pancake

func NewPancakeSmileyStack(orientations ...bool) Stack {
	stack := Stack{}
	for _, orientation := range orientations {
		stack = append(stack, NewPancakeSmiley(orientation))
	}
	return stack
}

func (s Stack) FlipRange(index int) {
	for i := 0; i < index; i++ {
		s[i].Flip()
	}
}

func (s Stack) Last() (int, Pancake) {
	lastIndex := len(s) - 1
	return lastIndex, s[lastIndex]
}

func (s Stack) String() string {
	var buffer bytes.Buffer
	for _, p := range s {
		buffer.WriteString(p.String())
	}
	return buffer.String()
}
