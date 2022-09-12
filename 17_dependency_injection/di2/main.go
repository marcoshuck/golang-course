package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

// Dependency injection
// Dependency = Service A depends on Service B. User service depends on the Email service.
// Injection = We pass service B to A.
//
// Rules:
// You should only depend on abstractions, not implementations.

type User struct {
	Name  string
	Email string
}

type UserService interface {
	Create(ctx context.Context, user User) error
}

type userService struct {
	emailSvc EmailService
	logger   *log.Logger
}

func NewUserService(emailSvc EmailService, logger *log.Logger) UserService {
	return &userService{
		emailSvc: emailSvc,
		logger:   logger,
	}
}

func (svc *userService) Create(ctx context.Context, user User) error {
	// err := svc.validatorSvc.Validate(ctx, user)
	// err := svc.repository.Create(ctx, user)
	err := svc.emailSvc.SendEmail(ctx, user.Email)
	if err != nil {
		svc.logger.Println("Failed to send the email:", err)
		return err
	}
	return nil
}

type EmailService interface {
	SendEmail(ctx context.Context, email string) error
}

type emailServiceSES struct {
	// Amazon Simple Email Service
	ses interface{}
}

func NewEmailServiceSES() EmailService {
	return &emailServiceSES{
		ses: nil, // Bla bla
	}
}

func (svc *emailServiceSES) SendEmail(ctx context.Context, email string) error {
	// svc.ses.SendEmail() // Arguments)
	return nil
}

type emailServiceFake struct {
	emailsReceived []string
}

func NewEmailServiceFake() EmailService {
	return &emailServiceFake{}
}

func (svc *emailServiceFake) SendEmail(ctx context.Context, email string) error {
	svc.emailsReceived = append(svc.emailsReceived, fmt.Sprintf("Hello, %s. Welcome to the jungle!", email))
	log.Println("Emails received up to now:", svc.emailsReceived)
	return nil
}

func main() {
	logger := log.New(os.Stdout, "DI Application", log.LstdFlags)

	emails := NewEmailServiceFake()
	users := NewUserService(emails, logger)

	// Imagine: Here comes the event I was waiting for:
	err := users.Create(context.Background(), User{
		Name:  "Marcos",
		Email: "marcos@huck.com.ar",
	})
	err = users.Create(context.Background(), User{
		Name:  "Marcos",
		Email: "marcos@huck.com.ar",
	})
	err = users.Create(context.Background(), User{
		Name:  "Marcos",
		Email: "marcos@huck.com.ar",
	})
	err = users.Create(context.Background(), User{
		Name:  "Marcos",
		Email: "marcos@huck.com.ar",
	})
	err = users.Create(context.Background(), User{
		Name:  "Marcos",
		Email: "marcos@huck.com.ar",
	})
	if err != nil {
		logger.Fatalln("Error:", err)
	}

}
