package device_type

import (
	"fmt"
	"smart_house_backend/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const TABLE_NAME = "device_types"

type queryBuilder struct {
	prefix string
}

func NewQueryBuilder(prefix string) *queryBuilder {
	return &queryBuilder{
		prefix: prefix,
	}
}

func (qb *queryBuilder) getTableName() string {
	return fmt.Sprintf("%s.%s", qb.prefix, TABLE_NAME)
}

func (qb *queryBuilder) prepareGet(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Select("id", "name", "photo").From(qb.getTableName()).Where("id = ?", id)
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepeareCreate(deviceType domain.DeviceType) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Insert(qb.getTableName()).Columns("id", "name", "photo").Values(deviceType.ID, deviceType.Name, deviceType.Photo).Suffix("RETURNING \"id\"")
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepareUpdate(deviceType domain.DeviceType) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Update(qb.getTableName()).Set("name", deviceType.Name).Set("photo", deviceType.Photo).Where("id = ?", deviceType.ID)
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepareDelete(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Delete(qb.getTableName()).Where("id = ?", id)
	return rawQuery.ToSql()
}
