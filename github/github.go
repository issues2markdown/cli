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

package github

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// User ...
type User struct {
	Login string
}

// Issue ...
type Issue struct {
	Title      *string
	Closed     bool
	HTMLURL    *string
	Repository string
}

// GithubProvider ...
type GithubProvider struct {
	APIClient *github.Client
	User      User
	Issues    []Issue
}

// NewGithubProvider ...
func NewGithubProvider() (*GithubProvider, error) {
	provider := &GithubProvider{}

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
		return provider, err
	}
	provider.User.Login = user.GetLogin()

	provider.APIClient = client

	log.Printf("Created authenticated data for user: %s\n", provider.User.Login)

	return provider, nil
}

// QueryIssues ...
func (gp *GithubProvider) QueryIssues() ([]Issue, error) {
	ctx := context.Background()

	// query issues
	options := &github.SearchOptions{
		Sort:  "created",
		Order: "asc",
	}
	query := fmt.Sprintf("is:issue author:%s state:open", gp.User.Login)
	searchResult, _, err := gp.APIClient.Search.Issues(ctx, query, options)
	if err != nil {
		return nil, err
	}

	// process query results
	for _, v := range searchResult.Issues {
		item := Issue{
			Title:   v.Title,
			HTMLURL: v.HTMLURL,
		}
		organization, repository, err := gp.getOrgAndRepoFromIssueURL(v.URL)
		if err != nil {
			return nil, err
		}
		item.Repository = fmt.Sprintf("%s/%s", organization, repository)
		gp.Issues = append(gp.Issues, item)
	}

	log.Printf("Fetched %d issues", len(gp.Issues))

	return gp.Issues, nil
}

// getOrgAndRepoFromIssueURL ...
func (gp *GithubProvider) getOrgAndRepoFromIssueURL(u *string) (string, string, error) {
	parsedU, err := url.Parse(*u)
	if err != nil {
		return "", "", err
	}
	parsedPartsPathU := strings.Split(parsedU.Path, "/")
	organization := parsedPartsPathU[2]
	repository := parsedPartsPathU[3]
	return organization, repository, nil
}
