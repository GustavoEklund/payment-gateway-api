package factory

import "github.com/GustavoEklund/payment-gateway-api/domain/repositories"

type SaveTransactionRepositoryFactory interface {
	Make() repositories.SaveTransactionRepository
}
