package entity

import (
	"context"
)

const (
	ProcessBucket = "process"
)

const (
	InProgressStatus = "inProgress"
)

type Process struct {
	ID                int64  `json:"id" db:"id"`
	UserID            *int64 `json:"user_id" db:"user_id"`
	OriginalFilePath  string `json:"original_file_path" db:"original_file_path"`
	ProcessedFilePath string `json:"processed_file_path" db:"processed_file_path"`
	UploadDatetime    int64  `json:"upload_datetime" db:"upload_datetime"`
	Status            string `json:"status" db:"status"`
}

type ProcessLogic interface {
	CreateProcess(ctx context.Context, file []byte, secondFile []byte, process Process) (err error)
	Upload(ctx context.Context, data []byte) (photoPath string, err error)
}

type ProcessRepository interface {
	CreateProcess(ctx context.Context, process Process) (err error)
}
