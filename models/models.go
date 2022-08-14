package models

type Person struct {
	Login string `json:"login" xml:"login" form:"login"`
	Pass  string `json:"pass" xml:"pass" form:"pass"`
}
