package main

import (
	"reflect"
	"testing"
)

func TestFindIntersections(t *testing.T) {
	tests := []struct {
		name       string
		a          []int
		b          []int
		wantHas    bool
		wantSliced []int
	}{
		{
			name:       "example from task",
			a:          []int{65, 3, 58, 678, 64},
			b:          []int{64, 2, 3, 43},
			wantHas:    true,
			wantSliced: []int{64, 3},
		},
		{
			name:       "no intersections",
			a:          []int{1, 2, 3},
			b:          []int{4, 5, 6},
			wantHas:    false,
			wantSliced: []int{},
		},
		{
			name:       "empty first slice",
			a:          []int{},
			b:          []int{1, 2, 3},
			wantHas:    false,
			wantSliced: []int{},
		},
		{
			name:       "empty second slice",
			a:          []int{1, 2, 3},
			b:          []int{},
			wantHas:    false,
			wantSliced: []int{},
		},
		{
			name:       "both slices empty",
			a:          []int{},
			b:          []int{},
			wantHas:    false,
			wantSliced: []int{},
		},
		{
			name:       "nil slices",
			a:          nil,
			b:          nil,
			wantHas:    false,
			wantSliced: []int{},
		},
		{
			name:       "duplicates in both slices",
			a:          []int{1, 1, 2, 3, 3},
			b:          []int{3, 3, 1, 1, 4},
			wantHas:    true,
			wantSliced: []int{3, 1},
		},
		{
			name:       "single common value",
			a:          []int{42},
			b:          []int{42},
			wantHas:    true,
			wantSliced: []int{42},
		},
		{
			name:       "negative numbers",
			a:          []int{-1, -2, 0, 5},
			b:          []int{-2, 5, 7},
			wantHas:    true,
			wantSliced: []int{-2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHas, gotSliced := findIntersections(tt.a, tt.b)

			if gotHas != tt.wantHas {
				t.Errorf("findIntersections() has = %v, want %v", gotHas, tt.wantHas)
			}
			if !reflect.DeepEqual(gotSliced, tt.wantSliced) {
				t.Errorf("findIntersections() intersections = %v, want %v", gotSliced, tt.wantSliced)
			}
		})
	}
}
