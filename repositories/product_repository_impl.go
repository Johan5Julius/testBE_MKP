package repositories

import (
	"context"
	"errors"
	"testMKP/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return &ProductRepositoryImpl{
		Collection: db.Collection("products"),
	}
}

func (p *ProductRepositoryImpl) Create(product models.ProductCreateRequest) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newProduct := models.Product{
		ID:    primitive.NewObjectID(),
		Name:  product.Name,
		Price: product.Price,
	}

	_, err := p.Collection.InsertOne(ctx, newProduct)
	if err != nil {
		return models.Product{}, err
	}
	return newProduct, nil
}

func (p *ProductRepositoryImpl) Update(product models.ProductUpdateRequest) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return models.Product{}, errors.New("invalid id format")
	}

	update := bson.M{
		"$set": bson.M{
			"price": product.Price,
			"name":  product.Name,
		},
	}

	result := p.Collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		update)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return models.Product{}, errors.New("product not found")
		}
		return models.Product{}, result.Err()
	}

	var updatedProduct models.Product

	err = p.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&updatedProduct)
	if err != nil {
		return models.Product{}, err
	}
	return updatedProduct, nil
}

func (p *ProductRepositoryImpl) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}

	result, err := p.Collection.DeleteOne(ctx, bson.M{"_id": objectID})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (p *ProductRepositoryImpl) FindAll() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var products []models.Product

	cursor, err := p.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if products == nil {
		products = []models.Product{}
	}
	return products, nil

}

func (p *ProductRepositoryImpl) FindById(id string) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Product{}, errors.New("invalid id format")
	}

	var product models.Product
	err = p.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Product{}, errors.New("product not found")
		}
		return models.Product{}, err
	}

	return product, nil
}
