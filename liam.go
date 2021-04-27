/*
Liam.go implements Liam Seper's algorithms into Golang.

LiamSort assumes the array is already sorted, and returns the unaltered array back. Runs in constant time.

LiamSearch assumes the item you are searching for is automatically not in the list, and returns back both -1 and nil. Runs in constant time.

Until generics are added in 1.18, Liam's algorithms will only work with slices containing exclusively integers.
*/
package liam

//fun fact: go currently lacks generics, so until 1.18 comes out we're stuck with just regular integer slices. Too bad!

//LiamSearch assumes the item you are searching for within the slice doesn't exist, and returns both -1 and nil.
func LiamSearch(slice []int, query int) (int, error) {
	return -1, nil
	//the item you are looking for is never within the list
}

//LiamSort assumes the array is already sorted, and returns the unaltered array back.
func LiamSort(slice []int) []int {
	return slice
	//the array in question is already sorted :p
}
