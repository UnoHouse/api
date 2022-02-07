package service

import (
	"context"
	"strconv"
	"time"

	"github.com/UnoHouse/api/model"
	mySql "github.com/UnoHouse/api/sql"
	"github.com/UnoHouse/api/utils"
	_ "github.com/go-sql-driver/mysql"
)

type SqlService interface {
	GetAppDescription(id int64) (*model.AppDescription, error)
	CreateAppDescription(*model.AppDescription) (int64, error)
	UpdateAppDescription(*model.AppDescription) (id int64, err error)
	DeleteAppDescription(id int64) (err error)

	GetNotificationsByUserId(userId int64) (notifications *model.Notifications, err error)
	CreateNotification(*model.Notification) (id int64, err error)
}

type SqlServiceImpl struct {
	db mySql.DB
}

func NewMySqlService(db mySql.DB) SqlServiceImpl {
	return SqlServiceImpl{db}
}

func (s *SqlServiceImpl) GetNotificationsByUserId(userId int64) (notifications *model.Notifications, err error) {
	notifications = &model.Notifications{}

	rows, err := s.db.QueryContext(getContext(), "SELECT `id`,`user_id`,`title`,`body`,`click_action`,"+
		"`channel_id`,`data`,`device_id`,`created_at`,`deleted_at` FROM `notification` WHERE `user_id` = ? AND `deleted_at` IS NULL ORDER BY `created_at` DESC", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var notification model.Notification
		err := rows.Scan(
			&notification.Id,
			&notification.UserId,
			&notification.Title,
			&notification.Body,
			&notification.ClickAction,
			&notification.ChannelId,
			&notification.Data,
			&notification.DeviceId,
			&notification.CreatedAt,
			&notification.DeletedAt,
		)

		if err != nil {
			return nil, err
		}

		notifications.Items = append(notifications.Items, &notification)
	}

	return notifications, nil
}

func (s *SqlServiceImpl) CreateNotification(notification *model.Notification) (id int64, err error) {
	sqlStatement := "INSERT INTO `notification`	(`user_id`,`title`,`body`,`click_action`," +
		"`channel_id`,`data`,`sent`,`created_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?,?);"
	rows, err := s.db.ExecContext(getContext(), sqlStatement,
		notification.UserId,
		notification.Title,
		notification.Body,
		notification.ClickAction,
		notification.ChannelId,
		notification.Data,
		notification.Sent,
		notification.CreatedAt,
		notification.DeletedAt,
	)

	if err != nil {
		return 0, err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func getContext() context.Context {
	ctx := context.Background()
	timeout, err := strconv.Atoi(utils.Env("SQL_REQUEST_TIMEOUT_SECONDS", "60"))
	if err == nil {
		ctx, _ = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	}
	return ctx
}
