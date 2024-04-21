package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"strings"
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

func getRoute() {
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
					Latitude:  43.271326,
					Longitude: 76.952299,
				},
			},
		},
	}

	// create the destination using a latitude and longitude
	destination := &routespb.Waypoint{
		LocationType: &routespb.Waypoint_Location{
			Location: &routespb.Location{
				LatLng: &latlng.LatLng{
					Latitude:  43.253762,
					Longitude: 76.932571,
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

	// execute rpc
	resp, err := client.ComputeRoutes(ctx, req)

	if err != nil {
		// "rpc error: code = InvalidArgument desc = Request contains an invalid
		// argument" may indicate that your project lacks access to Routes
		log.Fatal(err)
	}
	respString := fmt.Sprintf("%v", resp)
	respString = strings.ReplaceAll(respString, "start_location", "\nstart_location\n")
	respString = strings.ReplaceAll(respString, "end_location", "\nend_location\n")

	fmt.Printf("Response: %v", respString)

}
