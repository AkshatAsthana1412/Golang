package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

const dateLayout = "2006-01-02"

func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.Format(dateLayout) + `"`), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

type Event struct {
	Name string `json:"name"`
	When Date   `json:"when"`
}

func main() {
	e := Event{Name: "ship", When: Date{time.Date(2026, 5, 9, 0, 0, 0, 0, time.UTC)}}

	data, _ := json.Marshal(e)
	fmt.Println("encoded:", string(data))

	var back Event
	_ = json.Unmarshal(data, &back)
	fmt.Printf("decoded: %+v\n", back)
}
