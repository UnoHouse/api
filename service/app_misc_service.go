package service

import (
	"database/sql"

	"github.com/UnoHouse/api/model"
)

func (s *SqlServiceImpl) GetAppDescription(id int64) (*model.AppDescription, error) {
	description := model.AppDescription{}

	row := s.db.QueryRow("SELECT `id`,`title`,`body`, `updated_at` " +
		"FROM `app_description` LIMIT 1")
	err := row.Scan(&description.Id,
		&description.Title,
		&description.Body)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &description, nil
}

func (s *SqlServiceImpl) CreateAppDescription(description *model.AppDescription) (int64, error) {
	sqlStatement :=
		"INSERT INTO `app_description` (`title`,`body`, `updated_at` VALUES (?,?,?)"
	rows, err := s.db.ExecContext(
		getContext(),
		sqlStatement,
		&description.Title,
		&description.Body,
		&description.UpdatedAt)

	if err != nil {
		return 0, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SqlServiceImpl) UpdateAppDescription(description *model.AppDescription) (id int64, err error) {

	return 0, nil
}

func (s *SqlServiceImpl) DeleteAppDescription(id int64) (err error) {

	sqlStatement := "DELETE FROM app_description WHERE `id` = ?"
	_, err = s.db.ExecContext(getContext(), sqlStatement, id)

	if err != nil {
		return err
	}
	return nil
}
