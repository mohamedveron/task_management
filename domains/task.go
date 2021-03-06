package domains

import "time"

type EnumtaskState string

const (
	EnumtaskStateInProgress EnumtaskState = "inProgress"
	EnumtaskStateActive  EnumtaskState = "active"
	EnumtaskStateDone    EnumtaskState = "done"
	EnumtaskStateFailed  EnumtaskState = "failed"
)

type Task struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Owner          User          `json:"owner"`
	CreatedAt      time.Time     `json:"created_at"`
	Estimation     float32       `json:"estimation"`
	ReminderPeriod float32       `json:"reminder_period"`
	State          EnumtaskState `json:"state" validate:"oneof=inProgress active done failed"`
}
