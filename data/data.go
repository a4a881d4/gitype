// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// 负责加data目录下的数据。
// 会调用github.com/issue9/logs包的内容，调用之前需要初始化该包。
package data

import (
	"html/template"

	"github.com/caixw/typing/vars"
)

// 客户保存的时间格式。
const parseDateFormat = "2006-01-02T15:04:05-0700"

// Data 结构体包含了数据目录下所有需要加载的数据内容。
type Data struct {
	path *vars.Path

	Config   *Config            // 配置内容
	URLS     *URLS              // 自定义URL
	Tags     []*Tag             // map对顺序是未定的，所以使用slice
	Links    []*Link            // 友情链接
	Template *template.Template // 当前主题模板
	Posts    []*Post            // 所有的文章列表
}

// Load 函数用于加载一份新的数据。
func Load(path *vars.Path) (*Data, error) {
	d := &Data{
		path: path,
	}

	if err := d.loadMeta(); err != nil {
		return nil, err
	}

	// 加载文章
	if err := d.loadPosts(path.DataPosts); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Data) loadMeta() error {
	// urls
	if err := d.loadURLS(d.path.DataURLS); err != nil {
		return err
	}

	// tags
	if err := d.loadTags(d.path.DataTags); err != nil {
		return err
	}

	// config
	if err := d.loadConfig(d.path.DataConf); err != nil {
		return err
	}

	// 加载主题的模板
	if err := d.loadTemplate(d.path.DataThemes); err != nil {
		return err
	}

	// links
	return d.loadLinks(d.path.DataLinks)
}
