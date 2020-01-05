package app

type Friend struct {
	ID        int     `json: "id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Gender    string  `json:"gender"`
	Language  string  `json:"language"`
	Emails    []Email `json:"emails"`
}

type Email struct {
	ID      int    `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
}
