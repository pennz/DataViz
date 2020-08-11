package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pennz/DataViz/viz"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Static("/Viz", "GraphVizOnline")
	r.POST("/compile", compileHandler)
	r.POST("/compile_debug", compileHandler_debug)

	//r.LoadHTMLGlob("GraphVizOnline/*.html")
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})
	return r
}

func main() {
	r := setupRouter()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r.Run(":" + port)
}

func readCloser2String(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	newStr := buf.String()
	return newStr
}

// runGoAsync(goSRCcode).then((d) => {
//     if (d.Events == null) {
//         show_status("compiling failed: " + d.Errors, 5500);
//     }
//     for (let i = d.Events.length - 1; i >= 1; i--) { // reverse
//         setGraphSteps(d.Events[i].Message, goSRCcode);
//     }
//     let ss0 = splitGraph(d.Events[0].Message)
//     graphSRC.setValue(ss0[0]);
//     for (let i = ss0.length; i >= 0; i--) { // reverse
//         let s = ss0[i];
//         updateStateHistory(
//             encodeURIComponent(s),
//             encodeURIComponent(goSRCcode));
//     }
// });
func readCloser2SVG(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	log.Println("Converting dot to SVG...")
	json_dot := buf.String()
	log.Println("json: ", json_dot)
	var result map[string]interface{}
	json.Unmarshal(buf, &result)
	log.Println("json parsed:", result)
	newStr := new(bytes.Buffer)
	newStr.WriteString(viz.Dot2SVG(result['Events'][0]['Message'][0]))
	log.Println("SVG parsed:")
	return newStr.String()
}

func read2buf(rc io.ReadCloser) *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	return buf
}

func compileHandler_debug(c *gin.Context) {
	version := c.PostForm("version")
	body := c.PostForm("body")
	withVet := c.PostForm("withVet")
	log.Println(version, body, withVet)
}

// compileHandler will relay the request to play.golang.org
func compileHandler(c *gin.Context) {
	//log.Printf("%v\n", readCloser2String(c.Request.Body))
	// https://github.com/gin-gonic/gin#try-to-bind-body-into-different-structs
	// The normal methods for binding request body consumes c.Request.Body and
	// they cannot be called multiple times.
	//buf := read2buf(c.Request.Body)
	// we can change the body in the go
	body := c.PostForm("body")
	s := fmt.Sprintf("version=%d&body=%s&withVet=%s", 2, url.QueryEscape(body), "true")
	//log.Println(s)
	buf := bytes.NewBufferString(s)
	var relay io.Reader = bytes.NewReader(buf.Bytes())
	response, err := http.Post("https://play.golang.org/compile", "application/x-www-form-urlencoded; charset=UTF-8", relay)
	if err == nil {
		if response.StatusCode == 200 && response.Body != nil {
			c.String(response.StatusCode, readCloser2SVG(response.Body))
		} else {
			c.String(response.StatusCode, "Error or cannot get response from play.golang.org.")
		}
	} else {
		c.String(404, "Cannot access play.golang.org.")
	}
}
