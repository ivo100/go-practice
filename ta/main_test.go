package main

import (
	"reflect"
	"testing"
)

func TestSeries_IsMonotonic(t *testing.T) {
	type fields struct {
		Data []float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
		want1  bool
	}{
		{"test1", fields{Data: []float64{1, 2, 2, 3, 4, 5}}, true, false},
		{"test2", fields{Data: []float64{4, 3, 3, 2, 2, 1}}, false, true},
		{"test3", fields{Data: []float64{4, 3, 5, 2, 2, 1}}, false, false},
		{"test4", fields{Data: []float64{4, 3, 2, 1, 2, 1}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Series{
				Data: tt.fields.Data,
			}
			got, got1 := s.IsMonotonic()
			if got != tt.want {
				t.Errorf("IsMonotonic() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsMonotonic() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLocalMinMax(t *testing.T) {
	type args struct {
		a []float64
	}
	tests := []struct {
		name     string
		args     args
		wantIMin []int
		wantIMax []int
	}{
		{
			name:     "test1",
			args:     args{a: []float64{1, 3, 2, 4, 3}},
			wantIMin: []int{2},
			wantIMax: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIMin, gotIMax := LocalMinMax(tt.args.a)
			if !reflect.DeepEqual(gotIMin, tt.wantIMin) {
				t.Errorf("LocalMinMax() gotIMin = %v, want %v", gotIMin, tt.wantIMin)
			}
			if !reflect.DeepEqual(gotIMax, tt.wantIMax) {
				t.Errorf("LocalMinMax() gotIMax = %v, want %v", gotIMax, tt.wantIMax)
			}
		})
	}
}
