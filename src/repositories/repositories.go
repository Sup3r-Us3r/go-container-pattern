package repositories

import (
	"github.com/Sup3r-Us3r/go-container-pattern/src/interfaces"
	"github.com/Sup3r-Us3r/go-container-pattern/src/repositories/installment"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func GetRepositories(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		InstallmentRepository: installment.NewInstallmentRepository(db),
	}
}
