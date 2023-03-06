package firebase

import (
	"context"
	"religion/config"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

const FirebaseAuth string = "FIREBASE"

func New(ctx context.Context, cfg *config.Config) (*auth.Client, error) {
	opt := option.WithCredentialsFile(cfg.Firebase.Path)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	return app.Auth(ctx)
}
