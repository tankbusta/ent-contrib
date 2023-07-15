// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package todo

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entgql/internal/todogotype/ent"
)

type NamedNode interface {
	IsNamedNode()
}

type CategoryTypes struct {
	Public *bool `json:"public"`
}

type CategoryTypesInput struct {
	Public *bool `json:"public"`
}

// CreateUserInput is used for create User object.
// Input was generated by ent.
type CreateUserInput struct {
	Name             *string                `json:"name"`
	Username         *string                `json:"username"`
	Password         *string                `json:"password"`
	RequiredMetadata map[string]interface{} `json:"requiredMetadata"`
	Metadata         map[string]interface{} `json:"metadata"`
	GroupIDs         []string               `json:"groupIDs"`
	FriendIDs        []string               `json:"friendIDs"`
}

type OneToMany struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	Field2   *string      `json:"field2"`
	Parent   *OneToMany   `json:"parent"`
	Children []*OneToMany `json:"children"`
}

func (OneToMany) IsNode() {}

// A connection to a list of items.
type OneToManyConnection struct {
	// A list of edges.
	Edges []*OneToManyEdge `json:"edges"`
	// Information to aid in pagination.
	PageInfo *entgql.PageInfo[string] `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type OneToManyEdge struct {
	// The item at the end of the edge.
	Node *OneToMany `json:"node"`
	// A cursor for use in pagination.
	Cursor entgql.Cursor[string] `json:"cursor"`
}

// Ordering options for OneToMany connections
type OneToManyOrder struct {
	// The ordering direction.
	Direction entgql.OrderDirection `json:"direction"`
	// The field by which to order OneToManies.
	Field OneToManyOrderField `json:"field"`
}

// OneToManyWhereInput is used for filtering OneToMany objects.
// Input was generated by ent.
type OneToManyWhereInput struct {
	Not *OneToManyWhereInput   `json:"not"`
	And []*OneToManyWhereInput `json:"and"`
	Or  []*OneToManyWhereInput `json:"or"`
	// id field predicates
	ID      *string  `json:"id"`
	IDNeq   *string  `json:"idNEQ"`
	IDIn    []string `json:"idIn"`
	IDNotIn []string `json:"idNotIn"`
	IDGt    *string  `json:"idGT"`
	IDGte   *string  `json:"idGTE"`
	IDLt    *string  `json:"idLT"`
	IDLte   *string  `json:"idLTE"`
	// name field predicates
	Name             *string  `json:"name"`
	NameNeq          *string  `json:"nameNEQ"`
	NameIn           []string `json:"nameIn"`
	NameNotIn        []string `json:"nameNotIn"`
	NameGt           *string  `json:"nameGT"`
	NameGte          *string  `json:"nameGTE"`
	NameLt           *string  `json:"nameLT"`
	NameLte          *string  `json:"nameLTE"`
	NameContains     *string  `json:"nameContains"`
	NameHasPrefix    *string  `json:"nameHasPrefix"`
	NameHasSuffix    *string  `json:"nameHasSuffix"`
	NameEqualFold    *string  `json:"nameEqualFold"`
	NameContainsFold *string  `json:"nameContainsFold"`
	// field2 field predicates
	Field2             *string  `json:"field2"`
	Field2neq          *string  `json:"field2NEQ"`
	Field2In           []string `json:"field2In"`
	Field2NotIn        []string `json:"field2NotIn"`
	Field2gt           *string  `json:"field2GT"`
	Field2gte          *string  `json:"field2GTE"`
	Field2lt           *string  `json:"field2LT"`
	Field2lte          *string  `json:"field2LTE"`
	Field2Contains     *string  `json:"field2Contains"`
	Field2HasPrefix    *string  `json:"field2HasPrefix"`
	Field2HasSuffix    *string  `json:"field2HasSuffix"`
	Field2IsNil        *bool    `json:"field2IsNil"`
	Field2NotNil       *bool    `json:"field2NotNil"`
	Field2EqualFold    *string  `json:"field2EqualFold"`
	Field2ContainsFold *string  `json:"field2ContainsFold"`
	// parent edge predicates
	HasParent     *bool                  `json:"hasParent"`
	HasParentWith []*OneToManyWhereInput `json:"hasParentWith"`
	// children edge predicates
	HasChildren     *bool                  `json:"hasChildren"`
	HasChildrenWith []*OneToManyWhereInput `json:"hasChildrenWith"`
}

// OrganizationWhereInput is used for filtering Workspace objects.
// Input was generated by ent.
type OrganizationWhereInput struct {
	Not *OrganizationWhereInput   `json:"not"`
	And []*OrganizationWhereInput `json:"and"`
	Or  []*OrganizationWhereInput `json:"or"`
	// id field predicates
	ID      *string  `json:"id"`
	IDNeq   *string  `json:"idNEQ"`
	IDIn    []string `json:"idIn"`
	IDNotIn []string `json:"idNotIn"`
	IDGt    *string  `json:"idGT"`
	IDGte   *string  `json:"idGTE"`
	IDLt    *string  `json:"idLT"`
	IDLte   *string  `json:"idLTE"`
	// name field predicates
	Name             *string  `json:"name"`
	NameNeq          *string  `json:"nameNEQ"`
	NameIn           []string `json:"nameIn"`
	NameNotIn        []string `json:"nameNotIn"`
	NameGt           *string  `json:"nameGT"`
	NameGte          *string  `json:"nameGTE"`
	NameLt           *string  `json:"nameLT"`
	NameLte          *string  `json:"nameLTE"`
	NameContains     *string  `json:"nameContains"`
	NameHasPrefix    *string  `json:"nameHasPrefix"`
	NameHasSuffix    *string  `json:"nameHasSuffix"`
	NameEqualFold    *string  `json:"nameEqualFold"`
	NameContainsFold *string  `json:"nameContainsFold"`
}

type Project struct {
	ID    string              `json:"id"`
	Todos *ent.TodoConnection `json:"todos"`
}

func (Project) IsNode() {}

// ProjectWhereInput is used for filtering Project objects.
// Input was generated by ent.
type ProjectWhereInput struct {
	Not *ProjectWhereInput   `json:"not"`
	And []*ProjectWhereInput `json:"and"`
	Or  []*ProjectWhereInput `json:"or"`
	// id field predicates
	ID      *string  `json:"id"`
	IDNeq   *string  `json:"idNEQ"`
	IDIn    []string `json:"idIn"`
	IDNotIn []string `json:"idNotIn"`
	IDGt    *string  `json:"idGT"`
	IDGte   *string  `json:"idGTE"`
	IDLt    *string  `json:"idLT"`
	IDLte   *string  `json:"idLTE"`
	// todos edge predicates
	HasTodos     *bool                 `json:"hasTodos"`
	HasTodosWith []*ent.TodoWhereInput `json:"hasTodosWith"`
}

// UpdateFriendshipInput is used for update Friendship object.
// Input was generated by ent.
type UpdateFriendshipInput struct {
	CreatedAt *time.Time `json:"createdAt"`
	UserID    *string    `json:"userID"`
	FriendID  *string    `json:"friendID"`
}

// UpdateUserInput is used for update User object.
// Input was generated by ent.
type UpdateUserInput struct {
	Name             *string                `json:"name"`
	Username         *string                `json:"username"`
	Password         *string                `json:"password"`
	ClearPassword    *bool                  `json:"clearPassword"`
	RequiredMetadata map[string]interface{} `json:"requiredMetadata"`
	Metadata         map[string]interface{} `json:"metadata"`
	ClearMetadata    *bool                  `json:"clearMetadata"`
	AddGroupIDs      []string               `json:"addGroupIDs"`
	RemoveGroupIDs   []string               `json:"removeGroupIDs"`
	ClearGroups      *bool                  `json:"clearGroups"`
	AddFriendIDs     []string               `json:"addFriendIDs"`
	RemoveFriendIDs  []string               `json:"removeFriendIDs"`
	ClearFriends     *bool                  `json:"clearFriends"`
}

// Properties by which OneToMany connections can be ordered.
type OneToManyOrderField string

const (
	OneToManyOrderFieldName OneToManyOrderField = "NAME"
)

var AllOneToManyOrderField = []OneToManyOrderField{
	OneToManyOrderFieldName,
}

func (e OneToManyOrderField) IsValid() bool {
	switch e {
	case OneToManyOrderFieldName:
		return true
	}
	return false
}

func (e OneToManyOrderField) String() string {
	return string(e)
}

func (e *OneToManyOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OneToManyOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OneToManyOrderField", str)
	}
	return nil
}

func (e OneToManyOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
