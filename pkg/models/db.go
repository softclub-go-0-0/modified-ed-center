package models

import (
	"fmt"
	"os"
)

type DB struct { // имитатор базы данных
	Tutors   []Person  // имитатор таблицы преподавателей (обычный слайс)
	Groups   []Group   // имитатор таблицы групп (обычный слайс)
	Students []Student // имитатор таблицы студентов (обычный слайс)
}

func (d DB) GetTutors() []Person {
	return d.Tutors
}

func (d DB) GetTutor(id uint) Person {
	for _, tutor := range d.Tutors {
		if tutor.ID == id {
			return tutor
		}
	}
	return Person{}
}

func (d DB) SearchTutor(query string) Person {
	for _, tutor := range d.Tutors {
		if tutor.Name == query || tutor.Surname == query || tutor.Phone == query {
			return tutor
		}
	}
	return Person{}
}

// group methods

func (d DB) GetGroups() []Group {
	return d.Groups
}

func (d DB) GetGroup(id uint) Group {
	for _, group := range d.Groups {
		if group.ID == id {
			return group
		}
	}
	return Group{}
}

func (d DB) SearchGroup(query string) Group {
	for _, group := range d.Groups {
		if group.Title == query {
			return group
		}
	}
	return Group{}
}

// student methods

func (d DB) GetStudents() []Student {
	return d.Students
}

func (d DB) GetStudent(id uint) Student {
	for _, student := range d.Students {
		if student.ID == id {
			return student
		}
	}
	return Student{}
}

func (d DB) SearchStudent(query string) Student {
	for _, student := range d.Students {
		if student.Name == query || student.Surname == query || student.Phone == query {
			return student
		}
	}
	return Student{}
}

func (d DB) ExportTutorsToFile(filename string) (totalNumber int, err error) {
	file, err := os.Create(filename)
	if err != nil {
		return totalNumber, err
	}
	defer file.Close()

	for _, tutor := range d.Tutors {
		tutorData := fmt.Sprintf("%d,%s,%s,%s\n", tutor.ID, tutor.Name, tutor.Surname, tutor.Phone)
		_, err = file.WriteString(tutorData)
		if err != nil {
			return
		}
	}

	return len(d.Tutors), nil
}

func (d DB) ExportGroupsToFile(filename string) (totalNumber int, err error) {
	file, err := os.Create(filename)
	if err != nil {
		return totalNumber, err
	}
	defer file.Close()

	for _, group := range d.Groups {
		groupData := fmt.Sprintf("%d,%s,%s\n", group.ID, group.Title, group.Tutor.Name+" "+group.Tutor.Surname)
		_, err = file.WriteString(groupData)
		if err != nil {
			return
		}
	}

	return len(d.Groups), nil
}
