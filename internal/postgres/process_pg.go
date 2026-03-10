package postgres

import (
	"context"

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

func (p *processRepository) CreateProcess(ctx context.Context, file []byte) (status string, err error) {
	return "ok", nil
}
