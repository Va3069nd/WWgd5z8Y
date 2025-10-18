// 代码生成时间: 2025-10-18 19:40:57
package main

import (
    "fmt"
    "time"
    "math/rand"
    "math"

    "github.com/kataras/iris/v12"
)

// Trade represents a single trade with its details.
type Trade struct {
    ID         string    `json:"id"`
    Symbol     string    `json:"symbol"`
    Price      float64   `json:"price"`
    Quantity   int       `json:"quantity"`
    Timestamp time.Time `json:"timestamp"`
}

// NewTrade creates a new trade with a random price and quantity.
func NewTrade(symbol string) Trade {
    price := rand.Float64() * 100.0 // Random price between 0 and 100
    quantity := rand.Intn(100) + 1       // Random quantity between 1 and 100
    return Trade{
        ID:        fmt.Sprintf("%d", rand.Int63()),
        Symbol:    symbol,
        Price:     price,
        Quantity:  quantity,
        Timestamp: time.Now(),
    }
}

// TradeHandler handles trade requests.
func TradeHandler(ctx iris.Context) {
    symbol := ctx.URLParam("symbol")
    if symbol == "" {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": "Symbol parameter is missing",
        })
        return
    }

    trade := NewTrade(symbol)
    ctx.JSON(iris.Map{
        "trade": trade,
    })
}

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html"))

    // Define route for trading.
    app.Get("/trade/{symbol}", TradeHandler)

    // Start the server.
    app.Listen(":8080")
}
