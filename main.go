package main

import (
    "database/sql"
    "fmt"
    "github.com/alessandroprudencio/Go-Hexagonal/application"
    db2 "github.com/alessandroprudencio/Go-Hexagonal/adapters/db"
    _ "github.com/mattn/go-sqlite3"
)


func main() {
    db, _ := sql.Open("sqlite3", "db.sqlite")

    productDbAdapter := db2.NewProductDb(db)

    productService := application.NewProductService(productDbAdapter)

    product, _ := productService.Create("Product Example", 34.00)

    productService.Enable(product)

    fmt.Println(product)

}
