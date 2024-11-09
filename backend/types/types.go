package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//main struct

type Company struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Company        string             `bson:"company,omitempty" json:"company,omitempty"`
	URL            string             `bson:"url,omitempty" json:"url,omitempty"`
	RemoteFriendly bool               `bson:"remoteFriendly,omitempty" json:"remoteFriendly,omitempty"`
	Market         string             `bson:"market,omitempty" json:"market,omitempty"`
	Size           string             `bson:"size,omitempty" json:"size,omitempty"`
	// AllJobs        []JobDetails       `bson:"allJobs,omitempty" json:"allJobs,omitempty"`
}

type SalaryRange struct {
	From     int32  `bson:"from,omitempty" json:"from,omitempty"`
	To       int32  `bson:"to,omitempty" json:"to,omitempty"`
	Currency string `bson:"currency,omitempty" json:"currency,omitempty"`
}

type Equity struct {
	From float64 `bson:"from,omitempty" json:"from,omitempty"`
	To   float64 `bson:"to,omitempty" json:"to,omitempty"`
}

type Job struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CompanyID   primitive.ObjectID `bson:"companyID,omitempty" json:"companyID,omitempty"`
	Position    string             `bson:"position,omitempty" json:"position,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	URL         string             `bson:"url,omitempty" json:"url,omitempty"`
	Type        string             `bson:"type,omitempty" json:"type,omitempty"`
	Posted      string             `bson:"posted,omitempty" json:"posted,omitempty"`
	Location    string             `bson:"location,omitempty" json:"location,omitempty"`
	Skills      []string           `bson:"skills,omitempty" json:"skills,omitempty"`
	SalaryRange SalaryRange        `bson:"salaryRange,omitempty" json:"salaryRange,omitempty"`
	Equity      Equity             `bson:"equity,omitempty" json:"equity,omitempty"`
	Perks       []string           `bson:"perks,omitempty" json:"perks,omitempty"`
	Apply       string             `bson:"apply,omitempty" json:"apply,omitempty"`
}

// type JobFilter struct {
// 	CompanyID int32    `json:"companyID,omitempty"`
// 	Title     string   `json:"title,omitempty"`
// 	Type      string   `json:"type,omitempty"`
// 	State     string   `json:"state,omitempty"`
// 	Skills    []string `json:"skills,omitempty"`
// }

type Mentor struct {
	Name       string `bson:"name,omitempty" json:"name,omitempty"`
	Title      string `bson:"title,omitempty" json:"title,omitempty"`
	Company    string `bson:"company,omitempty" json:"company,omitempty"`
	ProfilePic string `bson:"profilePic,omitempty" json:"profilePic,omitempty"`
}
