package bexio

const (
	// ContactTypeCompany identifies a company.
	ContactTypeCompany int = 1

	// ContactTypePerson identifiers a person.
	ContactTypePerson int = 2
)

// Contact is the main contact model.
type Contact struct {
	ID          int    `json:"id,omitempty"`
	ContactType int    `json:"contact_type_id"`
	OwnerID     int    `json:"owner_id"`
	UserID      int    `json:"user_id"`
	Mail        string `json:"mail"`
	Name1       string `json:"name_1"`
	Name2       string `json:"name_2"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Postcode    string `json:"postcode"`
	CountryID   int    `json:"country_id,omitempty"`
	PhoneFixed  string `json:"phone_fixed"`
	PhoneMobile string `json:"phone_mobile"`
}
