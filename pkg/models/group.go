package models

type Group struct {
	ID       uint      // уникальный идентификатор
	Title    string    // название группы
	Tutor    Person    // преподаватель
	Students []Student // студенты группы
}
