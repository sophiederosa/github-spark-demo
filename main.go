package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/subjects", subjectsPageHandler)
	http.HandleFunc("/quiz", quizPageHandler)

	http.ListenAndServe(":8080", nil)
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "main.html", nil)
}

func subjectsPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "subjects.html", nil)
}

func quizPageHandler(w http.ResponseWriter, r *http.Request) {
	questions := generateMultipleChoiceQuestions()
	templates.ExecuteTemplate(w, "quiz.html", questions)
}

func generateQuizQuestions() []string {
	rand.Seed(time.Now().UnixNano())
	allQuestions := []string{
		"What is the keyword used to declare a variable in Golang?",
		"How do you create a new Goroutine?",
		"What is the purpose of the `defer` statement in Golang?",
		"How do you handle errors in Golang?",
		"What is a struct in Golang?",
		"What is a slice in Golang?",
		"How do you create a map in Golang?",
		"What is the difference between a pointer and a value in Golang?",
		"How do you write a test in Golang?",
		"What is the `interface` keyword used for in Golang?",
	}

	rand.Shuffle(len(allQuestions), func(i, j int) {
		allQuestions[i], allQuestions[j] = allQuestions[j], allQuestions[i]
	})

	return allQuestions[:5]
}

func generateQuizQuestionsFromSubjects() []string {
	subjects := []string{
		"Introduction to Golang",
		"Golang Syntax",
		"Concurrency in Golang",
		"Web Development with Golang",
		"Testing in Golang",
	}

	questions := []string{}
	for _, subject := range subjects {
		questions = append(questions, "What did you learn about "+subject+"?")
	}

	return questions
}

func generateMultipleChoiceQuestions() []map[string]interface{} {
	rand.Seed(time.Now().UnixNano())
	allQuestions := []map[string]interface{}{
		{
			"question": "What is the keyword used to declare a variable in Golang?",
			"choices":  []string{"var", "let", "const", "func"},
			"answer":   "var",
		},
		{
			"question": "How do you create a new Goroutine?",
			"choices":  []string{"go func()", "new Goroutine()", "create Goroutine()", "start Goroutine()"},
			"answer":   "go func()",
		},
		{
			"question": "What is the purpose of the `defer` statement in Golang?",
			"choices":  []string{"To handle errors", "To create a new Goroutine", "To delay the execution of a function", "To declare a variable"},
			"answer":   "To delay the execution of a function",
		},
		{
			"question": "How do you handle errors in Golang?",
			"choices":  []string{"try-catch", "if err != nil", "throw-catch", "handle error"},
			"answer":   "if err != nil",
		},
		{
			"question": "What is a struct in Golang?",
			"choices":  []string{"A collection of methods", "A collection of variables", "A collection of functions", "A collection of constants"},
			"answer":   "A collection of variables",
		},
		{
			"question": "What is a slice in Golang?",
			"choices":  []string{"A fixed-size array", "A dynamic array", "A map", "A struct"},
			"answer":   "A dynamic array",
		},
		{
			"question": "How do you create a map in Golang?",
			"choices":  []string{"make(map[keyType]valueType)", "new map[keyType]valueType", "create map[keyType]valueType", "map[keyType]valueType{}"},
			"answer":   "make(map[keyType]valueType)",
		},
		{
			"question": "What is the difference between a pointer and a value in Golang?",
			"choices":  []string{"A pointer holds the address of a value", "A pointer holds the value itself", "A value holds the address of a pointer", "A value holds the pointer itself"},
			"answer":   "A pointer holds the address of a value",
		},
		{
			"question": "How do you write a test in Golang?",
			"choices":  []string{"Using the `test` keyword", "Using the `testing` package", "Using the `assert` package", "Using the `check` package"},
			"answer":   "Using the `testing` package",
		},
		{
			"question": "What is the `interface` keyword used for in Golang?",
			"choices":  []string{"To define a set of methods", "To define a set of variables", "To define a set of constants", "To define a set of functions"},
			"answer":   "To define a set of methods",
		},
	}

	rand.Shuffle(len(allQuestions), func(i, j int) {
		allQuestions[i], allQuestions[j] = allQuestions[j], allQuestions[i]
	})

	return allQuestions[:5]
}
