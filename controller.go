package main

import "gopkg.in/mgo.v2"

// Controller is ...
type Controller struct {
	session        *mgo.Session
	collectionName string
}

// GetCollection is...
func (c *Controller) GetCollection() *mgo.Collection {
	collection := c.session.DB(DatabaseName).C(c.collectionName)
	return collection
}
