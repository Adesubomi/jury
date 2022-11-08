package jury

type Rules []Rule

type Rule interface {
	Validate(string) string
}
