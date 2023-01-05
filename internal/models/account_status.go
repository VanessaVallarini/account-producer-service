package models

type AccountStatus int32

const (
	Active AccountStatus = iota
	Disabled
)

func (as AccountStatus) String() string {
	switch as {
	case Active:
		return "ACTIVE"
	case Disabled:
		return "DESABLE"
	}

	return "DESABLE"
}

func AccountStatusString(status string) AccountStatus {
	switch status {
	case "ACTIVE":
		return Active
	case "DESABLE":
		return Disabled
	}

	return Disabled
}
