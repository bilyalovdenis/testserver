package model


type Product struct{
	ID int
	Name string
	Description string
	Price float64
	Category string
	Quantity int
	Photo string
}

// func ConvertAndValidateId(id string) (int, error){

// }

// func (p *Product) Validate() error{
// 	return validation.ValidateStruct(p, validation.Field(&p.ID, validation.Required))
// }
