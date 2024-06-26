package entity

import "errors"

type OrderRepositoryInterface interface {
	Save(order *Order) error
}

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

// ISSO AQUI É UM METODO, ELE RECEBE UM >RECEPTOR< QUANDO É DEFINIDO E,
// REALIZA ALGUMA ACAO NO RECEPTOR, COMO É UM PONTEIRO O RECEPTOR, ELE MODIFICA NA INSTANCIA QUE
// O METODO FOI INVOCADO

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}

	if o.Price == 0 {
		return errors.New("invalid price")
	}

	if o.Tax == 0 {
		return errors.New("invalid tax")
	}
	return nil
}
