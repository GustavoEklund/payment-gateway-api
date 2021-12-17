package sqlite

import "database/sql"

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) Make() *TransactionSqliteRepository {
	return NewTransactionSqliteRepository(r.DB)
}
