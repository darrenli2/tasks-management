package assignment

import (
	"fmt"
	"math"
)

// Volunteer represents a volunteer with an ID, name, and a list of interested tasks
type Volunteer struct {
	ID              int
	Name            string
	InterestedTasks []*Task
}

// AddInterestedTask adds a task to the volunteer's list of interested tasks
func (v *Volunteer) AddInterestedTask(task *Task) {
	v.InterestedTasks = append(v.InterestedTasks, task)
}

// RemoveInterestedTask removes a task from the volunteer's list of interested tasks
func (v *Volunteer) RemoveInterestedTask(task *Task) {
	for i, t := range v.InterestedTasks {
		if t.ID == task.ID {
			v.InterestedTasks = append(v.InterestedTasks[:i], v.InterestedTasks[i+1:]...)
			return
		}
	}
}

// IsInterested checks if the volunteer is interested in a given task
func (v *Volunteer) IsInterested(task *Task) bool {
	for _, t := range v.InterestedTasks {
		if t.ID == task.ID {
			return true
		}
	}
	return false
}

// InterestScore returns the interest score for a given task
func (v *Volunteer) InterestScore(task *Task) int {
	for i, t := range v.InterestedTasks {
		if t.ID == task.ID {
			return int(math.Max(float64(4-i), 1))
		}
	}
	return 0
}

// SatisfactionScorePerTask returns the satisfaction score for a given task
func (v *Volunteer) SatisfactionScorePerTask(task *Task) int {
	for i, t := range v.InterestedTasks {
		if t.ID == task.ID {
			// If the task is in the top 3 choices, return the score 4 3 2
			// else return 1
			return int(math.Max(float64(4-i), 1))
		}
	}
	// If the task is not in the list, return -1
	return -1
}

func (v Volunteer) String() string {
	return fmt.Sprintf("Volunteer #%d: %s", v.ID, v.Name)
}
