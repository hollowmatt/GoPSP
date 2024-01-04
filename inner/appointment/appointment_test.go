package appointment

import (
	"reflect"
	"testing"
	"time"
)

func TestAppointments(t *testing.T) {
	type testCase struct {
		useCase string
		input   []*Appointment
		want    []int
	}

	// helper function to create an appointment
	mkappts := func(start, end int64) *Appointment {
		appt, _ :=
			NewAppointment(time.Unix(start, 0), time.Unix(end, 0))
		return appt
	}

	// tests to run - based on testCase struct
	tests := []testCase{
		{
			useCase: "empty appts",
		},
		{
			useCase: "no conflict",
			input:   []*Appointment{mkappts(10, 15), mkappts(20, 25)},
		},
		{
			useCase: "identical - conflict",
			input:   []*Appointment{mkappts(10, 15), mkappts(10, 15)},
			want:    []int{0, 1},
		},
		{
			useCase: "overlap - conflict",
			input:   []*Appointment{mkappts(5, 10), mkappts(10, 20), mkappts(15, 25), mkappts(25, 30)},
			want:    []int{1, 2},
		},
	}

	//run tests
	for _, test := range tests {

		FindConflicts(test.input)
		var got []int
		for i, appts := range test.input {
			if appts.HasConflict {
				got = append(got, i)
			}
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("%s: results do not match, got %v, want %v", test.useCase, got, test.want)
		}
	}

}
