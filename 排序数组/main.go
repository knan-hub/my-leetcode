package main

func sortArray(nums []int) []int {
	// 冒泡排序
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] > nums[j] {
	// 			nums[i], nums[j] = nums[j], nums[i]
	// 		}
	// 	}
	// }
	// return nums

	// 插入排序
	// for i := 1; i < len(nums); i++ {
	//     j := i - 1
	//     num := nums[i]
	//     for j >= 0 && nums[j] > num {
	//         nums[j + 1] = nums[j]
	//         j--
	//     }
	//     nums[j + 1] = num
	// }
	// return nums

	// 选择排序
	// for i := 1; i < len(nums); i++ {
	//     minIndex := i - 1
	//     minNum := nums[minIndex]
	//     for j := i; j < len(nums); j++ {
	//         if (nums[j] < minNum) {
	//             minIndex = j
	//             minNum = nums[j]
	//         }
	//     }
	//     if minIndex != i - 1 {
	//         nums[minIndex], nums[i - 1] = nums[i - 1], nums[minIndex]
	//     }
	// }
	// return nums

	// 快速排序
	// var sort func(nums []int, left, right int)
	// sort = func(nums []int, left, right int) {
	// 	if left >= right {
	// 		return
	// 	}
	// 	newLeft := left
	// 	newRight := right
	// 	num := nums[newLeft]
	// 	for newLeft < newRight {
	// 		for newLeft < newRight && nums[newRight] >= num {
	// 			newRight--
	// 		}
	// 		nums[newLeft] = nums[newRight]
	// 		for newLeft < newRight && nums[newLeft] <= num {
	// 			newLeft++
	// 		}
	// 		nums[newRight] = nums[newLeft]
	// 	}
	// 	nums[newLeft] = num
	// 	sort(nums, left, newLeft-1)
	// 	sort(nums, newLeft+1, right)
	// }
	// sort(nums, 0, len(nums)-1)
	// return nums

	// 归并排序
	var sort func(nums []int, left, right int)
	tempNums := make([]int, len(nums))
	sort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		middle := (left + right) >> 1
		sort(nums, left, middle)
		sort(nums, middle+1, right)
		newLeft := left
		newRight := middle + 1
		index := left
		for newLeft <= middle && newRight <= right {
			if nums[newLeft] <= nums[newRight] {
				tempNums[index] = nums[newLeft]
				newLeft++
			} else {
				tempNums[index] = nums[newRight]
				newRight++
			}
			index++
		}
		for newLeft <= middle {
			tempNums[index] = nums[newLeft]
			index++
			newLeft++
		}
		for newRight <= right {
			tempNums[index] = nums[newRight]
			index++
			newRight++
		}
		for i := left; i <= right; i++ {
			nums[i] = tempNums[i]
		}
	}
	sort(nums, 0, len(nums)-1)
	return nums
}
