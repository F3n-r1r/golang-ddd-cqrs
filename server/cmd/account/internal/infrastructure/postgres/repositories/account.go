/*------------------------------------------------------------------------------------*/
//
/*------------------------------------------------------------------------------------*/
package repositories

import (
	"context"
	"golang-ddd-cqrs/cmd/account/internal/domain/model/aggregates"
	"golang-ddd-cqrs/cmd/account/internal/infrastructure/postgres/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// var (
// 	ErrQueryingAccount           = errors.New("could not fetch accounts from the database")
// 	ErrCollectingUsersInSlice    = errors.New("error collecting the results of the accounts query into a slice of the infrastructure model Account")
// 	ErrConvertingToDomainAccount = errors.New("error trying to convert the infrastructure account object to the domain account object ")
// )

type AccountRepository struct {
	db *pgxpool.Pool
}

func NewAccountRepository(db *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Create(ctx context.Context, a *aggregates.Account) error {
	// Map domain object to data object
	dbAccount := models.ToDBAccount(a)

	row, err := r.db.Query(context.Background(),
		`INSERT INTO account (id, email, password) VALUES ($1, $2, $3)`,
		dbAccount.ID, dbAccount.Email, dbAccount.Password)
	defer row.Close()

	if err != nil {
		return err
	}

	return nil
}
