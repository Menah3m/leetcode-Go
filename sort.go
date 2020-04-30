func sortArray(nums []int) []int {

    if len(nums) < 2 {
        return nums
    }

    before := make([]int, 0)
    after := make([]int, 0)
    var pivot int
    pivot = nums[len(nums)-1]   //取最后一个元素
    nums = nums[:len(nums)-1]   //删除最后一个元素

    for _, v := range nums {
        if v < pivot {
            before = append(before, v)
        } else {
            after = append(after, v)
        }
    }
    
    r1 := sortArray(before)
    r2 := sortArray(after)
    ret := append(r1,pivot)
    ret = append(ret, r2...)
    return ret
    
}