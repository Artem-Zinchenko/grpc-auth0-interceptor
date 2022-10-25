package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type JWTManager interface {
	Validate(ctx context.Context) (context.Context, error)
}

type auth0JwtManager struct {
	jwks string
	aud  string
	iss  string
}

// NewAuth0JWTManager returns a new jwt manager to parse and validate jwt token
func NewAuth0JWTManager(jwks string, aud string, iss string) JWTManager {
	return &auth0JwtManager{
		jwks: jwks,
		aud:  aud,
		iss:  iss,
	}
}

// Validate extracts and validate JWT token from the Authorization bearer
func (a auth0JwtManager) Validate(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil || token == "" {
		return nil, status.Errorf(codes.Unauthenticated, "bearer token is required: %v", err)
	}

	//parse and validate token
	parsedToken, err := jwt.Parse(token, a.validationKeyGetter)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "error parsing token: %w", err)
	}

	if !parsedToken.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "token is invalid")
	}

	grpc_ctxtags.Extract(ctx).Set("auth.sub", parsedToken.Claims.(jwt.StandardClaims).Audience)
	return ctx, nil
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// validationKeyGetter receives the parsed token and return the key for validating
func (a auth0JwtManager) validationKeyGetter(token *jwt.Token) (interface{}, error) {
	// Verify 'aud' claim
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(a.aud, false)
	if !checkAud {
		return token, errors.New("invalid audience")
	}
	// Verify 'iss' claim
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(a.iss, false)
	if !checkIss {
		return token, errors.New("invalid issuer")
	}

	cert, err := a.getCert(token)
	if err != nil {
		panic(err.Error())
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func (a auth0JwtManager) getCert(token *jwt.Token) (string, error) {
	cert := ""
	//request jwks key set
	resp, err := http.Get(a.jwks)
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	//decode response
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return cert, err
	}

	//find certification
	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
