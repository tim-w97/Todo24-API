package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/tim-w97/my-awesome-Todo-API/db"
	"github.com/tim-w97/my-awesome-Todo-API/types"
	"github.com/tim-w97/my-awesome-Todo-API/util"
	"github.com/tim-w97/my-awesome-Todo-API/validation"
	"log"
	"net/http"
)

func ShareTodo(context *gin.Context) {
	var sharedTodo types.SharedTodo

	if bindErr := context.BindJSON(&sharedTodo); bindErr != nil {
		context.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"message": "can't convert body to share todo"},
		)

		log.Print(bindErr)
		return
	}

	sharedTodo.TodoID = context.GetInt("todoID")

	sql, err := util.ReadSQLFile("share_todo.sql")

	if err != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't read SQL"},
		)

		log.Print(err)
		return
	}

	result, insertErr := db.Database.Exec(
		sql,
		sharedTodo.TodoID,
		sharedTodo.UserID,
	)

	if insertErr != nil {
		context.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "can't insert shared todo"},
		)

		log.Print(insertErr)
		return
	}

	if ok := validation.ValidateSQLResult(result, context); !ok {
		return
	}

	context.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "shared todo successfully"},
	)
}
