// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/caixw/typing/core"
	"github.com/issue9/mux"
	"github.com/issue9/orm"
)

var (
	db  *orm.DB
	opt *core.Options
)

func Init(options *core.Options, database *orm.DB) {
	opt = options
	db = database
}

func InitRoute(admin *mux.Prefix) {
	admin.PostFunc("/login", adminPostLogin).
		Delete("/login", loginHandlerFunc(adminDeleteLogin)).
		Put("/password", loginHandlerFunc(adminChangePassword)).
		Get("/state", loginHandlerFunc(adminGetState)).
		Put("/sitemap", loginHandlerFunc(adminPutSitemap))

	admin.Get("/themes", loginHandlerFunc(adminGetThemes)).
		Get("/themes/current", loginHandlerFunc(adminGetCurrentTheme)).
		Put("/themes/current", loginHandlerFunc(adminPutCurrentTheme))

	// options
	admin.Get("/options/{key}", loginHandlerFunc(adminGetOption)).
		Patch("/options/{key}", loginHandlerFunc(adminPatchOption))

	// tags
	admin.Put("/tags/{id:\\d+}", loginHandlerFunc(adminPutTag)).
		Delete("/tags/{id:\\d+}", loginHandlerFunc(adminDeleteTag)).
		Get("/tags/{id:\\d+}", loginHandlerFunc(adminGetTag)).
		Post("/tags", loginHandlerFunc(adminPostTag)).
		Get("/tags", loginHandlerFunc(adminGetTags))

	// comments
	admin.Get("/comments", loginHandlerFunc(adminGetComments)).
		Get("/comments/count", loginHandlerFunc(adminGetCommentsCount)).
		Post("/comments", loginHandlerFunc(adminPostComment)).
		Put("/comments/{id:\\d+}", loginHandlerFunc(adminPutComment)).
		Delete("comments/{id:\\d+}", loginHandlerFunc(adminDeleteComment)).
		Post("/comments/{id:\\d+}/waiting", loginHandlerFunc(adminSetCommentWaiting)).
		Post("/comments/{id:\\d+}/spam", loginHandlerFunc(adminSetCommentSpam)).
		Post("/comments/{id:\\d+}/approved", loginHandlerFunc(adminSetCommentApproved))

	// posts
	admin.Get("/posts", loginHandlerFunc(adminGetPosts)).
		Get("/posts/count", loginHandlerFunc(adminGetPostsCount)).
		Post("/posts", loginHandlerFunc(adminPostPost)).
		Get("/posts/{id:\\d+}", loginHandlerFunc(adminGetPost)).
		Delete("/posts/{id:\\d+}", loginHandlerFunc(adminDeletePost)).
		Put("/posts/{id:\\d+}", loginHandlerFunc(adminPutPost)).
		Post("/posts/{id:\\d+}/draft", loginHandlerFunc(adminSetPostDraft)).
		Post("/posts/{id:\\d+}/published", loginHandlerFunc(adminSetPostPublished))
}
