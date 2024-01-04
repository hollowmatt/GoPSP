package appointment

import (
	"fmt"
	"time"
)

// define an appointment
type Appointment struct {
	Start, End  time.Time
	HasConflict bool
}

// simple constructor, rule out invalid start/end
func NewAppointment(start time.Time, end time.Time) (*Appointment, error) {
	if start.After(end) {
		return nil, fmt.Errorf("Invalid: start time must be prior to end... (%v > %v)", start, end)
	}
	return &Appointment{Start: start, End: end}, nil
}

// func (a *Appointment) Equals(b *Appointment) bool {
// 	return a.Start == b.Start && a.End == b.End
// }

// func FindConflicts(appointments []*Appointment) {
// 	for _, a := range appointments {
// 		for _, b := range appointments {
// 			if a.Equals(b) {
// 				a.HasConflict = true
// 				b.HasConflict = true
// 				continue
// 			}
// 			if a.Start.Before(b.Start) && b.End.Before(a.End) {
// 				a.HasConflict = true
// 				b.HasConflict = true
// 			}
// 		}
// 	}
// }

func FindConflicts(appointments []*Appointment) {
	var latest *Appointment
	for _, cur := range appointments {
		if latest != nil && cur.Start.Before(latest.End) {
			cur.HasConflict = true
			latest.HasConflict = true
		}
		if latest == nil || cur.End.After(latest.End) {
			latest = cur
		}
	}
}
