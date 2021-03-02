package main

import "fmt"


func partition(arr []int, start, end int) int {
    i := start -1
    pivot := arr[end]

    for j := start; j < end; j++ {
        if arr[j] <  pivot {
            i++
            arr[j],arr[i] = arr[i],arr[j]

        }
    }
    arr[i+1] ,arr[end] = arr[end],arr[i+1]
    return i+1
}


func quickSort(arr []int, start int, end int) {
    if start < end {
        pi := partition(arr, start, end)
        quickSort(arr, start, pi -1)
        quickSort(arr, pi + 1, end)

    }
}



func main() {
    a := []int{19,9,1,5,20,60,30}
    quickSort(a,0, len(a)-1)
    fmt.Println(a)

}
