package service

import (
	"banking-system/dto"
	"banking-system/entity"
	"banking-system/repository"
	"github.com/mashingan/smapping"
	"log"
)

type ManagerService interface {
	Insert(p dto.ManagerCreateDTO) entity.Password
	Update(p dto.ManagerUpdateDTO) entity.Password
	Delete(p entity.Password)
	All() []entity.Password
}

type managerService struct {
	managerRepository repository.ManagerRepository
}

func NewManagerService(managerRepo repository.ManagerRepository) ManagerService {
	return &managerService{
		managerRepository: managerRepo,
	}
}

func (service *managerService) Insert(b dto.ManagerCreateDTO) entity.Password {
	manager := entity.Password{}
	err := smapping.FillStruct(&manager, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.managerRepository.InsertPassword(manager)
	return res
}

func (service *managerService) Update(b dto.ManagerUpdateDTO) entity.Password {
	manager := entity.Password{}
	err := smapping.FillStruct(&manager, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.managerRepository.UpdatePassword(manager)
	return res
}

func (service *managerService) Delete(b entity.Password) {
	service.managerRepository.DeletePassword(b)
}

func (service *managerService) All() []entity.Password {
	return service.managerRepository.GetAll()
}
