package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	im := m - 1
	in := n - 1
	for i := 0; i < m+n; i++ {
		if im >= 0 && in >= 0 {
			if nums1[im] >= nums2[in] {
				nums1[len(nums1)-i-1] = nums1[im]
				im--
			} else {
				nums1[len(nums1)-i-1] = nums2[in]
				in--
			}
			continue
		}
		if im < 0 {
			nums1[len(nums1)-i-1] = nums2[in]
			in--
		} else {
			nums1[len(nums1)-i-1] = nums1[im]
			im--
		}
	}
}
