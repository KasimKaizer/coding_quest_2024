package connect

import "math/big"

func FindPath(steps []int, max int) *big.Int {
	mem := make(map[int]*big.Int)
	return find(steps, max, mem)
}

func find(steps []int, max int, mem map[int]*big.Int) *big.Int {

	if x, ok := mem[max]; ok {
		return x
	}
	if max == 0 {
		return big.NewInt(1)
	}
	if max < 0 {
		return big.NewInt(0)
	}
	paths := big.NewInt(0)
	for _, step := range steps {
		paths.Add(paths, find(steps, max-step, mem))
	}
	mem[max] = paths
	return paths
}
