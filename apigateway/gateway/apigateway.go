package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

const podcastServiceURL = "https://601f1754b5a0e9001706a292.mockapi.io"

type Podcast struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CategoryName string `json:"categoryName"`
}

func (server *Server) getAllPodcasts(ctx *gin.Context) {

	searchQuery := ctx.Query("search")
	titleQuery := ctx.Query("title")
	categoryNameQuery := ctx.Query("categoryName")
	pageQuery := ctx.Query("page")
	limitQuery := ctx.Query("limit")

	// Build the URL with the provided query parameters
	url := podcastServiceURL + "/podcasts"
	if searchQuery != "" {
		url += fmt.Sprintf("?search=%s", searchQuery)
	} else if titleQuery != "" {
		url += fmt.Sprintf("?title=%s", titleQuery)
	} else if categoryNameQuery != "" {
		url += fmt.Sprintf("?categoryName=%s", categoryNameQuery)
	}

	// Convert page and limit to integers with default values
	page, _ := strconv.Atoi(pageQuery)
	limit, _ := strconv.Atoi(limitQuery)
	// Add pagination parameters to the URL
	if page > 0 {
		url += fmt.Sprintf("&page=%d", page)
	}
	if limit > 0 {
		url += fmt.Sprintf("&limit=%d", limit)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error while calling podcast service : ", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var podcasts []Podcast
	err = json.Unmarshal(body, &podcasts)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, podcasts)
}
