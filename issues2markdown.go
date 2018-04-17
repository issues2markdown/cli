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
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	// DefaultIssueTemplate ...
	DefaultIssueTemplate = `{{ range . }}- [{{ if eq .State "closed" }}x{{ else }} {{ end }}] {{ .Organization }}/{{ .Repository }} : [{{ .Title }}]({{ .HTMLURL }})
{{ end }}`
)

// QueryOptions ...
type QueryOptions struct {
	Organization string
	Repository   string
	State        string
}

// NewQueryOptions ...
func NewQueryOptions(username string) *QueryOptions {
	options := &QueryOptions{
		Organization: username,
		State:        "all",
	}
	return options
}

// BuildQuey ...
func (qo *QueryOptions) BuildQuey() string {
	query := strings.Builder{}
	_, _ = query.WriteString("type:issue") // whe only want issues
	if qo.Repository == "" {
		_, _ = query.WriteString(fmt.Sprintf(" org:%s", qo.Organization)) // organization
	}
	if qo.Repository != "" {
		_, _ = query.WriteString(fmt.Sprintf(" repo:%s/%s", qo.Organization, qo.Repository)) // repository
	}
	if qo.State != "" { // issue status
		if qo.State == "all" {
			_, _ = query.WriteString(fmt.Sprintf(" state:open state:closed"))
		} else {
			_, _ = query.WriteString(fmt.Sprintf(" state:%s", qo.State))
		}
	}
	return query.String()
}

// RenderOptions ...
type RenderOptions struct {
	TemplateSource string
}

// NewRenderOptions ...
func NewRenderOptions() *RenderOptions {
	options := &RenderOptions{
		TemplateSource: DefaultIssueTemplate,
	}
	return options
}

// IssuesToMarkdown ...
type IssuesToMarkdown struct {
	client      *github.Client
	GithubToken string
	Username    string
}

// NewIssuesToMarkdown ...
func NewIssuesToMarkdown(githubToken string) (*IssuesToMarkdown, error) {
	i2md := &IssuesToMarkdown{
		GithubToken: githubToken,
	}
	if i2md.GithubToken == "" {
		return nil, fmt.Errorf("A valid Github Token is required")
	}

	ctx := context.Background()

	// create github client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: i2md.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	i2md.client = github.NewClient(tc)

	// get user information
	user, _, err := i2md.client.Users.Get(ctx, "")
	if err != nil {
		log.Printf("ERROR: %s", err)
		return nil, err
	}
	i2md.Username = user.GetLogin()
	log.Printf("Created authenticated github API client for user: %s\n", i2md.Username)

	return i2md, nil
}

// Query ...
func (im *IssuesToMarkdown) Query(options *QueryOptions) ([]Issue, error) {
	ctx := context.Background()

	// query issues
	query := options.BuildQuey()
	githubOptions := &github.SearchOptions{}
	listResult, _, err := im.client.Search.Issues(ctx, query, githubOptions)
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
	t := template.Must(template.New("issueslist").Parse(options.TemplateSource))
	_ = t.Execute(&compiled, issues)
	result := compiled.String()
	result = strings.TrimRight(result, "\n") // trim the last linebreak
	return result, nil
}
