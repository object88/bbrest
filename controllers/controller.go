package controllers

import "gopkg.in/mgo.v2"

// // Controller is ...
// type Controller interface {
// 	Create(d *dtos.BaseDto) *dtos.BaseDto
// 	Get(id string) *dtos.BaseDto
// }

// BaseController is...
type BaseController struct {
	session        *mgo.Session
	databaseName   string
	collectionName string
}

// GetCollection is...
func (c *BaseController) GetCollection() *mgo.Collection {
	collection := c.session.DB(c.databaseName).C(c.collectionName)
	return collection
}
