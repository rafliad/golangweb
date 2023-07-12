package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() (status string) {
	if p.Stock < 3 {
		status = "- stock hampir habis"
	} else if p.Stock < 10 {
		status = "- stock terbatas"
	}
	return
}
