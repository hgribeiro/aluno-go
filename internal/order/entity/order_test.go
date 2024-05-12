package entity_test

import (
	"testing"

	"github.com/hgribeiro/aluno-go/order/entity"
	"github.com/stretchr/testify/assert"
)

// // NOME DO PACOTE + NOME DA STRUCK + NOME DO QUE VAI FAZER
// É UMA POSSIBILIDADE, MAS VOU POSSO USAR O
// given_when_then
// ENTREGUE XXXX PARAMENTRO -- QUANDO QUAL ACONTECER -- ENTÃO O QUE DEVE ACONTECER
func TestGivingAnEmpyId_WhenCreateANewOrder_ThemShouldReceiveAnError(t *testing.T) {

	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivingAnEmpyPrice_WhenCreateANewOrder_ThemShouldReceiveAnError(t *testing.T) {

	order := entity.Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivingAnEmpyTax_WhenCreateANewOrder_ThemShouldReceiveAnError(t *testing.T) {

	order := entity.Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenValidParams_WhenCreateANewOrder_ThemShouldNotReceiveCreateORderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 1)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
}

func TestGivenAValidParams_WhenCallculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPricePropety(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 1)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 11.0, order.FinalPrice)
}
