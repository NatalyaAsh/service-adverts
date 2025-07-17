package pgsql

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"

	"service-advert/internal/config"
	"service-advert/internal/models"
)

var db *sql.DB

func InitUser(cfg *config.Config) error {
	var err error
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		cfg.PGS.User, cfg.PGS.Name, cfg.PGS.Password, cfg.PGS.Host)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		// panic(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		// panic(err)
		return err
	}
	slog.Info("Start db PostgreSQL: Успешное подключение к базе данных!")

	_, err = db.Exec(models.Schema_user)
	if err != nil {
		// panic(err)
		return err
	}
	slog.Info("Start db PostgreSQL: Таблица USERS успешно создана или уже была!")

	return nil
}

func InitAdvert(cfg *config.Config) error {
	var err error
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		cfg.PGS.User, cfg.PGS.Name, cfg.PGS.Password, cfg.PGS.Host)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		// panic(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		// panic(err)
		return err
	}
	slog.Info("Start db PostgreSQL: Успешное подключение к базе данных!")

	// _, err = db.Exec(models.Schema_advert)
	// if err != nil {
	// 	// panic(err)
	// 	return err
	// }
	// slog.Info("Start db PostgreSQL: Таблица ADVERTS успешно создана или уже была!")

	return nil
}

func CloseDB() {
	db.Close()
}

func PostUser(user *models.User) (int64, error) {
	slog.Info("pgsql.PostUser")
	// Стартуем транзакцию
	// tx, err := db.Begin()
	// if err != nil {
	// 	return 0, err
	// }
	// defer tx.Rollback()
	// slog.Info("pg Post Begin transaction")

	// query := `INSERT INTO users (name, email, phon) VALUES ('` + user.Name + `', '` + user.Email + `', '` + user.Phon + `');`
	// slog.Info("pgsql.PostUser", "query", query)
	//	res, err := db.Exec(query)

	var id int64
	query := `INSERT INTO users (name, email, phon) VALUES ($1, $2, $3) RETURNING id;`
	err := db.QueryRow(query, user.Name, user.Email, user.Phon).Scan(&id)
	//res, err := db.Exec(query, user.Name, user.Email, user.Phon)
	slog.Info("pgsql.PostUser Exec: insert")
	if err != nil {
		slog.Error(err.Error())
		//tx.Rollback()
		return 0, err
	}

	// slog.Info("pgsql.PostUser: get id")
	// //id, err := res.LastInsertId()
	// err = db.QueryRow(query, column1, column2).Scan(&id)
	// if err != nil {
	// 	slog.Error(err.Error())
	// 	//tx.Rollback()
	// 	return 0, err
	// }

	slog.Info("pgsql.PostUser", "id", id)
	// //Завершаем транзакцию коммитом
	// if err = tx.Commit(); err != nil {
	// 	return 0, err
	// }
	return id, nil
}

func GetUserById(id int) (models.User, error) {
	slog.Info("PostgreSQL: GetAdvert", "id", id)
	row := db.QueryRow(`SELECT id, name, email, phon FROM users WHERE id=$1`, id)
	if row == nil {
		return models.User{}, fmt.Errorf("user not found")
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phon)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUsers(limit, offset int) (*[]models.User, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset <= 0 {
		offset = 10
	}
	slog.Info("PostgreSQL: GetUsers")
	rows, err := db.Query(`SELECT id, name, email, phon FROM users LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		slog.Error(err.Error())
		return &[]models.User{}, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phon)
		if err != nil {
			slog.Error(err.Error())
			return &[]models.User{}, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		slog.Error(err.Error())
		return &[]models.User{}, err
	}
	return &users, nil
}

// func Post(good *modeldb.Goods) (int64, error) {
// 	// Проверка валидности данных
// 	// Старт транзакции
// 	// Вычисляем приоритет max+1
// 	// Insert
// 	// Commit transaction

// 	// Стартуем транзакцию
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer tx.Rollback()
// 	slog.Info("pg Post Begin transaction")

// 	good.Priority = GetGoodPriority(good.ProjectId) + 1
// 	query := `INSERT INTO goods (project_id, name, priority) VALUES ($1, $2, $3)`
// 	res, err := db.Exec(query, good.ProjectId, good.Name, good.Priority)
// 	slog.Info("pg Post Exec: insert")
// 	if err != nil {
// 		tx.Rollback()
// 		return 0, err
// 	}

// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		tx.Rollback()
// 		return 0, err
// 	}

// 	// Завершаем транзакцию коммитом
// 	if err = tx.Commit(); err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }

// func Patch(good *modeldb.Goods) error {
// 	// Старт транзакции
// 	// Update
// 	// Commit transaction

// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()
// 	slog.Info("pg Patch Begin transaction")

// 	query := `UPDATE goods SET name=$1, description=$2 WHERE id=$3 and project_id=$4`
// 	res, err := db.Exec(query, good.Name, good.Description, good.ID, good.ProjectId)
// 	if err != nil {
// 		return err
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		slog.Error(err.Error())
// 		return err
// 	}
// 	if count == 0 {
// 		return fmt.Errorf(`incorrect id for updating task`)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func Delete(good *modeldb.Goods) error {
// 	// Старт транзакции
// 	// Removed = true
// 	// Commit transaction

// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()
// 	slog.Info("pg Delete Begin transaction")

// 	query := `UPDATE goods SET removed=TRUE WHERE id=$1 and project_id=$2`
// 	res, err := db.Exec(query, good.ID, good.ProjectId)
// 	if err != nil {
// 		return err
// 	}
// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		slog.Error(err.Error())
// 		return err
// 	}
// 	if count == 0 {
// 		return fmt.Errorf(`incorrect id for updating task`)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return err
// 	}
// 	return nil
// }

func GetAdverts(limit, offset int) (*[]models.Advert, error) {
	slog.Info("PostgreSQL: GetAdverts")
	rows, err := db.Query(`SELECT * FROM adverts LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		slog.Error(err.Error())
		return &[]models.Advert{}, err
	}
	defer rows.Close()

	var descRaw sql.NullString // ******
	adverts := []models.Advert{}
	for rows.Next() {
		var advert models.Advert
		err := rows.Scan(&advert.ID, &advert.Name, &advert.Description, &advert.Price, &advert.Author, &advert.Contacts, &advert.Removed, &advert.CreatedAt)
		if err != nil {
			slog.Error(err.Error())
			return &[]models.Advert{}, err
		}
		if descRaw.Valid {
			advert.Description = descRaw.String
		} else {
			advert.Description = ""
		}

		adverts = append(adverts, advert)
	}
	if err = rows.Err(); err != nil {
		slog.Error(err.Error())
		return &[]models.Advert{}, err
	}
	return &adverts, nil
}

// func GetGoodPriority(id int) int {
// 	// Проверяем сначала данные в Redis
// 	// Если нету, то берём из PostgreSQL и инвалидируем в Redis

// 	slog.Info("pgsql GetGoodPriority")
// 	row := db.QueryRow(`SELECT max(priority) FROM goods WHERE project_id=$1`, id)
// 	if row == nil {
// 		return 0
// 	}

// 	var prior sql.NullInt64
// 	err := row.Scan(&prior)
// 	if err != nil {
// 		return 0
// 	}

// 	if prior.Valid {
// 		return int(prior.Int64)
// 	}
// 	return 0
// }

// func GetGoodsCount() (int, error) {
// 	slog.Info("PostgreSQL: GetGoodCount")
// 	row := db.QueryRow(`SELECT count(id) AS cnt FROM goods`)
// 	if row == nil {
// 		return 0, fmt.Errorf("good not found")
// 	}

// 	//var good modeldb.Goods
// 	//var descRaw sql.NullString
// 	var cnt int
// 	err := row.Scan(&cnt)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return cnt, nil
// }

// func GetGoodsCountRemoved() (int, error) {
// 	slog.Info("PostgreSQL: GetGoodCountRemoved")
// 	row := db.QueryRow(`SELECT count(id) AS cnt FROM goods WHERE removed=true`)
// 	if row == nil {
// 		return 0, fmt.Errorf("good not found")
// 	}

// 	//var good modeldb.Goods
// 	//var descRaw sql.NullString
// 	var cnt int
// 	err := row.Scan(&cnt)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return cnt, nil
// }
