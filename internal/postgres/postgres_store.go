package postgres

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kashifsoofi/payment-gateway/internal"
	"github.com/kashifsoofi/payment-gateway/internal/config"
)

const driverName = "pgx"

type postgresStore struct {
	cfg config.Database
	dbx *sqlx.DB
}

func NewPostgresStore(cfg config.Database) *postgresStore {
	postgresStore := &postgresStore{
		cfg: cfg,
	}

	return postgresStore
}

func (s *postgresStore) Connect(ctx context.Context) error {
	dbx, err := sqlx.ConnectContext(ctx, driverName, s.cfg.DatabaseURL)
	if err != nil {
		return err
	}

	s.dbx = dbx
	return nil
}

func (s *postgresStore) Close() error {
	return s.dbx.Close()
}

//go:embed sql/get.sql
var getSql string

func (s *postgresStore) Get(ctx context.Context, id uuid.UUID) (*internal.Payment, error) {
	err := s.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.Close()

	var payment internal.Payment
	if err := s.dbx.GetContext(
		ctx,
		&payment,
		getSql,
		id); err != nil {
		if err != sql.ErrNoRows {
			return nil, fmt.Errorf("could not get payment, err: %w", err)
		}

		return nil, errors.New("not found")
	}

	return &payment, nil
}

//go:embed sql/list.sql
var listSql string

func (s *postgresStore) List(ctx context.Context, merchantId uuid.UUID) ([]*internal.Payment, error) {
	err := s.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer s.Close()

	var payments []*internal.Payment
	if err := s.dbx.SelectContext(
		ctx,
		&payments,
		listSql); err != nil {
		return nil, fmt.Errorf("could not list movies, err: %w", err)
	}
	return payments, nil
}

//go:embed sql/create.sql
var createSql string

func (s *postgresStore) Create(ctx context.Context, payment *internal.Payment) error {
	err := s.Connect(ctx)
	if err != nil {
		return err
	}
	defer s.Close()

	if _, err := s.dbx.NamedExecContext(
		ctx,
		createSql,
		payment); err != nil {
		return fmt.Errorf("count not insert payment, err: %w", err)
	}

	return nil
}
