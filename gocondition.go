package gocondition

type Rule int

type IRule interface {
	Value(context interface{}) bool
}

const (
	RuleAnd = Rule(1)
	RuleOr  = Rule(2)
)

type RuleSet struct {
	Type     Rule
	Children []IRule
}

func (set *RuleSet) IsOr() bool {

	return set.Type == RuleOr
}

func (set *RuleSet) IsAnd() bool {

	return set.Type == RuleAnd
}

func (set *RuleSet) Value(context interface{}) bool {

	if set.Type == RuleOr {

		value := false

		for _, child := range set.Children {

			value = value || child.Value(context)
		}
		return value

	} else {

		value := true

		for _, child := range set.Children {

			value = value && child.Value(context)
		}
		return value
	}
}
