package service

import "github.com/jmoiron/sqlx"

type ProbabilityLinkService struct {
	DB *sqlx.DB
}

func NewProbabilityLinkService(db *sqlx.DB) *ProbabilityLinkService {
	return &ProbabilityLinkService{
		DB: db,
	}
}

func (p *ProbabilityLinkService) GetProbabilityLink(shortCode string) (string, error) {
	var link string
	err := p.DB.Get(&link, "SELECT link FROM probability_links WHERE short_code=$1", shortCode)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (p *ProbabilityLinkService) CreateProbabilityLink(link_a string, link_b, probability float64) (string, error) {

}
