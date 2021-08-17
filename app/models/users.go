package models

import (
	"log"
	"time"
)


type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos     []Todo
}
// Userのstrcutを作成。これをベースにuserの作成をする。

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int 
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	// User型のメソッドして作成。errを返り値に。
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`
		// userを作成するコマンドを作成。
	_, err = Db.Exec(cmd,
		// Dbはmodelsパッケージ内に存在するからパッケージを指定しなくても使える。cmdを渡す。
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())
		// ↑上記でユーザー情報を取得。Encryptでパスワードをハッシュにしている。
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	// userを取得する関数。idを引数に。User型とerrorを返り値に。
	user = User{}
	// 構造体のユーザーを渡している。
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	// コマンドを作成。
	err = Db.QueryRow(cmd, id).Scan(
		// QueryRowを使ってcmdとidを渡す。
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
		// ポインタ型で宣言しているから&をつけないといけない。
	)
	return user, err
} 

func (u *User) UpdateUser() (err error) {
	// User型のメソッドして作成。errを返り値に。
	cmd := `update users set name = ?, email = ? where id = ? `
	// cmdを作成。指定されたデータを更新するように実装している。
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	// cmdを実行。
	if err != nil {
		log.Fatalln(err)
	}
	// エラーハンドリング
	return err
}

func (u *User) DeleteUser() (err error) {
	// User型のメソッドして作成。errを返り値に。
	cmd := `delete from users where id = ?`
	// cmdを作成usersのidが一致するものを削除
	_, err = Db.Exec(cmd, u.ID)
	// Db.Execでidが一致するものを削除。
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID, 
		&user.UUID, 
		&user.Name, 
		&user.Email, 
		&user.PassWord, 
		&user.CreatedAt)
	
	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid, 
		email, 
		user_id, 
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at
	from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID, 
		&session.Email, 
		&session.UserID, 
		&session.CreatedAt)

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at
	from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID, 
		&sess.UUID,
		&sess.Email, 
		&sess.UserID, 
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err 
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, created_at FROM users
	where id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID, 
		&user.UUID, 
		&user.Name, 
		&user.Email,
		&user.CreatedAt)

	return user, err
}