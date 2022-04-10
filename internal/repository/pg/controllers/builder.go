package controllers

import (
	"smart_house_backend/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const TABLE_NAME = "controllers"

func prepareGet(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Select("id", "controller_type_id", "ip").From(TABLE_NAME).Where("id = ?", id)
	return rawQuery.ToSql()
}

func prepeareCreate(controller domain.Controller) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Insert(TABLE_NAME).Columns("id", "controller_type_id", "ip").Values(controller.ID, controller.ControllerTypeId, controller.Ip).Suffix("RETURNING \"id\"")
	return rawQuery.ToSql()
}

func prepareUpdate(controller domain.Controller) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Update(TABLE_NAME).Set("controller_type_id", controller.ControllerTypeId).Set("ip", controller.Ip).Where("id = ?", controller.ID)
	return rawQuery.ToSql()
}

func prepareDelete(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Delete(id)
	return rawQuery.ToSql()
}
