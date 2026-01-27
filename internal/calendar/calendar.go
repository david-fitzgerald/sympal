package calendar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/david-fitzgerald/sympal/internal/auth"
	"github.com/david-fitzgerald/sympal/internal/config"
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
	accesstoken, refreshtoken, err := keyring.LoadTokens()
	if err != nil {
		return nil, err
	}
	// Build HTTP request
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	timeMin := startOfDay.Format(time.RFC3339)
	timeMax := endOfDay.Format(time.RFC3339)

	requestURL := fmt.Sprintf("https://www.googleapis.com/calendar/v3/calendars/primary/events?timeMin=%s&timeMax=%s&singleEvents=true", url.QueryEscape(timeMin), url.QueryEscape(timeMax))
	req, err := http.NewRequest("GET", requestURL, nil)
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
	// check if error code = 200, then refresh using NewToken func and save in keyring
	if resp.StatusCode == http.StatusUnauthorized {
		token, err := auth.NewToken(refreshtoken, config.Current.Google.ClientID, config.Current.Google.ClientSecret)
		if err != nil {
			return nil, err
		}
		err = keyring.SaveTokens(token.AccessToken, refreshtoken)
		if err != nil {
			return nil, err
		}
		// if 200 error, after refresh/save, try HTTP again
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
		client = &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			return nil, err
		}

	}
	// check if error code != 200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Calendar API error: %s", resp.Status)
	}

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
	// Return events
	return events, nil

}
