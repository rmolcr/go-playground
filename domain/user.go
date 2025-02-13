package domain

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"mail"`
	Number  int    `json:"number"`
}
