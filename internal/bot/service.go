package bot

import "context"

type Access interface {
	Start(ctx context.Context, user UserProfile, inviteToken string) (Session, error)
	CreateInvitation(ctx context.Context, mentorID int64) (string, error)
	DisconnectStudent(ctx context.Context, studentID int64) error
}

type Learning interface {
	GetLearningPlan(ctx context.Context, studentID int64) (LearningPlan, error)
	OpenTopic(ctx context.Context, studentID int64, topicID string) (TopicDetails, error)
	ListTasks(ctx context.Context, studentID int64, topicID string) (TopicTasks, error)
	GetTask(ctx context.Context, studentID int64, taskID string) (TaskDetails, error)
	GetCurrentTask(ctx context.Context, studentID int64) (TaskDetails, error)
	MoveTask(ctx context.Context, studentID int64, direction TaskDirection) (TaskDetails, error)
}

type Submissions interface {
	SubmitSolution(ctx context.Context, studentID int64, taskID string, source []byte) (SubmissionSummary, error)
	ListPendingSubmissions(ctx context.Context, mentorID int64) ([]SubmissionSummary, error)
	ListStudentPendingSubmissions(ctx context.Context, mentorID, studentID int64) ([]SubmissionSummary, error)
	GetSubmission(ctx context.Context, mentorID int64, submissionID string) (SubmissionDetails, error)
	ReviewSubmission(ctx context.Context, mentorID int64, submissionID string, decision ReviewDecision) error
}

type Quizzes interface {
	StartOrResumeQuiz(ctx context.Context, studentID int64, topicID string) (TestQuestion, error)
	AnswerQuestion(ctx context.Context, studentID int64, questionID, optionID string) (TestStep, error)
	ExitQuiz(ctx context.Context, studentID int64) error
}

type Mentoring interface {
	ListStudents(ctx context.Context, mentorID int64) ([]StudentSummary, error)
	FindStudentByUsername(ctx context.Context, mentorID int64, username string) (StudentSummary, error)
	GetStudentProgress(ctx context.Context, mentorID, studentID int64) (StudentProgress, error)
	RemoveStudent(ctx context.Context, mentorID, studentID int64, comment string) error

	ListDigestTimes(ctx context.Context, mentorID int64) ([]DigestTime, error)
	AddDigestTime(ctx context.Context, mentorID int64, hour, minute int) (DigestTime, error)
	RemoveDigestTime(ctx context.Context, mentorID int64, digestTimeID string) error
}

type Help interface {
	RequestHelp(ctx context.Context, studentID int64, comment string) error
}
