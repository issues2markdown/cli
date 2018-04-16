// Copyright 2018 The issues2markdown Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package issues2markdown

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	issuesTemplate = `{{ range . }}- [{{ if eq .State "closed" }}*{{ else }} {{ end }}] {{ .Organization }}/{{ .Repository }} : [{{ .Title }}]({{ .HTMLURL }})
{{ end }}`
)

// QueryOptions ...
type QueryOptions struct {
}

// RenderOptions ...
type RenderOptions struct {
}

// IssuesToMarkdown ...
type IssuesToMarkdown struct {
	User User
}

// NewIssuesToMarkdown ...
func NewIssuesToMarkdown() *IssuesToMarkdown {
	i2md := &IssuesToMarkdown{}
	return i2md
}

// Query ...
func (im *IssuesToMarkdown) Query(options *QueryOptions) ([]Issue, error) {
	ctx := context.Background()

	// create github client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// get user information
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Printf("ERROR: %s", err)
		return nil, err
	}
	im.User.Login = user.GetLogin()
	log.Printf("Created authenticated github API client for user: %s\n", im.User.Login)

	// query issues
	query := "type:issue state:open state:closed author:repejota"
	githubOptions := &github.SearchOptions{}
	listResult, _, err := client.Search.Issues(ctx, query, githubOptions)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return nil, err
	}
	log.Printf("Search Query: %s\n", query)
	log.Printf("Total results: %d\n", *listResult.Total)

	// process results
	var result []Issue
	for _, v := range listResult.Issues {
		item := Issue{
			Title:   *v.Title,
			State:   *v.State,
			URL:     *v.URL,
			HTMLURL: *v.HTMLURL,
		}
		result = append(result, item)
	}

	return result, nil
}

// Render ...
func (im *IssuesToMarkdown) Render(issues []Issue, options *RenderOptions) (string, error) {
	var compiled bytes.Buffer
	t := template.Must(template.New("issueslist").Parse(issuesTemplate))
	_ = t.Execute(&compiled, issues)
	result := compiled.String()
	result = strings.TrimRight(result, "\n") // trim the last linebreak
	return result, nil
}
