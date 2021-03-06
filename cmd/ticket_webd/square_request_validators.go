package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	root "github.com/timhugh/ticket_service"
)

type SquareRequest struct {
	Body      string
	Signature string
	URL       string
	Event     Event
}

type SquareRequestValidator struct {
	LocationRepository locationFinder
}

type locationFinder interface {
	Find(string) (*root.Location, error)
}

func (s SquareRequestValidator) Validate(req *SquareRequest) error {
	location, err := s.LocationRepository.Find(req.Event.LocationID)
	if err != nil {
		return fmt.Errorf("Error finding location '%s': %s", req.Event.LocationID, err)
	}

	if !s.verifySignature(req.URL, location.SignatureKey, req.Body, req.Signature) {
		return fmt.Errorf("Request failed signature validation")
	}

	return nil
}

func (s SquareRequestValidator) verifySignature(webhookURL string, webhookSignatureKey string, body string, signature string) bool {
	mac := hmac.New(sha1.New, []byte(webhookSignatureKey))
	mac.Write([]byte(webhookURL + body))
	expectedMAC := mac.Sum(nil)
	expectedSignature := base64.StdEncoding.EncodeToString([]byte(expectedMAC))
	return expectedSignature == signature
}
