package services

type MessageProcessor struct {
	messageService MessageService
}

func (mp MessageProcessor) Process(msg string) bool {
	return mp.messageService.Send(msg) // comment this line to simulate "failure" in the tests
	// return true
}

func NewMessageProcessor(msgService MessageService) MessageProcessor {
	return MessageProcessor{
		messageService: msgService,
	}
}
