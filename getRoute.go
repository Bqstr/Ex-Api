package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"

	routespb "google.golang.org/genproto/googleapis/maps/routing/v2"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	fieldMask  = "routes.duration,routes.distanceMeters,routes.legs"
	apiKey     = "AIzaSyCQ3aLkzejyUJsYPbfUP8DowJdOsy6VOgg"
	serverAddr = "routes.googleapis.com:443"
)

func getRoute(context_gin *gin.Context) {

	start_latitudeStr := context_gin.Query("start_latitude")
	start_longitudeStr := context_gin.Query("start_longitude")
	end_latitudeStr := context_gin.Query("end_latitude")
	end_longitudeStr := context_gin.Query("end_longitude")

	// Convert latitude and longitude to float64
	start_latitude, err := strconv.ParseFloat(start_latitudeStr, 64)
	if err != nil {
		context_gin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
		return
	}
	start_longitude, err := strconv.ParseFloat(start_longitudeStr, 64)
	if err != nil {
		context_gin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
		return
	}

	end_latitude, err := strconv.ParseFloat(end_latitudeStr, 64)
	if err != nil {
		context_gin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
		return
	}
	end_longitude, err := strconv.ParseFloat(end_longitudeStr, 64)
	if err != nil {
		context_gin.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
		return
	}

	config := tls.Config{}
	conn, err := grpc.Dial(serverAddr,
		grpc.WithTransportCredentials(credentials.NewTLS(&config)))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := routespb.NewRoutesClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Api-Key", apiKey)
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Goog-Fieldmask", fieldMask)
	defer cancel()

	// create the origin using a latitude and longitude
	origin := &routespb.Waypoint{
		LocationType: &routespb.Waypoint_Location{
			Location: &routespb.Location{
				LatLng: &latlng.LatLng{
					Latitude:  start_latitude,
					Longitude: start_longitude,
				},
			},
		},
	}

	// create the destination using a latitude and longitude
	destination := &routespb.Waypoint{
		LocationType: &routespb.Waypoint_Location{
			Location: &routespb.Location{
				LatLng: &latlng.LatLng{
					Latitude:  end_latitude,
					Longitude: end_longitude,
				},
			},
		},
	}

	req := &routespb.ComputeRoutesRequest{
		Origin:                   origin,
		Destination:              destination,
		TravelMode:               routespb.RouteTravelMode_DRIVE,
		RoutingPreference:        routespb.RoutingPreference_TRAFFIC_AWARE,
		ComputeAlternativeRoutes: true,
		Units:                    routespb.Units_METRIC,
		RouteModifiers: &routespb.RouteModifiers{
			AvoidTolls:    false,
			AvoidHighways: true,
			AvoidFerries:  true,
		},
		PolylineQuality: routespb.PolylineQuality_OVERVIEW,
	}

	resp, err := client.ComputeRoutes(ctx, req)

	if err != nil {

		log.Fatal(err)
	}
	respString := fmt.Sprintf("%v", resp)
	//respString = strings.ReplaceAll(respString, "start_location", "\nstart_location\n")
	//respString = strings.ReplaceAll(respString, "end_location", "\nend_location\n")

	context_gin.String(http.StatusOK, respString)

}
