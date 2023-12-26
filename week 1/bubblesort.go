package main 

import(
	"fmt"
)

func main(){
	sli := make([]int, 10, 10)
	fmt.Println("Enter 10 integers separated by space")
	fmt.Scanln(&sli[0], &sli[1], &sli[2], &sli[3], &sli[4], &sli[5], &sli[6], &sli[7], &sli[8], &sli[9])
	BubbleSort(sli)
	fmt.Println("sorted array .............")
	fmt.Println(sli)
}

func BubbleSort(arr []int ){
	for i:=0; i<len(arr)-1; i++ {
		for j:=0; j<len(arr)-i-1; j++ {
			if (arr[j] > arr[j+1]){
				swap(arr, j)
			}
		}
	}
}

func swap(arr []int, ind int){
	temp := arr[ind]
	arr[ind] = arr[ind+1]
	arr[ind+1] = temp
}