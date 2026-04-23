package logic

import (
	"context"
	"fmt"

	"github.com/DmitriyChubarov/processing/internal/entity"
	"github.com/google/uuid"
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

func (p *processLogic) CreateProcess(ctx context.Context, firstPhoto []byte, secondPhoto []byte, process entity.Process) (err error) {
	//Сохраняем снимки в MinIO
	process.OriginalFilePath, err = p.Upload(ctx, firstPhoto)
	if err != nil {
		return fmt.Errorf("cant upload first photo: %w", err)
	}

	process.ProcessedFilePath, err = p.Upload(ctx, secondPhoto)
	if err != nil {
		return fmt.Errorf("cant upload second photo: %w", err)
	}

	//Отправляем данные о процессе для сохранения в PostrgeSQL
	err = p.processRepository.CreateProcess(ctx, process)
	if err != nil {
		return fmt.Errorf("cant create process: %w", err)
	}

	return nil
}

func (p *processLogic) Upload(ctx context.Context, data []byte) (photoPath string, err error) {
	photoName := uuid.New().String()
	err = p.minIORepository.Put(ctx, entity.ProcessBucket, photoName, data, "image/jpeg")
	if err != nil {
		return "", err
	}
	photoPath = "http://localhost:9000/" + entity.ProcessBucket + "/" + photoName
	return photoPath, nil
}

func (p *processLogic) Download(ctx context.Context, bucket, name string) ([]byte, string, error) {
	return p.minIORepository.Get(ctx, bucket, name)
}
