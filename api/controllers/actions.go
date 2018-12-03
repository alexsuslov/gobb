package controllers

import (
	"fmt"
	"github.com/alexsuslov/gobb/api/models"
	"github.com/alexsuslov/gobb/api/utils"
	"github.com/lib/pq"
	"net/http"
	"strconv"
	"time"
)

func ActionMarkAllRead(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if user == nil {
		http.NotFound(w, r)
		return
	}

	db := models.GetDbSession()
	user.LastUnreadAll = pq.NullTime{Time: time.Now(), Valid: true}
	db.Update(user)

	http.Redirect(w, r, "/", http.StatusFound)
}

func ActionStickThread(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if !user.CanModerate() {
		http.NotFound(w, r)
		return
	}

	thread_id_str := r.FormValue("post_id")
	thread_id, err := strconv.Atoi(thread_id_str)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	db := models.GetDbSession()
	obj, err := db.Get(&models.Post{}, thread_id)
	thread := obj.(*models.Post)

	if thread == nil || err != nil {
		http.NotFound(w, r)
		return
	}

	thread.Sticky = !(thread.Sticky)
	db.Update(thread)

	http.Redirect(w, r, fmt.Sprintf("/board/%d/%d", thread.BoardId, thread.Id), http.StatusFound)
}

func ActionLockThread(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	if !user.CanModerate() {
		http.NotFound(w, r)
		return
	}

	thread_id_str := r.FormValue("post_id")
	thread_id, err := strconv.Atoi(thread_id_str)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	db := models.GetDbSession()
	obj, err := db.Get(&models.Post{}, thread_id)
	thread := obj.(*models.Post)

	if thread == nil || err != nil {
		http.NotFound(w, r)
		return
	}

	thread.Locked = !(thread.Locked)
	db.Update(thread)

	http.Redirect(w, r, fmt.Sprintf("/board/%d/%d", thread.BoardId, thread.Id), http.StatusFound)
}

func ActionDeleteThread(w http.ResponseWriter, r *http.Request) {
	user := utils.GetCurrentUser(r)
	thread_id_str := r.FormValue("post_id")
	thread_id, err := strconv.Atoi(thread_id_str)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	db := models.GetDbSession()
	obj, err := db.Get(&models.Post{}, thread_id)
	thread := obj.(*models.Post)

	if thread == nil || err != nil {
		http.NotFound(w, r)
		return
	}

	if (thread.AuthorId != user.Id.String()) && !user.CanModerate() {
		http.NotFound(w, r)
		return
	}

	redirect_board := true
	if thread.ParentId.Valid {
		redirect_board = false
	}

	thread.DeleteAllChildren()
	db.Delete(thread)

	if redirect_board {
		http.Redirect(w, r, fmt.Sprintf("/board/%s", thread.BoardId), http.StatusFound)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/board/%s/%s", thread.BoardId, thread.ParentId.String), http.StatusFound)
	}

}

func ActionMoveThread(w http.ResponseWriter, r *http.Request) {
	current_user := utils.GetCurrentUser(r)
	if current_user == nil || !current_user.CanModerate() {
		http.NotFound(w, r)
		return
	}

	thread_id_str := r.FormValue("post_id")
	board_id_str := r.FormValue("to")

	op, err := models.GetPost(thread_id_str)
	boards, _ := models.GetBoards()

	if op == nil || err != nil {
		http.NotFound(w, r)
		return
	}

	if board_id_str != "" {
		db := models.GetDbSession()
		new_board, _ := models.GetBoard(board_id_str)
		if new_board == nil {
			http.NotFound(w, r)
			return
		}

		_, err := db.Exec("UPDATE posts SET board_id=$1 WHERE parent_id=$2", new_board.Id, op.Id)
		op.BoardId = new_board.Id.String()
		db.Update(op)
		if err != nil {
			http.NotFound(w, r)
			fmt.Printf("Error moving post: %s\n", err.Error())
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/board/%d/%d", op.BoardId, op.Id), http.StatusFound)
	}

	board, err := models.GetBoard(op.BoardId)

	utils.RenderTemplate(w, r, "action_move_thread.html", map[string]interface{}{
		"board":  board,
		"thread": op,
		"boards": boards,
	}, nil)
}
