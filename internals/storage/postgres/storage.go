package postgres

import "gorm.io/gorm"

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(db *gorm.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (st *PostgresStorage) Close() error {
	sqlDB, err := st.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
