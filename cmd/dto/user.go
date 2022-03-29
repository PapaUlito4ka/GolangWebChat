package dto

type UserDto struct {
	Username string   `json:"username"`
	Friends  []uint64 `json:"friends"`
}

type UsersDto struct {
	Users []UserDto `json:"users"`
}
