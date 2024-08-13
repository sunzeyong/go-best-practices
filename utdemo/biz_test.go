package utdemo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

// 使用testify库
func TestQuickSort(t *testing.T) {
	// 表驱动ut
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"case1", []int{7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"case2", []int{1, 6, 3, 1, 5, 4, 1}, []int{1, 1, 1, 3, 4, 5, 6}},
		{"case3", []int{3, 4, 0, -1, -1, 5}, []int{-1, -1, 0, 3, 4, 5}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			QuickSort(c.input)
			// 使用断言
			assert.Equal(t, c.expected, c.input)
		})
	}
}

// 使用goconvey, 提供了更好的结果展示和UI
func TestHeapSort(t *testing.T) {
	Convey("基础测试", t, func() {
		tt := []struct {
			name     string
			input    []int
			expected []int
		}{
			{"case1", []int{7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7}},
			{"case2", []int{1, 6, 3, 1, 5, 4, 1}, []int{1, 1, 1, 3, 4, 5, 6}},
			{"case3", []int{3, 4, 0, -1, -1, 5}, []int{-1, -1, 0, 3, 4, 5}},
		}

		for _, tc := range tt {
			Convey(tc.name, func() {
				HeapSort(tc.input)
				So(tc.input, ShouldResemble, tc.expected)
			})
		}
	})
}
