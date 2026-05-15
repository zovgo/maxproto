package protocol

type Contact struct {
	RegistrationTime Timestamp     `json:"registrationTime"`
	AccountStatus    int           `json:"accountStatus"`
	Country          string        `json:"country"`
	BaseUrl          string        `json:"baseUrl"`
	Names            []ContactName `json:"names"`
	Phone            int64         `json:"phone"`
	Options          []string      `json:"options"`
	PhotoID          int64         `json:"photoId"`
	UpdateTime       Timestamp     `json:"updateTime"`
	ID               int64         `json:"id"`
	BaseRawURL       string        `json:"baseRawUrl"`
}

type ContactName struct {
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Type      string `json:"type"`
}
