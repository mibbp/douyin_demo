package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")                                 //获取点击关注的人的token
	temp_user_id := c.Query("to_user_id")                     //获取被关注的人id
	to_user_id, err := strconv.ParseInt(temp_user_id, 10, 64) //获取被关注的人token
	if err != nil {

	}

	to_user_token := id_to_token[to_user_id] //获取被关注的人token
	to_user := usersLoginInfo[to_user_token] //根据token获取user
	//user 是点关注的那个人 touser是被点关注的人
	if user, exist := usersLoginInfo[token]; exist {
		if token != to_user_token { //增加关注逻辑不能关注自己
			to_user.IsFollow = true
			//点关注的人更新关注列表
			user.FollowList = append(user.FollowList, to_user)
			user.FollowCount++
			//被关注的人更新粉丝列表
			to_user.FansList = append(to_user.FansList, user)
			to_user.FollowerCount++
			//更新user信息
			usersLoginInfo[token] = user
			usersLoginInfo[to_user_token] = to_user
			fmt.Println(user.FollowList)
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "关注成功"})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "关注失败，不能关注自己"})
		}

		//user.FollowList.app(usersLoginInfo[to_user_token])

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	user := usersLoginInfo[token]
	fmt.Println(user.FollowList)
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: user.FollowList,
	})

}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	user := usersLoginInfo[token]
	fmt.Println(user.FollowList)
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: user.FansList,
	})
}
