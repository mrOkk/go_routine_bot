package bot

type TopicSummary struct {
	ID          string
	Title       string
	IsAvailable bool // Ученик может получить данные темы.
	IsCompleted bool // Ученик прошёл тему.
	HasTasks    bool
	HasTest     bool
}

type TopicProgress struct {
	Topic           TopicSummary
	BestTestPercent *int // nil: результата ещё нет; 0: есть попытка с 0%.
}

type StudentProgress struct {
	CompletedTopics []TopicProgress
	WaitingReview   []SubmissionSummary
}

type LearningPlan struct {
	Topics         []TopicSummary
	CurrentTopicID string
}

type TopicDetails struct {
	Topic     TopicSummary
	TheoryURL string
}

type TopicTasks struct {
	Tasks         []TaskSummary
	CurrentTaskID string // Пустая строка, если текущей задачи нет.
}

type TaskDirection string

const (
	TaskPrevious TaskDirection = "previous"
	TaskNext     TaskDirection = "next"
)
