package findmediansortedarrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{"case 1", []int{1, 3}, []int{2}, 2.0},
		{"case 2", []int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
