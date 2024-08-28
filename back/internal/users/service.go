package users

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/nallanos/fire/internal/db"
	"github.com/nallanos/fire/internal/errdefs"
	"github.com/nrednav/cuid2"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	db *db.Queries
}

func NewService(db *db.Queries) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetUserById(ctx context.Context, id string) (db.User, error) {
	user, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.User{}, errdefs.ErrNotFound
		}
		return db.User{}, fmt.Errorf("error while getting user by id in bdd: %w", err)
	}
	return user, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	user, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return db.User{}, errdefs.ErrNotFound
		}
		return db.User{}, fmt.Errorf("error while getting user by email in bdd: %w", err)
	}
	return user, nil
}
func (s *Service) GetSession(ctx context.Context, token string) (db.Token, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(token))
	if err != nil {
		return db.Token{}, fmt.Errorf("error while writing token in hash: %w", err)
	}

	hashValue := hash.Sum([]byte{})
	token = string(hashValue)
	session, err := s.db.GetSession(ctx, token)
	if err != nil {
		return db.Token{}, fmt.Errorf("error while getting session by token: %w", err)
	}
	return session, nil
}
func (s *Service) CreateUser(ctx context.Context, name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error while generating hash from password: %w", err)
	}

	// Verifying if user already exist
	_, err = s.GetUserByEmail(ctx, email)
	if err != nil && err != errdefs.ErrNotFound {
		return fmt.Errorf("error while getting user by email: %w", err)
	}
	err = s.db.CreateUser(ctx, db.CreateUserParams{
		ID:       cuid2.Generate(),
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	})

	if err != nil {
		return fmt.Errorf("error while creating user in database : %w", err)
	}

	return nil
}

func (s *Service) verifyAuth(ctx context.Context, email, password string) (bool, db.User, error) {
	user, err := s.GetUserByEmail(ctx, email)
	if err != nil {
		if err == errdefs.ErrNotFound {
			return false, db.User{}, fmt.Errorf("error while getting user by email: %w", err)
		}

		return false, db.User{}, fmt.Errorf("error while getting user by email: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, db.User{}, fmt.Errorf("error while comparing hash and password: %w", err)
	}

	return true, db.User{}, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	ok, user, err := s.verifyAuth(ctx, email, password)
	if err != nil {
		return "", fmt.Errorf("error while verifying auths: %w", err)
	}

	if !ok {
		return "", errdefs.ErrUnauthorized
	}

	token, err := s.createSession(ctx, user.ID)
	if err != nil {
		return "", fmt.Errorf("error while creating session: %w", err)
	}
	return token, nil
}

func (s *Service) createSession(ctx context.Context, userID string) (string, error) {
	token, err := generateToken()
	if err != nil {
		return "", fmt.Errorf("error while generating token : %w", err)
	}

	hash := sha256.New()

	_, err = hash.Write([]byte(token))
	if err != nil {
		return "", fmt.Errorf("error while writing token in hash: %w", err)
	}
	hashValue := hash.Sum([]byte{})

	err = s.db.CreateSession(ctx, db.CreateSessionParams{
		UserID:    userID,
		Token:     string(hashValue),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	})

	if err != nil {
		return "", fmt.Errorf("error while creating user session: %w", err)
	}

	return token, nil
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("error while generating token:%w", err)
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func (s *Service) GetToken(ctx context.Context, email string) (db.Token, error) {
	userID, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		return db.Token{}, fmt.Errorf("error while getting usser by email in database: %w", err)
	}
	token, err := s.db.GetSessionByUserId(ctx, userID.ID)
	if err != nil {
		return db.Token{}, fmt.Errorf("error while getting session by userID in database: %w", err)
	}
	return token, nil
}
