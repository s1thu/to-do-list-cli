package todo

import (
	"testing"
	"time"
)

func TestNewTest(t *testing.T) {
	task := NewTask(1, "Test Task")
	if task.ID != 1 {
		t.Errorf("Expected ID 1, got %d", task.ID)
	}
	if task.Name != "Test Task" {
		t.Errorf("Expected Name 'Test Task', got '%s'", task.Name)
	}
	if task.Completed {
		t.Errorf("Expected Completed false, got true")
	}
	if time.Since(task.CreatedAt) > time.Second {
		t.Errorf("CreatedAt time is not recent")
	}
	if time.Since(task.UpdatedAt) > time.Second {
		t.Errorf("UpdatedAt time is not recent")
	}
}

func TestTaskString(t *testing.T) {
	task := Task{
		ID:        1,
		Name:      "Test Task",
		Completed: true,
		CreatedAt: time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2023, 10, 1, 12, 5, 0, 0, time.UTC),
	}
	expected := "[x] 1: Test Task (Created: 2023-10-01 12:00) (Updated: 2023-10-01 12:05)"
	if task.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, task.String())
	}
}
