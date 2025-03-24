package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mojahige/test-templ/components/blocks"
	"github.com/mojahige/test-templ/components/pages"
	"github.com/mojahige/test-templ/response"
	"github.com/mojahige/test-templ/services"
)

var validate = validator.New()

type PostsHandler struct{}

func NewPostsHandler() *PostsHandler {
	return &PostsHandler{}
}

func (h *PostsHandler) Get(c echo.Context) error {
	posts, err := services.GetAllPosts()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "投稿の取得に失敗しました")
	}

	return response.HTML(c, http.StatusOK, pages.PostsPage(posts))
}

func (h *PostsHandler) Post(c echo.Context) error {
	var requestBody struct {
		Title string `form:"title" validate:"required,min=3,max=100"`
		Body  string `form:"body" validate:"required,min=10"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "無効なリクエスト形式です",
		})
	}

	if err := validate.Struct(requestBody); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, formatValidationError(err))
		}

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "バリデーションエラー",
			"details": errorMessages,
		})
	}

	id, err := services.CreatePost(requestBody.Title, requestBody.Body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "投稿の作成に失敗しました",
		})
	}

	post, err := services.GetPost(int(id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "作成後のデータ取得に失敗しました",
		})
	}

	return response.HTML(c, http.StatusOK, blocks.Post(blocks.PostProps{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}))
}

func (h *PostsHandler) Patch(c echo.Context) error {
	var requestBody struct {
		Id    int    `json:"id" validate:"required"`
		Title string `json:"title" validate:"required,min=3,max=100"`
		Body  string `json:"body" validate:"required,min=10"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "無効なリクエスト形式です",
		})
	}

	if err := validate.Struct(requestBody); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, formatValidationError(err))
		}

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "バリデーションエラー",
			"details": errorMessages,
		})
	}

	err := services.UpdatePost(requestBody.Id, requestBody.Title, requestBody.Body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "投稿の更新に失敗しました",
		})
	}

	post, err := services.GetPost(requestBody.Id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "更新後のデータ取得に失敗しました",
		})
	}

	return response.HTML(c, http.StatusOK, blocks.Post(blocks.PostProps{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}))
}

func (h *PostsHandler) Delete(c echo.Context) error {
	var requestBody struct {
		Id int `json:"id" validate:"required"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "無効なリクエスト形式です",
		})
	}

	if err := validate.Struct(requestBody); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, formatValidationError(err))
		}

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "バリデーションエラー",
			"details": errorMessages,
		})
	}

	err := services.DeletePost(requestBody.Id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "投稿の削除に失敗しました",
		})
	}

	return c.NoContent(204)
}

func formatValidationError(err validator.FieldError) string {
	field := err.Field()

	switch field {
	case "Title":
		field = "タイトル"
	case "Body":
		field = "本文"
	}

	switch err.Tag() {
	case "required":
		return field + "は必須です"
	case "min":
		return field + "は" + err.Param() + "文字以上必要です"
	case "max":
		return field + "は" + err.Param() + "文字以下にしてください"
	default:
		return field + "は" + err.Tag() + "ルールに違反しています"
	}
}
