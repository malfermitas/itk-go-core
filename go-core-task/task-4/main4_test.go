package main

import (
	"reflect"
	"testing"
)

func TestSliceDifference(t *testing.T) {
	tests := []struct {
		name   string
		first  []string
		second []string
		want   []string
	}{
		{
			name:   "example from task",
			first:  []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			second: []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:   "no common elements",
			first:  []string{"a", "b"},
			second: []string{"c", "d"},
			want:   []string{"a", "b"},
		},
		{
			name:   "all elements are present in second",
			first:  []string{"a", "b"},
			second: []string{"b", "a"},
			want:   []string{},
		},
		{
			name:   "second slice is empty",
			first:  []string{"a", "b"},
			second: []string{},
			want:   []string{"a", "b"},
		},
		{
			name:   "first slice is empty",
			first:  []string{},
			second: []string{"a", "b"},
			want:   []string{},
		},
		{
			name:   "nil slices",
			first:  nil,
			second: nil,
			want:   []string{},
		},
		{
			name:   "duplicates in first are preserved",
			first:  []string{"a", "a", "b", "c", "a"},
			second: []string{"c"},
			want:   []string{"a", "a", "b", "a"},
		},
		{
			name:   "duplicates in second do not affect result",
			first:  []string{"a", "b", "c"},
			second: []string{"b", "b", "b"},
			want:   []string{"a", "c"},
		},
		{
			name:   "empty strings are handled as regular values",
			first:  []string{"", "a", ""},
			second: []string{"a"},
			want:   []string{"", ""},
		},
		{
			name:   "case sensitive comparison",
			first:  []string{"Apple", "apple"},
			second: []string{"apple"},
			want:   []string{"Apple"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceDifference(tt.first, tt.second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceDifference(%v, %v) = %v, want %v", tt.first, tt.second, got, tt.want)
			}
		})
	}
}

func TestSliceDifferenceDoesNotModifyInput(t *testing.T) {
	first := []string{"apple", "banana", "cherry"}
	second := []string{"banana"}
	firstCopy := append([]string(nil), first...)
	secondCopy := append([]string(nil), second...)

	SliceDifference(first, second)

	if !reflect.DeepEqual(first, firstCopy) {
		t.Errorf("first slice was modified: got %v, want %v", first, firstCopy)
	}
	if !reflect.DeepEqual(second, secondCopy) {
		t.Errorf("second slice was modified: got %v, want %v", second, secondCopy)
	}
}

func TestSliceDifferenceReturnsNonNilEmptySlice(t *testing.T) {
	got := SliceDifference([]string{"a"}, []string{"a"})
	if got == nil {
		t.Fatal("expected non-nil empty slice, got nil")
	}
	if len(got) != 0 {
		t.Errorf("expected empty slice, got %v", got)
	}
}
