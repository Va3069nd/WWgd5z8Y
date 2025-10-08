// 代码生成时间: 2025-10-09 01:57:22
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"
    "github.com/kataras/iris/v12/middleware/logger"
)

// Product represents a product for live streaming commerce
type Product struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

// ProductsAPI handles product-related operations
type ProductsAPI struct {
    products map[string]Product
}

// NewProductsAPI creates a new instance of ProductsAPI
func NewProductsAPI() *ProductsAPI {
    return &ProductsAPI{
        products: make(map[string]Product),
    }
}

// AddProduct adds a new product to the products map
func (api *ProductsAPI) AddProduct(ctx iris.Context, product Product) error {
    api.products[product.ID] = product
    return ctx.JSON(iris.StatusOK, product)
}

// GetProduct retrieves a product by ID
func (api *ProductsAPI) GetProduct(ctx iris.Context) {
    productID := ctx.Params().Get("id")
    if product, exists := api.products[productID]; exists {
        ctx.JSON(iris.StatusOK, product)
    } else {
        ctx.StatusCode(iris.StatusNotFound)
        ctx.JSON(iris.Map{
            "error": "Product not found",
        })
    }
}

func main() {
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())

    // Register the routes
    productsAPI := NewProductsAPI()
    app.Post("/products", func(ctx iris.Context) {
        var product Product
        if err := ctx.ReadJSON(&product); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        if err := productsAPI.AddProduct(ctx, product); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
    })
    app.Get("/products/{id}", func(ctx iris.Context) {
        productsAPI.GetProduct(ctx)
    })

    // Start the server
    fmt.Println("Server is running on http://localhost:8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        panic(err)
    }
}
