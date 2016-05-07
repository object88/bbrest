package controllers

import "gopkg.in/mgo.v2"

// Controller is ...
type Controller struct {
	session        *mgo.Session
	databaseName   string
	collectionName string
}

// GetCollection is...
func (c *Controller) GetCollection() *mgo.Collection {
	collection := c.session.DB(c.databaseName).C(c.collectionName)
	return collection
}
