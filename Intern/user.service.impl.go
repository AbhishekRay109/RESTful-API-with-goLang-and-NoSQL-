package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *User) error { //calling the create function to create data
	_, err := u.usercollection.InsertOne(u.ctx, user) //Insertion of Data serially
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*User, error) {
	var user *User
	query := bson.D{bson.E{Key: "name", Value: name}} // Finding the data based on the user_name
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*User, error) {
	var users []*User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}}) // retrieving all the data in bson
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *User) error {
	filter := bson.D{primitive.E{Key: "name", Value: user.Name}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "id", Value: user.ID}, primitive.E{Key: "user_name", Value: user.Name}, primitive.E{Key: "user_age", Value: user.Age}, primitive.E{Key: "user_address", Value: user.Address}, primitive.E{Key: "user_desc", Value: user.Desc}, primitive.E{Key: "user_created", Value: user.CreatedAt}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update) // replacing the old data with the new data
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}} // Deleting the whole entry having the perticular user_name

	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
