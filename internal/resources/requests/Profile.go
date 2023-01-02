package requests

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	OtherData string `json:"other_data"`
	Address   string `json:"address"`
}
