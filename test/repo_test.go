// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// +build integration

package test

import (
	"fmt"
	"math/rand"
	"net/url"
	"testing"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/util/idutil"
	"openpitrix.io/openpitrix/test/client/repo_manager"
	"openpitrix.io/openpitrix/test/models"
)

//var clientConfig = &ClientConfig{}
//
//func init() {
//	clientConfig = GetClientConfig()
//	log.Printf("Got Client Config: %+v", clientConfig)
//}

//func TestMain(m *testing.M) {
//	os.Exit(m.Run())
//}

var (
	REPO_URL = "https://kubernetes-charts.storage.googleapis.com"
)

func TestRepo(t *testing.T) {
	client := GetClient(clientConfig)

	// test validate repo
	repoType := "https"
	credential := "{}"
	validateParams := repo_manager.NewValidateRepoParams()
	validateParams.SetType(&repoType)
	validateParams.SetURL(&REPO_URL)
	validateParams.SetCredential(&credential)
	validateResp, err := client.RepoManager.ValidateRepo(validateParams)
	if err != nil {
		t.Fatal(err)
	}
	if validateResp.Payload.Ok != true {
		t.Fatal("validate repo failed")
	}

	// delete old repo
	testRepoName := "e2e_test_repo"
	describeParams := repo_manager.NewDescribeReposParams()
	describeParams.SetName([]string{testRepoName})
	describeParams.SetStatus([]string{constants.StatusActive})
	describeResp, err := client.RepoManager.DescribeRepos(describeParams)
	if err != nil {
		t.Fatal(err)
	}
	repos := describeResp.Payload.RepoSet
	for _, repo := range repos {
		deleteParams := repo_manager.NewDeleteRepoParams()
		deleteParams.SetBody(
			&models.OpenpitrixDeleteRepoRequest{
				RepoID: repo.RepoID,
			})
		_, err := client.RepoManager.DeleteRepo(deleteParams)
		if err != nil {
			t.Fatal(err)
		}
	}
	// create repo
	createParams := repo_manager.NewCreateRepoParams()
	createParams.SetBody(
		&models.OpenpitrixCreateRepoRequest{
			Name:        testRepoName,
			Description: "description",
			Type:        "https",
			URL:         REPO_URL,
			Credential:  `{}`,
			Visibility:  "public",
			Providers:   []string{constants.ProviderKubernetes},
		})
	createResp, err := client.RepoManager.CreateRepo(createParams)
	if err != nil {
		t.Fatal(err)
	}
	repoId := createResp.Payload.Repo.RepoID
	// modify repo
	modifyParams := repo_manager.NewModifyRepoParams()
	modifyParams.SetBody(
		&models.OpenpitrixModifyRepoRequest{
			RepoID:      repoId,
			Description: "cc",
			Type:        "https",
			URL:         REPO_URL,
			Credential:  `{}`,
			Visibility:  "private",
		})
	modifyResp, err := client.RepoManager.ModifyRepo(modifyParams)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(modifyResp)
	// describe repo
	describeParams.WithRepoID([]string{repoId})
	describeResp, err = client.RepoManager.DescribeRepos(describeParams)
	if err != nil {
		t.Fatal(err)
	}
	repos = describeResp.Payload.RepoSet
	if len(repos) != 1 {
		t.Fatalf("failed to describe repos with params [%+v]", describeParams)
	}
	if repos[0].Name != testRepoName || repos[0].Description != "cc" || repos[0].URL != REPO_URL {
		t.Fatalf("failed to modify repo [%+v]", repos[0])
	}
	// delete repo
	deleteParams := repo_manager.NewDeleteRepoParams()
	deleteParams.WithBody(&models.OpenpitrixDeleteRepoRequest{
		RepoID: repoId,
	})
	deleteResp, err := client.RepoManager.DeleteRepo(deleteParams)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(deleteResp)
	// describe deleted repo
	describeParams.WithRepoID([]string{repoId})
	describeParams.WithStatus([]string{constants.StatusDeleted})
	describeParams.WithName(nil)
	describeResp, err = client.RepoManager.DescribeRepos(describeParams)
	if err != nil {
		t.Fatal(err)
	}
	repos = describeResp.Payload.RepoSet
	if len(repos) != 1 {
		t.Fatalf("failed to describe repos with params [%+v]", describeParams)
	}
	repo := repos[0]
	if repo.RepoID != repoId {
		t.Fatalf("failed to describe repo")
	}
	if repo.Status != constants.StatusDeleted {
		t.Fatalf("failed to delete repo, got repo status [%s]", repo.Status)
	}

	t.Log("test repo finish, all test is ok")
}

func generateRepoLabels(length int) (labels []*models.OpenpitrixRepoLabel) {
	i := 0
	for i < length {
		labels = append(labels, &models.OpenpitrixRepoLabel{LabelKey: getRandomKey(), LabelValue: idutil.GetUuid("")})
		i++
	}
	return labels
}

func getRandomNumber() int {
	return rand.Intn(10) + 1
}

func getRandomKey() string {
	return fmt.Sprintf("key%d", getRandomNumber())
}

func getRepoLabel(labels []*models.OpenpitrixRepoLabel) *string {
	v := url.Values{}
	for _, label := range labels {
		v.Add(label.LabelKey, label.LabelValue)
	}
	label := v.Encode()
	return &label
}

func generateRepoSelectors(length int) (labels []*models.OpenpitrixRepoSelector) {
	i := 0
	for i < length {
		labels = append(labels, &models.OpenpitrixRepoSelector{SelectorKey: getRandomKey(), SelectorValue: idutil.GetUuid("")})
		i++
	}
	return labels
}

func getRepoSelector(labels []*models.OpenpitrixRepoSelector) *string {
	v := url.Values{}
	for _, label := range labels {
		v.Add(label.SelectorKey, label.SelectorValue)
	}
	label := v.Encode()
	return &label
}

func testDescribeReposWithLabelSelector(t *testing.T,
	repoId string,
	labels []*models.OpenpitrixRepoLabel,
	selectors []*models.OpenpitrixRepoSelector) {
	client := GetClient(clientConfig)

	describeParams := repo_manager.NewDescribeReposParams()
	describeParams.SetLabel(getRepoLabel(labels))
	describeParams.SetSelector(getRepoSelector(selectors))
	describeParams.SetStatus([]string{constants.StatusActive})
	describeResp, err := client.RepoManager.DescribeRepos(describeParams)
	if err != nil {
		t.Fatal(err)
	}
	if describeResp.Payload.RepoSet[0].RepoID != repoId {
		t.Fatalf("describe repo with filter failed")
	}
	repo := describeResp.Payload.RepoSet[0]
	for i, label := range repo.Labels {
		if label.LabelKey != labels[i].LabelKey {
			t.Fatalf("repo label key not matched")
		}
		if label.LabelValue != labels[i].LabelValue {
			t.Fatalf("repo label value not matched")
		}
	}
	for i, selector := range repo.Selectors {
		if selector.SelectorKey != selectors[i].SelectorKey {
			t.Fatalf("repo selector key not matched")
		}
		if selector.SelectorValue != selectors[i].SelectorValue {
			t.Fatalf("repo selector value not matched")
		}
	}
}

func TestRepoLabelSelector(t *testing.T) {
	client := GetClient(clientConfig)
	// Create a test repo that can attach label and selector on it
	testRepoName := "e2e_test_repo"
	labels := generateRepoLabels(6)
	selectors := generateRepoSelectors(6)
	createParams := repo_manager.NewCreateRepoParams()
	createParams.SetBody(
		&models.OpenpitrixCreateRepoRequest{
			Name:        testRepoName,
			Description: "description",
			Type:        "https",
			URL:         REPO_URL,
			Credential:  `{}`,
			Visibility:  "public",
			Labels:      labels,
			Selectors:   selectors,
		})
	createResp, err := client.RepoManager.CreateRepo(createParams)
	if err != nil {
		t.Fatal(err)
	}
	repoId := createResp.Payload.Repo.RepoID
	testDescribeReposWithLabelSelector(t, repoId, labels, selectors)

	i := 0
	for i < 10 {
		i++
		newLabels := generateRepoLabels(getRandomNumber())
		newSelectors := generateRepoSelectors(getRandomNumber())
		modifyParams := repo_manager.NewModifyRepoParams()
		modifyParams.SetBody(
			&models.OpenpitrixModifyRepoRequest{
				RepoID:    repoId,
				Labels:    newLabels,
				Selectors: newSelectors,
			},
		)
		_, err := client.RepoManager.ModifyRepo(modifyParams)
		if err != nil {
			t.Fatal(err)
		}
		testDescribeReposWithLabelSelector(t, repoId, newLabels, newSelectors)
	}

	// delete repo
	deleteParams := repo_manager.NewDeleteRepoParams()
	deleteParams.WithBody(&models.OpenpitrixDeleteRepoRequest{
		RepoID: repoId,
	})
	deleteResp, err := client.RepoManager.DeleteRepo(deleteParams)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(deleteResp)

	t.Log("test repo label and selector finish, all test is ok")
}
