package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (d DB) ExportStudentsToFile(filename string) (totalNumber int, err error) {
	file, err := os.Create(filename)
	if err != nil {
		return totalNumber, err
	}
	defer file.Close()

	for _, student := range d.Students {
		studentData := fmt.Sprintf("%d,%s,%s,%s,%d\n", student.ID, student.Name, student.Surname, student.Phone, student.GroupID)
		_, err = file.WriteString(studentData)
		if err != nil {
			return
		}
	}

	return len(d.Students), nil
}

func (d DB) ImportTutorsFromFile(filename string) (tutors []Person, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tutorDataString := scanner.Text()
		tutorData := strings.Split(tutorDataString, ",")
		id, err := strconv.Atoi(tutorData[0])
		if err != nil {
			return nil, err
		}
		tutor := Person{
			ID:      uint(id),
			Name:    tutorData[1],
			Surname: tutorData[2],
			Phone:   tutorData[3],
		}
		tutors = append(tutors, tutor)
	}

	return
}
