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
	"net/url"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	issuesTemplate = `{{ range . }}- [ ] {{ .Organization }}/{{ .Repository }} : [{{ .Title }}]({{ .HTMLURL }})
{{ end }}`
)

// User ...
type User struct {
	Login string
}

// Issue ...
type Issue struct {
	Title   string
	URL     string
	HTMLURL string
}

// Organization ...
func (i *Issue) Organization() (string, error) {
	parsedU, _ := url.Parse(i.URL)
	parsedPartsPathU := strings.Split(parsedU.Path, "/")
	organization := parsedPartsPathU[2]
	return organization, nil
}

// Repository ...
func (i *Issue) Repository() (string, error) {
	parsedU, _ := url.Parse(i.URL)
	parsedPartsPathU := strings.Split(parsedU.Path, "/")
	repository := parsedPartsPathU[3]
	return repository, nil
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
func (im *IssuesToMarkdown) Query() ([]Issue, error) {
	// create github client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "f8cb77e12d26827e5d235e93cae6bda4796236e1"},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	// get user information
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	im.User.Login = user.GetLogin()
	log.Printf("Created authenticated data for user: %s\n", im.User.Login)

	// query issues
	options := &github.SearchOptions{}
	query := fmt.Sprintf("state:open type:issue author:%s archived:false", im.User.Login)
	searchResult, _, err := client.Search.Issues(ctx, query, options)
	if err != nil {
		return nil, err
	}

	// process results
	var result []Issue
	for _, v := range searchResult.Issues {
		item := Issue{
			Title:   *v.Title,
			URL:     *v.URL,
			HTMLURL: *v.HTMLURL,
		}
		result = append(result, item)
	}

	return result, nil
}

// Render ...
func (im *IssuesToMarkdown) Render(issues []Issue) (string, error) {
	var compiled bytes.Buffer
	t := template.Must(template.New("issueslist").Parse(issuesTemplate))
	_ = t.Execute(&compiled, issues)
	result := compiled.String()
	result = strings.TrimRight(result, "\n") // trim the last linebreak
	return result, nil
}

// GetOrgAndRepoFromIssueURL ...
func (im *IssuesToMarkdown) GetOrgAndRepoFromIssueURL(u string) (string, string, error) {
	parsedU, _ := url.Parse(u)
	parsedPartsPathU := strings.Split(parsedU.Path, "/")
	organization := parsedPartsPathU[2]
	repository := parsedPartsPathU[3]
	return organization, repository, nil
}
