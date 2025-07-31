package handler

import (
	"net/http"
	"strconv"
	"vision-service/entity"
	"vision-service/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostHandler struct {
	usecase usecase.PostUsecase
}

func NewPostHandler(u usecase.PostUsecase) *PostHandler {
	return &PostHandler{u}
}

var validate = validator.New()

func (h *PostHandler) CreateArticle(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(post); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				if err.Tag() == "required" {
					errors["Title"] = "Title wajib diisi"
				} else if err.Tag() == "min" {
					errors["Title"] = "Minimal 20 karakter"
				}
			case "Content":
				if err.Tag() == "required" {
					errors["Content"] = "Content wajib diisi"
				} else if err.Tag() == "min" {
					errors["Content"] = "Minimal 200 karakter"
				}
			case "Category":
				if err.Tag() == "required" {
					errors["Category"] = "Category wajib diisi"
				} else if err.Tag() == "min" {
					errors["Category"] = "Minimal 3 karakter"
				}
			case "Status":
				if err.Tag() == "required" {
					errors["Status"] = "Status wajib diisi"
				} else if err.Tag() == "oneof" {
					errors["Status"] = "Hanya boleh bernilai: publish, draft, atau thrash"
				}
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": errors})
		return
	}

	if err := h.usecase.CreateArticle(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article berhasil dibuat", "data": post})
}

func (h *PostHandler) GetAllArticle(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err1 := strconv.Atoi(limitStr)
	offset, err2 := strconv.Atoi(offsetStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit dan offset harus berupa angka"})
		return
	}

	posts, total, err := h.usecase.GetAllArticle(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   posts,
		"limit":  limit,
		"offset": offset,
		"total":  total,
	})
}

func (h *PostHandler) GetArticleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := h.usecase.GetArticleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) UpdateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updateData entity.Post
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updateData.ID = uint(id)

	if err := validate.Struct(updateData); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				if err.Tag() == "required" {
					errors["Title"] = "Title wajib diisi"
				} else if err.Tag() == "min" {
					errors["Title"] = "Minimal 20 karakter"
				}
			case "Content":
				if err.Tag() == "required" {
					errors["Content"] = "Content wajib diisi"
				} else if err.Tag() == "min" {
					errors["Content"] = "Minimal 200 karakter"
				}
			case "Category":
				if err.Tag() == "required" {
					errors["Category"] = "Category wajib diisi"
				} else if err.Tag() == "min" {
					errors["Category"] = "Minimal 3 karakter"
				}
			case "Status":
				if err.Tag() == "required" {
					errors["Status"] = "Status wajib diisi"
				} else if err.Tag() == "oneof" {
					errors["Status"] = "Hanya boleh bernilai: publish, draft, atau thrash"
				}
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": errors})
		return
	}

	if err := h.usecase.UpdateArticle(&updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article berhasil diupdate", "data": updateData})
}

func (h *PostHandler) DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.usecase.DeleteArticle(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus Article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article berhasil dihapus"})
}
