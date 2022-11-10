package products

import (
	"context"
	"database/sql"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bootcamp-go/go-bDs/internal/domain"
	"github.com/bootcamp-go/go-bDs/pkg/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// func TestGetByName2(t *testing.T) {
// 	product := domain.Product{
// 		Name: "test",
// 	}
// 	db, _, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	repository := NewRepository(db)

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	productResult, err := repository.GetByName(ctx, "test")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	assert.Equal(t, product.Name, productResult.Name)
// }

func TestGetByName(t *testing.T) {
	product := domain.Product{
		ID:    2,
		Name:  "Colchón",
		Type:  "Juegos de Sommier y Colchón",
		Count: 1,
		Price: 65_447,
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "name", "type", "count", "price"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(product.ID, product.Name, product.Type, product.Count, product.Price)
	mock.ExpectQuery(regexp.QuoteMeta(GET_BY_NAME)).WithArgs(product.Name).WillReturnRows(rows)

	repository := NewRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	productReturned, err := repository.GetByName(ctx, product.Name)

	assert.NoError(t, err)
	assert.Equal(t, product.Name, productReturned.Name)
	assert.Equal(t, product.ID, productReturned.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStore_txdb(t *testing.T) { // Realizamos mock sobre la transacción de la base de datos
	productStub := domain.Product{
		Name:  "Alexa",
		Type:  "Zapatillas",
		Count: 1,
		Price: 17_999,
	}

	os.Setenv("DBNAME", "storage")
	os.Setenv("DBPASS", "Meli_Sprint#123")
	os.Setenv("DBUSER", "meli_sprint_user")
	_, db, err := db.ConnectMockDatabase()
	assert.NoError(t, err)

	repository := NewRepository(db)

	ctx := context.TODO()

	productIDReturned, err := repository.Store(ctx, productStub)
	assert.NoError(t, err)
	productStub.ID = productIDReturned

	productReturned, err := repository.Get(ctx, productIDReturned)

	assert.NoError(t, err)
	assert.Equal(t, productStub, productReturned)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	id := 1

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE))
	mock.ExpectExec(regexp.QuoteMeta(DELETE)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewRepository(db)

	ctx := context.TODO()

	err = repository.Delete(ctx, id)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery(regexp.QuoteMeta(GET_BY_ID)).WillReturnError(sql.ErrNoRows)
	_, err = repository.Get(ctx, id)

	assert.ErrorContains(t, sql.ErrNoRows, err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}
