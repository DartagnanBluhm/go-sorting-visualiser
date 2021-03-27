package sorting

import (
	"errors"
	"fmt"
	"govisualiser/api/util"
	"math/rand"
	"time"
)

func Sort(arr []int, algorithm string) ([]util.SortChanges, error) {
	switch algorithm {
	case "insertion":
		fmt.Printf("%s - Insertion sort request processing...", time.Now().Local().Format("2006/01/02 15:04:05"))
		changes := insertionSort(arr)
		fmt.Print("done.\n")
		return changes, nil
	case "quick":
		fmt.Printf("%s - Quicksort request processing...", time.Now().Local().Format("2006/01/02 15:04:05"))
		changes := quickSort(&arr, 0, len(arr)-1)
		fmt.Print("done.\n")
		return changes, nil
	case "merge":
		fmt.Printf("%s - Mergesort request processing...", time.Now().Local().Format("2006/01/02 15:04:05"))
		_, changes := mergeSort(arr, 0)
		fmt.Print("done.\n")
		return changes, nil
	case "bubble":
		fmt.Printf("%s - Bubble sort request processing...", time.Now().Local().Format("2006/01/02 15:04:05"))
		changes := bubbleSort(arr)
		fmt.Print("done.\n")
		return changes, nil
	case "selection":
		fmt.Printf("%s - Selection sort request processing...", time.Now().Local().Format("2006/01/02 15:04:05"))
		changes := selectionSort(arr)
		fmt.Print("done.\n")
		return changes, nil
	case "heap":
		fmt.Printf("%s - Heapsort request processing...", time.Now().Local().Format("2006/01/02 15:04:05"))
		changes := heapSort(arr)
		fmt.Print("done.\n")
		return changes, nil
	case "radix":
		fmt.Printf("%s - Radixsort request processing...", time.Now().Local().Format("2006/01/02 15:04:05"))
		changes := radixSort(arr)
		fmt.Print("done.\n")
		return changes, nil
	case "bogo":
		if len(arr) > 10 {
			return []util.SortChanges{}, errors.New("Bogosort can only handle an array size of maximum 10")
		}
		fmt.Printf("%s - Bogosort request processing...", time.Now().Local().Format("2006/01/02 15:04:05"))
		changes := bogoSort(arr)
		fmt.Print("done.\n")
		return changes, nil
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

func quickSort(arr *[]int, low, high int) []util.SortChanges {
	var res []util.SortChanges
	if low < high {
		p, changes := partition(arr, low, high)
		res = append(res, changes...)
		res = append(res, quickSort(arr, low, p-1)...)
		res = append(res, quickSort(arr, p+1, high)...)
	}
	return res
}

func partition(arr *[]int, low, high int) (int, []util.SortChanges) {
	var res []util.SortChanges
	p := (*arr)[high]
	i := low - 1
	for j := low; j < high; j++ {
		if (*arr)[j] < p {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			res = append(res, util.SortChanges{
				FirstIndex:  i,
				SecondIndex: j,
				FirstValue:  (*arr)[i],
				SecondValue: (*arr)[j],
			})
		}
	}
	(*arr)[i+1], (*arr)[high] = (*arr)[high], (*arr)[i+1]
	res = append(res, util.SortChanges{
		FirstIndex:  i + 1,
		SecondIndex: high,
		FirstValue:  (*arr)[i+1],
		SecondValue: (*arr)[high],
	})
	return i + 1, res
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
	arr := make([]int, length)
	for k := 0; k < length; k++ {
		if i >= len(l) {
			arr[k] = r[j]
			res = append(res, util.SortChanges{
				FirstIndex: offset + k,
				FirstValue: arr[k],
			})
			j++
		} else if j >= len(r) {
			arr[k] = l[i]
			res = append(res, util.SortChanges{
				FirstIndex: offset + k,
				FirstValue: arr[k],
			})
			i++
		} else if l[i] < r[j] {
			arr[k] = l[i]
			res = append(res, util.SortChanges{
				FirstIndex: offset + k,
				FirstValue: arr[k],
			})
			i++
		} else {
			arr[k] = r[j]
			res = append(res, util.SortChanges{
				FirstIndex: offset + k,
				FirstValue: arr[k],
			})
			j++
		}
	}
	return arr, res
}

func bubbleSort(arr []int) []util.SortChanges {
	var res []util.SortChanges
	for i := range arr {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				res = append(res, util.SortChanges{
					FirstIndex:  j,
					SecondIndex: j + 1,
					FirstValue:  arr[j],
					SecondValue: arr[j+1],
				})
			}
		}
	}
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
	count := make([]int, 10)
	output := make([]int, len(*arr))
	for i := range *arr {
		count[((*arr)[i]/e)%10]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	for j := len(*arr) - 1; j >= 0; j-- {
		res = append(res, util.SortChanges{
			FirstIndex: count[((*arr)[j]/e)%10] - 1,
			FirstValue: (*arr)[j],
		})
		output[count[((*arr)[j]/e)%10]-1] = (*arr)[j]
		count[((*arr)[j]/e)%10]--
	}
	*arr = output
	return res
}

func bogoSort(arr []int) []util.SortChanges {
	var res []util.SortChanges
	for !sorted(arr) {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(arr), func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
			res = append(res, util.SortChanges{
				FirstIndex:  i,
				SecondIndex: j,
				FirstValue:  arr[i],
				SecondValue: arr[j],
			})
		})
	}
	return res
}

func sorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
