package nolataskmodels

import "time"

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DueDate     *time.Time `json:"due_date"`
	DoDate      *time.Time `json:"do_date"`
	CompletedAt *time.Time `json:"completed_at"`
	Tags        []int      `json:"tag_id"`

	// Is the task due at any time during the given date?
	DueAllDay bool `json:"due_all_day"`
	DoAllDay  bool `json:"do_all_day"`

	// 3 - low priority; 2 - medium priority; 1 - high priority
	Priority *int `json:"priority"`
}

type AtomicTag struct {
	Pinned      int       `json:"pinned"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CompoundTag struct {
	Pinned      int       `json:"pinned"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	QueryString string    `json:"query_string"`
	// json for QueryAST
	QueryAST string `json:"query_ast"`
}

type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type UserCredentials struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"RefreshToken"`
}

type SignUpInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
