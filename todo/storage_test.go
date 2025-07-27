package todo

import "testing"

func TestSaveTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task 1", Completed: false},
		{ID: 2, Name: "Test Task 2", Completed: true},
	}

	err := SaveTasks("test_tasks.json", tasks)
	if err != nil {
		t.Fatalf("Failed to save tasks: %v", err)
	}
}

func TestLoadTask(t *testing.T) {
	tasks, err := LoadTasks("test_tasks.json")
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	if len(tasks) != 2 {
		t.Fatalf("Expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].Name != "Test Task 1" || tasks[1].Name != "Test Task 2" {
		t.Fatalf("Loaded tasks do not match expected values")
	}
}
