package sqlite

import (
	"github.com/GustavoEklund/payment-gateway-api/domain/entities"
	"github.com/GustavoEklund/payment-gateway-api/infra/repositories/sqlite/fixtures"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTransaction_Insert(t *testing.T) {
	migrationsDir := os.DirFS("./fixtures/sql")
	db := fixtures.Up(migrationsDir)
	defer fixtures.Down(db, migrationsDir)

	sut := NewTransactionSqliteRepository(db)
	err := sut.Save("1", "1", 12.1, entities.STATUS_APPROVED, "")

	assert.Nil(t, err)
}
