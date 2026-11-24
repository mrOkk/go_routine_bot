package bot

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotRegistered = errors.New("user not registered")
	ErrInvalidInvitation = errors.New("invalid invitation")

	ErrThemeNotAvailable = errors.New("topic is not available")

	ErrRejectionReasonRequired = errors.New("a comment is required to reject a submission")

	ErrStudentNotFound           = errors.New("student not found")
	ErrRegisteredWithOtherMentor = errors.New("student is already registered with another mentor")

	ErrNoCurrentTask       = errors.New("no current task")
	ErrTaskMoveUnavailable = errors.New("no task in the requested direction")
	ErrTaskNotSubmittable  = errors.New("task does not accept submissions in its current status")

	ErrSubmissionAlreadyReviewed = errors.New("submission already reviewed")
	ErrDigestTimeAlreadyExists   = errors.New("digest time already exists")
	ErrHelpRequestTooSoon        = errors.New("help request sent too recently")
)

// UserNotRegisteredError возвращается при отказе в доступе.
// UserID — числовой Telegram ID отправителя.
type UserNotRegisteredError struct {
	UserId int64
}

func (e *UserNotRegisteredError) Error() string {
	return fmt.Sprintf("user %d is not registered", e.UserId)
}

func (e *UserNotRegisteredError) Unwrap() error {
	return ErrUserNotRegistered
}

// TopicUnavailableError описывает условия, которые ещё предстоит
// выполнить ученику, чтобы открыть запрошенную тему.
type TopicUnavailableError struct {
	TopicID          string
	PendingTasks     []TaskSummary
	UnpassedTest     *TopicSummary
	UnreceivedTheory *TopicSummary
}

func (e *TopicUnavailableError) Error() string {
	return fmt.Sprintf("topic %q is not available", e.TopicID)
}

func (e *TopicUnavailableError) Unwrap() error {
	return ErrThemeNotAvailable
}
