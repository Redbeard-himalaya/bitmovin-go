package services

import (
	"encoding/json"
	"fmt"

	"github.com/bitmovin/bitmovin-go/bitmovin"
	// "github.com/bitmovin/bitmovin-go/models"
	"../models"
)

type PrewarmedEncoderPoolService struct {
	RestService *RestService
}

const (
	PrewarmedEncoderPoolEndpoint string = "encoding/infrastructure/prewarmed-encoder-pools"
)

func NewPrewarmedEncoderPoolService(bitmovin *bitmovin.Bitmovin) *PrewarmedEncoderPoolService {
	r := NewRestService(bitmovin)
	return &PrewarmedEncoderPoolService{RestService: r}
}

func (s *PrewarmedEncoderPoolService) Create(i *models.CreatePrewarmedEncoderPoolRequest) (*models.PrewarmedEncoderPoolResponse, error) {
	b, err := json.Marshal(*i)
	if err != nil {
		return nil, err
	}
	o, err := s.RestService.Create(PrewarmedEncoderPoolEndpoint, b)
	if err != nil {
		return nil, err
	}
	var r models.PrewarmedEncoderPoolResponse
	if err := json.Unmarshal(o, &r); err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}

func (s *PrewarmedEncoderPoolService) Retrieve(id string) (*models.PrewarmedEncoderPoolResponse, error) {
	path := PrewarmedEncoderPoolEndpoint + "/" + id
	o, err := s.RestService.Retrieve(path)
	if err != nil {
		return nil, err
	}
	var r models.PrewarmedEncoderPoolResponse
	if err := json.Unmarshal(o, &r); err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}

func (s *PrewarmedEncoderPoolService) Delete(id string) (*models.PrewarmedEncoderPoolResponse, error) {
	path := PrewarmedEncoderPoolEndpoint + "/" + id
	o, err := s.RestService.Delete(path)
	if err != nil {
		return nil, err
	}
	var r models.PrewarmedEncoderPoolResponse
	if err := json.Unmarshal(o, &r); err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}

func (s *PrewarmedEncoderPoolService) List(offset int64, limit int64) (*[]models.PrewarmedEncoderPoolDetail, error) {
	o, err := s.RestService.List(PrewarmedEncoderPoolEndpoint, offset, limit)
	if err != nil {
		return nil, err
	}
	var r models.PrewarmedEncoderPoolListResponse
	if err := json.Unmarshal(o, &r); err != nil {
		return nil, err
	} else {
		return &r.Data.Result.Items, nil
	}
}

func (s *PrewarmedEncoderPoolService) Start(id string) (*models.PrewarmedEncoderPoolResponse, error) {
	path := PrewarmedEncoderPoolEndpoint + "/" + id + "/start"
	o, err := s.RestService.Create(path, nil)
	if err != nil {
		return nil, err
	}
	var r models.PrewarmedEncoderPoolResponse
	if err := json.Unmarshal(o, &r); err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}

func (s *PrewarmedEncoderPoolService) Stop(id string) (*models.PrewarmedEncoderPoolResponse, error) {
	path := PrewarmedEncoderPoolEndpoint + "/" + id + "/stop"
	o, err := s.RestService.Create(path, nil)
	if err != nil {
		return nil, err
	}
	var r models.PrewarmedEncoderPoolResponse
	if err := json.Unmarshal(o, &r); err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}
