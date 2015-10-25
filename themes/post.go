// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package themes

import (
	"strconv"
	"time"

	"github.com/caixw/typing/core"
	"github.com/caixw/typing/models"
	"github.com/issue9/logs"
	"github.com/issue9/orm/fetch"
)

// 文章的详细内容
type Post struct {
	ID           int64
	Name         string
	Title        string
	Summary      string
	Content      string
	Author       string // 作者名称
	CommentsSize int    // 评论数量
	Created      int64  // 创建时间
	Modified     int64  // 修改时间
	AllowComment bool   // 是否允许评论
}

func (p *Post) CreatedFormat() string {
	return time.Unix(p.Created, 0).Format(opt.DateFormat)
}

func (p *Post) ModifiedFormat() string {
	return time.Unix(p.Modified, 0).Format(opt.DateFormat)
}

// 返回文章的摘要或是具体内容。
func (p *Post) Entry() string {
	if len(p.Summary) > 0 {
		return p.Summary
	}
	return p.Content
}

// 返回文章的链接
func (p *Post) Permalink() string {
	if len(p.Name) > 0 {
		return opt.SiteURL + "/posts/" + p.Name + opt.Suffix
	}

	return opt.SiteURL + "/posts/" + strconv.FormatInt(p.ID, 10) + opt.Suffix
}

// 获取与当前文章相关的标签。
func (p *Post) Tags() []*Tag {
	sql := `SELECT t.{name} AS Name, t.{title} AS Text FROM #relationships AS r
	LEFT JOIN #tags AS t on t.{id}=r.{tagID}
	WHERE r.{postID}=?`

	rows, err := db.Query(true, sql, p.ID)
	if err != nil {
		logs.Error("themes.Post.Tags:", err)
		return nil
	}
	defer rows.Close()

	tags := make([]*Tag, 0, 5)
	if _, err = fetch.Obj(&tags, rows); err != nil {
		logs.Error("themes.Post.Tags:", err)
		return nil
	}
	return tags
}

// 返回文章的评论信息。
func (p *Post) Comments(page int) []*Comment {
	if page < 1 {
		page = 1
	}

	sql := `SELECT {id} AS ID, {created} AS Created, {agent} AS Agent, {content} AS Content,
	{isAdmin} AS IsAdmin, {authorName} AS AuthorName,{authorURL} AS AuthorURL
	FROM #comments
	WHERE {postID}=? AND {state}=?
	ORDER BY {created} `
	if opt.CommentOrder == core.CommentOrderDesc {
		sql += `DESC `
	}
	sql += `LIMIT ? OFFSET ?`

	rows, err := db.Query(true, sql, p.ID, models.CommentStateApproved, opt.PageSize, opt.PageSize*page)
	if err != nil {
		logs.Error("themes.Post.Comment:", err)
		return nil
	}
	defer rows.Close()

	comments := make([]*Comment, 0, opt.PageSize)
	if _, err := fetch.Obj(&comments, rows); err != nil {
		logs.Error("themes.Post.Comment:", err)
		return nil
	}
	for _, c := range comments {
		c.post = p
	}
	return comments
}