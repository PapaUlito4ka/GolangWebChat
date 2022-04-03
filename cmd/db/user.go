package db

import (
	"database/sql"

	"github.com/papaulito4ka/golangwebchat/cmd/dao"
	"golang.org/x/crypto/bcrypt"
)

type UserDB struct {
	Db *sql.DB
}

func NewUserDB(Db *sql.DB) UserDB {
	return UserDB{Db: Db}
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
		user.Friends, err = userDb.Friends(user.Id)
		if err != nil {
			return []dao.UserDao{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userDb *UserDB) FindById(id int64) (dao.UserDao, error) {
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

func (userDb *UserDB) Friends(id int64) ([]dao.UserDao, error) {
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

func (userDb *UserDB) Chats(id int64) ([]dao.ChatDao, error) {
	query := `	select
					chat.chat_id,
					case when chat.user_id = $1 then u2.user_id else u1.user_id end as id,
					case when chat.user_id = $1 then u2.username else u1.username end as username,
					case when chat.user_id = $1 then u2.password else u1.password end as password
				from chat
				join "user" u1 on u1.user_id = chat.user_id
				join "user" u2 on u2.user_id = chat.user_id_
				where chat.user_id = $1 or chat.user_id_ = $1`

	stmt, err := userDb.Db.Prepare(query)
	if err != nil {
		return []dao.ChatDao{}, nil
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return []dao.ChatDao{}, nil
	}

	chats := []dao.ChatDao{}
	for rows.Next() {
		chat := dao.ChatDao{}

		if err := rows.Scan(&chat.Id, &chat.Friend.Id, &chat.Friend.Username, &chat.Friend.Password); err != nil {
			return []dao.ChatDao{}, nil
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

func (userDb *UserDB) Create(username, password string) (int64, error) {
	query := `INSERT INTO public."user" (user_id, username, password)
				VALUES (DEFAULT, $1, $2) returning user_id`
	stmt, err := userDb.Db.Prepare(query)
	if err != nil {
		return 0, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return 0, err
	}

	row := stmt.QueryRow(username, hashedPassword)
	if err != nil {
		return 0, err
	}

	userId := int64(0)
	err = row.Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (userDb *UserDB) Find(username, password string) (dao.UserDao, error) {
	query := `select user_id, username, password from "user" where username=$1`
	stmt, err := userDb.Db.Prepare(query)
	if err != nil {
		return dao.UserDao{}, err
	}

	if err != nil {
		return dao.UserDao{}, err
	}

	row := stmt.QueryRow(username)

	user := dao.UserDao{}
	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return dao.UserDao{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return dao.UserDao{}, err
	}

	user.Password = password

	return user, nil
}

func (userDb *UserDB) AddFriend(userId, friendId int64) (int64, error) {
	query := `	INSERT INTO public.chat (chat_id, user_id, user_id_)
				VALUES (DEFAULT, $1, $2) returning chat_id`
	stmt, err := userDb.Db.Prepare(query)
	if err != nil {
		return 0, err
	}

	row := stmt.QueryRow(userId, friendId)

	chatId := int64(0)
	if err := row.Scan(&chatId); err != nil {
		return 0, err
	}

	return chatId, nil
}
