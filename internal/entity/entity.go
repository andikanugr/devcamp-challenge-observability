package entity

type Priority int

const (
	PriorityLow Priority = iota + 1
	PriorityMedium
	PriorityHigh
)

type Todo struct {
	ID          int      `json:"id,omitempty" db:"id" swaggerignore:"true""`
	Title       string   `json:"title" db:"title"`
	Description string   `json:"description" db:"description"`
	Priority    Priority `json:"priority" db:"priority"`
	IsDone      bool     `json:"is_done" db:"is_done"`
}

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "Low"
	case PriorityMedium:
		return "Medium"
	case PriorityHigh:
		return "High"
	default:
		return "Unknown"
	}
}
