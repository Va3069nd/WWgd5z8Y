// 代码生成时间: 2025-10-03 01:55:18
Usage:

1. Install the required package: go get github.com/shirou/gopsutil/cpu
2. Run the program and navigate to http://localhost:8080/cpuusage

Example:

curl http://localhost:8080/cpuusage

The program will return the CPU usage percentage as a JSON object.
*/

package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
    "github.com/shirou/gopsutil/cpu"
)

// CpuUsageHandler handles the CPU usage request and returns the CPU usage percentage.
func CpuUsageHandler(ctx iris.Context) {
    // Get the CPU usage percentage
    usage, err := cpu.Percent(0, false)
    if err != nil {
        ctx.StatusCode(http.StatusInternalServerError)
        ctx.JSON(iris.Map{{"error": fmt.Sprintf("Failed to get CPU usage: %v", err)}})
        return
    }

    // Return the CPU usage percentage as a JSON object
    ctx.JSON(iris.Map{{"cpu_usage": usage[0]}})
}

func main() {
    app := iris.New()

    // Register the CPU usage handler
    app.Get("/cpuusage", CpuUsageHandler)

    // Start the server
    app.Listen(":8080")
}
