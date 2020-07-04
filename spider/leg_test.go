package spider

import (
	"math"
	"testing"
)

func TestSetAngleDegrees(t *testing.T) {
	var tests = []struct {
		x, y, z float64
		w, k, h float64
	}{
		{44.9, 44.9, -80.0, 91.09, 180.02, 45.00},
		{44.9, 44.9, 50.0, 51.51, 27.65, 45.00},
		{80.0, 20.0, 50.0, 71.53, 58.93, 75.96},
		{20.0, 80.0, 50.0, 71.53, 58.93, 14.04},
		{80.0, 80.0, 0.0, 90.96, 116.19, 45.0},
		{10.0, 10.0, -100.0, 111.75, 227.95, 45.0},
	}

	for _, tt := range tests {
		l := NewLeg()
		l.SetAbsolutePos(tt.x, tt.y, tt.z)
		wrist, knee, hip := l.GetAngles()
		if math.Abs(wrist-tt.w) > 0.1 || math.Abs(knee-tt.k) > 0.1 || math.Abs(hip-tt.h) > 0.1 {
			t.Errorf("SetAbsolutePos(%f, %f, %f) -> GetAngles(%f, %f, %f); want (%f, %f, %f)", tt.x, tt.y, tt.z, wrist, knee, hip, tt.w, tt.k, tt.h)
		}
	}
}
