package database

import "modified-ed-center/pkg/models"

var DB = models.DB{
	Tutors:   tutors,
	Groups:   groups,
	Students: append(students, groups[2].Students...),
}

var tutors = []models.Person{
	{
		ID:      1,
		Name:    "Nurullah",
		Surname: "Sulaymonov",
		Phone:   "987654300",
	},
	{
		ID:      2,
		Name:    "Muhammadjon",
		Surname: "Mirzoev",
		Phone:   "987654301",
	},
	{
		ID:      3,
		Name:    "Siyovush",
		Surname: "Shorakhimzoda",
		Phone:   "987654302",
	},
}

var groups = []models.Group{
	{
		ID:    1,
		Title: "C#",
		Tutor: tutors[0],
		Students: []models.Student{
			{
				Person: models.Person{
					ID:      4,
					Name:    "John",
					Surname: "Doe",
					Phone:   "987654303",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 78,
					"week2": 85,
					"week3": 82,
					"week4": 56,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      5,
					Name:    "Oliver",
					Surname: "William",
					Phone:   "987654304",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 43,
					"week2": 56,
					"week3": 68,
					"week4": 77,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      6,
					Name:    "Jack",
					Surname: "Henry",
					Phone:   "987654305",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 98,
					"week2": 89,
					"week3": 90,
					"week4": 77,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      7,
					Name:    "Jackson",
					Surname: "Mateo",
					Phone:   "987654306",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 35,
					"week2": 62,
					"week3": 48,
					"week4": 80,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      8,
					Name:    "Daniel",
					Surname: "Logan",
					Phone:   "987654307",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 96,
					"week2": 42,
					"week3": 80,
					"week4": 84,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      9,
					Name:    "Samuel",
					Surname: "Jacob",
					Phone:   "987654308",
				},
				GroupID: 1,
				Grades: map[string]int{
					"week1": 43,
					"week2": 85,
					"week3": 29,
					"week4": 88,
				},
				TopicsToWorkOn: nil,
			},
		},
	},
	{
		ID:    2,
		Title: "C++",
		Tutor: tutors[1],
		Students: []models.Student{
			{
				Person: models.Person{
					ID:      10,
					Name:    "John",
					Surname: "Joseph",
					Phone:   "987654309",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 54,
					"week2": 85,
					"week3": 33,
					"week4": 56,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      11,
					Name:    "David",
					Surname: "Hudson",
					Phone:   "987654310",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 39,
					"week2": 65,
					"week3": 82,
					"week4": 78,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      12,
					Name:    "Jack",
					Surname: "Henry",
					Phone:   "987654311",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 44,
					"week2": 63,
					"week3": 82,
					"week4": 79,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      13,
					Name:    "Leo",
					Surname: "Matthew",
					Phone:   "987654312",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 56,
					"week2": 85,
					"week3": 90,
					"week4": 67,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      14,
					Name:    "Daniel",
					Surname: "Luke",
					Phone:   "987654313",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 68,
					"week2": 85,
					"week3": 82,
					"week4": 78,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      15,
					Name:    "Carter",
					Surname: "Jacob",
					Phone:   "987654314",
				},
				GroupID: 2,
				Grades: map[string]int{
					"week1": 90,
					"week2": 65,
					"week3": 82,
					"week4": 67,
				},
				TopicsToWorkOn: nil,
			},
		},
	},
	{
		ID:    3,
		Title: "Go",
		Tutor: tutors[2],
		Students: []models.Student{
			{
				Person: models.Person{
					ID:      16,
					Name:    "Alisher",
					Surname: "Yoqubov",
					Phone:   "987654315",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 86,
					"week2": 80,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      17,
					Name:    "Amir",
					Surname: "Arifjonov",
					Phone:   "987654316",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 76,
					"week2": 68,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      18,
					Name:    "Behruz",
					Surname: "Shodiev",
					Phone:   "987654317",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 68,
					"week2": 62,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      19,
					Name:    "Farhod",
					Surname: "Olimzoda",
					Phone:   "987654318",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 88,
					"week2": 86,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      20,
					Name:    "Foziljon",
					Surname: "Muminov",
					Phone:   "987654319",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 85,
					"week2": 88,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      21,
					Name:    "Alijon",
					Surname: "Boboyorov",
					Phone:   "987654320",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 89,
					"week2": 74,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      22,
					Name:    "Khurshed",
					Surname: "Rahimov",
					Phone:   "987654321",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 96,
					"week2": 99,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      23,
					Name:    "Mehrdod",
					Surname: "Rahmatov",
					Phone:   "987654322",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 96,
					"week2": 94,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      24,
					Name:    "Muhammad",
					Surname: "Hoshimov",
					Phone:   "987654323",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 83,
					"week2": 56,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      25,
					Name:    "Muhammad",
					Surname: "Khujaev",
					Phone:   "987654324",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 91,
					"week2": 88,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      26,
					Name:    "Nozim",
					Surname: "Safarov",
					Phone:   "987654325",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 99,
					"week2": 90,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      27,
					Name:    "Sunatullo",
					Surname: "Gafurov",
					Phone:   "987654326",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 88,
					"week2": 97,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      28,
					Name:    "Tamim",
					Surname: "Orif",
					Phone:   "987654327",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 89,
					"week2": 71,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
			{
				Person: models.Person{
					ID:      29,
					Name:    "Zohira",
					Surname: "Furmal",
					Phone:   "987654328",
				},
				GroupID: 3,
				Grades: map[string]int{
					"week1": 85,
					"week2": 74,
					"week3": 33,
					"week4": 1,
				},
				TopicsToWorkOn: nil,
			},
		},
	},
}

var students = append(groups[0].Students, groups[1].Students...)
