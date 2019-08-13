package core

//Total contains a summary of all transactions
type Total struct {
	Completed int64
	Cancelled int64
	Active    int64
}
