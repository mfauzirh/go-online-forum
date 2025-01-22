package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/mfauzirh/go-online-forum/internal/model/memberships"
	"github.com/mfauzirh/go-online-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, refreshToken memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not exists")
	}

	token, err := jwt.CreateToken(userID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}

	return token, nil
}
