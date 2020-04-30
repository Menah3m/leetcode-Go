func exchange(nums []int) []int {
    odd := make([]int, 0)
    even := make([]int, 0)
    for _, v := range nums {
        if v % 2 == 0 {
            even = append(even, v)
        } else {
            odd = append(odd, v)
        }
    }
    return append(odd, even...)
}