package domain

type Order struct {
	ID         string
	CustomerID string
	Items      []Item
}

type Item struct {
	ID    string
	Title string
	Price float64
}
