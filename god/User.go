package main

type User struct {
	Name      string `json:"name"`
	Id        int    `json:"id"`
	Password  string
	Shortname string
}

type IUser struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type UserList struct {
	UserList []IUser `json:"userList"`
}

var users = map[int]User{
	1: {
		Name:     "rassoul",
		Id:       1,
		Password: "manche3ter",
		Shortname: "Rss",
	},
	2: {
		Name:     "ali",
		Id:       2,
		Password: "ali",
		Shortname: "Ali",
	},
	3: {
		Name:     "roya",
		Id:       3,
		Password: "roya",
		Shortname: "Rye",
	},
	4: {
		Name:     "hossein",
		Id:       4,
		Password: "hossein",
		Shortname: "Hsn",
	},
	5: {
		Name:     "heravi",
		Id:       5,
		Password: "heravi",
		Shortname: "MHr",
	},
	6: {
		Name:      "gorbani",
		Id:        6,
		Password:  "gorbani",
		Shortname: "HRZ",
	},
	7: {
		Name:      "mazdak",
		Id:        7,
		Password:  "mazdak",
		Shortname: "Mzd",
	},
}

var usersNameIndex = map[string]int{
	"rassoul": 1,
	"ali":     2,
	"roya":    3,
	"hossein": 4,
	"heravi":  5,
	"sanati": 6,
	"mazdak":  7,
}

func (user User) checkPassword(password string) bool {
	return true
	//return user.Password == password
}

func (user User) transform() IUser {
	return IUser{
		Id:   user.Id,
		Name: user.Name,
	}
}
