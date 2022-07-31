package mysql

const (
	fetchProductQuery = `SELECT id, name, stock, updated_at, created_at FROM product ORDER BY created_at`

	insertProduct = `INSERT product SET name=?, stock=?`

	deleteProduct = `DELETE FROM product WHERE id=?`

	updateProduct = `UPDATE product SET name=?, stock=? WHERE id=?`
)
