package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type User struct {
// 	Name     string `json:"_name"`
// 	LastName string `json:"_lastName"`
// 	ID       int    `json:"_id"`
// }

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection("movies")
	var title string = "Back to the Future"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

	// r := mux.NewRouter()
	// r.HandleFunc("/", homeHandler)
	// r.HandleFunc("/page", pageHandler)
	// http.Handle("/", r)

	// http.ListenAndServe(":3000", nil)
}

func (u models.User) userf() User {
	return User{Name: "jose", LastName: "calderon", ID: 123}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	u := User{Name: "jose", LastName: "Calderon", ID: 123}

	fmt.Println(u.userf())
	w.Write([]byte("Hello"))
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Another Page"))
}
