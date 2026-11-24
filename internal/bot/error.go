package bot

import "errors"

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

	ErrDigestTimeAlreadyExists = errors.New("digest time already exists")
	ErrHelpRequestTooSoon      = errors.New("help request sent too recently")
)
