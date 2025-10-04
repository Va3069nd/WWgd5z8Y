// 代码生成时间: 2025-10-05 01:51:19
package main

import (
    "crypto/sha1"
    "encoding/hex"
    "fmt"
    "io"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
)

// HashCalculator 结构体用于处理哈希值计算
type HashCalculator struct{}

// CalculateSHA1 计算字符串的SHA1哈希值
func (h *HashCalculator) CalculateSHA1(s string) (string, error) {
    hash := sha1.New()
    if _, err := io.WriteString(hash, s); err != nil {
        return "", err
    }
    return hex.EncodeToString(hash.Sum(nil)), nil
}

func main() {
    app := iris.New()
    app.Get("/hash", func(ctx iris.Context) {
        query := ctx.URLParam("text")
        if query == "" {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.WriteString("Please provide a text to hash.")
            return
        }

        calculator := HashCalculator{}
        sha1Hash, err := calculator.CalculateSHA1(query)
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.WriteString("Error calculating hash: " + err.Error())
            return
        }

        ctx.JSON(http.StatusOK, iris.Map{
            "hash": sha1Hash,
        })
    })

    // 启动服务器
    app.Listen(":8080")
}
