package api

import (
	. "DIaaS/diaas/constants"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func NewRegistry() *Registry {
	return &Registry{
		Items:   make(map[string]*Deployment),
		Aliases: make(map[string]*Alias),
	}
}

func NewDeployment(source DeployBody, lang string) *Deployment {
	hasher := sha256.New()
	hasher.Write([]byte(source.Code))
	hasher.Write([]byte(lang))
	sum := hasher.Sum(nil)
	id := hex.EncodeToString(sum)
	return &Deployment{
		Source:      []byte(source.Code),
		ID:          id,
		Name:        source.FunctionName,
		Description: source.Description,
		Image:       fmt.Sprintf(DeploymentImage, id),
		Language:    lang,
		Date:        time.Now(),
		URL:         TriggerPrefix + id,
		User:        source.User,
		Ready:       false,
	}
}
