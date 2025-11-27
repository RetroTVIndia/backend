package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Returns list of all category names
func CategoriesHandler(c *gin.Context) {
	var categoryNames []string
	for _, cat := range EraData.Categories {
		categoryNames = append(categoryNames, cat.Name)
	}
	c.JSON(http.StatusOK, categoryNames)
}

// Returns all shows in a specific category
func CategoryVideosHandler(c *gin.Context) {
	categoryName := c.Query("name")
	if categoryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name required"})
		return
	}

	for _, cat := range EraData.Categories {
		if cat.Name == categoryName {
			c.JSON(http.StatusOK, cat.Videos)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}

// Returns a random video from a category (or any if not provided)
func RandomVideoHandler(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	categoryName := c.Query("category")
	var videos []Video

	if categoryName != "" {
		// pick videos from specific category
		for _, cat := range EraData.Categories {
			if cat.Name == categoryName {
				videos = cat.Videos
				break
			}
		}
	} else {
		// all videos from all categories
		for _, cat := range EraData.Categories {
			videos = append(videos, cat.Videos...)
		}
	}

	if len(videos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No videos found"})
		return
	}

	// pick a random video
	video := videos[rand.Intn(len(videos))]

	// pick a random YouTube URL if multiple
	if len(video.YoutubeURLs) > 0 {
		video.YoutubeURLs = []string{video.YoutubeURLs[rand.Intn(len(video.YoutubeURLs))]}
	}

	c.JSON(http.StatusOK, video)
}
