package api

import (
	"net/http"
	"time"
)

type Deployment struct {
	ID          string
	Name        string
	Description string
	Source      []byte `json:"-"`
	Image       string
	Date        time.Time
	Language    string
	URL         string
	Ready       bool
	BuildLogs   string `json:"-"`
	User        string
}

type Alias struct {
	Name string
	For  string
	URL  string
	Date time.Time
}

type Registry struct {
	Items   map[string]*Deployment
	Aliases map[string]*Alias
}

type Server struct {
	http.ServeMux
	Registry *Registry
}

type TriggerRequestBody struct {
	Duration string `json:"duration"`
}

type DeployBody struct {
	Code         string `json:"code"`
	Description  string `json:"description"`
	FunctionName string `json:"functionName"`
	Lang         string `json:"lang"`
	User         string `json:"user"`
}
