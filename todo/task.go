package todo

import (
	"fmt"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask(id int, name string) Task {
	return Task{
		ID:        id,
		Name:      name,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t Task) String() string {
	status := " "
	if t.Completed {
		status = "x"
	}
	return fmt.Sprintf("[%s] %d: %s (Created: %s) (Updated: %s)", status, t.ID, t.Name, t.CreatedAt.Format("2006-01-02 15:04"), t.UpdatedAt.Format("2006-01-02 15:04"))
}
