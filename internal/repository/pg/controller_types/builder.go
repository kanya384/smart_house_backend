package controller_types

import (
	"fmt"
	"smart_house_backend/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const TABLE_NAME = "controller_types"

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
	rawQuery := psqlSq.Select("id", "name", "photo", "digital_pin_cnt", "analog_pin_cnt").From(qb.getTableName()).Where("id = ?", id)
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepeareCreate(controllerType domain.ControllerType) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Insert(qb.getTableName()).Columns("id", "name", "photo", "digital_pin_cnt", "analog_pin_cnt").Values(controllerType.ID, controllerType.Name, controllerType.Photo, controllerType.DigitalPinCnt, controllerType.AnalogPinCnt).Suffix("RETURNING \"id\"")
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepareUpdate(controllerType domain.ControllerType) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Update(qb.getTableName()).Set("name", controllerType.Name).Set("photo", controllerType.Photo).Set("digital_pin_cnt", controllerType.DigitalPinCnt).Set("digital_pin_cnt", controllerType.AnalogPinCnt).Where("id = ?", controllerType.ID)
	return rawQuery.ToSql()
}

func (qb *queryBuilder) prepareDelete(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Delete(qb.getTableName()).Where("id = ?", id)
	return rawQuery.ToSql()
}
