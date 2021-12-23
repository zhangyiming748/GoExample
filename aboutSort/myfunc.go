package aboutSort

func SelectionSort(nums []int) []int {
	length := len(nums)
	var minIndex, tmp int
	for i := 0; i < length-1; i++ {
		minIndex = i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[minIndex] { // 寻找最小的数
				minIndex = j // 将最小数的索引保存
			}
		}
		tmp = nums[i]
		nums[i] = nums[minIndex]
		nums[minIndex] = tmp
	}
	return nums
}
func BinarySearch(start, end, target int, nums []int) int {
	if start >= end {
		return -1
	}
	mid := (start + end) / 2
	if target > nums[mid] {
		BinarySearch(mid, end, target, nums)
	}
	if target < nums[mid] {
		BinarySearch(start, mid, target, nums)
	}
	return mid
}
