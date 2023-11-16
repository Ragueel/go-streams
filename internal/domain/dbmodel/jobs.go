package dbmodel

import "github.com/google/uuid"

type JobStatus int

const (
	Pending    JobStatus = iota
	Processing JobStatus = iota
	Done       JobStatus = iota
	Failed     JobStatus = iota
)

type ConversionProperties struct {
	TargetFormat string
}

type ConversionJob struct {
	Id           uuid.UUID
	ResourcePath string
	TargetPath   string
	Properties   ConversionProperties
	Status       JobStatus
}

type ResizeProperties struct {
	TargetWidth  int
	TargetHeight int
}

type ResizeJob struct {
	Id           uuid.UUID
	ResourcePath string
	TargetPath   string
	Status       JobStatus
}
