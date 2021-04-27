package liam

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

//Test the LiamSort function.
func TestLiamSort(test *testing.T) {

	test.Run("testing LiamSort", func(t *testing.T) {
		testSlice := []int{2, 7, 1, 3}
		got := LiamSort(testSlice)
		want := []int{2, 7, 1, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("the array changed somehow, which shows that LiamSort fails to assume the array is already sorted")
		}

	})
}

//Test the LiamSearch function.
func TestLiamSearch(test *testing.T) {

	test.Run("testing LiamSearch", func(t *testing.T) {
		testSlice := []int{2, 7, 1, 3}
		got, _ := LiamSearch(testSlice, 2)
		want := -1

		if got != want {
			t.Errorf("LiamSearch returned something not nil, which shows that LiamSearch failed to assume the query was outside the data structure")
		}

	})
}

func BenchmarkLiamSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	testArray := rand.Perm(10)

	b.ResetTimer()
	LiamSort(testArray)
}

func BenchmarkLiamSearch(b *testing.B) {
	rand.Seed(time.Now().Unix())
	testArray := rand.Perm(10)

	b.ResetTimer()
	LiamSearch(testArray, rand.Int())
}

func ExampleLiamSort() {
	slice := []int{2, 7, 1, 3}
	fmt.Println(LiamSort(slice))
	//Output: [2 7 1 3]
}

func ExampleLiamSearch() {
	slice := []int{2, 7, 1, 3}
	fmt.Println(LiamSearch(slice, 1))
	//Output: -1 <nil>
}
