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

	"github.com/google/go-github/github"
)

// Issue ...
type Issue struct {
	Title      *string
	Closed     bool
	HTMLURL    *string
	Repository string
}

// IssuesList ...
type IssuesList struct {
	Issues []Issue
}

const (
	issuesTemplate = `{{ range .Issues }}- [ ] {{ .Repository }} : [{{ .Title }}]({{ .HTMLURL }})
{{ end }}`
)

// Render ...
func Render(issues IssuesList) (bytes.Buffer, error) {
	var result bytes.Buffer
	t := template.Must(template.New("issueslist").Parse(issuesTemplate))
	_ = t.Execute(&result, issues)
	return result, nil
}

// Fetch ...
func Fetch() (IssuesList, error) {
	var data IssuesList
	ctx := context.Background()
	client := github.NewClient(nil)
	options := &github.SearchOptions{
		Sort:  "created",
		Order: "asc",
	}
	result, _, err := client.Search.Issues(ctx, "is:issue state:open", options)
	if err != nil {
		return data, err
	}
	for _, v := range result.Issues {
		issue := Issue{
			Title:   v.Title,
			HTMLURL: v.HTMLURL,
		}
		organization, repository, err := GetOrgAndRepoFromIssueURL(*v.HTMLURL)
		if err != nil {
			return data, err
		}
		issue.Repository = fmt.Sprintf("%s/%s", organization, repository)
		data.Issues = append(data.Issues, issue)
	}
	return data, nil
}
