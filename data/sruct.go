package data

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Company  string   `json:"company"`
	Country  string   `json:"country"`
	Job      string   `json:"job"`
	Browsers []string `json:"browsers"`
}
