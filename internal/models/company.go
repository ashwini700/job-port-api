package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Jobs     []Job  `json:"jobs,omitempty" gorm:"foreignKey:CompanyId"`
}
type NewCompany struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type Job struct {
	gorm.Model
	CompanyId       uint64          `json:"companyId"`
	JobRole         string          `json:"job_role"`
	Salary          uint            `json:"salary"`
	MinNotice       uint            `json:"minnotice"`
	MaxNotice       uint            `json:"maxnotice"`
	Budget          uint            `json:"budget"`
	JobLocations    []Loc           `gorm:"many2many:job_location;"`
	TechnologyStack []Tech_stack    `gorm:"many2many:tech_stack;"`
	Description     string          `json:"desc"`
	MinExp          uint            `json:"min_exp"`
	MaxMax          uint            `json:"max_exp"`
	Qualification   []Qualification `gorm:"many2many:qualification;"`
}

type NewJob struct {
	JobRole         string          `json:"job_role"`
	Salary          uint         `json:"salary"`
	MinNotice       uint            `json:"minnotice"`
	MaxNotice       uint            `json:"maxnotice"`
	Budget          uint            `json:"budget"`
	JobLocations    []Loc           `json:"joblocs"`
	TechnologyStack []Tech_stack    `json:"techstack"`
	Description     string          `json:"desc"`
	MinExp          uint            `json:"min_exp"`
	MaxMax          uint            `json:"max_exp"`
	Qualification   []Qualification `json:"qualification"`
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
