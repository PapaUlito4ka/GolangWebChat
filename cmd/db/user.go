package db

import (
	"database/sql"
	"errors"

	"github.com/papaulito4ka/golangwebchat/cmd/dao"
	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	Db *sql.DB
}

func (userDb *UserDB) FindAll() ([]dao.UserDao, error) {
	rows, err := userDb.Db.Query("select user_id, username, password from \"user\"")
	if err != nil {
		return []dao.UserDao{}, err
	}

	users := []dao.UserDao{}
	for rows.Next() {
		user := dao.UserDao{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			return []dao.UserDao{}, err
		}
		user.Friends, err = userDb.Friends(int(user.Id))
		if err != nil {
			return []dao.UserDao{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userDb *UserDB) FindById(id int) (dao.UserDao, error) {
	query := `select * from "user" where user_id = $1`
	stmt, err := userDb.Db.Prepare(query)
	if err != nil {
		return dao.UserDao{}, nil
	}

	row := stmt.QueryRow(id)

	user := dao.UserDao{}
	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return dao.UserDao{}, nil
	}

	user.Friends, err = userDb.Friends(id)
	if err != nil {
		return dao.UserDao{}, nil
	}

	return user, nil
}

func (userDb *UserDB) Friends(id int) ([]dao.UserDao, error) {
	query := `	select
					case when chat.user_id = $1 then u2.user_id else u1.user_id end as id,
					case when chat.user_id = $1 then u2.username else u1.username end as username,
					case when chat.user_id = $1 then u2.password else u1.password end as password
				from chat
				join "user" u1 on u1.user_id = chat.user_id
				join "user" u2 on u2.user_id = chat.user_id_
				where chat.user_id = $1 or chat.user_id_ = $1`

	stmt, err := userDb.Db.Prepare(query)
	if err != nil {
		return []dao.UserDao{}, nil
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return []dao.UserDao{}, nil
	}

	friends := []dao.UserDao{}
	for rows.Next() {
		friend := dao.UserDao{}

		if err := rows.Scan(&friend.Id, &friend.Username, &friend.Password); err != nil {
			return []dao.UserDao{}, nil
		}

		friends = append(friends, friend)
	}

	return friends, nil
}

func (userDb *UserDB) Create(username, password string) (int64, error) {
	query := `INSERT INTO public."user" (user_id, username, password)
				VALUES (DEFAULT, $1, $2)`
	stmt, err := userDb.Db.Prepare(query)
	if err != nil {
		return 0, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(username, hashedPassword)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 1 {
		id, _ := res.LastInsertId()
		return id, nil
	}
	return 0, errors.New("unknown error during sql execution")
}

func (userDb *UserDB) Find(username, password string) (dao.UserDao, error) {
	query := `select username, password from "users" where username=$1 and password=$2`
	stmt, err := userDb.Db.Prepare(query)
	if err != nil {
		return dao.UserDao{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return dao.UserDao{}, err
	}

	row := stmt.QueryRow(username, hashedPassword)

	user := dao.UserDao{}
	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return dao.UserDao{}, err
	}

	return user, nil
}
