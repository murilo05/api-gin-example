package entities

type Error struct {
	ErrorMessage string `json:"errorMessage"`
	Code         int    `json:"code"`
}

type Env struct {
	HOST        string
	PORT        string
	DBHOST      string
	DBPORT      string
	DBUSER      string
	DBPASSAWORD string
	DBNAME      string
	TIMEOUT     int
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
}
