package house_part

import (
	"smart_house_backend/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const TABLE_NAME = "house_parts"

func prepareGet(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Select("id", "name", "photo", "digital_pin_cnt", "analog_pin_cnt").From(TABLE_NAME).Where("id = ?", id)
	return rawQuery.ToSql()
}

func prepeareCreate(housePart domain.HousePart) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Insert(TABLE_NAME).Columns("id", "name", "house_id").Values(housePart.ID, housePart.Name, housePart.HouseID).Suffix("RETURNING \"id\"")
	return rawQuery.ToSql()
}

func prepareUpdate(housePart domain.HousePart) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Update(TABLE_NAME).Set("name", housePart.Name).Set("house_id", housePart.HouseID).Where("id = ?", housePart.ID)
	return rawQuery.ToSql()
}

func prepareDelete(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Delete(id)
	return rawQuery.ToSql()
}
