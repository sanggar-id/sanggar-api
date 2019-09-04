package models

import (
	"gopkg.in/mgo.v2/bson"
	"sanggar-api/app/models/mongodb"
	"time"
)

type Post struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Title     string        `json:"title" bson:"title"`
	Content   string        `json:"content" bson:"content"`
	Slug      string        `json:"slug" bson:"slug"`
	Date      string        `json:"date" bson:"date"`
	Banner    string        `json:"banner" bson:"banner"`
	Category  int           `json:"category" bson:"category"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func newPostCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("posts")
}

// AddPost insert a new Post into database and returns
// last inserted post on success.
func AddPost(m Post) (post Post, err error) {
	c := newPostCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdatePost update a Post into database and returns
// last nil on success.
func (m Post) UpdatePost() error {
	c := newPostCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"title": m.Title, "content": m.Content, "slug": m.Slug, "date": m.Date, "banner": m.Banner, "category": m.Category, "updatedAt": time.Now()},
	})
	return err
}

// DeletePost Delete Post from database and returns
// last nil on success.
func (m Post) DeletePost() error {
	c := newPostCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetPosts Get all Post from database and returns
// list of Post on success
func GetPosts() ([]Post, error) {
	var (
		posts []Post
		err   error
	)

	c := newPostCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&posts)
	return posts, err
}

// GetPost Get a Post from database and returns
// a Post on success
func GetPost(id bson.ObjectId) (Post, error) {
	var (
		post Post
		err  error
	)

	c := newPostCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&post)
	return post, err
}
