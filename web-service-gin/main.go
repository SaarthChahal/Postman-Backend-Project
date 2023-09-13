package main

// create struct for Headphone

type headphone struct {
	ID    string  `json:"id"`
	Item  string  `json:"item"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}


// headphones record to see headphones data
var headphones = []headphone{
	{ID: "1", Item: "Wireless Headphone", Brand: " Sennheiser", Price: 8990},
	{ID: "2", Item: "Foldable Headphone", Brand: "Zebronics", Price: 783},
	{ID: "3", Item: " Wired Headphone", Brand: "JBL", Price: 599},
	{ID: "4", Item: "Bluetooth Headphone", Brand: "boAT", Price: 999},
	{ID: "5", Item: "Noise Cancellation Headphones", Brand: "Sony", Price: 9990},
}
