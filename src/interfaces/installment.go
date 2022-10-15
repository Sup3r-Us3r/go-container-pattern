package interfaces

import "github.com/Sup3r-Us3r/go-container-pattern/src/structs"

type InstallmentRepository interface {
	Create(installment structs.Installment) error
}

type InstallmentService interface {
	Create(installment structs.Installment) error
}
