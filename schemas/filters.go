package schemas

//list all your parameter search here
type ParamSearch struct {
	ID          uint64
	Offset      int
	Limit       int
	Search      string
	Name        string
	Email       string
	PhoneNumber string
	Pagination  bool
}
