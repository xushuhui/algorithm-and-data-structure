package sort

import (
	"fmt"
	"testing"

	"data-structures/utils"

	"golang.org/x/sync/errgroup"
)

func TestSlice(t *testing.T) {
	n := 10
	s0 := utils.GenerateRandomArray(n, 0, n)
	fmt.Println(s0)
	BubbleSort(s0)
	fmt.Println(s0)
}

func TestSort(t *testing.T) {
	n := 100000
	s0 := utils.GenerateRandomArray(n, 0, n)

	var eg errgroup.Group
	eg.Go(func() error {
		utils.TimeSpent("BubbleSort", BubbleSort, utils.CopyArray(s0, n))
		return nil
	})
	eg.Go(func() error {
		utils.TimeSpent("QuickSort", QuickSort, utils.CopyArray(s0, n))
		return nil
	})
	eg.Go(func() error {
		utils.TimeSpent("InsertionSort", InsertionSort, utils.CopyArray(s0, n))
		return nil
	})
	eg.Go(func() error {
		utils.TimeSpent("InsertionSortAdvance", InsertionSortAdvance, utils.CopyArray(s0, n))
		return nil
	})
	eg.Go(func() error {
		utils.TimeSpent("SelectionSort", SelectionSort, utils.CopyArray(s0, n))
		return nil
	})
	err := eg.Wait()
	if err != nil {
		return
	}
}

func sort(s []int, n int) {
	s5 := utils.CopyArray(s, n)
	s6 := utils.CopyArray(s, n)

	// utils.TimeSpent("SelectionSort", algorithm.SelectionSort, s4)
	utils.TimeSpent("InsertionSort", InsertionSort, s5)
	utils.TimeSpent("InsertionSortAdvance", InsertionSortAdvance, s6)
}
