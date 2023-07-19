package main

import "time"

type Task struct {
	TaskName  string
	StartTime time.Time
}

var currentTask Task = Task{}

func AddTask(taskName string) Task {
	return Task{TaskName: taskName, StartTime: time.Now()}
}

func PrintTask(task Task) string {
	// Check for empty struct
	if (Task{} == task) {
		return "Current task is empty"
	} else {
		return "Task - " + task.TaskName
	}
}
