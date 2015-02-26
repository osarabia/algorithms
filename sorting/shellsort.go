package main

import "fmt"

func shellSort(elements []int) []int {
	arrayLen := len(elements)
	gap := arrayLen / 2
	fmt.Println("original array:",elements)
	for( gap > 0 ){
		elements = gapInsertSort(gap, elements)
		fmt.Println("gap:",gap)
		fmt.Println("elements:",elements)
		gap = gap / 2
	}
    return elements
}

func shellSortFixedGaps(elements []int, gaps []int) []int{
    fmt.Println("original array:",elements)
    for _, gap := range gaps {
        elements = gapInsertSort(gap, elements)
		fmt.Println("gap:",gap)
		fmt.Println("elements:",elements)
    }

    return elements
}

func gapInsertSort(gap int, elements [] int) [] int{
   swaps := 0
   for outerIndex:=gap; outerIndex < len(elements);outerIndex++ {
	   evaluationElement := elements[outerIndex]
	   innerIndex := outerIndex
	   for innerIndex-gap >= 0 && elements[innerIndex-gap] > evaluationElement {
           elements[innerIndex] = elements[innerIndex-gap]
           innerIndex = innerIndex - gap
           swaps++
	   }
	   elements[innerIndex] = evaluationElement
   }
   fmt.Println("number of swaps:",swaps)
   return elements
}


func main(){
   elements := [] int{20,2,15,10,3,8,1,7,5,11,14,16,21,45,22,34,23,44,555}
   gaps := [] int{7,3,1}
   fmt.Println("sorted:",shellSortFixedGaps(elements, gaps))
}
