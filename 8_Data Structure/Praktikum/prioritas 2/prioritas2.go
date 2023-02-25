package main
import "fmt"
func main()  {
	fmt.Println(pairSum([]int{1,2,3,4,6},6))
	fmt.Println(pairSum([]int{2,5,9,11},11))
	fmt.Println(pairSum([]int{1,3,5,7},12))
	fmt.Println(pairSum([]int{1,4,6,8},10))
	fmt.Println(pairSum([]int{1,5,6,7},6))
}
func pairSum(arr []int, target int)[]int  {
	result:=[]int{}
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			sum := arr[i]+arr[j]
			// fmt.Println(sum)
			if sum==target {
				result = append(result,i,j )

			}
		}
	}
	return result
}