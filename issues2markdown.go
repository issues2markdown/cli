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
	"html/template"
	"strings"

	"github.com/repejota/issues2markdown/github"
)

const (
	issuesTemplate = `{{ range . }}- [ ] {{ .Repository }} : [{{ .Title }}]({{ .HTMLURL }})
{{ end }}`
)

// IssuesToMarkdown ...
type IssuesToMarkdown struct {
}

// NewIssuesToMarkdown ...
func NewIssuesToMarkdown() *IssuesToMarkdown {
	i2md := &IssuesToMarkdown{}
	return i2md
}

// Query ...
func (i *IssuesToMarkdown) Query() ([]github.Issue, error) {
	// create authenticated client
	provider, err := github.NewGithubProvider()
	if err != nil {
		return nil, err
	}
	// query issues
	issuesList, err := provider.Query()
	if err != nil {
		return nil, err
	}
	return issuesList, nil
}

// Render ...
func (i *IssuesToMarkdown) Render(issues []github.Issue) (string, error) {
	var compiled bytes.Buffer
	t := template.Must(template.New("issueslist").Parse(issuesTemplate))
	_ = t.Execute(&compiled, issues)
	result := compiled.String()
	result = strings.TrimRight(result, "\n") // trim the last linebreak
	return result, nil
}
