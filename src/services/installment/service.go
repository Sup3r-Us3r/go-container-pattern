package installment

import (
	"errors"
	"time"

	"github.com/Sup3r-Us3r/go-container-pattern/src/interfaces"
	"github.com/Sup3r-Us3r/go-container-pattern/src/repositories"
	"github.com/Sup3r-Us3r/go-container-pattern/src/structs"
)

type InstallmentService struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func NewInstallmentService(repositoryContainer repositories.RepositoryContainer) *InstallmentService {
	return &InstallmentService{
		InstallmentRepository: repositoryContainer.InstallmentRepository,
	}
}

func (is *InstallmentService) Create(installment structs.Installment) error {
	if installment.DueDay < time.Now().Day() {
		return errors.New("invalid installment")
	}

	return is.InstallmentRepository.Create(installment)
}
