package main

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"test/templates"
)

//go:embed style.css
var Styles string

var SomeGlobal = "lol"

const YourName = "lois"

func main() {
	fmt.Println("Hello World!")

	// var lois *Person = NewPerson("lois")
	lois := NewPerson("lois")

	lois.privateParts = append(lois.privateParts, "lol")

	rnd1, ok := lois.Random["laptop"]
	if !ok {
		fmt.Println("didn't find a laptop 1")
	} else {
		fmt.Println("Laptop has", rnd1)
	}

	lois.Random["laptop"] = 0

	rnd1, ok = lois.Random["laptop"]
	if !ok {
		fmt.Println("didn't find a laptop 2")
	} else {
		fmt.Println("Laptop has", rnd1)
	}

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(Styles))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "Hello world")

		templates.Index("Lois").Render(context.Background(), w)
	})

	//var err error
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// class Person { public Name: string; private privateParts: string[] }
type Person struct {
	Name         string
	Age          int
	Savings      []int64
	privateParts []string
	Random       map[string]int
	ComplexStuff complex128 // used for 3d renderng for co
}

func NewPerson(name string) *Person {
	complexStuff := complex(12, 0)

	return &Person{
		Name:         name,
		privateParts: make([]string, 0),
		ComplexStuff: complexStuff,
		Random:       make(map[string]int),
	}
}
