package main

import (
	"context"
	proto "guestbook-web/proto"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
)

var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("s3cre7"),
})

func main() {

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/guest/checkin", checkin)
	e.GET("/guest/visit/list", list)
	e.PUT("/guest/checkout", checkout)
	// Server
	e.Logger.Fatal(e.Start(":1323"))

	// conn, err := grpc.Dial("127.0.0.1:5009", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// client := proto.NewVisitServiceClient(conn)
	// res, err := client.Add(context.Background(), &proto.VisitRequest{BuildingId: "B002", UserId: "test@gmail.com"})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// log.Printf("Response from server: %s", res.GetToken())

	// client.Signup(context.Background(), &proto.SignupRequest{Username: "Carl", Email: "carl@gmail.com", Password: "examplestring"})
}

type user struct {
	Token string `json:"token" xml:"name"`
}
type checkinRequest struct {
	BuildingID string `json:"building_id" xml:"building_id"`
	UserID     string `json:"user_id" xml:"userId"`
}

//Checkin
func checkin(c echo.Context) (err error) {
	// Get name and email

	cr := new(checkinRequest)
	if err = c.Bind(cr); err != nil {
		return
	}
	log.Printf("BuildingID: %s", cr.BuildingID)
	conn, err := grpc.Dial("127.0.0.1:5009", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	client := proto.NewVisitServiceClient(conn)
	res, err := client.Add(context.Background(), &proto.VisitRequest{BuildingId: cr.BuildingID, UserId: cr.UserID})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Response from server: %s", res.GetToken())
	// u := controller.GetValue()
	u := &user{
		Token: res.GetToken(),
	}
	return c.JSON(http.StatusOK, u)

}

//Checkin
func list(c echo.Context) (err error) {
	// Get name and email

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return
	}
	conn, err := grpc.Dial("127.0.0.1:5009", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	client := proto.NewVisitServiceClient(conn)
	res, err := client.List(context.Background(), &proto.VisitListRequest{Token: token})
	if err != nil {
		log.Fatal(err.Error())
	}

	// u := controller.GetValue()
	// u := &visitList{
	// 	List: res.GetVisit(),
	// }
	return c.JSON(http.StatusOK, res.GetVisit())

}

//Checkin
func checkout(c echo.Context) (err error) {
	// Get name and email

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return
	}
	conn, err := grpc.Dial("127.0.0.1:5009", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	client := proto.NewVisitServiceClient(conn)
	client.Checkout(context.Background(), &proto.CheckoutRequest{Token: token})

	// u := controller.GetValue()
	// u := &visitList{
	// 	List: res.GetVisit(),
	// }
	return c.NoContent(http.StatusNoContent)

}
