package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample1(t *testing.T) {
	vm := VM{
		c:       9,
		program: []int{2, 6},
	}

	vm.Run()
	require.Equal(t, 1, vm.b)
}

func TestExample2(t *testing.T) {
	vm := VM{
		a:       10,
		program: []int{5, 0, 5, 1, 5, 4},
	}

	vm.Run()
	require.Equal(t, "0,1,2", vm.String())
}

func TestExample3(t *testing.T) {
	vm := VM{
		a:       2024,
		program: []int{0, 1, 5, 4, 3, 0},
	}

	vm.Run()
	require.Equal(t, "4,2,5,6,7,7,7,7,3,1,0", vm.String())
	require.Equal(t, 0, vm.a)
}

func TestExample4(t *testing.T) {
	vm := VM{
		b:       29,
		program: []int{1, 7},
	}

	vm.Run()
	require.Equal(t, 26, vm.b)
}

func TestExample5(t *testing.T) {
	vm := VM{
		b:       2024,
		c:       43690,
		program: []int{4, 0},
	}

	vm.Run()
	require.Equal(t, 44354, vm.b)
}

func TestLambda(t *testing.T) {
	vm := VM{
		a:       51571418,
		b:       0,
		c:       0,
		program: []int{2, 4, 1, 1, 7, 5, 0, 3, 1, 4, 4, 5, 5, 5, 3, 0},
	}
	vm.Run()

	require.Equal(t, vm.String(), λ(51571418))
}
