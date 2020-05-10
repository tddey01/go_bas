package dao

import (
	"go_bas/go_web/day05/bookstore0612/model"
	"go_bas/go_web/day05/bookstore0612/utils"
)

//AddCartItem 向购物项表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	//写sql
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemByBookID 根据图书的id和购物车的id获取对应的购物项
func GetCartItemByBookID(bookID string, cartID string) (*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	//创建cartItem
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

//GetCartItemsByCartID 根据购物车的id获取购物车中所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,cart_id from cart_items where cart_id = ?"
	//执行
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		//创建cartItem
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
		if err2 != nil {
			return nil, err2
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}
