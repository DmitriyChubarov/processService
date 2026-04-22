package logic

import (
	"context"

	"github.com/DmitriyChubarov/processing/internal/entity"
)

const (
	processBucket = "process"
)

type processLogic struct {
	processRepository entity.ProcessRepository
	minIORepository   entity.MinIo
}

func NewProcessLogic(processRepository entity.ProcessRepository, minIORepository entity.MinIo) entity.ProcessLogic {
	return &processLogic{
		processRepository: processRepository,
		minIORepository:   minIORepository,
	}
}

func (p *processLogic) CreateProcess(ctx context.Context, file []byte) (status string, err error) {
	err = p.Upload(
		ctx,
		processBucket,
		"test",
		file,
		"image/jpeg",
	)
	if err != nil {
		return "cant upload file", err
	}

	return "ok", nil
}

func (p *processLogic) Upload(ctx context.Context, bucket, name string, data []byte, contentType string) error {
	return p.minIORepository.Put(ctx, bucket, name, data, contentType)
}

func (p *processLogic) Download(ctx context.Context, bucket, name string) ([]byte, string, error) {
	return p.minIORepository.Get(ctx, bucket, name)
}
