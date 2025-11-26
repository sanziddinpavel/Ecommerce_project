package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(ProductID int) error
	Update(p Product) (*Product, error)
}
type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateIntialProduct(repo)
	return repo

}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}

	return nil, nil

}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil

}

func (r *productRepo) Delete(ProductID int) error {
	var tempList []*Product
	for _, p := range r.productList {
		if p.ID != ProductID {
			tempList = append(tempList, p)

		}
	}
	r.productList = tempList
	return nil

}

func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil

}
func generateIntialProduct(r *productRepo) {
	pd1 := &Product{
		ID:          1,
		Title:       "mango",
		Description: "I love this",
		Price:       1234,
		ImgUrl:      "https://png.pngtree.com/png-vector/20250913/ourmid/pngtree-ripe-fresh-yellow-mango-fruit-png-image_17426358.webp",
	}
	pd2 := &Product{
		ID:          2,
		Title:       "Banana",
		Description: "Banana is boring",
		Price:       14332,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTnlrPIEqTWTlxd8qaZzLTpdt8jYNrgJT25oA&s",
	}
	pd3 := &Product{
		ID:          3,
		Title:       "Apple",
		Description: "I like apple",
		Price:       6467,
		ImgUrl:      "https://www.collinsdictionary.com/images/full/apple_158989157.jpg",
	}
	r.productList = append(r.productList, pd1, pd2, pd3)

}
