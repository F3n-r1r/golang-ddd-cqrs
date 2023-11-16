/*------------------------------------------------------------------------------------*/
// View repository can be accesed from both application and domain layer
/*------------------------------------------------------------------------------------*/
package repositories

import (
	"context"
	"errors"
	"golang-ddd-cqrs/cmd/account/internal/domain/model/aggregates"
	"golang-ddd-cqrs/cmd/account/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrQueryingAccount           = errors.New("could not fetch accounts from the database")
	ErrCollectingUsersInSlice    = errors.New("error collecting the results of the accounts query into a slice of the infrastructure model Account")
	ErrConvertingToDomainAccount = errors.New("error trying to convert the infrastructure account object to the domain account object ")
)

type AccountViewRepository struct {
	db *pgxpool.Pool
}

func NewAccountViewRepository(db *pgxpool.Pool) *AccountViewRepository {
	return &AccountViewRepository{
		db: db,
	}
}

func (r *AccountViewRepository) GetByID(ctx context.Context, id uuid.UUID) (*aggregates.Account, error) {
	row, err := r.db.Query(context.Background(), `SELECT * FROM public.users WHERE id = $1}`, id)
	defer row.Close()

	if err != nil {
		return &aggregates.Account{}, ErrQueryingAccount
	}

	data, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByName[models.Account])
	if err != nil {
		return nil, err
	}

	user, err := models.ToDomainAccount(data)
	if err != nil {
		return &aggregates.Account{}, ErrConvertingToDomainAccount
	}

	return user, nil
}

func (r *AccountViewRepository) GetByEmail(ctx context.Context, email string) (*aggregates.Account, error) {
	row, err := r.db.Query(context.Background(), `SELECT * FROM public.users WHERE email = $1}`, email)
	defer row.Close()

	if err != nil {
		return &aggregates.Account{}, ErrQueryingAccount
	}

	data, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByName[models.Account])
	if err != nil {
		return nil, err
	}

	user, err := models.ToDomainAccount(data)
	if err != nil {
		return &aggregates.Account{}, ErrConvertingToDomainAccount
	}

	return user, nil
}

func (r *AccountViewRepository) GetAll(ctx context.Context) ([]*aggregates.Account, error) {
	rows, err := r.db.Query(context.Background(), `SELECT * FROM public.accounts`)
	defer rows.Close()

	if err != nil {
		return []*aggregates.Account{}, ErrQueryingAccount
	}

	data, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Account])
	if err != nil {
		return []*aggregates.Account{}, ErrCollectingUsersInSlice
	}

	users, err := models.ToDomainAccounts(data)
	if err != nil {
		return []*aggregates.Account{}, ErrConvertingToDomainAccount
	}

	return users, nil
}
