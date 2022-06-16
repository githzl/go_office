package pgmodel

const (
	TYPE_APP = 0
	TYPE_DOCUMENT = 1
)
type Tags struct {
	Id int 	`json:"id" structs:"id"`
	Name string `json:"name" structs:"name"`
	Logo string `json:"logo" structs:"logo"`
}


