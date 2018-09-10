package token

var Name = "test token"
var Symbol = "tt"
var TotalAmount = 100000000

func Token() (string, string, int) {
	return Name, Symbol, TotalAmount
}

func AddPublish(amount int) {
	TotalAmount += amount
}
