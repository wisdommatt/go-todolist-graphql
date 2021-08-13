package todo

import "time"

type Todo struct {
	ID        int `gorm:"autoIncrement,primaryKey"`
	Task      string
	Status    string
	TimeAdded time.Time
}
