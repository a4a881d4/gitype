// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package client

import (
	"net/http"

	"github.com/caixw/typing/vars"
)

func (client *Client) initOpensearch() error {
	if client.data.Config.Opensearch == nil {
		return nil
	}

	if err := client.buildOpensearch(); err != nil {
		return err
	}

	conf := client.data.Config
	client.patterns = append(client.patterns, conf.Opensearch.URL)
	client.mux.GetFunc(conf.Opensearch.URL, client.prepare(func(w http.ResponseWriter, r *http.Request) {
		setContentType(w, conf.Opensearch.Type)
		w.Write(client.opensearch)
	}))

	return nil
}

// 用于生成一个符合 atom 规范的 XML 文本。
func (client *Client) buildOpensearch() error {
	w := newWrite()
	o := client.data.Config.Opensearch

	w.writeStartElement("OpenSearchDescription", map[string]string{
		"xmlns": "http://a9.com/-/spec/opensearch/1.1/",
	})

	w.writeElement("InputEncoding", "UTF-8", nil)
	w.writeElement("OutputEncoding", "UTF-8", nil)
	w.writeElement("ShortName", o.ShortName, nil)
	w.writeElement("Description", o.Description, nil)

	if len(o.LongName) > 0 {
		w.writeElement("LongName", o.LongName, nil)
	}

	if o.Image != nil {
		w.writeElement("Image", o.Image.URL, map[string]string{
			"type": o.Image.Type,
		})
	}

	w.writeCloseElement("Url", map[string]string{
		"type":     "text/html",
		"template": vars.SearchURL("{searchTerms}", 0),
	})

	w.writeElement("Developer", vars.AppName, nil)
	w.writeElement("Language", client.data.Config.Language, nil)

	w.writeEndElement("OpenSearchDescription")

	bs, err := w.bytes()
	if err != nil {
		return err
	}
	client.opensearch = bs

	return nil
}