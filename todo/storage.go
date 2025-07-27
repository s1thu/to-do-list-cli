package todo

import (
	"encoding/json"
	"os"
)

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadTasks(filename string) ([]Task, error) {
	_, err := os.Stat(filename) // Ensure the file exists
	if os.IsNotExist(err) {
		return nil, nil // Return empty slice if file does not exist
	}
	if err != nil {
		return nil, err // Return error if there is another issue
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil // 🛠️ <- Fix: Handle empty file
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
