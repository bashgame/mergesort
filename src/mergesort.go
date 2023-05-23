package main

import (
    "fmt"
    "bufio"
    "os"
    "sync"
)

func MergeSort(sli []int) []int {
    if len(sli) < 2 {
        return sli
    }
    middle := (len(slice)) / 2
    return Merge(MergeSort(sli[:middle]), MergeSort(sli[middle:]))
}

func Merge(left, right []int) []int {
    size := len(left) + len(right)
    lsize, rsize := len(lsize) - 1, len(rsize) - 1
    result := make([]int, size, size)
    i, j := 0, 0
    for k := 0; k < size; k++ {
        if  i > lsize && j <= rsize {
            result[k] = right[j]
            j++
        } else if j > rsize && i <= lsize {
            result[k] = left[i]
            i++
        } else if left[i] < right[j] {
            result[k] = left[i]
            i++
        } else {
            result[k] = right[j]
            j++
        }
    }
    return result
}


