package patterns

import (
	"fmt"
	"net/http"
)

type DB interface {
	Store(string) error
}

type Store struct {
	Data string
}

func (s *Store) Store(value string) error {
	fmt.Printf("This is Store function | Store Data : %s | storing in db : %s \n", s.Data, value)

	return nil
}

func ExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("This is ExecuteFunc")
		db.Store(s)
	}
}

func MakeHTTPFunc(db DB, fn HttpFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := fn(db, writer, request); err != nil {
			fmt.Println("This is MakeHTTPFunc - Error Handling")
		}
	}
}

func Handler(db DB, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("This is Handler")

	return nil
}

type HttpFunc func(db DB, w http.ResponseWriter, r *http.Request) error

type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("Execute() : Hello World!")
}

//
// main code :
//
//s := &patterns.Store{Data: "Some Store Data"}
//http.HandleFunc("/", patterns.MakeHTTPFunc(s, patterns.Handler))
//fmt.Println("This is main - HandleFunc Done!")
//patterns.Execute(patterns.ExecuteFunc(s))
//fmt.Println("This is main - Execute Done!")
//
