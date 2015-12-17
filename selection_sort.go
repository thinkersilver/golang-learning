package main 

import "fmt"


func selection_sort(arr []int){
	for i:=len(arr)-1;i>=0;i-- {
		max :=i
		for j:=i;j >=0;j--{
				
				if arr[max] < arr[j]{
					max = j 
				}
				
		}
		
		if i != max {
			temp := arr[i]
			arr[i] = arr[max]
			arr[max] = temp
		}
		
	}
}


func main() {
	x := []int{12,3,2,15,17,23,12,1,8,7,9}
	fmt.Println("unsorted:",x)
	selection_sort(x)
	fmt.Println("sorted:",x)
}