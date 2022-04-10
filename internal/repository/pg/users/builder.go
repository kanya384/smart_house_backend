package users

import (
	"smart_house_backend/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const TABLE_NAME = "users"

func prepareGet(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Select("id", "name", "surname").From(TABLE_NAME).Where("id = ?", id)
	return rawQuery.ToSql()
}

func prepeareCreate(user domain.User) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Insert(TABLE_NAME).Columns("id", "name", "surname").Values(user.ID, user.Name, user.Surname).Suffix("RETURNING \"id\"")
	return rawQuery.ToSql()
}

func prepareUpdate(user domain.User) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Update(TABLE_NAME).Set("name", user.Name).Set("surname", user.Surname).Where("id = ?", user.ID)
	return rawQuery.ToSql()
}

func prepareDelete(id string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	rawQuery := psqlSq.Delete(id)
	return rawQuery.ToSql()
}
