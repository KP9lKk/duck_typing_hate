package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	Conn *pgx.Conn
}

func New(url string, ctx context.Context) (*Postgres, error) {
	pg, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgx.Connect: %w", err)
	}
	return &Postgres{Conn: pg}, nil
}

func (p *Postgres) Close(ctx context.Context) {
	if p.Conn != nil {
		p.Conn.Close(ctx)
	}
}
