package main

import "fmt"
func main()  {
	arr:=[][]int{{1,2,3},{4,5,6},{9,8,9}}
	var dlr, drl,diff int
	j:=len(arr)-1
	for i := 0; i < len(arr); i++ {
		for k := 0; k < len(arr[i]); k++ {
			fmt.Print(arr[i][k] ," ")
		}
		fmt.Println()
		dlr = dlr + arr[i][i]
		drl = drl + arr[i][j]
		j=j-1
	}
	if drl<dlr {
		diff=dlr-drl
	}else{
		diff=drl-dlr
	}
	fmt.Print("|",dlr," - ",drl,"| = ",diff)
}