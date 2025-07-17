package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"service-advert/internal/config"
	"service-advert/internal/database/pgsql"
	"service-advert/internal/models"
	//"service-advert/internal/tasks"
	// "service-demo/internal/config"
	// "service-demo/internal/database/pgsql"
	// "service-demo/internal/database/redis"
	// modeldb "service-demo/internal/models"
)

// type Meta struct {
// 	Total   int `json:"total"`
// 	Removed int `json:"removed"`
// 	Limit   int `json:"limit"`
// 	Offset  int `json:"offset"`
// }

// type StructGetGoods struct {
// 	Meta  Meta             `json:"meta"`
// 	Goods *[]modeldb.Goods `json:"goods"`
// }

func Init(mux *http.ServeMux, cfg *config.Config) {
	mux.HandleFunc("POST /advert", PostAdvert)
	mux.HandleFunc("GET /advert", GetAdvert)
	mux.HandleFunc("DELETE /advert", DeleteAdvert)
	mux.HandleFunc("POST /user", PostUser)
}

func writeJson(w http.ResponseWriter, data any, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	msg, _ := json.Marshal(data)
	io.Writer.Write(w, msg)
}

func PostAdvert(w http.ResponseWriter, r *http.Request) {
	slog.Info("Api.PostAdvert")
	var buf bytes.Buffer
	// var task models.Task

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeJson(w, models.ResponseErr{Error: "ошибка передачи данных"}, http.StatusBadRequest)
		return
	}
	slog.Info("Api.PostAdvert", "body", buf.Bytes())
	// if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
	// 	writeJson(w, models.ResponseErr{Error: "ошибка конвертации данных"}, http.StatusBadRequest)
	// 	return
	// }

	// fmt.Printf("task: %v\n", task)
	// id := tasks.PostTask(task)
	// // Запускаем горутину на время выполнения задачи
	// go WorkTask(w, id)

	writeJson(w, models.ResponseId{ID: 111}, http.StatusOK)

}

func GetAdvert(w http.ResponseWriter, r *http.Request) {
	slog.Info("Api.GetTask")

	IdRaw := r.URL.Query().Get("id")
	// if IdRaw == "" {
	// 	writeJson(w, models.ResponseErr{Error: "не указан идентификатор задачи"}, http.StatusBadRequest)
	// 	return
	// }
	// id, err := strconv.Atoi(IdRaw)
	// if err != nil {
	// 	writeJson(w, models.ResponseErr{Error: "не корректный идентификатор задачи"}, http.StatusBadRequest)
	// 	return
	// }

	slog.Info("Api.GetTask", "Id", IdRaw)
	// task, err := tasks.GetTask(id)
	// if err != nil {
	// 	writeJson(w, models.ResponseErr{Error: err.Error()}, http.StatusBadRequest)
	// 	return
	// }
	// writeJson(w, task, http.StatusOK)
}

func DeleteAdvert(w http.ResponseWriter, r *http.Request) {
	//slog.Info("Api.DeleteTask")

	// // IdRaw := r.URL.Query().Get("id")
	// // if IdRaw == "" {
	// // 	writeJson(w, models.ResponseErr{Error: "не указан идентификатор задачи"}, http.StatusBadRequest)
	// // 	return
	// // }
	// // id, err := strconv.Atoi(IdRaw)
	// // if err != nil {
	// // 	writeJson(w, models.ResponseErr{Error: "не корректный идентификатор задачи"}, http.StatusBadRequest)
	// // 	return
	// // }

	// // //slog.Info("Api.DeleteTask", "Id", id)
	// // task, err := tasks.GetTask(id)
	// // if err != nil {
	// // 	writeJson(w, models.ResponseErr{Error: err.Error()}, http.StatusBadRequest)
	// // 	return
	// // }
	// // if task.Status != models.StatusFinished {
	// // 	writeJson(w, models.ResponseErr{Error: "задача ещё выполняется, её нельзя удалить"}, http.StatusOK)
	// // 	return
	// }

	// err = tasks.DeleteTask(id)
	// if err != nil {
	// 	writeJson(w, models.ResponseErr{Error: err.Error()}, http.StatusBadRequest)
	// 	return
	// }
	// writeJson(w, "", http.StatusOK)
}

func WorkTask(w http.ResponseWriter, id int) {
	// task, err := tasks.GetTask(id)
	// if err != nil {
	// 	writeJson(w, models.ResponseErr{Error: err.Error()}, http.StatusBadRequest)
	// 	return
	// }
	// time.Sleep(time.Duration(task.Lasting) * time.Second)
	// tasks.TaskFinished(task.ID)
	// slog.Info("Api.WorkTask Finished", "id", task.ID)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("Api.PostUser")
	var buf bytes.Buffer
	var user models.User

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeJson(w, models.ResponseErr{Error: "ошибка передачи данных"}, http.StatusBadRequest)
		return
	}
	slog.Info("Api.PostUser", "body", buf.Bytes())
	if err = json.Unmarshal(buf.Bytes(), &user); err != nil {
		writeJson(w, models.ResponseErr{Error: "ошибка конвертации данных"}, http.StatusBadRequest)
		slog.Error("ошибка конвертации данных")
		return
	}

	slog.Info("Api.PostUser", "user", user)
	id, err := pgsql.PostUser(&user)
	if err != nil {
		return
	}
	// // Запускаем горутину на время выполнения задачи
	// go WorkTask(w, id)

	writeJson(w, models.ResponseId{ID: id}, http.StatusOK)

}
