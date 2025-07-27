package todo

import (
	"errors"
	"slices"
	"sort"
	"time"
)

var ErrTaskNotFound = errors.New("task not found")

type Manager struct {
	tasks  []Task
	nextId int
}

func NewManager() *Manager {
	return &Manager{
		tasks:  []Task{},
		nextId: 1,
	}
}

func (m *Manager) AddTask(name string) Task {
	task := NewTask(m.nextId, name)
	m.tasks = append(m.tasks, task)
	m.nextId++
	return task
}

func (m *Manager) GetTasks() []Task {
	sortedTasks := make([]Task, len(m.tasks))
	copy(sortedTasks, m.tasks)
	sort.Slice(sortedTasks, func(i, j int) bool {
		return sortedTasks[i].ID < sortedTasks[j].ID
	})
	return sortedTasks
}

func (m *Manager) CompletedTask(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks[i].Completed = true
			m.tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return ErrTaskNotFound
}

func (m *Manager) DeleteTask(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks = slices.Delete(m.tasks, i, i+1)
			if m.nextId > 1 {
				m.nextId--
			}
			return nil
		}
	}
	return ErrTaskNotFound
}
