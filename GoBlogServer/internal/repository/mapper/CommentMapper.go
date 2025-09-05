package mapper

import (
	"GoBlogServer/internal/dtos"
	"GoBlogServer/internal/repository/model"
	"gorm.io/gorm"
	"time"
)

type CommentMapper struct {
	db *gorm.DB
}

// New构造
func NewCommentMapper(db *gorm.DB) *CommentMapper {
	return &CommentMapper{db: db}
}

/*
*
保存
*/
func (c *CommentMapper) SaveNew(optUserId int, postId int64, content string) (int, error) {
	var t = time.Now()
	var bean = model.Comments{Content: content, UserId: optUserId, PostId: postId, CreateTime: &t, State: 1}
	r := c.db.Create(&bean)
	if r.Error != nil {
		return -1, r.Error
	}
	return bean.Id, nil
}

/*
*
删除评论
*/
func (c *CommentMapper) UpdateDisable(id int) (bool, error) {
	r := c.db.Model(&model.Comments{}).Where("id=?").UpdateColumn("state", 0)
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

/*
*
根据文章id分页查询评论列表
*/
func (c *CommentMapper) GetListByPosts(postId int64, pageOffset int, pageSize int) (*dtos.BasePageDto[dtos.PostCommentDto], error) {
	varSqlCount := "select bc.id,bu.user_name as userName,bc.content,bc.create_time as createTime " +
		"from blog_comments bc " +
		"JOIN blog_users bu on bu.id=bc.user_id " +
		"where bc.post_id=? and bc.state=1 " +
		"order by bc.create_time asc"
	var totalNum int64
	c.db.Raw(varSqlCount, postId).Scan(&totalNum)
	if totalNum <= 0 {
		return &dtos.BasePageDto[dtos.PostCommentDto]{}, nil
	}
	varSqlQuery := "select bc.id,bu.user_name as userName,bc.content,bc.create_time as createTime " +
		"from blog_comments bc " +
		"JOIN blog_users bu on bu.id=bc.user_id " +
		"where bc.post_id=? and bc.state=1 " +
		"order by bc.create_time asc " +
		"limit ?,?"
	var list = dtos.PostCommentDto{}
	tx := c.db.Raw(varSqlQuery, postId, pageOffset, pageSize).Scan(&list)
	if tx.Error != nil {
		return &dtos.BasePageDto[dtos.PostCommentDto]{}, tx.Error
	}
	var result = dtos.BasePageDto[dtos.PostCommentDto]{TotalCount: totalNum, List: &list}
	return &result, nil
}
