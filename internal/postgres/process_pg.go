package postgres

import (
	"context"
	"fmt"

	"github.com/DmitriyChubarov/processing/internal/entity"
	"github.com/gocraft/dbr/v2"
)

type processRepository struct {
	session *dbr.Session
}

func NewProcessRepository(session *dbr.Session) entity.ProcessRepository {
	return &processRepository{
		session: session,
	}
}

func (p *processRepository) CreateProcess(ctx context.Context, process entity.Process) (err error) {
	query := "INSERT INTO processing (user_id, original_file_path, processed_file_path) VALUES ($1, $2, $3)"

	_, err = p.session.Exec(query, process.UserID, process.OriginalFilePath, process.ProcessedFilePath)
	if err != nil {
		return fmt.Errorf("cant save process: %w", err)
	}

	return nil
}
