

func findRepeatNumber(nums []int) int {
    m := make(map[int]int, 0)     //声明并初始化map
    n := make([]int, 0)          // 声明并初始化一个存放结果的slice
    for k, v := range nums {     // 遍历nums
		_, ok := m[v]            // 判断每次遍历的数是否存在于map中，
        if ok {                  //若存在，则说明重复，添加该数字到结果slice中
            n = append(n, v) 
        } else {                 //若不存在，则更新map
            m[v] = k
        }
    }
    return n[0]                  //返回结果slice的第一个值
}

