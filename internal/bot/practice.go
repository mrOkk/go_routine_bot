package bot

type TaskStatus int

const (
	StatusTodo TaskStatus = iota
	StatusTesting
	StatusInReview
	StatusDone
)

func (t TaskStatus) String() string {
	switch t {
	case StatusTodo:
		return "TODO"
	case StatusTesting:
		return "TESTING"
	case StatusInReview:
		return "IN REVIEW"
	case StatusDone:
		return "DONE"
	default:
		return "NOT DEFINED"
	}
}

type TaskSummary struct {
	ID     string
	Title  string
	Status TaskStatus // Статус этой задачи у конкретного ученика.
}

type TaskDetails struct {
	Task        TaskSummary
	Description string // Условие задачи.
}

type SubmissionSummary struct {
	ID      string // ID конкретной отправки.
	Task    TaskSummary
	Student StudentSummary // Полезно в общем списке работ ментора.
}

type SubmissionDetails struct {
	Submission SubmissionSummary
	SourceCode string // Текст присланного .go-файла.
}

type ReviewVerdict string

const (
	ReviewAccepted ReviewVerdict = "accepted"
	ReviewRejected ReviewVerdict = "rejected"
)

type ReviewDecision struct {
	Verdict ReviewVerdict
	Comment string // Может быть пустым при принятии.
}
