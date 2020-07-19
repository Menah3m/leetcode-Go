//二分法求平方根
func mySqrt(x int) int {
    
    left := 0
    right := x
    res := 0
    
    for i:=0; i<=x; i++{
        if left > right{
            break
        } else {
            mid := (left + right) / 2
            if mid * mid <= x{
                res = mid
                left = mid + 1
            } else{
                right = mid -1
            }
        }
    }
    return res  
    }
