package appointment

import (
	"testing"
	"time"
)

func TestConflicts(t *testing.T) {
	type testCase struct {
		name  string
		input []*Appointment
		want  []int
	}
	mkappointments := func(start, end int64) *Appointment {
		a, _ := NewAppointment(time.Unix(start, 0), time.Unix(end, 0))
		return a
	}
	tests := []testCase{
		{
			name: "empty input",
		},
		{
			name:  "no conflict",
			input: []*Appointment{mkappointments(10, 20), mkappointments(25, 30)},
		},
		{
			name:  "back to back -> no conflict",
			input: []*Appointment{mkappointments(10, 20), mkappointments(20, 30)},
		},
		{
			name:  "identical -> conflict",
			input: []*Appointment{mkappointments(10, 20), mkappointments(10, 20)},
			want:  []int{0, 1},
		},
		{
			name:  "partial overlap -> conflict",
			input: []*Appointment{mkappointments(10, 20), mkappointments(15, 30)},
			want:  []int{0, 1},
		},
		{
			name: "nested",
			input: []*Appointment{
				mkappointments(10, 100),
				mkappointments(15, 95),
				mkappointments(20, 90),
				mkappointments(25, 85),
				mkappointments(50, 55),
				mkappointments(105, 110),
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "double nested",
			input: []*Appointment{
				mkappointments(10, 100),
				mkappointments(20, 30),
				mkappointments(40, 50),
			},
			want: []int{0, 1, 2},
		},
		{
			name: "two separate conflicts",
			input: []*Appointment{
				mkappointments(10, 20),
				mkappointments(15, 25),
				mkappointments(30, 40),
				mkappointments(60, 70),
				mkappointments(62, 64),
			},
			want: []int{0, 1, 3, 4},
		},
	}
	for _, test := range tests {
		FindConflicts(test.input)
		var got []int
		for i, appointments := range test.input {
			if HasConflict {
				got = append(got, i)
			}
		}
	}
}
