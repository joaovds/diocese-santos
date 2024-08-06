package parish

type Parish struct {
	*Church
}

type Chapel struct {
	*Church
	ParishID int `json:"parish_id"`
}

type Church struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	ImageURL string   `json:"image_url"`
	Address  *Address `json:"address"`
	Contact  *Contact `json:"contact"`
}

type Address struct {
	Street       string  `json:"street"`
	Neighborhood string  `json:"neighborhood"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	PostalCode   string  `json:"postal_code"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type Contact struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
	Site  string `json:"site"`
}
