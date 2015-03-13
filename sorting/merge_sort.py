#!/usr/bin/python


def mergeSort(arr):
    if len(arr) > 1:
        middle = len(arr) / 2
        firstHalf = arr[:middle]
        secondHalf = arr[middle:]

        mergeSort(firstHalf)
        mergeSort(secondHalf)

        i = 0
        j = 0
        k = 0

        #merging
        while i < len(firstHalf) and j <len(secondHalf):
            if firstHalf[i] < secondHalf[j]:
                arr[k] = firstHalf[i]
                i = i + 1
            else:
                arr[k] = secondHalf[j]
                j = j + 1
            k = k + 1

        while i < len(firstHalf):
            arr[k] = firstHalf[i]
            i = i + 1
            k = k + 1

        while j < len(secondHalf):
            arr[k] = secondHalf[j]
            j = j + 1
            k = k + 1

    return arr

if __name__ == "__main__":
    print mergeSort([7,1,9,11])
    print mergeSort([7,2,9,11,1])
    print mergeSort([44,11,93,17,77,31,22,55,1])
