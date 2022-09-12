package main

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
)

type User struct {
	Name  string `json:"name" validate:"required,alpha"`
	Email string `json:"email" validate:"required,email"`
}

func main() {
	// OOP -> Class => Object
	// OOP -> Constructor is a special method in a Class
	// let v: Validator = new Validator(args);
	// Constructor functions in Go are functions that return an object.
	// NewHelmUpdater() -> Constructor => A certain implementation (Updater using Helm)
	v := validator.New()

	emailService := EmailService{
		validator: v,
	}

	userService := UserService{
		validator:    v,
		emailService: &emailService,
	}

	u := User{
		Name:  "Marcos",
		Email: "marcos@huck.com.ar",
	}

	u, err := userService.Create(context.Background(), u)
	log.Println("Error:", err)

	b, err := json.Marshal(u)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	log.Println("JSON:", string(b))
}

// UserService manages users.
// Service: Contains business logic
type UserService struct {
	validator   *validator.Validate
	logger      *log.Logger
	emailSender EmailSender
}

func NewUserService(v *validator.Validate, logger *log.Logger, emailSender EmailSender) *UserService {
	return &UserService{
		validator:   v,
		logger:      logger,
		emailSender: emailSender,
	}
}

func (svc *UserService) Create(ctx context.Context, u User) (User, error) {
	// validate the user
	if err := svc.validate(ctx, u); err != nil {
		return User{}, err
	}

	// Save the record in the database
	// And return the new representation of the user
	// Or do some other things
	if err := svc.emailSender.SendEmail(ctx, u.Email, "Congrats!!!"); err != nil {
		return User{}, err
	}

	return u, nil
}

func (svc *UserService) validate(ctx context.Context, u User) error {
	if err := svc.validator.StructCtx(ctx, u); err != nil {
		// Logging of the error
		return err
	}
	return nil
}

type EmailSender interface {
	SendEmail(ctx context.Context, email string, msg string) error
}

type FakeEmailService struct {
	validator    *validator.Validate
	emailPrinter *log.Logger
}

func (svc FakeEmailService) SendEmail(ctx context.Context, email string, msg string) error {
	err := svc.validator.Var(email, "email")
	if err != nil {

	}
	svc.emailPrinter.Println("The email has been sent!")
	return nil
}

type EmailService struct {
	validator *validator.Validate
	ses       *aws.SES
}

func (svc EmailService) SendEmail(ctx context.Context, email string, msg string) error {
	err := svc.validator.Var(email, "email")
	if err != nil {

	}
	svc.ses.SendEmail(ctx, email, msg)
	return nil
}

func NewEmailService(v *validator.Validate) *EmailService {
	return &EmailService{validator: v}
}
