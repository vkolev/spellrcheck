package main

import (
	"log"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/trustmaster/go-aspell"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", getHome)
	r.POST("/check", postCheck)
	r.POST("/suggest", postSuggest)

	r.Run("0.0.0.0:8080")
}

func getHome(c *gin.Context) {
	baseUrl := strings.Join([]string{"http://", c.Request.Host}, "")

	routes := make(map[string]string)

	routes["/check"] = strings.Join([]string{baseUrl, "check"}, "/")
	routes["/suggest"] = strings.Join([]string{baseUrl, "suggest"}, "/")

	content := gin.H{
		"base_url": baseUrl,
		"routes": routes,
	}
	c.JSON(200, content)
}

func postCheck(c *gin.Context) {
	lang := c.PostForm("lang")
	text := c.PostForm("text")

	speller, err := aspell.NewSpeller(map[string]string{
		"lang": lang,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer speller.Delete()

	test := strings.Fields(text)

	result := make(map[string][]string)

	for i := range test {
		if speller.Check(stripchars(test[i], ".,!?+=")) != true {
			result[stripchars(test[i], ".,!?+=")] = speller.Suggest(test[i])
		}
	}

	c.JSON(200, gin.H{"lang": lang, "misspelled": result})
}


func postSuggest(c *gin.Context) {
	lang := c.PostForm("lang")
	text := strings.TrimSpace(c.PostForm("word"))

	log.Println(text)

	speller, err := aspell.NewSpeller(map[string]string{
		"lang": lang,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer speller.Delete()
	result := make(map[string][]string)

	result[text] = speller.Suggest(stripchars(text, ".,!?+="))

	c.JSON(200, gin.H{"lang": lang,"result": result})
}


func stripchars(str, chr string) string {
    return strings.Map(func(r rune) rune {
        if strings.IndexRune(chr, r) < 0 {
            return r
        }
        return -1
    }, str)
}
