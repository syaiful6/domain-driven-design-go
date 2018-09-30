package cart

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/user"
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	FindOrCreateCart(tx *gorm.DB, user user.User) (*Cart, e.Exception)
	FindAllCartItems(tx *gorm.DB, cart *Cart) ([]*CartItem, e.Exception)
	AddCartItem(tx *gorm.DB, cart *Cart, currItemCount int, quantity uint, productID uint, productOptionIDList []uint) (*CartItem, e.Exception)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) FindOrCreateCart(tx *gorm.DB, user user.User) (*Cart, e.Exception) {
	record := &Cart{
		UserID:     user.ID,
		TotalPrice: 0,
	}

	err := tx.Where("user_id = ?", user.ID).First(record).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewInternalServerException(err, "Failed to get user cart")
		}
	}

	// if exists return it
	if err == nil {
		return record, nil
	}

	if err = tx.Create(record).Error; err != nil {
		return nil, e.NewInternalServerException(err, "Failed to create Cart")
	}

	return record, nil
}

func (r *repositoryImpl) FindAllCartItems(tx *gorm.DB, cart *Cart) ([]*CartItem, e.Exception) {
	var cartItemList []*CartItem

	if err := tx.Model(cart).Related(&cartItemList).Error; err != nil {
		ex := e.NewInternalServerException(err, "Failed to find CartItem list")
		return nil, ex
	}

	for i := range cartItemList {
		cartItem := cartItemList[i]
		var cartItemOptionList []*CartItemOption

		if err := tx.Model(cartItem).Related(&cartItemOptionList).Error; err != nil {
			ex := e.NewInternalServerException(err, "Failed to find CartItemOption")
			return nil, ex
		}
	}

	return cartItemList, nil
}

func (r *repositoryImpl) AddCartItem(tx *gorm.DB, cart *Cart, currItemCount int, quantity uint, productID uint, productOptionIDList []uint) (*CartItem, e.Exception) {

	cartItem := CartItem{
		Index: uint(currItemCount + 1),
		Quantity: quantity,

		CartID: cart.ID,
		ProductID: productID,
	}

	if err := tx.Create(&cartItem).Error; err != nil {
		ex := e.NewInternalServerException(err, "Failed to create CartItem")
		return nil, ex
	}

	return &cartItem, nil
}
