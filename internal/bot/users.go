package bot

type UserProfile struct {
	ID       int64  // Telegram ID отправителя.
	Name     string // Имя для показа в боте.
	Username string // Пустая строка, если никнейма нет.
}

type Role string

const (
	RoleStudent Role = "student"
	RoleMentor  Role = "mentor"
)

type Session struct {
	Role       Role
	MentorName string // Имя для приветствия после регистрации ученика.
}

type StudentSummary struct {
	ID       int64  // Числовой Telegram ID, если используем его как ID ученика.
	Name     string // Имя для показа в списке.
	Username string // Пустая строка, если никнейма нет.
}

type DigestTime struct {
	ID     string // Нужен, если ментор удаляет время по ID.
	Hour   int    // 0–23.
	Minute int    // 0–59.
}

type HelpRequest struct {
	Student StudentSummary
	Comment string
}
