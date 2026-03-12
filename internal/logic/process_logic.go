package logic

import (
	"context"

	"github.com/DmitriyChubarov/processing/internal/entity"
)

type processLogic struct {
	processRepository entity.ProcessRepository
}

func NewProcessLogic(processRepository entity.ProcessRepository) entity.ProcessLogic {
	return &processLogic{
		processRepository: processRepository,
	}
}

func (p *processLogic) CreateProcess(ctx context.Context, file []byte) (status string, err error) {
	return "ok", nil
}
