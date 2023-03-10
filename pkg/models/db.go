package models

type DB struct { // имитатор базы данных
	Tutors   []Person  // имитатор таблицы преподавателей (обычный слайс)
	Groups   []Group   // имитатор таблицы групп (обычный слайс)
	Students []Student // имитатор таблицы студентов (обычный слайс)
}

func (d *DB) GetTutors() []Person {
	return d.Tutors
}

func (d *DB) GetTutor(id uint) Person {
	for _, tutor := range d.Tutors {
		if tutor.ID == id {
			return tutor
		}
	}
	return Person{}
}

func (d *DB) SearchTutor(query string) Person {
	for _, tutor := range d.Tutors {
		if tutor.Name == query || tutor.Surname == query || tutor.Phone == query {
			return tutor
		}
	}
	return Person{}
}

// group methods

func (d *DB) GetGroups() []Group {
	return d.Groups
}

func (d *DB) GetGroup(id uint) Group {
	for _, group := range d.Groups {
		if group.ID == id {
			return group
		}
	}
	return Group{}
}

func (d *DB) SearchGroup(query string) Group {
	for _, group := range d.Groups {
		if group.Title == query {
			return group
		}
	}
	return Group{}
}

// student methods

func (d *DB) GetStudents() []Student {
	return d.Students
}

func (d *DB) GetStudent(id uint) Student {
	for _, student := range d.Students {
		if student.ID == id {
			return student
		}
	}
	return Student{}
}

func (d *DB) SearchStudent(query string) Student {
	for _, student := range d.Students {
		if student.Name == query || student.Surname == query || student.Phone == query {
			return student
		}
	}
	return Student{}
}
