package handlers

//Services for handle IAPs
type Services interface {
	Verify() map[string]interface{}
}
