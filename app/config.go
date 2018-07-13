package app

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/teyushen/star-go/arrays"
)

type User struct {
	Token string
}

type RepoConfig struct {
	Owner     string
	RepoNames []string
}

func SaveUser(user User) {

	writeToConfig(".config", user)
}

func GetUser() User {

	content, err := ioutil.ReadFile(".config")
	if err != nil {
		log.Fatal(err)
	}

	user := User{}
	json.Unmarshal(content, &user)

	log.Printf("Filename: [.config] -> %s", user)

	return user
}

func SaveRepos(repos []RepoConfig) {

	repoConfigs := make([]RepoConfig, 0)
	for _, repo := range repos {
		needMerge := mergeRepo(repoConfigs, repo)
		if !needMerge {
			repoConfigs = append(repoConfigs, repo)
		}
	}
	writeToConfig(".star-go", repoConfigs)
}

func mergeRepo(repos []RepoConfig, repo RepoConfig) bool {

	for index, value := range repos {
		if value.Owner == repo.Owner {
			for _, name := range repo.RepoNames {
				if !arrays.Contains(value.RepoNames, name) {
					repos[index].RepoNames = append(repos[index].RepoNames, name)
				}
			}
			return true
		}
	}
	return false
}

func AppendRepos(repos ...RepoConfig) []RepoConfig {

	repoConfigs := GetRepos()
	for _, repo := range repos {
		needMerge := mergeRepo(repoConfigs, repo)
		if !needMerge {
			repoConfigs = append(repoConfigs, repo)
		}
	}
	SaveRepos(repoConfigs)
	return repoConfigs
}

func GetRepos() []RepoConfig {

	content, err := ioutil.ReadFile(".star-go")
	if err != nil {
		log.Fatal(err)
	}

	repos := make([]RepoConfig, 0)
	json.Unmarshal(content, &repos)

	log.Printf("Filename: [.star-go] -> %s", content)

	return repos
}

func writeToConfig(filename string, str interface{}) {
	b, _ := json.Marshal(str)

	err := ioutil.WriteFile(filename, b, 0777)

	if err != nil {
		panic(err)
	}
	return
}
