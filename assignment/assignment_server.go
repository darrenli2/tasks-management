package assignment

import (
	"fmt"
	"sort"
)

// AssignmentServer handles the assignment of tasks to volunteers
type AssignmentServer struct {
	Tasks             map[int]*Task
	Volunteers        []*Volunteer
	Assignments       map[int]*Volunteer
	SatisfactionScore int // satisfaction score of the assignment
}

func NewAssignmentServer(tasks map[int]*Task, volunteers []*Volunteer) *AssignmentServer {
	return &AssignmentServer{
		Tasks:       tasks,
		Volunteers:  volunteers,
		Assignments: make(map[int]*Volunteer),
	}
}

// GetInterestedVolunteers returns a list of volunteers interested in a specific task
func (as *AssignmentServer) GetInterestedVolunteers(task *Task) []*Volunteer {
	var interestedVolunteers []*Volunteer
	for _, volunteer := range as.Volunteers {
		if volunteer.IsInterested(task) {
			interestedVolunteers = append(interestedVolunteers, volunteer)
		}
	}
	return interestedVolunteers
}

// AssignTasks assigns tasks to volunteers based on their interests
func (as *AssignmentServer) AssignTasks() {
	for _, task := range as.Tasks {
		interestedVolunteers := as.GetInterestedVolunteers(task)

		// Assign task to volunteers needed by looping through the number of volunteers needed
		for i := 0; i < task.VolunteerNeeded; i++ {
			if len(interestedVolunteers) > 0 {
				// If there are interested volunteers, assign the task to the first one
				as.Assignments[task.ID] = interestedVolunteers[0]
				// Add the satisfaction score for the task
				as.SatisfactionScore += interestedVolunteers[0].SatisfactionScorePerTask(task)
			} else if len(as.Volunteers) > 0 {
				// assign the task to first volunteer
				as.Assignments[task.ID] = as.Volunteers[0]
				// Add the satisfaction score for the task
				as.SatisfactionScore += as.Volunteers[0].SatisfactionScorePerTask(task)
			}
		}
	}
}

// PrintAssignments prints the assignments of tasks to volunteers
func (as *AssignmentServer) PrintAssignments() {
	// Extract and sort task IDs
	taskIDs := make([]int, 0, len(as.Tasks))
	for id := range as.Tasks {
		taskIDs = append(taskIDs, id)
	}
	sort.Ints(taskIDs)

	// Print assignments sorted by task ID
	for _, id := range taskIDs {
		task := as.Tasks[id]
		fmt.Println(task)
		if assignee, ok := as.Assignments[task.ID]; ok {
			fmt.Printf("    Assigned to %s\n", assignee)
		} else {
			fmt.Println("    Unassigned")
		}
		fmt.Println()
	}
}

// PrintSatisfactionScore prints the overall satisfaction score
func (as *AssignmentServer) PrintSatisfactionScore() {
	fmt.Printf("Overall satisfaction score: %d\n", as.SatisfactionScore)
}

// For task 3 I would like to reassign the task based on the volunteer's interest score
// for a tasks interestedVolunteers, calculate the interest score of each volunteer
// then reassign the task to the volunteer with the highest interest score

// ReassignUnsatisfiedTask reassigns the task to the next interested volunteer
// func (as *AssignmentServer) ReassignUnsatisfiedTasks {

// }
