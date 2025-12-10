package models

type Product struct {
	ID    string `json:"ID"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

func SampleProducts() []Product {
	var products1 = []Product{
	  {
      ID:    "1",
      Name:  "Laptop",
      Stock: 50,
    },   
  }
	return products1
}
