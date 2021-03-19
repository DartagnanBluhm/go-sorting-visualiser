package sorting

import (
	"errors"
	"fmt"
	"govisualiser/api/util"
)

func Sort(arr []int, algorithm string) ([]util.SortChanges, error) {
	fmt.Println(arr)
	switch algorithm {
	case "insertion":
		return insertionSort(arr), nil
	case "quick":
		return quickSort(arr), nil
	case "merge":
		_, changes := mergeSort(arr, 0)
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

func mergeSort(arr []int, offset int) ([]int, []util.SortChanges) {
	var res []util.SortChanges
	if len(arr) > 1 {
		m := len(arr) / 2
		sliceL, changes := mergeSort(arr[:m], offset)
		res = append(res, changes...)
		sliceR, changes := mergeSort(arr[m:], offset+m)
		res = append(res, changes...)
		slice, changes := merge(sliceL, sliceR, offset)
		res = append(res, changes...)
		return slice, res
	}
	return arr, []util.SortChanges{}
}

func merge(l, r []int, offset int) ([]int, []util.SortChanges) {
	var res []util.SortChanges
	length, i, j := len(l)+len(r), 0, 0
	all := make([]int, 0)
	all = append(all, l...)
	all = append(all, r...)
	arr := make([]int, length)
	for k := 0; k < length; k++ {
		fmt.Println("arr", arr, "l", l, "r", r, "all", all, "i", i, "j", j, "k", k, "offset", offset)
		if i >= len(l) {
			arr[k] = r[j]
			all[len(l)+j] = all[k]
			fmt.Print("opt 1-")
			all[k] = arr[k]
			j++
		} else if j >= len(r) {
			arr[k] = l[i]
			if i != k && all[k] < l[i] {
				res = append(res, util.SortChanges{
					FirstIndex:  offset + k,
					SecondIndex: offset + i,
					FirstValue:  arr[k],
					SecondValue: l[i-1],
				})
				fmt.Println(res, "opt2")
				all[k] = arr[k]
			}
			fmt.Print("opt 2-")
			i++
		} else if l[i] < r[j] {
			arr[k] = l[i]
			if i != k && all[k] < all[i] {
				res = append(res, util.SortChanges{
					FirstIndex:  offset + k,
					SecondIndex: offset + i,
					FirstValue:  arr[k],
					SecondValue: all[k],
				})
				fmt.Println(res, "opt3")
				all[i] = all[k]
				all[k] = arr[k]
			}
			fmt.Print("opt 3-")
			i++
		} else {
			arr[k] = r[j]
			if len(l)+j != k {
				res = append(res, util.SortChanges{
					FirstIndex:  offset + k,
					SecondIndex: offset + len(l) + j,
					FirstValue:  all[len(l)+j],
					SecondValue: all[k],
				})
				fmt.Println(res, "opt4")
				temp := all[len(l)+j]
				all[len(l)+j] = all[k]
				all[k] = temp
			}

			j++
		}
	}
	fmt.Println(arr, all, "\n")
	return arr, res
}

func bubbleSort(arr []int) []util.SortChanges {
	var res []util.SortChanges

	return res
}

func selectionSort(arr []int) []util.SortChanges {
	var res []util.SortChanges
	for i := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		res = append(res, util.SortChanges{
			FirstIndex:  i,
			SecondIndex: min,
			FirstValue:  arr[i],
			SecondValue: arr[min],
		})
	}
	fmt.Println(arr)
	return res
}

func heapSort(arr []int) []util.SortChanges {
	var res []util.SortChanges
	for i := len(arr)/2 - 1; i >= 0; i-- {
		c := heap(&arr, len(arr), i)
		res = append(res, c...)
	}
	for i := len(arr) - 1; i > 0; i-- {
		temp := arr[0]
		arr[0] = arr[i]
		arr[i] = temp
		res = append(res, util.SortChanges{
			FirstIndex:  0,
			SecondIndex: i,
			FirstValue:  arr[0],
			SecondValue: arr[i],
		})
		c := heap(&arr, i, 0)
		res = append(res, c...)
	}
	fmt.Println(arr)
	return res
}

func heap(arr *[]int, s int, i int) []util.SortChanges {
	var res []util.SortChanges
	big := i
	l := 2*i + 1
	r := 2*i + 2
	if l < s && (*arr)[l] > (*arr)[big] {
		big = l
	}
	if r < s && (*arr)[r] > (*arr)[big] {
		big = r
	}
	if big != i {
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[big]
		(*arr)[big] = temp
		res = append(res, util.SortChanges{
			FirstIndex:  i,
			SecondIndex: big,
			FirstValue:  (*arr)[i],
			SecondValue: (*arr)[big],
		})
		c := heap(arr, s, big)
		res = append(res, c...)
	}
	return res
}

func radixSort(arr []int) []util.SortChanges {
	var res []util.SortChanges
	max := radixGetMax(arr)
	for exp := 1; max/exp > 0; exp *= 10 {
		c := radixCountingSort(&arr, exp)
		res = append(res, c...)
	}
	return res
}

func radixGetMax(arr []int) int {
	max := arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func radixCountingSort(arr *[]int, e int) []util.SortChanges {
	var res []util.SortChanges

	return res
}

func bogoSort(arr []int) []util.SortChanges {
	var res []util.SortChanges

	return res
}
