// 暴力遍历
func searchMatrix(matrix [][]int, target int) bool {
    for _, v1 := range matrix {
        for _, v2 := range v1 {
            if v2 == target {
                return true
            } 
        }
    }
    return false   
}