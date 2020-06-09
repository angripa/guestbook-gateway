package controller

import (
	"context"
	proto "guestbook-web/proto"
	"log"

	"google.golang.org/grpc"
)

type visitResponse struct {
	Token string `json:"token"`
}

func getValue() string {
	return "Hello from this another package"
}

//create user
func checkin(buildingID string, userID string) string {
	conn, err := grpc.Dial("127.0.0.1:5009", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	client := proto.NewVisitServiceClient(conn)
	res, err := client.Add(context.Background(), &proto.VisitRequest{BuildingId: buildingID, UserId: userID})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Response from server: %s", res.GetToken())
	return res.GetToken()

}
