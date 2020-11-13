package gts

import (
	"log"
	"testing"

	"github.com/twpayne/go-geom"
)

func TestLocatePointInLinearRing(t *testing.T) {
	square1x1 := geom.NewLinearRing(geom.XY)
	square1x1, err := square1x1.SetCoords([]geom.Coord{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	})

	if err != nil {
		log.Println("SetCoords", err)
		t.Fail()
	}

	type args struct {
		point *geom.Coord
		ring  *geom.LinearRing
	}
	tests := []struct {
		name string
		args args
		want Location
	}{
		{
			name: "Test interior",
			args: args{
				point: &geom.Coord{0.5, 0.5},
				ring:  square1x1,
			},
			want: Interior,
		},
		{
			name: "Test exterior",
			args: args{
				point: &geom.Coord{1.5, 0.5},
				ring:  square1x1,
			},
			want: Exterior,
		},
		{
			name: "Test boundary",
			args: args{
				point: &geom.Coord{1, 0.5},
				ring:  square1x1,
			},
			want: Boundary,
		},
		{
			name: "Test boundary horizontal",
			args: args{
				point: &geom.Coord{0.5, 0},
				ring:  square1x1,
			},
			want: Boundary,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LocatePointInLinearRing(tt.args.point, tt.args.ring); got != tt.want {
				t.Errorf("LocatePointInLinearRing() = %v, want %v", got, tt.want)
			}
		})
	}
}
