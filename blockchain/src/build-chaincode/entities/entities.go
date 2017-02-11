package entities

type ECertResponse struct {
	OK string `json:"OK"`
}

type TestData struct {
	Users  		[]User 	 `json:"users"`
	Things 		[]Thing  `json:"things"`
}

type TestDataElement interface {
	ID() string
}

type User struct {
	TestDataElement 	`json:"-"`
	UserID   	string 	`json:"userID"`
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
	Salt     	string 	`json:"salt"`
	Hash     	string 	`json:"hash"`
}

type Thing struct {
	TestDataElement    	`json:"-"`
	ThingID      	string 	`json:"thingID"`
	SomeProperty 	string 	`json:"someProperty"`
	UserID    	string 	`json:"userID"`
}

type UserAuthenticationResult struct {
	User        	User
	Authenticated 	bool
}

type Users struct {
	Users []User `json:"users"`
}

func (t *User) ID() string {
	return t.Username
}

func (t *Thing) ID() string {
	return t.ThingID
}

//======================================================================================================================
// model 
//======================================================================================================================

type Project struct {
	ProjectID   			string 		`json:"projectId"`
	Name					string		`json:"name"`
	Description 			string 		`json:"description"`
	Location     			string 		`json:"location"`
	Tags 					[]string 	`json:"tags"`
	VoteRestrictionField	string		`json:"voteRestrictionField"`
	VoteRestrictionValues	[]string	`json:"voteRestrictionValues"`
	ExpiryDate				uint64		`json:"expiryDate"`
	Cost     				uint64 		`json:"cost"`
	CostCovered				uint64		`json:"costCovered"`
	PictureID				uint64		`json:"pictureId"`
}

type Projects struct {
	Projects []Project `json:"projects"`
}

type Voter struct {
	VoterId				string		`json:"voterId"`
	Location     		string 		`json:"location"`
	Gender				string		`json:"gender"`
	Dob					uint64		`json:"dob"`
	ProjectIDs   		[]string 	`json:"projectIds"`
}

type Vote struct {
	VoterId				string		`json:"voterId"`
	ProjectID   		string 		`json:"projectId"`
	VotePercent			float32		`json:"votePercent"`
}

type ProjectVotes struct {
	ProjectID   		string 		`json:"projectId"`
	Votes				[]Vote		`json:"votes"`
}


