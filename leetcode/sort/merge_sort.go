package main


import "fmt"




func mergeSort(arr []int) []int{
    if len(arr) <= 1 {
        return arr
    }

    mid := len(arr) / 2
    l , r := 0, 0
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])

    ll, lr := len(left), len(right)

    res := []int{}
    for l < ll && r <  lr  {
        if arr[l] < arr[r] {
            res = append(res, arr[l])
            l++
        } else {
            res = append(res, arr[r])
            r++
        }
  }
    res = append(res, arr[l:]...)
    res = append(res, arr[r:]...)
    return res
}






func main() {
    a := []int{10, 8, 4, 30,5,7,10,20}
    c := mergeSort(a)
    fmt.Println(c)
}
