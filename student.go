package main

import (
	"fmt"
	"strings"
)

type StudentList []Student

type Student struct {
	UserID              int64   `json:"UserID"`
	FirstName           string  `json:"FirstName"`
	LastName            string  `json:"LastName"`
	PublishUserProfile  bool    `json:"PublishUserProfile"`
	EmailBad            bool    `json:"EmailBad"`
	PublishEmail        bool    `json:"PublishEmail"`
	AddressLine1        string  `json:"AddressLine1,omitempty"`
	City                string  `json:"City,omitempty"`
	State               string  `json:"State,omitempty"`
	Zip                 string  `json:"Zip,omitempty"`
	MyContactsID        int64   `json:"MyContactsID"`
	HasRelationships    bool    `json:"HasRelationships"`
	Nickname            string  `json:"Nickname,omitempty"`
	GradYear            string  `json:"GradYear"`
	Grade               string  `json:"Grade"`
	GradeDisplay        string  `json:"GradeDisplay"`
	Department          string  `json:"Department"`
	SpouseID            int64   `json:"SpouseID"`
	GradeNumeric        int64   `json:"GradeNumeric"`
	GradeNumericDisplay string  `json:"GradeNumericDisplay"`
	TotalCount          int64   `json:"TotalCount"`
	Deceased            bool    `json:"Deceased"`
	IsStudentInd        bool    `json:"IsStudentInd"`
	PreferredAddressID  int64   `json:"PreferredAddressId"`
	PreferredAddressLat float64 `json:"PreferredAddressLat"`
	PreferredAddressLng float64 `json:"PreferredAddressLng"`
	JobTitle            string  `json:"JobTitle"`
	PhotoEditSettings   string  `json:"PhotoEditSettings,omitempty"`
	Email               string  `json:"Email,omitempty"`
	HomePhone           string  `json:"HomePhone,omitempty"`
	LargeFileName       string  `json:"LargeFileName,omitempty"`
}

func CleanLine(v *string) {
	*v = strings.TrimSpace(*v)
}

type StudentFile struct {
	Records StudentList `json:"students"`
}

func (r *StudentList) Clean() {
	for _, stud := range *r {
		CleanLine(&stud.AddressLine1)
		CleanLine(&stud.HomePhone)
		CleanLine(&stud.City)
		CleanLine(&stud.Email)
		CleanLine(&stud.City)
		CleanLine(&stud.GradeDisplay)
	}
}

// WhateverToString returns what as a string
// only works on intigers and numbers and stuff
func WhateverToString(what interface{}) string {
	s, ok := what.(string)
	if !ok {
		return fmt.Sprintf("%v", what)
	}

	return fmt.Sprintf("%s", s)
}

// ToCSVSlice will turn the student in a CSV element!
func (s *Student) ToCSVSlice() []string {
	return []string{
		WhateverToString(s.UserID),
		WhateverToString(s.Nickname),
		WhateverToString(s.FirstName),
		WhateverToString(s.LastName),
		WhateverToString(s.Email),
		WhateverToString(s.GradeDisplay),
		WhateverToString(s.AddressLine1),
		WhateverToString(s.PreferredAddressLat),
		WhateverToString(s.PreferredAddressLng),
		WhateverToString(s.City),
		WhateverToString(s.State),
		WhateverToString(s.Zip),
		WhateverToString(s.HomePhone),
	}
}

// ToCSVSlice will convert the entire student list into a csv slice
func (sl *StudentList) ToCSVSlice() [][]string {
	studs := make([][]string, 0) // couldn't figure out how to get the length of this properly, whatever...

	// the first line in any csv file is the heading
	studs = append(studs, []string{
		"user_id",
		"nickname",
		"first_name",
		"last_name",
		"email_address",
		"grade_display",
		"address_line_1",
		"address_lat",
		"address_lng",
		"address_city",
		"address_state",
		"address_zip",
		"home_phone",
	})

	// add all the students
	for _, s := range *sl {
		studs = append(studs, s.ToCSVSlice())
	}

	return studs
}
