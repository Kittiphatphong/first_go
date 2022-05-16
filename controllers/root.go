package controllers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type RootController interface {
	Index(ctx *fiber.Ctx) error
	PostData(ctx *fiber.Ctx) error
	GetDataByID(ctx *fiber.Ctx) error
	PostFile(ctx *fiber.Ctx) error
}
type rootController struct {
}
//PostDataStruct struct to describe postdata attributes.
type FileStruct struct {
	File string `json:"file"`
}

type Error struct {
	Error string `json:"error"`
}

// PostFile function
// @Summary Post file
// @Description Post file
// @Tags ROOT
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "File"
// @Success 200 {object} controllers.FileStruct
// @Failure 400 {object} controllers.Error
// @Router /profile/upload [post]
func (r rootController) PostFile(ctx *fiber.Ctx) error {
	multipartForm, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	FILE := multipartForm.File["file"]
	if FILE != nil {
		fileExt := strings.Split(FILE[0].Filename, ".")[1]
		if FILE[0].Size > 100000 {
			return errors.New("File size is too big")
		}
		if fileExt != "jpg" && fileExt != "png" && fileExt != "jpeg" {
			return errors.New("File extension is not supported")
		}
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": FILE[0].Filename,
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "File is empty",
	})
}

//PostDataStruct struct to describe postdata attributes.
type PostDataStruct struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

// GetDataByID funtion
// @Summary Get data by id
// @Description Get data by id
// @Tags ROOT
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} controllers.PostDataStruct
// @Router /{id} [get]
func (r rootController) GetDataByID(ctx *fiber.Ctx) error {
	paramsInt, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"my paramsInt id ": paramsInt,
	})
}

// PostData func for test post data.
// @Summary post data to server.
// @Description post data to server.
// @Tags ROOT
// @Accept json
// @Produce json
// @Param PostDataStruct body PostDataStruct true "PostDataStruct"
// @Success 200 {object} controllers.PostDataStruct
// @Router / [post]
func (r rootController) PostData(ctx *fiber.Ctx) error {
	postDataStruct := PostDataStruct{}
	err := ctx.BodyParser(&postDataStruct)
	if err != nil {
		return err
	}
	fmt.Println("postData Name:", postDataStruct.Name)
	fmt.Println("postData Age:", postDataStruct.Age)
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"name": postDataStruct.Name,
		"age":  postDataStruct.Age,
	})
}

// Index godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags ROOT
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (r rootController) Index(ctx *fiber.Ctx) error {
	err := ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Hello World",
	})
	if err != nil {
		return err
	}
	return nil
}

func NwRootController() RootController {
	return &rootController{}
}
