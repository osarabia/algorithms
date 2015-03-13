#!/usr/bin/python


def merge(x, y):

    result = []
    i = 0
    j = 0

    while i < len(x) and j < len(y):
        if x[i] < y[j]:
            result.append(x[i])
            i = i + 1
        else:
            result.append(y[j])
            j = j + 1

    if i == len(x):
        result = result + y[j:]

    if j == len(y):
        result = result + x[i:]

    return result

if __name__ == "__main__":
    merge([1,2],[2,3])
    merge([1,9],[3,5])

