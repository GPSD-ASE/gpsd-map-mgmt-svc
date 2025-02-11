package router

import (
	"disaster-response-map-api/internal/handlers"
	"disaster-response-map-api/pkg/database"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and routes
func SetupRouter(db *database.Database) *gin.Engine {
	r := gin.Default()

	// Create disaster zone handler (pass db)
	disasterZoneHandler := handlers.NewDisasterZoneHandler(db)

	// Define API routes
	r.POST("/zones", disasterZoneHandler.CreateDisasterZone)
	r.GET("/zones", disasterZoneHandler.GetDisasterZones)
	r.DELETE("/zones/:id", disasterZoneHandler.DeleteDisasterZone)

	return r
}
