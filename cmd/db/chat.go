package db

import (
	"database/sql"

	"github.com/papaulito4ka/golangwebchat/cmd/dao"
)

type ChatDB struct {
	Db *sql.DB
}

func NewChatDb(Db *sql.DB) ChatDB {
	return ChatDB{
		Db: Db,
	}
}

func (chatDb *ChatDB) Find(userId, friendId int64) (dao.ChatDao, error) {
	query := `	select
					chat.chat_id,
					u1.user_id,
					u1.username,
					u1.password,
					u2.user_id,
					u2.username,
					u2.password
				from chat
				join "user" u1 on chat.user_id = u1.user_id
				join "user" u2 on chat.user_id_ = u2.user_id
				where 	(chat.user_id = $1 or chat.user_id = $2) and
						(chat.user_id_ = $1 or chat.user_id_ = $2);`
	stmt, err := chatDb.Db.Prepare(query)
	if err != nil {
		return dao.ChatDao{}, err
	}
	row := stmt.QueryRow()

	chat := dao.ChatDao{}

	if err := row.Scan(&chat.Id, &chat.User, &chat.Friend); err != nil {
		return dao.ChatDao{}, err
	}

	return chat, nil
}

func (chatDb *ChatDB) SendMessage(chatId, userId int64, message string) (int64, error) {
	query := `	INSERT INTO public.message (message_id, user_id, chat_id, message, time)
				VALUES (DEFAULT, $1, $2, $3, DEFAULT) returning message_id`
	stmt, err := chatDb.Db.Prepare(query)
	if err != nil {
		return 0, err
	}

	row := stmt.QueryRow(userId, chatId, message)

	messageId := int64(0)
	if err := row.Scan(&messageId); err != nil {
		return 0, err
	}

	return messageId, nil
}

func (chatDb *ChatDB) FindChatMessages(chatId int64) ([]dao.MessageDao, error) {
	query := `	select
					u.user_id,
					u.username,
					u.password,
					message.message,
					message.time
				from message
				join "user" u on u.user_id = message.user_id
				where chat_id = $1
				order by time;`
	stmt, err := chatDb.Db.Prepare(query)
	if err != nil {
		return []dao.MessageDao{}, err
	}

	rows, err := stmt.Query(chatId)
	if err != nil {
		return []dao.MessageDao{}, err
	}

	messages := []dao.MessageDao{}

	for rows.Next() {
		message := dao.MessageDao{}

		if err := rows.Scan(
			&message.User.Id,
			&message.User.Username,
			&message.User.Password,
			&message.Message,
			&message.Time); err != nil {
			return []dao.MessageDao{}, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}
