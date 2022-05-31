package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/azka-kargo/kargo-trucks/graph/generated"
	"github.com/azka-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SaveTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	truck := &model.Truck{
		ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
		PlateNo: plateNo,
	}
	r.Trucks = append(r.Trucks, truck)
	return truck, nil
}

func (r *mutationResolver) SaveShipment(ctx context.Context, id *string, name string, origin string, destination string, deliveryDate string, truckID *string) (*model.Shipment, error) {
	shipment := &model.Shipment{
		ID:           fmt.Sprintf("TRUCK-%d", len(r.Shipment)+1),
		Name:         name,
		Origin:       origin,
		Destination:  destination,
		DeliveryDate: deliveryDate,
		TruckID:      *id,
	}
	r.Shipment = append(r.Shipment, shipment)
	return shipment, nil
}

func (r *mutationResolver) SendTruckDataToEmail(ctx context.Context, email string) (*model.Truck, error) {
	from := "from#gmail.com"
	password := os.Getenv("EMAILPASSWORD")

	to := []string{
		"sender@example.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("this is a message")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("Email Sent Successfully!")
	return nil, err
}

func (r *queryResolver) PaginatedTrucks(ctx context.Context, id *string, plateNo *string) ([]*model.Truck, error) {
	return r.Trucks, nil
}

func (r *queryResolver) PaginatedShipments(ctx context.Context, page *int, first *int) ([]*model.Shipment, error) {
	return r.Shipment, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
