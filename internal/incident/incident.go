package incident

import "golang.org/x/exp/slices"

type Incident struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func (i *Incident) Check() bool {
	return i.Topic != "" && i.checkStatus()
}

func (i *Incident) checkStatus() bool {
	allowedStatuses := []string{"closed", "active"}
	return slices.Contains(allowedStatuses, i.Status)
}