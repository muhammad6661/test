package types


//Money for Balance
type Money int64
//Currency for type money
type Currency string

const (
	TJS Currency="TJS"
	USD Currency="USD"
	RUB Currency="RUB"
  EUR Currency="EUR"
)
type PAN string

type Card struct {
  ID int 
  PAN PAN
  Balance Money
  Currency Currency
  Color string 
  Name string 
  Active bool
}