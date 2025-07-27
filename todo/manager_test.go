package todo

import "testing"

func TestNewManager(t *testing.T) {
	manager := NewManager()

	if manager == nil {
		t.Error("Expected a new Manager instance, got nil")
	} else {
		if len(manager.tasks) != 0 {
			t.Errorf("Expected tasks slice to be empty, got %d tasks", len(manager.tasks))
		}
	}
}

func TestAddTask(t *testing.T) {
	manager := NewManager()
	task := manager.AddTask("Test Task")

	if task.ID != 1 {
		t.Errorf("Expected task ID 1, got %d", task.ID)
	}
	if task.Name != "Test Task" {
		t.Errorf("Expected task name 'Test Task', got '%s'", task.Name)
	}
	if len(manager.tasks) != 1 {
		t.Errorf("Expected 1 task in manager, got %d", len(manager.tasks))
	}
}

func TestGetTasks(t *testing.T) {
	manager := NewManager()
	manager.AddTask("Task 1")
	manager.AddTask("Task 2")

	tasks := manager.GetTasks()

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
	if tasks[0].Name != "Task 1" || tasks[1].Name != "Task 2" {
		t.Errorf("Tasks are not in the expected order")
	}
}

func TestCompletedTask(t *testing.T) {
	manager := NewManager()
	task := manager.AddTask("Task to Complete")

	err := manager.CompletedTask(task.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !manager.tasks[0].Completed {
		t.Error("Expected task to be marked as completed")
	}
	if manager.tasks[0].UpdatedAt.Before(task.CreatedAt) {
		t.Error("UpdatedAt should be after CreatedAt")
	}

	err = manager.CompletedTask(999) // Non-existent task
	if err != ErrTaskNotFound {
		t.Errorf("Expected ErrTaskNotFound, got %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	manager := NewManager()
	task1 := manager.AddTask("Task 1")
	task2 := manager.AddTask("Task 2")

	err := manager.DeleteTask(task1.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(manager.tasks) != 1 || manager.tasks[0].ID != task2.ID {
		t.Error("Expected only Task 2 to remain after deletion")
	}

	err = manager.DeleteTask(999) // Non-existent task
	if err != ErrTaskNotFound {
		t.Errorf("Expected ErrTaskNotFound, got %v", err)
	}
}
