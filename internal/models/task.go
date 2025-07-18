package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// TaskSummary represents a summary of a task, including its type and the count of occurrences.
type TaskSummary struct {
	Type  string // TaskType indicates the type of the task.
	Count int    // Count represents the number of times the task has occurred.
}

// ActiveTask represents a task that is currently active. It contains
// the unique identifier, a brief description associated with the task.
type ActiveTask struct {
	ID          int    // ID is the unique identifier for the task.
	Description string // Description provides a brief overview of the task.
}

// TaskDetails represents the details of a task in the system.
// It includes information such as the task's ID, type, creation and closing dates,
// description, address, customer details, and any comments associated with the task.
type TaskDetails struct {
	ID            int           // Unique identifier for the task
	Type          string        // Type of the task
	CreationDate  time.Time     // Date when the task was created
	ClosingDate   time.Time     // Date when the task was closed
	Description   string        // Description of the task
	Address       string        // Address related to the task
	CustomerName  string        // Name of the customer associated with the task
	CustomerLogin string        // Login identifier of the customer
	Executors     []string      // List of executors related to the task
	Comments      []string      // List of comments related to the task
	Latitude      pgtype.Float8 // Latitude indicates the geographical latitude of the task.
	Longitude     pgtype.Float8 // Longitude indicates the geographical longitude of the task.
}
