package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// Role is a type a user may have.
type Role string

const (
	// What is a student?
	//
	// A student is primarily a person enrolled in a school or other educational institution and who is under learning with goals of acquiring knowledge, developing professions and achieving employment at desired field. In the broader sense, a student is anyone who applies themselves to the intensive intellectual engagement with some matter necessary to master it as part of some practical affair in which such mastery is basic or decisive.
	//
	// (from [Wikipedia](https://en.wikipedia.org/wiki/Student))
	RoleStudent Role = "STUDENT"
	// Teacher is a teacher, who teaches the students.
	RoleTeacher Role = "TEACHER"
)

// unexportedResponse is returned by unexported on success.
type unexportedResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User unexportedUser `json:"user"`
}

// unexportedUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type unexportedUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// UserQueryInput is the argument to Query.users.
//
// Ideally this would support anything and everything!
// Or maybe ideally it wouldn't.
// Really I'm just talking to make this documentation longer.
type userQueryInput struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	// id looks the user up by ID.  It's a great way to look up users.
	Id    testutil.ID `json:"id"`
	Role  Role        `json:"role"`
	Names []string    `json:"names"`
}

func unexported(
	client graphql.Client,
	query userQueryInput,
) (*unexportedResponse, error) {
	variables := map[string]interface{}{
		"query": query,
	}

	var retval unexportedResponse
	err := client.MakeRequest(
		nil,
		"unexported",
		`
query unexported ($query: UserQueryInput) {
	user(query: $query) {
		id
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}

