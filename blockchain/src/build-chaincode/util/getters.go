package util

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"errors"
	"encoding/json"
	"build-chaincode/entities"
)

func GetCurrentBlockchainUser(stub shim.ChaincodeStubInterface) (entities.User, error) {
	userIDAsBytes, err := stub.ReadCertAttribute("userID")
	if err != nil {
		return entities.User{}, errors.New("Could not retrieve user by certificate. Reason: " + err.Error())
	}

	return GetUser(stub, string(userIDAsBytes))
}

func GetProjectsForVoter(stub shim.ChaincodeStubInterface, voterID string) ([]entities.Project, error) {
	projectsIndex, err := GetIndex(stub, ProjectsIndexName)
	if err != nil {
		return []entities.Project{}, errors.New("Unable to retrieve projectsIndex, reason: " + err.Error())
	}

	projects := []entities.Project{}
	for _, projectID := range projectsIndex {
		projectAsBytes, err := stub.GetState(projectID)
		if err != nil {
			return []entities.Project{}, errors.New("Could not retrieve project for ID " + projectID + " reason: " + err.Error())
		}

		var project entities.Project
		err = json.Unmarshal(projectAsBytes, &project)
		if err != nil {
			return []entities.Project{}, errors.New("Error while unmarshalling projectAsBytes, reason: " + err.Error())
		}
		
		if ValidateProjectForVoterId(stub, project, voterID) {
			
			//projects = append(projects, string(projectAsBytes))
			projects = append(projects, project)
		}
		
	}

	return projects, nil
}

func GetVoter(stub shim.ChaincodeStubInterface, voterID string) (entities.Voter, error) {
	
	voterAsBytes, err := stub.GetState(VoterIndexPrefix + voterID)
	if err != nil {
		return entities.Voter{}, errors.New("Could not retrieve Voter for ID " + voterID + " reason: " + err.Error())
	}

	var voter entities.Voter
	err = json.Unmarshal(voterAsBytes, &voter)
	if err != nil {
		return entities.Voter{}, errors.New("Error while unmarshalling voterAsBytes, reason: " + err.Error())
	}
	
	return voter, nil

}


func ValidateProjectForVoterId(stub shim.ChaincodeStubInterface, project entities.Project, voterID string) (bool) {
	
	voter, err := GetVoter(stub, voterID)

	if err != nil {
		return false
	}

	return ValidateProjectForVoter(project, voter)
	
}

func ValidateProjectForVoter(project entities.Project, voter entities.Voter) (bool) {
	
	//validation 1: check if the user is voting twice!
	for _, votedProjectID := range voter.ProjectIDs {
		if votedProjectID == project.ProjectID {
			//user already voted. double spend!
			return false
		}
	}

	//validation 2: check if the user is allowed to vote on this project
	if project.VoteRestrictionField == "LOCATION" {
	
		return ContainsInList(project.VoteRestrictionValues, voter.Location)
		
	} else if project.VoteRestrictionField == "GENDER" {
	
		return ContainsInList(project.VoteRestrictionValues, voter.Gender)
		
	}
	
	return false; 
	
}

func ContainsInList(list []string, item string) (bool) {
	
	for _, listItem := range list {
		if listItem == item {
			return true
		}
	}
	
	return false;
	
}


func GetThingsByUserID(stub shim.ChaincodeStubInterface, userID string) ([]string, error) {
	thingsIndex, err := GetIndex(stub, ThingsIndexName)
	if err != nil {
		return []string{}, errors.New("Unable to retrieve thingsIndex, reason: " + err.Error())
	}

	thingIDs := []string{}
	for _, thingID := range thingsIndex {
		thingAsBytes, err := stub.GetState(thingID)
		if err != nil {
			return []string{}, errors.New("Could not retrieve thing for ID " + thingID + " reason: " + err.Error())
		}

		var thing entities.Thing
		err = json.Unmarshal(thingAsBytes, &thing)
		if err != nil {
			return []string{}, errors.New("Error while unmarshalling thingAsBytes, reason: " + err.Error())
		}

		if thing.UserID == userID {
			thingIDs = append(thingIDs, thing.ThingID)
		}
	}

	return thingIDs, nil
}

func GetUser(stub shim.ChaincodeStubInterface, username string) (entities.User, error) {
	userAsBytes, err := stub.GetState(username)
	if err != nil {
		return entities.User{}, errors.New("Could not retrieve information for this user")
	}

	var user entities.User
	if err = json.Unmarshal(userAsBytes, &user); err != nil {
		return entities.User{}, errors.New("Cannot get user, reason: " + err.Error())
	}

	return user, nil
}

func GetAllUsers(stub shim.ChaincodeStubInterface) ([]entities.User, error) {
	usersIndex, err := GetIndex(stub, UsersIndexName)
	if err != nil {
		return []entities.User{}, errors.New("Could not retrieve userIndex, reason: " + err.Error())
	}

	var users []entities.User
	for _, userID := range usersIndex {
		userAsBytes, err := stub.GetState(userID)
		if err != nil {
			return []entities.User{}, errors.New("Could not retrieve user with ID: " + userID + ", reason: " + err.Error())
		}

		var user entities.User
		err = json.Unmarshal(userAsBytes, &user)
		if err != nil {
			return []entities.User{}, errors.New("Error while unmarshalling user, reason: " + err.Error())
		}

		users = append(users, user)
	}

	return users, nil
}
