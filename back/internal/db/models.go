// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Application struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Image  string `json:"image"`
	Port   string `json:"port"`
}

type Deployment struct {
	ID         string       `json:"id"`
	AppID      string       `json:"app_id"`
	Status     string       `json:"status"`
	CreatedAt  time.Time    `json:"created_at"`
	FinishedAt sql.NullTime `json:"finished_at"`
}