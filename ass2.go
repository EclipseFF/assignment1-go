package main

import (
	"fmt"
)

type observer interface {
	handleEvent(vacancies []string)
}

type observable interface {
	subscribe(o observer)
	unsubscribe(o observer)
	sendAll()
}

type Person struct {
	name string
}

func (p *Person) handleEvent(vacancies []string) {
	fmt.Println("Hello " + p.name)
	fmt.Println("Vacancies updated: ")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
}

type JobWebsite struct {
	name        string
	vacancies   []string
	subscribers []observer
}

func (w *JobWebsite) addVacancy(vacancy string) {
	w.vacancies = append(w.vacancies, vacancy)
	w.sendAll()
}

func (w *JobWebsite) removeVacancy(vacancy string) {
	index := 0
	for sindex, wvacancy := range w.vacancies {
		if wvacancy == vacancy {
			index = sindex
		}
	}
	w.vacancies = append(w.vacancies[:index], w.vacancies[index+1:]...)
	w.sendAll()
}

func (w *JobWebsite) subscribe(o observer) {
	w.subscribers = append(w.subscribers, o)
}

func (w *JobWebsite) unsubscribe(o observer) {
	index := 0
	for sindex, subscribe := range w.subscribers {
		if subscribe == o {
			index = sindex
		}
	}
	w.subscribers = append(w.subscribers[:index], w.subscribers[index+1:]...)
}

func (w *JobWebsite) sendAll() {
	for _, subscriber := range w.subscribers {
		subscriber.handleEvent(w.vacancies)
	}
}

func main() {
	bob := Person{name: "Bob"}
	hhKZ := JobWebsite{name: "hhkz"}
	hhKZ.subscribe(&bob)
	hhKZ.addVacancy("1")
	hhKZ.addVacancy("2")
	pop := Person{name: "Pip"}
	hhKZ.subscribe(&pop)
	hhKZ.addVacancy("3")
	fmt.Println("--------------------")
	hhKZ.unsubscribe(&pop)
	hhKZ.removeVacancy("1")
}
