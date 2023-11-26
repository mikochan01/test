package main

import (
	"fmt"
	"time"
)

// Model: Event struct sebagai representasi event konser
type Event struct {
	ID          int
	Title       string
	Description string
	Date        time.Time
	Location    string
	Tickets     int
	Price       float64
}

// Controller: EventManager sebagai pengelola event
type EventManager struct {
	Events []Event
}

// View: Tampilkan informasi event
func (e Event) DisplayEventInfo() {
	fmt.Printf("Event ID: %d\nTitle: %s\nDescription: %s\nDate: %s\nLocation: %s\nTickets Available: %d\nPrice: $%.2f\n\n",
		e.ID, e.Title, e.Description, e.Date.Format("02-01-2006"), e.Location, e.Tickets, e.Price)
}

// Controller: Menambahkan event baru
func (em *EventManager) AddEvent(event Event) {
	em.Events = append(em.Events, event)
}

// Controller: Cari event berdasarkan ID
func (em EventManager) FindEventByID(id int) (Event, error) {
	for _, event := range em.Events {
		if event.ID == id {
			return event, nil
		}
	}
	return Event{}, fmt.Errorf("Event ID %d not found", id)
}

func main() {
	// Inisialisasi EventManager
	eventManager := EventManager{}

	// Menambahkan beberapa event
	eventManager.AddEvent(Event{
		ID:          1,
		Title:       "Concert A",
		Description: "An amazing concert experience!",
		Date:        time.Date(2023, 12, 10, 18, 0, 0, 0, time.UTC),
		Location:    "Venue X",
		Tickets:     100,
		Price:       50.0,
	})
	eventManager.AddEvent(Event{
		ID:          2,
		Title:       "Concert B",
		Description: "A night to remember!",
		Date:        time.Date(2023, 12, 15, 19, 30, 0, 0, time.UTC),
		Location:    "Venue Y",
		Tickets:     50,
		Price:       75.0,
	})

	// Menampilkan informasi event
	for _, event := range eventManager.Events {
		event.DisplayEventInfo()
	}

	// Contoh pencarian event berdasarkan ID
	foundEvent, err := eventManager.FindEventByID(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Event found:")
		foundEvent.DisplayEventInfo()
	}
}
