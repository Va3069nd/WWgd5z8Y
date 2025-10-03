// 代码生成时间: 2025-10-04 03:13:29
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "net/http"
)

// Define a struct to represent a table row with sortable fields
type TableRow struct {
    ID      int    "json:"id" xml:"id" form:"id" query:"id" schema:"id""
    Name    string "json:"name" xml:"name" form:"name" query:"name" schema:"name""
    Balance float64 "json:"balance" xml:"balance" form:"balance" query:"balance" schema:"balance"
}

// rows is a slice of TableRow to demonstrate sorting and filtering
var rows = []TableRow{
    {ID: 1, Name: "John", Balance: 1000},
    {ID: 2, Name: "Jane", Balance: 2000},
    {ID: 3, Name: "Mike", Balance: 1500},
    {ID: 4, Name: "Anna", Balance: 2500},
    {ID: 5, Name: "Chris", Balance: 1200},
}

// sortRows sorts the rows based on the given field and order
func sortRows(rows []TableRow, field string, order string) ([]TableRow, error) {
    switch field {
    case "id":
        switch order {
        case "asc":
            return sort.SliceStable(rows, func(i, j int) bool { return rows[i].ID < rows[j].ID }), nil
        case "desc":
            return sort.SliceStable(rows, func(i, j int) bool { return rows[i].ID > rows[j].ID }), nil
        }
    case "name":
        switch order {
        case "asc":
            return sort.SliceStable(rows, func(i, j int) bool { return rows[i].Name < rows[j].Name }), nil
        case "desc":
            return sort.SliceStable(rows, func(i, j int) bool { return rows[i].Name > rows[j].Name }), nil
        }
    case "balance":
        switch order {
        case "asc":
            return sort.SliceStable(rows, func(i, j int) bool { return rows[i].Balance < rows[j].Balance }), nil
        case "desc":
            return sort.SliceStable(rows, func(i, j int) bool { return rows[i].Balance > rows[j].Balance }), nil
        }
    }
    return nil, fmt.Errorf("invalid sort field: %s", field)
}

// filterRows filters the rows based on the given name filter
func filterRows(rows []TableRow, filter string) []TableRow {
    var filtered []TableRow
    for _, row := range rows {
        if strings.Contains(strings.ToLower(row.Name), strings.ToLower(filter)) {
            filtered = append(filtered, row)
        }
    }
    return filtered
}

func main() {
    app := iris.New()
    app.Get("/table", func(ctx iris.Context) {
        var sortField string
        var sortOrder string
        var filter string
        // Retrieve query parameters
        if err := ctx.ReadQuery/Application/json; err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.JSON(iris.Map{"error": "Failed to read query parameters"})
            return
        }
        ctx.ReadQuery("sortField", &sortField)
        ctx.ReadQuery("sortOrder", &sortOrder)
        ctx.ReadQuery("filter", &filter)

        // Filter the rows if a filter is provided
        if filter != "" {
            rows = filterRows(rows, filter)
        }

        // Sort the rows if sort parameters are provided
        if sortField != "" && sortOrder != "" {
            rows, _ = sortRows(rows, sortField, sortOrder)
        }

        // Return the sorted and/or filtered rows as JSON
        ctx.JSON(rows)
    })

    // Start the Iris web server
    if err := app.Run(iris.Addr(":8080")); err != nil {
        iris.StatusInternalServerError
        fmt.Printf("Error starting server: %s
", err)
    }
}