package products

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/bootcamp-go/go-bDs/internal/domain"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
	Store(ctx context.Context, product domain.Product) (int, error)
	Get(ctx context.Context, id int) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	DELETE      = "DELETE FROM products WHERE id = ?;"
	GET_ALL     = "SELECT id, name, type, count, price, id_warehouse FROM products;"
	GET_BY_ID   = "SELECT id, name, type, count, price, id_warehouse FROM products WHERE id = ?;"
	GET_BY_NAME = "SELECT id, name, type, count, price, id_warehouse FROM products WHERE name = ?;"
	INSERT      = "INSERT INTO products (name, type, count, price, id_warehouse) VALUES (?, ?, ?, ?, ?);"
)

func (r *repository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	// statement, err := r.db.Prepare(GET_BY_NAME)
	// if err != nil {
	// 	return domain.Product{}, fmt.Errorf("error preparing query - error %v", err)
	// }
	// defer statement.Close()

	row := r.db.QueryRowContext(ctx, GET_BY_NAME, name)

	var productToReturn domain.Product
	if err := /*statement.QueryRowContext(ctx, name)*/ row.Scan(&productToReturn.ID, &productToReturn.Name, &productToReturn.Type, &productToReturn.Count, &productToReturn.Price, &productToReturn.WarehouseId); err != nil {
		log.Println(err.Error())
		return domain.Product{}, fmt.Errorf("there are no records for %s - error %v", name, err)
	}

	return productToReturn, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	rows, err := r.db.Query(GET_ALL)
	if err != nil {
		return []domain.Product{}, err
	}

	for rows.Next() {
		var product domain.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price, &product.WarehouseId)
		if err != nil {
			return []domain.Product{}, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	statement, err := r.db.Prepare(DELETE)
	if err != nil {
		return err
	}
	defer statement.Close()

	sQLResult, err := statement.Exec(id)
	if err != nil {
		return err
	}

	affectedRows, err := sQLResult.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows != 1 {
		return errors.New("error: no rows affected")
	}

	return nil
}

func (r *repository) Store(ctx context.Context, product domain.Product) (int, error) {
	statement, err := r.db.Prepare(INSERT)
	if err != nil {
		return 0, fmt.Errorf("error preparing query - error %v", err)
	}
	defer statement.Close()

	sQLResult, err := statement.ExecContext(ctx, product.Name, product.Type, product.Count, product.Price, product.WarehouseId)
	if err != nil {
		return 0, fmt.Errorf("error executing query - error %v", err)
	}

	id, err := sQLResult.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last id - error %v", err)
	}

	return int(id), nil
}

// func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
// 	var productToReturn domain.Product

// 	statement, err := r.db.Prepare(GET_BY_ID)
// 	if err != nil {
// 		return domain.Product{}, err
// 	}
// 	defer statement.Close()

// 	err = statement.QueryRowContext(ctx, id).Scan(&productToReturn.ID, &productToReturn.Name, &productToReturn.Type, &productToReturn.Count, &productToReturn.Price, &productToReturn.WarehouseId)
// 	if err != nil {
// 		return domain.Product{}, err
// 	}

// 	return productToReturn, nil
// }

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product
	rows, err := r.db.QueryContext(ctx, GET_BY_ID, id)

	if err != nil {
		return product, err
	}
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return product, err
		}
	}
	return product, nil
}
