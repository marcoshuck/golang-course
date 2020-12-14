package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

/* this is the data model example for a get location
"result": [
                {
                        "update_id": 524107517, // text message id for editMessage to work
                        "message": {
                                "message_id": 205,
                                "from": {
                                        "id": <<int>>,
                                        "is_bot": false,
                                        "first_name": <<string>>,
                                        "username": <<string>>,
                                        "language_code": "en"
                                },
                                "chat": {
                                        "id": <<int>>,
                                        "first_name": <<string>>,
                                        "username": <<string>>,
                                        "type": "private"
                                },
                                "date": 1607481539, // stored in int
                                "location": {
                                        "latitude": 1.327737,
                                        "longitude": 103.889892
                                }
                        }
                }
		]
*/

// Post struct represents the structure of the data to Post.
type Post struct {
	ChatID    int64     `json:"-" bson:"ChatID,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt,omitempty" bson:"expiresAt"`
	Locations Location  `json:"locations,omitempty" bson:"locations"`
	Status    string    `json:"status,omitempty" bson:"status,omitempty"`
}

// Location is the data associated with the Post struct
type Location struct {
	Lat  float64 `json:"lat" bson:"lat"`
	Lng  float64 `json:"lng" bson:"lng"`
	Name string  `json:"name" bson:"name,omitempty"`
}

// PostRepository sets the methods for manipulation of data
type PostRepository interface {
	GetAll(ctx context.Context) ([]Post, error)
	Create(ctx context.Context, post Post) (*Post, error)
}

// MongoDB ...
type MongoDB struct {
	collection *mongo.Collection
}

func (m *MongoDB) GetAll(ctx context.Context) ([]Post, error) {
	// bson.D{}, pass empty filter to get all the data.
	cur, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// defer after execution of a function until the surrounding function returns.
	// runs cur.Close() process after cur.Next().
	defer cur.Close(ctx)

	// iterate through the cursor and deocode each entry
	var posts []Post
	for cur.Next(ctx) {

		// initializer to store the data
		var result Post

		// decodes the bson.D and maps it to the initializer
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		posts = append(posts, result)
	}

	fmt.Println(posts)
	return posts, nil

}

func (m *MongoDB) Create(ctx context.Context, post Post) (*Post, error) {

	post.CreatedAt = time.Now()                          // logs time of creation
	post.ExpiresAt = time.Now().Add(getExpirationTime()) // adds 15 days from creation
	// Insert post to mongodb
	insertResult, err := m.collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	fmt.Println("Inserted Post: ", insertResult.InsertedID)
	return &post, nil
}

func NewMongoDBPostRepository(collection *mongo.Collection) PostRepository {
	return &MongoDB{collection: collection}
}

func getExpirationTime() time.Duration {
	return time.Hour * 24 * 15
}
