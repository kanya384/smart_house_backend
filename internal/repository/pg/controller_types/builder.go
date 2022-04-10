package controller_types

import (
	"smart_house_backend/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const TABLE_NAME = "controller_types"

func prepareGet(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Select("id", "name", "photo", "digital_pin_cnt", "analog_pin_cnt").From(TABLE_NAME).Where("id = ?", id)
	return rawQuery.ToSql()
}

func prepeareCreate(controllerType domain.ControllerType) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Insert(TABLE_NAME).Columns("id", "name", "photo", "digital_pin_cnt", "analog_pin_cnt").Values(controllerType.ID, controllerType.Name, controllerType.Photo, controllerType.DigitalPinCnt, controllerType.AnalogPinCnt).Suffix("RETURNING \"id\"")
	return rawQuery.ToSql()
}

func prepareUpdate(controllerType domain.ControllerType) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Update(TABLE_NAME).Set("name", controllerType.Name).Set("photo", controllerType.Photo).Set("digital_pin_cnt", controllerType.DigitalPinCnt).Set("digital_pin_cnt", controllerType.AnalogPinCnt).Where("id = ?", controllerType.ID)
	return rawQuery.ToSql()
}

func prepareDelete(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Delete(id)
	return rawQuery.ToSql()
}
