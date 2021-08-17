package models

import (
	"log"
	"time"
)

type Todo struct {
	ID int 
	Content string
	UserID int
	CreatedAt time.Time
}
// Todoのstructを作成

func (u *User) CreateTodo(content string) (err error) {
	// User型のメソッドして作成。引数をcontentに。errを返り値に。
	cmd := `insert into todos (
		content,
	    user_id, 
		created_at) values (?, ?, ?)`
	// cmdを作成。content, user_id, created_atのデータを渡す。
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	// Todoを取得する関数を作成。引数はid。返り値はTodoとerror
	cmd := `select id, content, user_id, created_at from todos
	where id = ?`
	// cmdを作成。id, content, user_id, created_atを渡す。
	todo = Todo{}
	// Todo{}型の宣言をする。

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID, 
		&todo.Content,
		&todo.UserID, 
		&todo.CreatedAt)
	// Todo情報を取得。 
	// cmdを実行する。
    return todo, err
}

func GetTodos() (todos []Todo, err error) {
	// 複数のTodoを作成する関数。返り値はスライスのTodoとerror
	cmd := `select id, content, user_id, created_at from todos`
	// idとcontentとuser_id情報を取得
	rows, err := Db.Query(cmd)
	// cmdをQueryに渡す。
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		// 変数Todoを宣言。
		err = rows.Scan(
			&todo.ID, 
			&todo.Content,
			&todo.UserID, 
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func(u *User) GetTodosByUser() (todos []Todo, err error) {
	// 特定のユーザーを取得する関数。userのメソッドとして作成。
	cmd := `select id, content, user_id, created_at from todos
	where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		// Todoを宣言。
		err = rows.Scan(
			&todo.ID,
		    &todo.Content,
			&todo.UserID, 
			&todo.CreatedAt)
		// 各データ情報ををスキャンする。
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
		// スキャンされたtodoをtodosにappendする。
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() error {
	// Todoのメソッドとして作成
	cmd := `update todos set content = ?, user_id = ? 
	where id = ?`
	// cmdでアップデートする情報を取得する。
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error{
	// Todoのメソッドとして作成
	cmd := `delete from todos where id = ?`
	// cmdでdeleteする情報を取得する。(idが一致するものを削除するようにする)
	_, err = Db.Exec(cmd, t.ID)
	// cmdを実行
	if err != nil {
		log.Fatalln(err)
	}
	return err
}