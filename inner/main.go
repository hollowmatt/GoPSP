package main

import (
	"fmt"
	"time"

	"hollowmatt.com/inner/appointment"
)

func main() {
	mkappointments := func(start, end int64) *appointment.Appointment {
		a, _ := appointment.NewAppointment(time.Unix(start, 0), time.Unix(end, 0))
		return a
	}

	appts := []*appointment.Appointment{mkappointments(10, 20), mkappointments(25, 30)}
	for _, appt := range appts {
		fmt.Println("Start time: ", appt.Start, ", End Time: ", appt.End, ", conflict: ", appt.HasConflict)
	}

	moreAppts := []*appointment.Appointment{
		mkappointments(10, 100),
		mkappointments(15, 95),
		mkappointments(20, 90),
		mkappointments(25, 85),
		mkappointments(50, 55),
		mkappointments(105, 110),
		mkappointments(125, 140),
		mkappointments(125, 140),
	}

	appointment.FindConflicts(moreAppts)
	for _, appt := range moreAppts {
		fmt.Println("Start time: ", appt.Start, ", End Time: ", appt.End, ", conflict: ", appt.HasConflict)
	}
}
