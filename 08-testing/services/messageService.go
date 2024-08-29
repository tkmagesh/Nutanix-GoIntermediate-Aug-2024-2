package services

//go:generate mockery --name MessageService
type MessageService interface {
	Send(msg string) bool
}
