// 代码生成时间: 2025-10-16 21:08:34
package main

import (
    "fmt"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "time"
    "github.com/kataras/iris/v12"
)

// SystemResourceMonitor defines the structure to hold system resource data
type SystemResourceMonitor struct {
    CpuUsage float64 `json:"cpu_usage"`
    MemoryUsage float64 `json:"memory_usage"`
}

func main() {
    // Initialize Iris
    app := iris.Default()

    // Define a route for system resource monitoring
    app.Get("/monitor", func(ctx iris.Context) {
        // Retrieve CPU and memory usage
        c, err := cpu.Percent(0, false)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        m, err := mem.VirtualMemory()
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        // Create a new SystemResourceMonitor instance and populate the data
        sr := SystemResourceMonitor{
            CpuUsage: c[0],
            MemoryUsage: m.UsedPercent,
        }

        // Return the system resource monitor data as JSON
        ctx.JSON(iris.Map{
            "system_resource_monitor": sr,
        })
    })

    // Start the Iris server
    fmt.Println("System Resource Monitor is running on http://localhost:8080")
    err := app.Listen(":8080")
    if err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
