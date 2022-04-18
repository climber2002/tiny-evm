package vm

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestLen(t *testing.T) {
	stack := newstack()
	if stack.len() != 0 {
		t.Error("Stack should be len 0 initially")
	}

	i, _ := uint256.FromHex("0x10")
	stack.push(i)
	if stack.len() != 1 {
		t.Error("Stack len should be 1")
	}
}

func TestPeek(t *testing.T) {
	stack := newstack()
	i, _ := uint256.FromHex("0x10")
	stack.push(i)

	if *stack.peek() != *i {
		t.Error("should peek correct value")
	}
}

func TestDup(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt(10))
	stack.push(uint256.NewInt(20))

	stack.dup(2)
	if stack.len() != 3 {
		t.Error("Stack len not correct")
	}

	if *stack.peek() != *uint256.NewInt(10) {
		t.Error("Dupped value not correct")
	}
}

func TestSwap(t *testing.T) {
	stack := newstack()
	stack.push(uint256.NewInt(10))
	stack.push(uint256.NewInt(20))

	stack.swap(2)

	if *stack.peek() != *uint256.NewInt(10) {
		t.Error("Dupped value not correct")
	}

	stack.pop()
	if *stack.peek() != *uint256.NewInt(20) {
		t.Error("Dupped value not correct")
	}
}
