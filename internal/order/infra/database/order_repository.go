package database

import (
	"database/sql"

	"github.com/hgribeiro/aluno-go/internal/order/entity"
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (o *OrderRepository) Save(order entity.Order) error {
	stmt, err := o.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	// ESSE UNDRELINE __ É CHAMADO DE BLANK IDENTIFYER
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
