package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Returns list of all category names with total YouTube URL links count
func CategoriesHandler(c *gin.Context) {
	type CategoryInfo struct {
		Name            string `json:"name"`
		YoutubeURLLinks int    `json:"youtube_url_links"`
	}
	var categories []CategoryInfo
	for _, cat := range EraData.Categories {
		totalLinks := 0
		for _, video := range cat.Videos {
			totalLinks += len(video.YoutubeURLs)
		}
		categories = append(categories, CategoryInfo{
			Name:            cat.Name,
			YoutubeURLLinks: totalLinks,
		})
	}
	c.JSON(http.StatusOK, categories)
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
	categoryNames := c.QueryArray("category")
	type videoWithCategory struct {
		Video        Video
		CategoryName string
	}
	var videos []videoWithCategory

	if len(categoryNames) > 0 {
		// pick videos from specific categories
		categorySet := make(map[string]struct{}, len(categoryNames))
		for _, name := range categoryNames {
			categorySet[name] = struct{}{}
		}

		for _, cat := range EraData.Categories {
			if _, ok := categorySet[cat.Name]; ok {
				for _, video := range cat.Videos {
					videos = append(videos, videoWithCategory{
						Video:        video,
						CategoryName: cat.Name,
					})
				}
			}
		}
	} else {
		// all videos from all categories
		for _, cat := range EraData.Categories {
			for _, video := range cat.Videos {
				videos = append(videos, videoWithCategory{
					Video:        video,
					CategoryName: cat.Name,
				})
			}
		}
	}

	if len(videos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No videos found"})
		return
	}

	// pick a random video
	selected := videos[rand.Intn(len(videos))]
	video := selected.Video
	categoryName := selected.CategoryName

	// pick a random YouTube URL if multiple
	if len(video.YoutubeURLs) > 0 {
		video.YoutubeURLs = []string{video.YoutubeURLs[rand.Intn(len(video.YoutubeURLs))]}
	}

	// Add category name to the response
	response := gin.H{
		"title":         video.Title,
		"years":         video.Years,
		"youtube_urls":  video.YoutubeURLs,
		"category_name": categoryName,
	}

	c.JSON(http.StatusOK, response)
}
