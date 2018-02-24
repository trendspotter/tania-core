package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	TaskCreatedCode   = "TaskCreated"
	TaskCompletedCode = "TaskCompleted"
	TaskCancelledCode = "TaskCancelled"
	TaskDueCode       = "TaskDue"
	TaskModifiedCode  = "TaskModified"
)

type TaskCreated struct {
	Title         string
	UID           uuid.UUID
	Description   string
	CreatedDate   time.Time
	DueDate       *time.Time
	Priority      string
	Status        string
	Domain        string
	DomainDetails TaskDomain
	Category      string
	IsDue         bool
	AssetID       *uuid.UUID
}

type TaskModified struct {
	UID           uuid.UUID
	Title         string
	Description   string
	Priority      string
	DueDate       *time.Time
	Domain        string
	DomainDetails TaskDomain
	Category      string
	AssetID       *uuid.UUID
}

type TaskCompleted struct {
	UID           uuid.UUID
	Title         string
	Description   string
	Priority      string
	DueDate       *time.Time
	Domain        string
	DomainDetails TaskDomain
	Category      string
	AssetID       *uuid.UUID
	CompletedDate *time.Time
}

type TaskCancelled struct {
	UID           uuid.UUID
	Title         string
	Description   string
	Priority      string
	DueDate       *time.Time
	Domain        string
	DomainDetails TaskDomain
	Category      string
	AssetID       *uuid.UUID
	CancelledDate *time.Time
}

type TaskDue struct {
	UID           uuid.UUID
	Title         string
	Description   string
	Priority      string
	DueDate       *time.Time
	Domain        string
	DomainDetails TaskDomain
	Category      string
	AssetID       *uuid.UUID
}
