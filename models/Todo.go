package models


type Todo struct{
	IdTodo int64 `xorm:"'id' pk autoincr" json:"id_todo" `
	Name string `xorm:"varchar(25) not null 'todo_name'" json:"todo_name" binding:"required"`
}

