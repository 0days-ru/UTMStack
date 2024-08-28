package service

import (
	"github.com/0days-ru/UTMStack/agent-manager/models"
	"github.com/0days-ru/UTMStack/agent-manager/repository"
)

type AgentModuleService struct {
	repo *repository.AgentModuleRepository
}

func NewAgentModuleService() *AgentModuleService {
	repo := repository.NewAgentModulesRepository()
	return &AgentModuleService{repo: repo}
}

func (s *AgentModuleService) UpdateModuleConfig(configs []*models.AgentModuleConfiguration) error {
	return s.repo.UpdateAgentModule(configs)
}

func (s *AgentModuleService) FindByID(id uint) (*models.AgentModule, error) {
	return s.repo.FindByModuleId(id)
}

func (s *AgentModuleService) FindAll() ([]*models.AgentModule, error) {
	return s.repo.FindAll()
}
