package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./chat_history.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Tạo bảng nếu chưa tồn tại
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user TEXT NOT NULL,
			content TEXT NOT NULL,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

// Lưu tin nhắn vào cơ sở dữ liệu
func SaveMessage(db *sql.DB, user, content string) error {
	_, err := db.Exec("INSERT INTO messages (user, content) VALUES (?, ?)", user, content)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// Lấy tất cả tin nhắn từ cơ sở dữ liệu
func GetMessages(db *sql.DB) ([]Message, error) {
	rows, err := db.Query("SELECT user, content, timestamp FROM messages ORDER BY timestamp ASC")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.User, &msg.Content, &msg.Timestamp)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
