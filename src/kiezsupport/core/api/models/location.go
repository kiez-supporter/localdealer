package models

import (
	"kiezsupport/core/api/models/orm"
	"time"
)

/*{
zip: 30455,
name: "Berlin",
shops: 36,
progress: 50,
links: ["www.google.de"],
polygon: [
{ lat: 12.364, lng: 53.207 },
{ lat: 15.364, lng: 53.207 },
{ lat: 15.364, lng: 58.207 },
{ lat: 12.364, lng: 58.207 }
]
},*/

type (
	Location struct {
		BaseModel
		Zip      string   `json:"zip"`
		Name     string   `json:"name"`
		Shops    string   `json:"shops"`
		Progress string   `json:"progress"`
		Links    []string `json:"links"`
		Polygon  []string `json:"polygon"`
	}

	LocationPaginationResponse struct {
		Meta orm.PaginationResponse `json:"meta"`
		Data []Location             `json:"data"`
	}

	// just use string type, since it will be use on query at DB layer
	LocationFilterable struct {
		Zip  string `json:"zip"`
		Name string `json:"name"`
	}
)

var (
	_page = 1
	_rp   = 25
)

// Callback before update location
func (m *Location) BeforeUpdate() (err error) {
	m.UpdatedAt = time.Now()
	return
}

// Callback before create location
func (m *Location) BeforeCreate() (err error) {
	m.CreatedAt = time.Now()
	return
}

// Create
func Create(m *Location) (*Location, error) {
	var err error
	err = orm.Create(&m)
	return m, err
}

// Update
func (m *Location) Update() error {
	var err error
	err = orm.Save(&m)
	return err
}

// Delete
func (m *Location) Delete() error {
	var err error
	err = orm.Delete(&m)
	return err
}

// FindLocationById
func FindLocationById(id int) (Location, error) {
	var (
		location Location
		err      error
	)
	err = orm.FindOneByID(&location, id)
	return location, err
}

// FindAllLocations
func FindAllLocations(page int, rp int, filters interface{}) (interface{}, error) {
	var (
		locations []Location
		err       error
	)

	resp, err := orm.FindAllWithPage(&locations, page, rp, filters)
	return resp, err
}
