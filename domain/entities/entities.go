package entities

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
}

type Error struct {
	ErrorMessage string `json:"errorMessage"`
	Code         int    `json:"code"`
}
