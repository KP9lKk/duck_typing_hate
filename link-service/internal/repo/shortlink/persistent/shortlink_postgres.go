package persistent

import (
	"context"
	"duck_typing_hate/link-service/entity"
	"duck_typing_hate/shared/common"
	"duck_typing_hate/shared/pkg/postgres"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const (
	tableName  = "short_link"
	columnId   = "short_link.id"
	allColumns = "short_link.id, short_link.owner, short_link.original_url, short_link.short_code, short_link.clicks"
)

type ShortlinkRepo struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *ShortlinkRepo {
	return &ShortlinkRepo{pg: pg}
}

func (r *ShortlinkRepo) GetById(ctx context.Context, id int) (*entity.ShortLink, error) {
	sl := &entity.ShortLink{}
	ctx, cncl := context.WithTimeout(ctx, common.DBRequestDurartion)
	defer cncl()
	q := fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1", allColumns, tableName, columnId)
	err := r.pg.Conn.QueryRow(ctx, q, id).Scan(&sl.ID, &sl.Owner, &sl.OriginalUrl, &sl.ShortCode, &sl.Clicks)
	if err != nil {
		return nil, err
	}
	return sl, nil
}

func (r *ShortlinkRepo) GetAll(ctx context.Context) (*[]entity.ShortLink, error) {
	ctx, cncl := context.WithTimeout(ctx, common.DBRequestDurartion)
	defer cncl()
	q := fmt.Sprintf("SELECT %s FROM %s", allColumns, tableName)
	rows, err := r.pg.Conn.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.ShortLink])
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ShortlinkRepo) Create(ctx context.Context, sl *entity.ShortLink) error {
	ctx, cncl := context.WithTimeout(ctx, common.DBRequestDurartion)
	defer cncl()
	q := fmt.Sprintf(`
	INSERT INTO %s 
	(short_link.owner, 
	short_link.original_url, 
	short_link.short_code, 
	short_link.clicks) 
	VALUES ($1, $2, $3, $4,)`, tableName)
	_, err := r.pg.Conn.Exec(ctx, q, sl.Owner, sl.OriginalUrl, sl.ShortCode, 0)
	if err != nil {
		return err
	}
	return nil
}

func (r *ShortlinkRepo) Delete(ctx context.Context, id int) error {
	ctx, cncl := context.WithTimeout(ctx, common.DBRequestDurartion)
	defer cncl()
	q := fmt.Sprintf(`
	DELETE FROM %s
	WHERE %s = $1
	`, tableName, columnId)

	_, err := r.pg.Conn.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ShortlinkRepo) Update(ctx context.Context, sl *entity.ShortLink) error {
	ctx, cncl := context.WithTimeout(ctx, common.DBRequestDurartion)
	defer cncl()
	q := fmt.Sprintf(`
	UPDATE %s SET 
	short_link.owner = $1, 
	short_link.original_url = $2, 
	short_link.short_code = $3, 
	short_link.clicks = $4
	`, tableName)

	_, err := r.pg.Conn.Exec(ctx, q, sl.Owner, sl.OriginalUrl, sl.ShortCode, sl.Clicks)
	if err != nil {
		return err
	}
	return nil
}

func (r *ShortlinkRepo) GetByCode(ctx context.Context, code string) (*entity.ShortLink, error) {
	sl := &entity.ShortLink{}
	ctx, cncl := context.WithTimeout(ctx, common.DBRequestDurartion)
	defer cncl()
	q := fmt.Sprintf(`
	SELECT %s 
	FROM %s
	WHERE short_link.short_code = $1
	`, allColumns, tableName)
	err := r.pg.Conn.QueryRow(ctx, q, code).Scan(&sl.ID, &sl.Owner, &sl.OriginalUrl, &sl.ShortCode, &sl.Clicks)
	if err != nil {
		return nil, err
	}
	return sl, nil
}
