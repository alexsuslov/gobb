package models

import (
	//"database/sql"
	//"database/sql"

	"errors"
	"fmt"
	"github.com/alexsuslov/gobb/api/sqlint"
	"github.com/alexsuslov/gobb/pkg/config"
	"github.com/coopernurse/gorp"
	"math"
	"time"
)

type Post struct {
	Id          int      `db:"id"`
	BoardId     int      `db:"board_id"`
	ParentId    sqlint.NullInt 		`db:"parent_id"`
	Author      *User          `db:"-"`
	AuthorId    int      `db:"author_id"`
	Title       string         `db:"title"`
	Content     string         `db:"content"`
	CreatedOn   time.Time      `db:"created_on"`
	LatestReply time.Time      `db:"latest_reply"`
	LastEdit    time.Time      `db:"last_edit"`
	Sticky      bool           `db:"sticky"`
	Locked      bool           `db:"locked"`
}

// Initializes a new struct, adds some data, and returns the pointer to it
func NewPost(author *User, board *Board, title, content string) *Post {
	post := &Post{
		BoardId:   board.Id,
		AuthorId:  author.Id,
		Title:     title,
		Content:   content,
		CreatedOn: time.Now(),
		Sticky:    false,
	}

	return post
}

func GetPost(id int) (*Post, error) {
	db := GetDbSession()
	obj, err := db.Get(&Post{}, id)
	if obj == nil {
		return nil, err
	}

	return obj.(*Post), err
}

// Returns a pointer to the OP and a slice of post pointers for the given
// page number in the thread.
func GetThread(parent_id, page_id int) (error, *Post, []*Post) {
	db := GetDbSession()

	op, err := db.Get(Post{}, parent_id)
	if err != nil || op == nil {
		fmt.Printf("Something weird is going on here: parent_id: %d, page_id: %d", parent_id, page_id)
		return errors.New(fmt.Sprintf("[error] Could not get parent (%d)", parent_id)), nil, nil
	}

	posts_per_page, err := config.Config.GetInt64("gobb", "posts_per_page")
	if err != nil {
		posts_per_page = 15
	}

	i_begin := (int64(page_id) * (posts_per_page)) - 1
	// The first page already has the OP, which isn't included
	if page_id == 0 {
		posts_per_page -= 1
		i_begin += 1
	}

	var child_posts []*Post
	db.Select(&child_posts, "SELECT * FROM posts WHERE parent_id=$1 ORDER BY created_on ASC LIMIT $2 OFFSET $3", parent_id, posts_per_page, i_begin)

	return nil, op.(*Post), child_posts
}

// Returns the number of posts (on every board/thread)
func GetPostCount() (int64, error) {
	db := GetDbSession()

	count, err := db.SelectInt("SELECT COUNT(*) FROM posts")
	if err != nil {
		fmt.Printf("[error] Error selecting post count (%s)\n", err.Error())
		return 0, errors.New("Database error: " + err.Error())
	}

	return count, nil
}

// Post-SELECT hook for gorp which adds a pointer to the author
// to the Post's struct
func (post *Post) PostGet(s gorp.SqlExecutor) error {
	db := GetDbSession()
	user, _ := db.Get(User{}, post.AuthorId)

	if user == nil {
		return errors.New("Could not find post's author")
	}

	post.Author = user.(*User)

	return nil
}

// Ensures that a post is valid
func (post *Post) Validate() error {
	if post.BoardId == -1 {
		return errors.New("Board does not exist")
	}

	if len(post.Content) <= 3 {
		return errors.New("Post must be longer than three characters")
	}

	//if post.ParentId && len(post.Title) <= 3 {
	if len(post.Title) <= 3 {
		return errors.New("Post title must be longer than three characters")
	}

	return nil
}

// This is used primarily for threads. It will find the latest
// post in a thread, allowing for things like "last post was 10
// minutes ago.
func (post *Post) GetLatestPost() *Post {
	db := GetDbSession()
	latest := &Post{}

	db.SelectOne(latest, "SELECT * FROM posts WHERE parent_id=$1 ORDER BY created_on DESC LIMIT 1", post.Id)

	return latest
}

// Returns the number of pages contained by a thread. This won't work on
// post structs that have ParentIds.
func (post *Post) GetPagesInThread() int {
	db := GetDbSession()
	count, err := db.SelectInt("SELECT COUNT(*) FROM posts WHERE parent_id=$1", post.Id)

	if err != nil {
		fmt.Printf("[error] Could not get post count (%s)\n", err.Error())
	}

	posts_per_page, err := config.Config.GetInt64("gobb", "posts_per_page")

	if err != nil {
		posts_per_page = 15
	}

	if count == posts_per_page {
		return 1
	}

	return int(math.Floor(float64(count) / float64(posts_per_page)))
}

// This function tells us which page this particular post is in
// within a thread based on the current value of posts_per_page
func (post *Post) GetPageInThread() int {
	posts_per_page, err := config.Config.GetInt64("gobb", "posts_per_page")
	if err != nil {
		posts_per_page = 15
	}

	db := GetDbSession()
	n, err := db.SelectInt(`
        WITH thread AS (
                SELECT posts.*,
                ROW_NUMBER() OVER(ORDER BY posts.id) AS position
                FROM posts WHERE parent_id=$1)
        SELECT 
            posts.position
        FROM 
            thread posts
        WHERE 
            posts.id=$2 AND 
            posts.parent_id=$1;
    `, post.ParentId, post.Id)

	return int(math.Floor(float64(n) / float64(posts_per_page)))
}

// Used when deleting a thread. This deletes all posts who are
// children of the OP.
func (post *Post) DeleteAllChildren() error {
	db := GetDbSession()

	_, err := db.Exec("DELETE FROM posts WHERE parent_id=$1", post.Id)
	return err
}

// Get the thread id for a post
func (post *Post) GetThreadId() int {
	if post.ParentId.Valid {
		return post.ParentId.Int
	} else {
		return post.Id
	}
}

// Generate a link to a post
func (post *Post) GetLink() string {
	return fmt.Sprintf("/board/%d/%d?page=%d#post_%d", post.BoardId, post.GetThreadId(), post.GetPageInThread(), post.Id)
}