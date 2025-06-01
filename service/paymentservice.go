package service

import (
	"github.com/google/uuid"
	"paymentService/model"
	"paymentService/repository"
)

type PaymentService struct {
	Repo *repository.PaymentRepository
}

func (os *PaymentService) GetPaymentByOrderId(id uuid.UUID) (*model.Payment, error) {
	return os.Repo.GetPaymentByOrderId(id)
}

func (os *PaymentService) GetPayments() ([]model.Payment, error) {
	return os.Repo.GetPayments()
}

//func (os *PaymentService) UpdateOrderStatus(id uuid.UUID, orderstatus string, paymentstatus string) (*model.Order, error) {
//	return os.Repo.UpdateStatus(id, orderstatus, paymentstatus) }
