package data

type User struct {
	Id          string `json:"id,omitempty"`
	FirstName   string `json:"firstName,omitempty" validate:"required"`
	LastName    string `json:"lastName,omitempty" validate:"required"`
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	Avatar      string `json:"avatar,omitempty"`
}

type Records struct {
	Records []User `json:"records,omitempty" validate:"required"`
}

type CreateResponse struct {
	Id string `json:"id,omitempty" validate:"required"`
}