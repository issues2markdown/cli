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
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/google/go-github/github"
	"github.com/repejota/issues2markdown"
)

const (
	// baseURLPath is a non-empty Client.BaseURL path to use during tests,
	// to ensure relative URLs are used for all endpoints. See issue #752.
	baseURLPath = "/api-v3"
)

type IssuesProvider interface {
}

func TestIntanceQueryOptions(t *testing.T) {
	options := issues2markdown.NewQueryOptions("username")

	expectedOrganization := "username"
	if options.Organization != expectedOrganization {
		t.Fatalf("Default Organization filter expected to be %q but got %q", expectedOrganization, options.Organization)
	}

	expectedRepository := ""
	if options.Repository != expectedRepository {
		t.Fatalf("Default Repository filter expected to be %q but got %q", expectedRepository, options.Repository)
	}

	expectedState := "all"
	if options.State != expectedState {
		t.Fatalf("Default State filter expected to be %q but got %q", expectedState, options.State)
	}
}

func TestBuildQueryQueryOptions(t *testing.T) {
	options := issues2markdown.NewQueryOptions("username")

	expectedQuery := "type:issue org:username state:open state:closed"
	query := options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("Default QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}

	options.Organization = "organization"
	expectedQuery = "type:issue org:organization state:open state:closed"
	query = options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}

	options.Repository = "repository"
	expectedQuery = "type:issue repo:organization/repository state:open state:closed"
	query = options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}

	options.State = "open"
	expectedQuery = "type:issue repo:organization/repository state:open"
	query = options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}

	options.State = "closed"
	expectedQuery = "type:issue repo:organization/repository state:closed"
	query = options.BuildQuey()
	if query != expectedQuery {
		t.Fatalf("QueryOptions query expected to be %q but got %q", expectedQuery, query)
	}
}

func TestInstanceRenderOptions(t *testing.T) {
	options := issues2markdown.NewRenderOptions()
	expectedTemplateSource := issues2markdown.DefaultIssueTemplate
	if options.TemplateSource != expectedTemplateSource {
		t.Fatalf("Default RenderOptions template source expected to be %q but got %q", expectedTemplateSource, options.TemplateSource)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

// providerSetup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func providerSetup(t *testing.T) (client *github.Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// We want to ensure that tests catch mistakes where the endpoint URL is
	// specified as absolute rather than relative. It only makes a difference
	// when there's a non-empty base URL path. So, use that. See issue #752.
	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "FAIL: Client.BaseURL path prefix is not preserved in the request URL:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tDid you accidentally use an absolute endpoint URL rather than relative?")
		fmt.Fprintln(os.Stderr, "\tSee https://github.com/google/go-github/issues/752 for information.")
		http.Error(w, "Client.BaseURL path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the GitHub client being tested and is
	// configured to use test server.
	client = github.NewClient(nil)
	url, _ := url.Parse(server.URL + baseURLPath + "/")
	client.BaseURL = url
	client.UploadURL = url

	return client, mux, server.URL, server.Close
}

func TestInstanceIssuesToMarkdown(t *testing.T) {
	issuesProvider, mux, _, teardown := providerSetup(t)
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"login": "username"}`)
	})
	defer teardown()

	_, err := issues2markdown.NewIssuesToMarkdown(issuesProvider)
	if err != nil {
		t.Fatal(err)
	}
}
