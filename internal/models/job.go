package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Company Company `json:"-" gorm:"foreignKey:cid"`
	Cid     uint    `json:"cid"`
	JobRole string  `json:"job_role"`
	Salary  string  `json:"salary"`
	MinNotice int
	MaxNotice int
	Budget[]string
}

// Min-NP
// Max-NP
// Budget
// JobLocations []
// Technology Stack[]
// WorkMode - [Remote,OnSite, Hybrid]
// Description
// MinExp
// MaxMax
// Qualification-[]
// Shift - [day, night, rotational]
// JobType - [full time, part time]