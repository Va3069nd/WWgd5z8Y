// 代码生成时间: 2025-10-22 00:52:15
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
)

// Product represents a product in the supply chain.
type Product struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

// ProductController handles routes related to products.
type ProductController struct {
    ctx iris.Context
}

// NewProductController creates a new instance of ProductController.
func NewProductController(ctx iris.Context) *ProductController {
    return &ProductController{ctx: ctx}
}

// GetProducts returns a list of products.
func (pc *ProductController) GetProducts() mvc.Result {
    products := []Product{
        {ID: "1", Name: "Widget", Price: 19.99},
        {ID: "2", Name: "Gadget", Price: 9.99},
    }
    return mvc.Json(products)
}

// AddProduct adds a new product to the supply chain.
func (pc *ProductController) AddProduct(product Product) mvc.Result {
    // In a real-world scenario, you would add validation and save the product to a database.
    // For simplicity, we're just printing the product to the console.
    fmt.Printf("Adding product: %+v
", product)
    return mvc.Json(iris.Map{
        "status":  "success",
        "message": "Product added successfully",
    })
}

func main() {
    app := iris.New()
    mvc.New(app).Register(ProductController{},
        mvc.Descriptor[Product]{"Get": "/products", "Post": "/products"},
    )

    // Start the server.
    app.Listen(":8080")
}
