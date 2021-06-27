package product

import (
	"context"
	"github.com/beuus39/product/internal/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type mongoRepository struct {
	client *mongo.Client
	db string
	timeout time.Duration
}

func (m *mongoRepository) Find(id int) <-chan RepositoryResult {
	output := make(chan RepositoryResult)

	go func() {
		defer close(output)
		ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
		defer cancel()

		product := &domain.Product{}
		collection := m.client.Database(m.db).Collection("items")
		filter := bson.M{"id": id}

		err := collection.FindOne(ctx, filter).Decode(product)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				output <- RepositoryResult{Error: errors.New("product not found")}
				return
			}
		}
		output <- RepositoryResult{Result: product}
	}()
	return output
}

func (m *mongoRepository) Save(product *domain.Product) <-chan RepositoryResult {
	output := make(chan RepositoryResult)

	go func() {
		defer close(output)

		ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
		defer cancel()

		collection := m.client.Database(m.db).Collection("items")
		_, err := collection.InsertOne(
			ctx,
			bson.M{
				"id":           product.ID,
				"categoryId":   product.CategoryID,
				"name":         product.Name,
				"description":  product.Description,
				"image":        product.Image,
				"stock":        product.Stock,
				"price":        product.Price,
			})
		if err != nil {
			output <- RepositoryResult{Error: errors.New("Error writing to repository")}
		}
		output <- RepositoryResult{Result: true}
	}()
	return output
}

func (m *mongoRepository) Update(product *domain.Product) <-chan RepositoryResult {
	output := make(chan RepositoryResult)

	go func() {
		defer close(output)
		ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
		defer cancel()

		collection := m.client.Database(m.db).Collection("items")
		_, err := collection.UpdateOne(
			ctx,
			bson.M{"id": product.ID},
			bson.D{
				{"$set", bson.D{
					{ "name", product.Name},
					{ "description", product.Description},
					{ "image", product.Image},
					{ "stock", product.Stock},
					{ "price", product.Price},
				}},
			})
		if err != nil {
			output <- RepositoryResult{Error: errors.New("update failed")}
		}
		output <- RepositoryResult{Result: nil}
	}()
	return output
}

func (m *mongoRepository) FindAll() <-chan RepositoryResult {
	output := make(chan RepositoryResult)

	go func() {
		defer close(output)

		var items []*domain.Product
		collection := m.client.Database(m.db).Collection("items")
		cur, err := collection.Find(context.Background(), bson.D{}.Map())

		if err != nil {
			output <- RepositoryResult{Error: errors.New("Find all failed")}
		}
		defer cur.Close(context.Background())
		for cur.Next(context.TODO()) {
			var item domain.Product
			if err := cur.Decode(&item); err != nil {
				output <- RepositoryResult{Error: errors.New("Error")}
			}
			items = append(items, &item)
		}
		output <- RepositoryResult{Result: items}
	}()
	return output
}

func (m *mongoRepository) Delete(id int) <-chan RepositoryResult {
	output := make(chan RepositoryResult)
	go func() {
		defer close(output)
		ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
		defer cancel()

		filter := bson.M{"id": id}
		collection := m.client.Database(m.db).Collection("items")

		_, err := collection.DeleteOne(ctx, filter)

		if err != nil {
			output <- RepositoryResult{Error: errors.New("Deleting failed")}
		}
		output <- RepositoryResult{Result: true}

	}()
	return output
}

func newMongoClient(mongoServerUrl string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout) *time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoServerUrl))

	if err != nil {
		return nil, err
	}
	return client, nil
}
func NewMongoRepository(mongoServerUrl, mongoDb string, timeout int) (Repository, error) {
	mongoClient, err := newMongoClient(mongoServerUrl, timeout)

	repo := &mongoRepository{
		client: mongoClient,
		db: mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}

	if err != nil {
		return nil, errors.Wrap(err, "Client error")
	}
	return repo, nil
}