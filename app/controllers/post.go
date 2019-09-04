package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"sanggar-api/app/models"
)

type PostController struct {
	*revel.Controller
}

func (c PostController) Index() revel.Result {
	var (
		posts []models.Post
		err   error
	)
	posts, err = models.GetPosts()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 200
	return c.RenderJson(posts)
}

func (c PostController) Show(id string) revel.Result {
	var (
		post   models.Post
		err    error
		postID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	postID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	post, err = models.GetPost(postID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}

	c.Response.Status = 200
	return c.RenderJson(post)
}

func (c PostController) Create() revel.Result {
	var (
		post models.Post
		err  error
	)

	err = c.Params.BindJSON(&post)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJson(errResp)
	}

	post, err = models.AddPost(post)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 201
	return c.RenderJson(post)
}

func (c PostController) Update() revel.Result {
	var (
		post models.Post
		err  error
	)
	err = c.Params.BindJSON(&post)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	err = post.UpdatePost()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	return c.RenderJson(post)
}

func (c PostController) Delete(id string) revel.Result {
	var (
		err    error
		post   models.Post
		postID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	postID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid post id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	post, err = models.GetPost(postID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	err = post.DeletePost()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 204
	return c.RenderJson(nil)
}
