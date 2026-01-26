package calendar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/david-fitzgerald/sympal/internal/keyring"
)

type Event struct {
	ID        string
	Title     string
	StartTime time.Time
	EndTime   time.Time
}

type calendarResponse struct {
	Items []calendarEvent `json:"items"`
}

type calendarEvent struct {
	ID      string    `json:"id"`
	Summary string    `json:"summary"`
	Start   eventTime `json:"start"`
	End     eventTime `json:"end"`
}

type eventTime struct {
	DateTime string `json:"dateTime"`
}

func GetTodayEvents() ([]Event, error) {
	// Load token
	accesstoken, _, err := keyring.LoadTokens()
	if err != nil {
		return nil, err
	}
	// Build HTTP request
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	timeMin := startOfDay.Format(time.RFC3339)
	timeMax := endOfDay.Format(time.RFC3339)

	url := fmt.Sprintf("https://www.googleapis.com/calendar/v3/calendars/primary/events?timeMin=%s&timeMax=%s&singleEvents=true", timeMin, timeMax)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accesstoken)
	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var calResp calendarResponse
	err = json.Unmarshal(body, &calResp)
	if err != nil {
		return nil, err
	}
	events := []Event{}
	for _, item := range calResp.Items {
		startTime, err := time.Parse(time.RFC3339, item.Start.DateTime)
		if err != nil {
			// log warning and skip this event
			continue
		}
		endTime, err := time.Parse(time.RFC3339, item.End.DateTime)
		if err != nil {
			// log warning and skip this event
			continue
		}

		events = append(events, Event{
			ID:        item.ID,
			Title:     item.Summary,
			StartTime: startTime,
			EndTime:   endTime,
		})
	}
	return events, nil

	// Return events
}
