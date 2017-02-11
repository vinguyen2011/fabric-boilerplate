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
	ProjectID   		string 		`json:"projectId"`
	Description 		string 		`json:"description"`
	Location     		string 		`json:"location"`
	Tags 				[]string 	`json:"tags"`
	VoteRestriction		string		`json:"voteRestriction"`
	ExpiryDate			uint64		`json:"expiryDate"`
	Cost     			uint64 		`json:"cost"`
	CostCovered			uint64		`json:"costCovered"`
}

type Projects struct {
	Projects []Project `json:"projects"`
}

type Voter struct {
	VoterId				string		`json:"voterId"`
	Location     		string 		`json:"location"`
	Gender				string		`json:"gender"`
	Age					string		`json:"age"`
}

type Vote struct {
	VoterId				string		`json:"voterId"`
	VotePercent			float32		`json:"voterPercent"`
}

type ProjectVotes struct {
	ProjectID   		string 		`json:"projectId"`
	Votes				[]Vote		`json:"votes"`
}


