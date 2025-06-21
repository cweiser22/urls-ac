package service

import (
	"crypto/rand"
	"fmt"
	"github.com/cweiser22/urls-ac/internal/models"
	"github.com/cweiser22/urls-ac/internal/repository"
	"math/big"
)

type FiftyFiftyLinkService struct {
	repo *repository.FiftyFiftyLinkRepository
}

func NewFiftyFiftyLinkService(repo *repository.FiftyFiftyLinkRepository) *FiftyFiftyLinkService {
	return &FiftyFiftyLinkService{
		repo: repo,
	}
}

// Create creates a new FiftyFiftyLink.
func (s *FiftyFiftyLinkService) Create(probability float64, urlA, urlB, shortCode string) (*models.FiftyFiftyLink, error) {
	dto := &repository.CreateFiftyFiftyLinkDTO{
		Probability: probability,
		URLa:        urlA,
		URLb:        urlB,
		ShortCode:   shortCode,
	}

	link, err := s.repo.Insert(dto)
	if err != nil {
		return nil, fmt.Errorf("create link: %w", err)
	}
	return link, nil
}

// GetByShortCode retrieves a link by its shortcode.
func (s *FiftyFiftyLinkService) GetByShortCode(shortCode string) (*models.FiftyFiftyLink, error) {
	link, err := s.repo.GetByShortCode(shortCode)
	if err != nil {
		return nil, fmt.Errorf("get link: %w", err)
	}
	return link, nil
}

func MustRandomFloat64() float64 {
	const precision = 1 << 53
	n, err := rand.Int(rand.Reader, big.NewInt(precision))
	if err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return float64(n.Int64()) / float64(precision)
}

// GetLink randomly chooses urlA or urlB based on the probability.
func (s *FiftyFiftyLinkService) GetLink(link *models.FiftyFiftyLink) string {
	if MustRandomFloat64() < link.Probability {
		return link.URLa
	}
	return link.URLb
}
