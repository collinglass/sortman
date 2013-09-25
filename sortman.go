package main

import "log"

func main() {
	u := []int{3, 2, 1, 6, 8, 5, 7}
	log.Println("Sorted:   ", bubbleSort(u))
}

func isSorted(u []int) bool {
	for key := 0; key < len(u)-1; key++ {
		if u[key] >= u[key+1] {
			return false
		}
	}
	return true
}

func bubbleSort(u []int) []int { // Bubble or Exchange sort
	log.Println("Unsorted: ", u)
	for !isSorted(u) {
		for key := 0; key < len(u)-1; key++ {
			if u[key] >= u[key+1] {
				temp := u[key]
				u[key] = u[key+1]
				u[key+1] = temp
				// log.Println(u)
			}
		}
	}
	return u
}

func insertionSort(u []int) []int {
	if !isSorted(u) {
		log.Println("Unsorted: ", u)
		for i := 1; i < len(u); i++ {
			value := u[i]
			j := i - 1
			for j >= 0 && u[j] > value {
				u[j+1] = u[j]
				j = j - 1
			}
			u[j+1] = value
			// log.Println(u)
		}
	}
	return u
}

func selectionSort(u []int) []int {
	tempU := []int{}
	if isSorted(u) {
		return u
	} else {
		log.Println("Unsorted: ", u)
		for len(u) > 0 {
			var j int
			temp := u[0]
			for i := 1; i < len(u); i++ {
				if temp > u[i] {
					temp = u[i]
					j = i
				}
			}
			tempU = append(tempU, temp)
			u = u[:j+copy(u[j:], u[j+1:])]
			// log.Print(tempU, u)
		}
	}
	return tempU
}
