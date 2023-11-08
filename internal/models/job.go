package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Company         Company      `json:"-" gorm:"foreignKey:cid"`
	Cid             uint         `json:"cid"`
	JobRole         string       `json:"job_role"`
	Salary          string       `json:"salary"`
	MinNotice       uint         `json:"minnotice"`
	MaxNotice       uint         `json:"maxnotice"`
	Budget          uint         `json:"budget"`
	JobLocations    []Loc        `gorm:"many2many:job_location;"`
	TechnologyStack []Tech_stack `gorm:"many2many:tech_stack;"`
	// WorkMode        string          `json:"work_mode"`
	Description   string          `json:"desc"`
	MinExp        uint            `json:"min_exp"`
	MaxMax        uint            `json:"max_exp"`
	Qualification []Qualification `gorm:"many2many:qualification;"`
	// Shift           string          `json:"shift"`
	// JobType         string          `json:"job_type"`
}

type Loc struct {
	gorm.Model
	City string `json:"city"`
}

type Tech_stack struct {
	gorm.Model
	Skills string `json:"skill"`
}

type Qualification struct {
	gorm.Model
	Graduation string `json:"grad"`
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
