package models

type Student struct {
	Person                        // студент является человеком, поэтому нам нужны все поля структуры Person
	GroupID        uint           // идентификатор группы, в котором учится студент
	Grades         map[string]int // баллы по неделям
	TopicsToWorkOn []string       // темы, над которыми студенту нужно поработать
}
