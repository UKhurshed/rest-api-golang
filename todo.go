package todo

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoList struct {
	Id primitive.ObjectID
	Title string
	Description string
}


//type TodoItem struct{
//	Id primitive.ObjectID
//	Title
//}