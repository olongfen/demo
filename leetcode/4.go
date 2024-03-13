package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	if len(nums1) == 0 {
		return cal(nums2)
	}
	if len(nums2) == 0 {
		return cal(nums1)
	}
	var (
		left, right int
		max         = nums1[len(nums1)-1]
	)
	for left < len(nums1) && right < len(nums2) {
		if nums1[left] >= nums2[right] {
			n := make([]int, len(nums1[left:]))
			copy(n, nums1[left:])
			nums1 = append(nums1[:left], nums2[right])
			nums1 = append(nums1, n...)
			right++
		} else if nums2[right] >= max {
			nums1 = append(nums1, nums2[right:]...)
			break
		} else {
			left++
		}
	}
	return cal(nums1)

}

func cal(nums1 []int) float64 {
	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2-1]) / 2
	}
	return float64(nums1[len(nums1)/2])
}
