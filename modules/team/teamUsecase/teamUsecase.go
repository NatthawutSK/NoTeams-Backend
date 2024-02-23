package teamUsecase

import (
	"github.com/NatthawutSK/NoTeams-Backend/config"
	"github.com/NatthawutSK/NoTeams-Backend/modules/team"
	"github.com/NatthawutSK/NoTeams-Backend/modules/team/teamRepository"
)

type ITeamUsecase interface {
	CreateTeam(req *team.CreateTeamReq) (*team.CreateTeamRes, error)
}

type teamUsecase struct {
	teamRepo teamRepository.ITeamRepository
	cfg      config.IConfig
}

func TeamUsecase(teamRepo teamRepository.ITeamRepository, cfg config.IConfig) ITeamUsecase {
	return &teamUsecase{
		teamRepo: teamRepo,
		cfg:      cfg,
	}
}

func (u *teamUsecase) CreateTeam(req *team.CreateTeamReq) (*team.CreateTeamRes, error) {
	result, err := u.teamRepo.CreateTeam(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
