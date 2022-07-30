package db

import "time"

type Results struct {
	ID          string
	Name        string
	Description string
	User        string
	Source      []byte `json:"-"`
	Image       string
	Date        time.Time
	Language    string
	URL         string
	Ready       bool
	BuildLogs   string `json:"-"`
}
