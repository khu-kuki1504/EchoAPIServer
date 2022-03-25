package v2

import (
	u "GoBeegoProject/apiHelpers"
	v2s "GoBeegoProject/services/api/v2"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//UserList function will give you the list of users
func UserList(c *gin.Context) {
	var userService v2s.UserService

	//decode the request body into struct and failed if any error occur
	err := json.NewDecoder(c.Request.Body).Decode(&userService.User)
	if err != nil {
		u.Respond(c.Writer, u.Message(1, "Invalid request"))
		return
	}

	//call service
	resp := userService.UserList()

	//return response using api helper
	u.Respond(c.Writer, resp)

}
