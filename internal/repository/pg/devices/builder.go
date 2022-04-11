package devices

import (
	"fmt"
	"smart_house_backend/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const TABLE_NAME = "devices"

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
	rawQuery := psqlSq.Select("id", "device_type_id", "house_part_id").From(qb.getTableName()).Where("id = ?", id)
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepeareCreate(device domain.Device) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Insert(qb.getTableName()).Columns("id", "device_type_id", "house_part_id").Values(device.ID, device.DeviceTypeId, device.HousePartId).Suffix("RETURNING \"id\"")
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepareUpdate(device domain.Device) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Update(qb.getTableName()).Set("device_type_id", device.DeviceTypeId).Set("house_part_id", device.HousePartId).Where("id = ?", device.ID)
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepareDelete(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Delete(id).From(qb.getTableName())
	return rawQuery.ToSql()
}
