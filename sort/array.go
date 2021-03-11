package sort

type StringArray []string
type IntArray []int

func (arr IntArray) Len() int {
	return len(arr)
}
func (arr IntArray) Less(i int, j int) bool {
	return arr[i] < arr[j]
}
func (arr IntArray) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr StringArray) Len() int {
	return len(arr)
}
func (arr StringArray) Less(i int, j int) bool {
	return arr[i] < arr[j]
}
func (arr StringArray) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
