package main

import "fmt"

func bubbleSort(elements []int) []int {
	lengthElements := len(elements)
	for i := 0;i<lengthElements;i++ {
		swapped := false
		for j:= 0; j < lengthElements - i - 1;j++ {
			fmt.Println("Start:",elements)
            if elements[j] > elements[j+1]{
				tmp := elements[j]
				elements[j] = elements[j+1]
				elements[j+1] = tmp
				swapped = true
				fmt.Println("Swapped:",elements)
            }
		}
		if swapped == false {
	        break
		}
	}
    return elements
}

func main(){
	elements := [] int{1,5,3,20,19,6,7,2,10}
    //elements := [] int{1,2,3,4,5,6,7,8}
    //elements := [] int{3,4,2,1,10,5,6,7,8}
	fmt.Println("Sorted:",bubbleSort(elements))
}
