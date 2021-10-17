package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Post struct {
	UserId   int    `json:"data"`
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Comments string `json:"comments"`
}

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func getPosts() []Post {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var posts []Post
	json.Unmarshal(body, &posts)

	return posts
}

func getComments() []Comment {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var comments []Comment
	json.Unmarshal(body, &comments)
	return comments
}

func addCommentstoPosts() {
	posts := getPosts()
	comments := getComments()
	for idx, p := range posts {
		for _, comment := range comments {
			if comment.PostId == p.Id {
				posts[idx].Comments = strings.Replace(comment.Body, "\n", "|", -1)
			}
		}
	}
	printPosts(posts)
}

func printPosts(posts []Post) {

	toCsv := func(posts []Post) {
		csvFile, err := os.Create("./posts.csv")
		if err != nil {
			fmt.Println(err)
		}
		defer csvFile.Close()

		writer := csv.NewWriter(csvFile)

		headerRow := []string{"Id", "Title", "Body", "Comments"}
		writer.Write(headerRow)
		for _, post := range posts {
			var row []string
			row = append(row, fmt.Sprint(post.Id))
			row = append(row, post.Title)
			row = append(row, post.Body)
			row = append(row, post.Comments)
			writer.Write(row)
		}
		// remember to flush!
		writer.Flush()
	}

	toCsv(posts)
}
func main() {
	addCommentstoPosts()

}
