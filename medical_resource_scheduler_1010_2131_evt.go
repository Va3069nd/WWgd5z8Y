// 代码生成时间: 2025-10-10 21:31:56
package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
)

// Resource represents a medical resource with its ID and name.
type Resource struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// Scheduler is a struct that contains methods to manage medical resources.
type Scheduler struct {
    resources map[string]Resource
}

// NewScheduler creates a new instance of Scheduler.
func NewScheduler() *Scheduler {
    return &Scheduler{
        resources: make(map[string]Resource),
    }
}

// AddResource adds a new resource to the scheduler.
func (s *Scheduler) AddResource(id string, name string) {
    s.resources[id] = Resource{ID: id, Name: name}
}

// RemoveResource removes a resource from the scheduler by ID.
func (s *Scheduler) RemoveResource(id string) error {
    if _, exists := s.resources[id]; exists {
        delete(s.resources, id)
        return nil
    }
    return fmt.Errorf("resource with id '%s' not found", id)
}

// GetResource retrieves a resource by ID.
func (s *Scheduler) GetResource(id string) (*Resource, error) {
    resource, exists := s.resources[id]
    if !exists {
        return nil, fmt.Errorf("resource with id '%s' not found", id)
    }
    return &resource, nil
}

// StartServer starts the HTTP server with the Iris framework.
func (s *Scheduler) StartServer(port string) {
    app := iris.New()
    app.Get("/resources", func(ctx iris.Context) {
        resources := make([]Resource, 0, len(s.resources))
        for _, resource := range s.resources {
            resources = append(resources, resource)
        }
        ctx.JSON(http.StatusOK, resources)
    })
    
    app.Get("/resources/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        resource, err := s.GetResource(id)
        if err != nil {
            ctx.StatusCode(http.StatusNotFound)
            ctx.Writef("Resource with id '%s' not found", id)
            return
        }
        ctx.JSON(http.StatusOK, resource)
    })
    
    app.Post("/resources", func(ctx iris.Context) {
        var resource Resource
        if err := ctx.ReadJSON(&resource); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.Writef("Error reading JSON: %s", err.Error())
            return
        }
        s.AddResource(resource.ID, resource.Name)
        ctx.JSON(http.StatusCreated, resource)
    })
    
    app.Delete("/resources/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        if err := s.RemoveResource(id); err != nil {
            ctx.StatusCode(http.StatusNotFound)
            ctx.Writef("Resource with id '%s' not found", id)
            return
        }
        ctx.StatusCode(http.StatusOK)
        ctx.Writef("Resource with id '%s' removed", id)
    })
    
    log.Printf("Starting server on :%s...", port)
    if err := app.Listen(":%s", iris.WithOptimizations()); err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
}

func main() {
    scheduler := NewScheduler()
    // Add some initial resources for demonstration purposes.
    scheduler.AddResource("r1", "Ambulance")
    scheduler.AddResource("r2", "Hospital Bed")
    scheduler.AddResource("r3", "Defibrillator")
    
    // Start the HTTP server on port 8080.
    scheduler.StartServer("8080")
}