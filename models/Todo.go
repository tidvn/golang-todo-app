package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"Title"`
	Content   string             `bson:"Content"`
	Status    int                `bson:"Status"`
	CreatedOn primitive.DateTime `bson:"CreatedOn"`
}

type TodoDB struct {
	collection *mongo.Collection
}

func TodoData() *TodoDB {
	var mongo, _ = Connect()
	var c = mongo.TodoCollection()
	return &TodoDB{
		collection: c,
	}
}

func (u *Todo) Init() {
	u.ID = primitive.NewObjectID()
	u.CreatedOn = primitive.DateTime(time.Now().Unix() * 1000)
	u.Title = "Untitled"
	u.Content = "Empty"
}
func (db *TodoDB) AddTodo(m *Todo) (t *Todo, err error) {
	ctx := context.Background()
	m.Init()

	_, err = db.collection.InsertOne(ctx, m)
	t = m
	if err == nil {
		return t, nil
	}
	return nil, err
}
func (db *TodoDB) GetAll() ([]Todo, error) {
	ctx := context.Background()
	cursor, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var listTodo []Todo
	if err = cursor.All(ctx, &listTodo); err != nil {
		return nil, err
	}

	return listTodo, nil
}

func (db *TodoDB) GetById(idstr string) (v *Todo, err error) {

	idOj, _ := primitive.ObjectIDFromHex(idstr)
	ctx := context.Background()
	err = db.collection.FindOne(ctx, bson.M{"_id": idOj}).Decode(&v)

	if err != nil {
		return nil, err
	}
	return v, nil
}

func (db *TodoDB) UpdateTodo(m *Todo) (Todo, error) {
	ctx := context.Background()
	u, e := db.GetById(m.ID.Hex())
	fmt.Print(u)
	if e == nil {
		if m.Title != "" {
			u.Title = m.Title
		}
		if m.Content != "" {
			u.Content = m.Content
		}
		u.Status = m.Status

	}
	filter := bson.M{"_id": m.ID}
	update := bson.M{"$set": bson.M{
		"Title":   u.Title,
		"Content": u.Content,
		"Status":  u.Status,
	}}
	_, err := db.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return Todo{}, err
	}
	return *u, nil
}

func (db *TodoDB) DeleteTodo(idstr string) (err error) {

	id, _ := primitive.ObjectIDFromHex(idstr)
	ctx := context.Background()
	filter := bson.M{"_id": id}
	_, err = db.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
