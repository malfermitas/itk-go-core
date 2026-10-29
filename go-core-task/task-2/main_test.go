package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"mixed", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"all odd", []int{1, 3, 5}, []int{}},
		{"all even", []int{2, 4, 8}, []int{2, 4, 8}},
		{"empty", []int{}, []int{}},
		{"with zero", []int{0, 1}, []int{0}},
		{"negatives", []int{-2, -3, -4}, []int{-2, -4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceExample(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSliceExampleInt64(t *testing.T) {
	in := []int64{1, 2, 3, 4}
	want := []int64{2, 4}
	if got := sliceExample(in); !reflect.DeepEqual(got, want) {
		t.Errorf("sliceExample() = %v, want %v", got, want)
	}
}

func TestSliceExampleOriginalUnchanged(t *testing.T) {
	in := []int{1, 2, 3, 4}
	_ = sliceExample(in)
	if !reflect.DeepEqual(in, []int{1, 2, 3, 4}) {
		t.Errorf("input slice was modified: %v", in)
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name string
		dest []int
		num  int
		want []int
	}{
		{"append to nil dest", nil, 5, []int{5}},
		{"append to empty", []int{}, 7, []int{7}},
		{"append to non-empty", []int{1, 2}, 3, []int{1, 2, 3}},
		{"append to nil", nil, 5, []int{5}},
		{"append negatives", []int{-1}, -2, []int{-1, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElements(tt.dest, tt.num)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addElements(%v, %d) = %v, want %v", tt.dest, tt.num, got, tt.want)
			}
		})
	}
}

func TestAddElementsOriginalUnchanged(t *testing.T) {
	dest := []int{1, 2}
	_ = addElements(dest, 99)
	if !reflect.DeepEqual(dest, []int{1, 2}) {
		t.Errorf("dest slice was modified: %v", dest)
	}
}

func TestCopySlice(t *testing.T) {
	in := []int{1, 2, 3}
	got := copySlice(in)
	if !reflect.DeepEqual(got, in) {
		t.Errorf("copySlice() = %v, want %v", got, in)
	}
	got[0] = -1
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Errorf("original slice was modified after copy: %v", in)
	}
}

func TestCopySliceEmptyAndNil(t *testing.T) {
	if got := copySlice([]int{}); len(got) != 0 {
		t.Errorf("copySlice(empty) = %v, want empty", got)
	}
	if got := copySlice[int](nil); len(got) != 0 {
		t.Errorf("copySlice(nil) = %v, want empty", got)
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		index int
		want  []int
	}{
		{"middle", []int{1, 2, 3, 4}, 1, []int{1, 3, 4}},
		{"first", []int{1, 2, 3}, 0, []int{2, 3}},
		{"last", []int{1, 2, 3}, 2, []int{1, 2}},
		{"single", []int{42}, 0, []int{}},
		{"index out of range too big", []int{1, 2}, 5, []int{1, 2}},
		{"negative index", []int{1, 2}, -1, []int{1, 2}},
		{"empty slice", []int{}, 0, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.input, tt.index)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement(%v, %d) = %v, want %v", tt.input, tt.index, got, tt.want)
			}
		})
	}
}

func TestRemoveElementOriginalUnchanged(t *testing.T) {
	in := []int{1, 2, 3}
	_ = removeElement(in, 0)
	if !reflect.DeepEqual(in, []int{1, 2, 3}) {
		t.Errorf("input slice was modified: %v", in)
	}
}
