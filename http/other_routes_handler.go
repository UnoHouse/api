package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/UnoHouse/api/mapper"
	"github.com/UnoHouse/api/model"
	"github.com/UnoHouse/api/utils/logger"
	"github.com/gorilla/mux"
)

func (h handler) postNotification() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}
		userId, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		notification, err := mapper.MapNotificationFormToNotification(r, model.Notification{})
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		notification.CreatedAt = time.Now()
		userDevice, err := h.userSqlService.GetLatestDevice(userId)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		notification.DeviceId = userDevice.Id
		id, err := h.sqlService.CreateNotification(notification)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		notification.Id = id

		userNotificationResponse(notification, w, r)
	}
}

func (h handler) getAppDescription() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		description, err := h.sqlService.GetAppDescription(id)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		appDescriptionResponse(description, w, r)
	}

}

func (h handler) postAppDescription() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		description, err := mapper.MapDescriptionFormToDescription(r, model.AppDescription{})
		id, err := h.sqlService.CreateAppDescription(description)
		appDescription, err := h.sqlService.GetAppDescription(id)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		appDescriptionResponse(appDescription, w, r)
	}

}
func (h handler) putAppDescription() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondBadRequest(w, "")
			return
		}

		appDescription, err := h.sqlService.GetAppDescription(id)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}
		appDescription, err = mapper.MapDescriptionFormToDescription(r, *appDescription)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}
		id, err = h.sqlService.UpdateAppDescription(appDescription)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		appDescription, err = h.sqlService.GetAppDescription(id)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}
		appDescriptionResponse(appDescription, w, r)
	}

}

func (h handler) deleteAppDescription() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		err = h.sqlService.DeleteAppDescription(id)
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
			return
		}

		successResponse(w, r)
	}

}
