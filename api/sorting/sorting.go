package sorting

import (
	"errors"
	"fmt"
	"govisualiser/api/util"
)

func Sort(arr []int, algorithm string) ([]util.SortChanges, error) {
	switch algorithm {
	case "insertion":
		return insertionSort(arr), nil
	case "quick":
		return quickSort(arr), nil
	case "merge":
		_, changes := mergeSort(arr)
		return changes, nil
	case "bubble":
		return bubbleSort(arr), nil
	case "selection":
		return selectionSort(arr), nil
	case "heap":
		return heapSort(arr), nil
	case "radix":
		return radixSort(arr), nil
	case "bogo":
		return bogoSort(arr), nil
	}
	return []util.SortChanges{}, errors.New("algorithm type not supported")
}

func insertionSort(arr []int) []util.SortChanges {
	var res []util.SortChanges
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			res = append(res, util.SortChanges{
				FirstIndex:  j + 1,
				SecondIndex: j,
				FirstValue:  arr[j+1],
				SecondValue: key,
			})
			j--
		}
		arr[j+1] = key
	}
	return res
}

func quickSort(arr []int) []util.SortChanges {
	var res []util.SortChanges
	return res
}

func mergeSort(arr []int) ([]int, []util.SortChanges) {
	var slice []int
	var res []util.SortChanges
	if len(arr) > 1 {
		m := len(arr) / 2
		sliceL, changes := mergeSort(arr[:m])
		res = append(res, changes...)
		sliceR, changes := mergeSort(arr[m:])
		res = append(res, changes...)
		slice, changes = merge(sliceL, sliceR)
		res = append(res, changes...)
		return slice, res
	}
	return arr, []util.SortChanges{}
}

func merge(l, r []int) ([]int, []util.SortChanges) {
	var res []util.SortChanges
	length, i, j := len(l)+len(r), 0, 0
	all := l
	all = append(all, r...)
	arr := make([]int, length)
	for k := 0; k < length; k++ {
		if i > len(l)-1 && j <= len(r)-1 {
			arr[k] = r[j]
			res = append(res, util.SortChanges{
				FirstIndex:  k,
				SecondIndex: len(l) + j,
				FirstValue:  all[k],
				SecondValue: all[len(l)+j],
			})
			all[len(l)+j] = all[k]
			all[k] = arr[k]
			j++
		} else if j > len(r)-1 && i <= len(l)-1 {
			arr[k] = l[i]
			res = append(res, util.SortChanges{
				FirstIndex:  k,
				SecondIndex: i,
				FirstValue:  all[k],
				SecondValue: all[i],
			})
			all[i] = all[k]
			all[k] = arr[k]
			i++
		} else if l[i] < r[j] {
			arr[k] = l[i]
			res = append(res, util.SortChanges{
				FirstIndex:  k,
				SecondIndex: i,
				FirstValue:  all[k],
				SecondValue: all[i],
			})
			all[i] = all[k]
			all[k] = arr[k]
			i++
		} else {
			arr[k] = r[j]
			res = append(res, util.SortChanges{
				FirstIndex:  k,
				SecondIndex: len(l) + j,
				FirstValue:  all[k],
				SecondValue: all[len(l)+j],
			})
			all[len(l)+j] = all[k]
			all[k] = arr[k]
			j++
		}
	}
	fmt.Println(arr)
	return arr, res
}

func bubbleSort(arr []int) []util.SortChanges {
	var res []util.SortChanges

	return res
}

func selectionSort(arr []int) []util.SortChanges {
	var res []util.SortChanges

	return res
}

func heapSort(arr []int) []util.SortChanges {
	var res []util.SortChanges

	return res
}

func radixSort(arr []int) []util.SortChanges {
	var res []util.SortChanges

	return res
}

func bogoSort(arr []int) []util.SortChanges {
	var res []util.SortChanges

	return res
}
