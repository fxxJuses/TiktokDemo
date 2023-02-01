package controller

import (
	"douyin/models"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CommentAction(c *gin.Context) {
	// 评论操作
	// TODO 鉴权操作
	token := c.Query("token")

	usi := service.UserServiceImpl{}
	_, user := usi.GetTableUserByToken(token)

	// 获取 POST 信息
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	commentText := c.Query("comment_text")
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64) // 要删除评论的id 在action=2的时候使用

	csi := service.CommentServiceImpl{}
	var comment models.Comment
	// 判断 action ， action = 1 -> 评论操作 ；  action = 2 -> 删除评论操作(只有发布者和视频拥有者可以删除)
	if action == 1 {
		// 发布评论 ， 一条视频可以有多条 同一个用户发布的评论 ， 所以直接 db.save就可以了
		comment = models.Comment{
			User:        user,
			UserId:      user.Id,
			VideoId:     videoId,
			Content:     commentText,
			CreatedDate: time.Now(),
		}
		err := csi.SaveComment(&comment)
		if err != nil {
			c.JSON(http.StatusOK, models.Response{
				StatusCode: 1,
				StatusMsg:  "评论失败",
			})
			return
		}

	} else {
		// 删除评论
		// 只有两类人能够删除评论，第一类是发布评论的人 ， 第二类是发布视频的人
		vsi := service.VideoServiceImpl{}
		video := vsi.FindVideoByVideoId(videoId)
		videoAuthorId := video.AuthorId

		comment, err := csi.FindCommentByCommentId(commentId)
		if err != nil {
			c.JSON(http.StatusOK, models.Response{
				StatusCode: 1,
				StatusMsg:  "评论不存在",
			})
			return
		}

		// 鉴权 id
		if user.Id == comment.UserId || user.Id == videoAuthorId {
			err := csi.DeletComment(&comment)
			if err != nil {
				c.JSON(http.StatusOK, models.Response{
					StatusCode: 1,
					StatusMsg:  "删除评论失败",
				})
				return
			}
			c.JSON(http.StatusOK, models.CommentActionResponse{
				Response: models.Response{StatusCode: 0},
				Comment:  comment,
			})

		} else {
			c.JSON(http.StatusOK, models.Response{
				StatusCode: 1,
				StatusMsg:  "该用户没有权限删除此评论",
			})
			return
		}
	}

	// 将评论数据上传至评论数据库中后， 还要需要更新video列表的评论数量
	// TODO 这个操作后期可以放到 Feed 流中，因为其他人的操作并不会立马同步，数据还是会由前端的缓存所决定。
	vsi := service.VideoServiceImpl{}
	vsi.UpdateCommentCountByVideoId(videoId)

	c.JSON(http.StatusOK, models.CommentActionResponse{
		Response: models.Response{StatusCode: 0},
		Comment:  comment,
	})
}

func CommentList(c *gin.Context) {
	// 展示评论列表

	// TODO 鉴权

	// 根据 video_id 查询所有相关的评论
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	csi := service.CommentServiceImpl{}
	comments, err := csi.FindAllCommentByVideoId(videoId)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "获取评论列表失败，请稍等",
		})
	}
	c.JSON(http.StatusOK, models.CommentListResponse{
		Response:    models.Response{StatusCode: 0},
		CommentList: comments,
	})
}
