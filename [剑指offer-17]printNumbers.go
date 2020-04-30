func printNumbers(n int) []int {
    max := 1
    for i:= 0; i < n; i++ {
        max = max * 10
    }
    re := make([]int, 0)
    for i := 1; i < max; i++ {
        re = append(re, i)
    }

    return re
}