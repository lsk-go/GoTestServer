package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {}

func ListContainers(c *gin.Context) {
	APIURL := "http://" + c.Param("serverIP") + ":" + c.Param("port") + "/containers/json"
	fmt.Println(APIURL)
	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(body))

	var responseMap []map[string]interface{}
	err = json.Unmarshal(body, &responseMap)

	c.JSON(http.StatusOK, gin.H{
		"message": responseMap,
	})
}
