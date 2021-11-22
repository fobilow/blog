package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	indexPath = "src/assets/data/posts/index.json"
	postsPath = "src/assets/data/posts"
)

var f = flag.String("f", "", "post function to perform e.g create, delete")
var id = flag.String("id", "", "id of post - used for updating or deleting posts")
var title = flag.String("t", "", "title of post")
var content = flag.String("b", "", "path to post content")

func main() {
	flag.Parse()
	switch *f {
	case "create":
		if len(*id) > 0 {
			if err := updatePost(*id, *title, *content); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("post updated")
			return
		}
		if err := createPost(*title, *content); err != nil {
			fmt.Println("Failed to create post", err.Error())
			os.Exit(1)
		}
		fmt.Println("post created")
	case "delete":
		if err := deletePost(*id); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("post deleted")
	default:
		flag.Usage()
	}
}

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func createPost(title, content string) error {
	// read contents
	data, err := os.ReadFile(content)
	if err != nil {
		return err
	}

	post, err := json.Marshal(Post{Title: title, Body: string(data)})
	if err != nil {
		return err
	}

	// write to src/assets/posts/{timestamp}.json
	postID := time.Now().Unix()
	if err := os.WriteFile(fmt.Sprintf("%s/%d.json", postsPath, postID), post, os.ModePerm); err != nil {
		return err
	}

	return updateIndex(postID, title)
}

func updatePost(idStr, title, content string) error {
	postFile := fmt.Sprintf("%s/%s.json", postsPath, idStr)
	if _, err := os.Stat(postFile); err != nil {
		return err
	}

	// read contents
	data, err := os.ReadFile(content)
	if err != nil {
		return err
	}

	post, err := json.Marshal(Post{Title: title, Body: string(data)})
	if err != nil {
		return err
	}

	// write to src/assets/posts/{timestamp}.json
	if err := os.WriteFile(postFile, post, os.ModePerm); err != nil {
		return err
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}

	return updateIndex(id, title)
}

func deletePost(idStr string) error {
	postFile := fmt.Sprintf("%s/%s.json", postsPath, idStr)
	if _, err := os.Stat(postFile); err != nil {
		return err
	}

	if err := os.Remove(postFile); err != nil {
		return err
	}

	// remove from index
	index, err := os.ReadFile(indexPath)
	if err != nil {
		return err
	}

	var ids map[int64]string
	if err := json.Unmarshal(index, &ids); err != nil {
		return err
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}

	delete(ids, id)

	index, err = json.Marshal(ids)
	if err != nil {
		return err
	}

	return os.WriteFile(indexPath, index, os.ModePerm)
}

func updateIndex(id int64, title string) error {
	// write to index
	ids := map[int64]string{}
	index, err := os.ReadFile(indexPath)
	if err == nil {
		if err := json.Unmarshal(index, &ids); err != nil {
			return err
		}
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	ids[id] = title
	index, err = json.Marshal(ids)
	if err != nil {
		return err
	}

	return os.WriteFile(indexPath, index, os.ModePerm)
}
