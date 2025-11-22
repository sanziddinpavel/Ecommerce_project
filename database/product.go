package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}
func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}

	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}
func Delete(ProductID int) {
	var tempList []Product
	for _, p := range productList {
		if p.ID != ProductID {
			tempList = append(tempList, p)

		}
	}
	productList = tempList
}

func init() {
	pd1 := Product{
		ID:          1,
		Title:       "mango",
		Description: "I love this",
		Price:       1234,
		ImgUrl:      "https://png.pngtree.com/png-vector/20250913/ourmid/pngtree-ripe-fresh-yellow-mango-fruit-png-image_17426358.webp",
	}
	pd2 := Product{
		ID:          2,
		Title:       "Banana",
		Description: "Banana is boring",
		Price:       14332,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTnlrPIEqTWTlxd8qaZzLTpdt8jYNrgJT25oA&s",
	}
	pd3 := Product{
		ID:          3,
		Title:       "Apple",
		Description: "I like apple",
		Price:       6467,
		ImgUrl:      "https://www.collinsdictionary.com/images/full/apple_158989157.jpg",
	}
	productList = append(productList, pd1, pd2, pd3)

}
