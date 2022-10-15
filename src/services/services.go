package services

import (
	"github.com/Sup3r-Us3r/go-container-pattern/src/interfaces"
	"github.com/Sup3r-Us3r/go-container-pattern/src/repositories"
	"github.com/Sup3r-Us3r/go-container-pattern/src/services/installment"
)

type ServiceContainer struct {
	InstallmentService interfaces.InstallmentService
}

func GetServices(repositoryContainer repositories.RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		InstallmentService: installment.NewInstallmentService(repositoryContainer),
	}
}
