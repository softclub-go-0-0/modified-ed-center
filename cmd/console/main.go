package main

import (
	"fmt"
	"log"
	"modified-ed-center/internal/database"
)

func main() {
	var command int
	for {
		fmt.Println("Введите команду 0-9")
		fmt.Printf("\t0 - выход\n")
		fmt.Printf("\t1 - список всех преподавателей\n\t2 - поиск преподавателя по ID\n\t3 - поиск преподавателя по другим параметрам\n")
		fmt.Printf("\t4 - список всех групп\n\t5 - поиск группы по ID\n\t6 - поиск группы по названию\n")
		fmt.Printf("\t7 - список всех студентов\n\t8 - поиск студента по ID\n\t9 - поиск студента по другим параметрам\n")

		fmt.Printf("\t10 - экспортировать список преподователей в csv файл\n")
		fmt.Printf("\t11 - экспортировать список групп в csv файл\n")
		fmt.Printf("\t12 - экспортировать список студентов в csv файл\n")

		fmt.Scan(&command)

		switch command {
		case 1:
			fmt.Println("Список всех преподавателей центра:")
			for _, tutor := range database.DB.GetTutors() {
				fmt.Println(tutor)
			}
		case 2:
			var id uint
			fmt.Println("Введите ID преподавателя:")
			fmt.Scan(&id)
			fmt.Println("Результат поиска:")
			fmt.Println(database.DB.GetTutor(id))
		case 3:
			var query string
			fmt.Println("Введите имя, фамилию или номера телефона преподавателя:")
			fmt.Scan(&query)
			fmt.Println("Результат поиска:")
			fmt.Println(database.DB.SearchTutor(query))
		case 4:
			fmt.Println("Список всех групп центра:")
			for _, group := range database.DB.GetGroups() {
				fmt.Println(group.ID, group.Title)
			}
		case 5:
			var id uint
			fmt.Println("Введите ID группы:")
			fmt.Scan(&id)
			fmt.Println("Результат поиска:")
			group := database.DB.GetGroup(id)
			fmt.Println(group.ID, group.Title)
			fmt.Println("Преподаватель:", group.Tutor)
			fmt.Println("Студенты:")
			for _, student := range group.Students {
				fmt.Println(student)
			}
		case 6:
			var query string
			fmt.Println("Введите название группы:")
			fmt.Scan(&query)
			fmt.Println("Результат поиска:")
			group := database.DB.SearchGroup(query)
			fmt.Println(group.ID, group.Title)
			fmt.Println("Преподаватель:", group.Tutor)
			fmt.Println("Студенты:")
			for _, student := range group.Students {
				fmt.Println(student)
			}
		case 7:
			fmt.Println("Список всех студентов центра:")
			for _, student := range database.DB.GetStudents() {
				fmt.Println(student)
			}
		case 8:
			var id uint
			fmt.Println("Введите ID студента:")
			fmt.Scan(&id)
			fmt.Println("Результат поиска:")
			fmt.Println(database.DB.GetStudent(id))
		case 9:
			var query string
			fmt.Println("Введите имя, фамилию или номера телефона студента:")
			fmt.Scan(&query)
			fmt.Println("Результат поиска:")
			fmt.Println(database.DB.SearchStudent(query))
		case 10:
			var filename string
			fmt.Print("Введите имя файла: ")
			fmt.Scan(&filename)
			numberOfTutors, err := database.DB.ExportTutorsToFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Успешно экспортировано", numberOfTutors, "преподавателей в файл", filename)
		case 11:
			var filename string
			fmt.Print("Введите имя файла: ")
			fmt.Scan(&filename)
			numberOfGroups, err := database.DB.ExportGroupsToFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Успешно экспортировано", numberOfGroups, "групп в файл", filename)
		case 12:
			var filename string
			fmt.Print("Введите имя файла: ")
			fmt.Scan(&filename)
			numberOfStudents, err := database.DB.ExportStudentsToFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Успешно экспортировано", numberOfStudents, "студентов в файл", filename)

		case 0:
			fmt.Println("Bye!")
			return
		}
	}
}
