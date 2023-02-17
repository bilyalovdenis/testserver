package store

import (
	"fmt"

	"github.com/bilyalovdenis/testserver/internal/app/model"
)

type ProductRepository struct{
	store *Store
}

// func (r *ProductRepository) Create(u *model.Product) (*model.Product, error){
// 	r.store.db.QueryRow("")

// 	return nil, nil
// }

func (r *ProductRepository) FindById(ID int) (*model.Product, error){
	p := &model.Product{}
	row := r.store.db.QueryRow("SELECT * FROM product WHERE id = ?", ID)
	if err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price,
	                   &p.Category, &p.Quantity, &p.Photo); err != nil{
		fmt.Println(err)
		return nil,err
	}
	return p, nil
}