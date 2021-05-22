package db_test

import (
    "database/sql"
    "log"
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/alessandroprudencio/Go-Hexagonal/adapters/db"
    "github.com/alessandroprudencio/Go-Hexagonal/application"
)

var Db *sql.DB

func setUp() {
    Db, _= sql.Open("sqlite3", ":memory:")
    createTable(Db)
    createProduct(Db)
}

func createTable(db *sql.DB) {
    table := `CREATE TABLE products (
        id VARCHAR PRIMARY KEY UNIQUE,
        name VARCHAR,
        price FLOAT,
        status VARCHAR
    );`
    stmt, err := db.Prepare(table)

    if err != nil {
        log.Fatal(err.Error())
    }

    stmt.Exec()
}

func createProduct(db *sql.DB) {
    insert := `
        INSERT INTO products VALUES (
        "e35a4239-65ea-4a30-8e82-2de27e89dfb0",
        "Product 1",
        100,
        "disabled"
    );`
    stmt, err := db.Prepare(insert)

    if err != nil {
        log.Fatal(err.Error())
    }

    stmt.Exec()
}

func TestProductDb_Get(t *testing.T){
    setUp()
    defer Db.Close()
    productDb := db.NewProductDb(Db)
    product, err := productDb.Get("e35a4239-65ea-4a30-8e82-2de27e89dfb0")
    require.Nil(t, err)

    require.Equal(t, "Product 1", product.GetName())
    require.Equal(t, 100.00, product.GetPrice())
    require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T){
    setUp()
    defer Db.Close()

    productDb := db.NewProductDb(Db)

    product := application.NewProduct()
    product.Name = "Product Test"
    product.Price = 25

    productResult, err := productDb.Save(product)
    require.Nil(t, err)
    require.Equal(t,product.Name, productResult.GetName())
    require.Equal(t,product.Price, productResult.GetPrice())
    require.Equal(t,product.Status, productResult.GetStatus())

    product.Status = "enabled"

    productResult, err = productDb.Save(product)
    require.Nil(t, err)
    require.Equal(t,product.Status, productResult.GetStatus())
}
