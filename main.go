package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"os"
	"time"
)

type Configuration struct {
	Repos    []string
	Users    []string
	Duration float64
}

func main() {
	config, err := getConfig()
	if err != nil {
		Warning("Cannot open config file. ", err)
		return
	}

	if len(config.Users) == 0 {
		Warning("No users in your config file.")
		return
	}

	if len(config.Repos) == 0 {
		Warning("No repos in your config file.")
		return
	}

	now := time.Now()
	var team map[string]bool

	for _, project := range config.Repos {
		team = getTeam(config.Users)
		Info("Checking Repo: %s", project)
		r, err := git.PlainOpen(project)
		if err != nil {
			Warning("Error opening %s, moving on.", project)
			continue
		}

		ref, err := r.Head()
		CheckIfError(err)

		commits, err := r.Log(&git.LogOptions{From: ref.Hash(), Order: git.LogOrderCommitterTime})
		CheckIfError(err)

		for {
			commit, err := commits.Next()
			CheckIfError(err)
			when := commit.Author.When
			committer := commit.Author.Name
			timeDiff := now.Sub(when)
			if timeDiff.Hours() <= config.Duration {
				if _, ok := team[committer]; ok {
					fmt.Println(committer, when, commit.Message)
					delete(team, committer)
				}
			} else {
				break
			}
		}
	}
}

func getTeam(teamMembers []string) map[string]bool {
	team := make(map[string]bool)
	for _, teamMember := range teamMembers {
		team[teamMember] = true
	}
	return team
}

func getConfig() (Configuration, error) {
	// Manage config in Golang to get variables from file and env variables
	// https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152

	var configuration = Configuration{}
	file, err := os.Open("config/config.json")
	if err != nil {
		return configuration, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, err
}