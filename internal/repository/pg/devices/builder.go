package devices

import (
	"smart_house_backend/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const TABLE_NAME = "devices"

func prepareGet(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Select("id", "device_type_id", "house_part_id").From(TABLE_NAME).Where("id = ?", id)
	return rawQuery.ToSql()
}

func prepeareCreate(device domain.Device) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Insert(TABLE_NAME).Columns("id", "device_type_id", "house_part_id").Values(device.ID, device.DeviceTypeId, device.HousePartId).Suffix("RETURNING \"id\"")
	return rawQuery.ToSql()
}

func prepareUpdate(device domain.Device) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Update(TABLE_NAME).Set("device_type_id", device.DeviceTypeId).Set("house_part_id", device.HousePartId).Where("id = ?", device.ID)
	return rawQuery.ToSql()
}

func prepareDelete(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Delete(id)
	return rawQuery.ToSql()
}
