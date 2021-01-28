package client

type SubFilterConnective string

type FilterRelation string

const (
	SubFilterConnectiveOr  SubFilterConnective = "OR"
	SubFilterConnectiveAnd SubFilterConnective = "AND"

	FilterRelationEqual        FilterRelation = "equal"
	FilterRelationUnequal      FilterRelation = "unequal"
	FilterRelationGreater      FilterRelation = "greater"
	FilterRelationLess         FilterRelation = "less"
	FilterRelationGreaterEqual FilterRelation = "greaterEqual"
	FilterRelationLessEqual    FilterRelation = "lessEqual"
)

type RequestFilter struct {
	Field    string          `json:"field,omitempty"`
	Value    string          `json:"value,omitempty"`
	Relation *FilterRelation `json:"relation,omitempty"`

	SubFilterConnective *SubFilterConnective `json:"subFilterConnective,omitempty"`
	SubFilter           *[]RequestFilter     `json:"subFilter,omitempty"`
}
