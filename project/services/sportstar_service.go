package services

import (
	"app/project/datamodels"
	"app/project/datasource"
	"app/project/repositories"
)

type ISportstarService interface {
	GetAll() []datamodels.StarInfo
	GetByID(id int) *datamodels.StarInfo
	DeleteByID(id int) error
	Update(user *datamodels.StarInfo, columns []string) error
	Create(user *datamodels.StarInfo) error
	Search(country string) []datamodels.StarInfo
}

type SportstarService struct {
	repository *repositories.SportstarRepository
}

func NewSportstarService() ISportstarService {
	return &SportstarService{repository: repositories.NewSportstarRepository(datasource.InstanceMaster())}
}

func (s *SportstarService) GetAll() []datamodels.StarInfo {
	return s.repository.GetAll()
}
func (s *SportstarService) GetByID(id int) *datamodels.StarInfo {
	return s.repository.Get(id)
}
func (s *SportstarService) DeleteByID(id int) error {
	return s.repository.Delete(id)
}
func (s *SportstarService) Update(user *datamodels.StarInfo, columns []string) error {
	return s.repository.Update(user, columns)
}
func (s *SportstarService) Create(user *datamodels.StarInfo) error {
	return s.repository.Create(user)
}
func (s *SportstarService) Search(country string) []datamodels.StarInfo {
	return s.repository.Search(country)
}