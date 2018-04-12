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

package issues2markdown_test

import (
	"testing"

	"github.com/repejota/issues2markdown"

	"github.com/repejota/issues2markdown/github"
)

func TestRender(t *testing.T) {
	var issues []github.Issue
	issue1 := github.Issue{
		Title:      "github issue 1",
		HTMLURL:    "https://github.com/organization/repository/issues/1",
		Repository: "organization/repository",
	}
	issues = append(issues, issue1)
	issue2 := github.Issue{
		Title:      "github issue 2",
		HTMLURL:    "https://github.com/organization/repository/issues/2",
		Repository: "organization/repository",
	}
	issues = append(issues, issue2)

	expected := `- [ ] organization/repository : [github issue 1](https://github.com/organization/repository/issues/1)
- [ ] organization/repository : [github issue 2](https://github.com/organization/repository/issues/2)
`
	i2md := issues2markdown.NewIssuesToMarkdown()
	result, err := i2md.Render(issues)
	if err != nil {
		t.Fatal(err)
	}
	if result.String() != expected {
		t.Fatalf("Expected result: \n%s\n-------\nBut got: \n%s\n", expected, result.String())
	}
}
