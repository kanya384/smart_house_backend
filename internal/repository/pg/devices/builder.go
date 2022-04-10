package devices

import (
	"smart_house_backend/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

type queryBuilder struct {
	tableName string
}

func NewQueryBuilder(tableName string) *queryBuilder {
	return &queryBuilder{
		tableName: tableName,
	}
}

func (qb *queryBuilder) prepareGet(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Select("id", "device_type_id", "house_part_id").From(qb.tableName).Where("id = ?", id)
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepeareCreate(device domain.Device) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Insert(qb.tableName).Columns("id", "device_type_id", "house_part_id").Values(device.ID, device.DeviceTypeId, device.HousePartId).Suffix("RETURNING \"id\"")
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepareUpdate(device domain.Device) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Update(qb.tableName).Set("device_type_id", device.DeviceTypeId).Set("house_part_id", device.HousePartId).Where("id = ?", device.ID)
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepareDelete(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Delete(id)
	return rawQuery.ToSql()
}
