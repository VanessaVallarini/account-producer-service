package models

type AccountStatus int32

const (
	Active AccountStatus = iota
	Inactive
)

func (fs AccountStatus) String() string {
	switch fs {
	case Active:
		return "ACTIVE"
	case Inactive:
		return "INACTIVE"
	}

	return "INACTIVE"
}

func FromString(name string) AccountStatus {
	switch name {
	case "ACTIVE":
		return Active
	case "INACTIVE":
		return Inactive
	}

	return Inactive
}
