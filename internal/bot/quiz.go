package bot

type AnswerOption struct {
	ID   string
	Text string
}

type TestQuestion struct {
	ID      string
	Text    string
	Options []AnswerOption // По требованиям ровно четыре варианта.
}

type QuestionResult struct {
	QuestionID string
	IsCorrect  bool
}

type TestResult struct {
	ScorePercent int // 0–100 за завершённую попытку.
	Passed       bool
	Questions    []QuestionResult
}

type TestStep struct {
	NextQuestion *TestQuestion
	Result       *TestResult
}
