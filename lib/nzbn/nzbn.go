// BusinessNZ NZBN API
// https://api.business.govt.nz/api/apis/info?name=NZBN

package nzbn

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/ryankurte/go-businessnz/lib/base"
)

// NZBN API Implementation
type NzbnApi struct {
	*base.Base
}

type BusinessEntity struct {
	EntityName              string        `json:"entityName"`
	EntityStatusCode        string        `json:"entityStatusCode"`
	Nzbn                    string        `json:"nzbn"`
	EntityTypeCode          string        `json:"entityTypeCode"`
	EntityTypeDescription   string        `json:"entityTypeDescription"`
	EntityStatusDescription string        `json:"entityStatusDescription"`
	TradingNames            []TradingName `json:"tradingNames"`
}

type TradingName struct {
	UniqueIdentifier string `json:"uniqueIdentifier"`
	Name             string `json:"name"`
}

type SearchQuery struct {
	SearchTerm   string       `url:"search-term"`
	EntityStatus EntityStatus `url:"entity-status,omitempty"`
	EntityType   EntityType   `url:"entity-type,omitempty"`
	IndustryCode string       `url:"industry-code,omitempty"`
	Page         uint         `url:"page,omitempty"`
	PageSize     uint         `url:"page,omitempty"`
}

type EntityStatus string

const (
	Registered                EntityStatus = "Registered"
	VoluntaryAdministration   EntityStatus = "VoluntaryAdministration"
	InReceivership            EntityStatus = "InReceivership"
	InLiquidation             EntityStatus = "InLiquidation"
	InStatutoryAdministration EntityStatus = "InStatutoryAdministration"
	Inactive                  EntityStatus = "Inactive"
	RemovedClosed             EntityStatus = "RemovedClosed"
)

type EntityType string

const (
	EntityTypeNZCompany                     EntityType = "NZCompany"
	EntityTypeOverseasCompany               EntityType = "OverseasCompany"
	EntityTypeSoleTrader                    EntityType = "SoleTrader"
	EntityTypePartnership                   EntityType = "Partnership"
	EntityTypeTrust                         EntityType = "Trust"
	EntityTypeBuildingSociety               EntityType = "BuildingSociety"
	EntityTypeCharitableTrust               EntityType = "CharitableTrust"
	EntityTypeCreditUnion                   EntityType = "CreditUnion"
	EntityTypeFriendlySociety               EntityType = "FriendlySociety"
	EntityTypeIncorporatedSociety           EntityType = "IncorporatedSociety"
	EntityTypeIndustrialAndProvidentSociety EntityType = "IndustrialAndProvidentSociety"
	EntityTypeLimitedPartnershipNz          EntityType = "LimitedPartnershipNz"
	EntityTypeLimitedPartnershipOverseas    EntityType = "LimitedPartnershipOverseas"
	EntityTypeSpecialBodies                 EntityType = "SpecialBodies"
	EntityTypeSpecialBody                   EntityType = "SpecialBody"
	EntityTypeTrading_Trust                 EntityType = "Trading_Trust"
	EntityTypeSole_Trader                   EntityType = "Sole_Trader"
	EntityTypeB                             EntityType = "B"
	EntityTypeI                             EntityType = "I"
	EntityTypeD                             EntityType = "D"
	EntityTypeF                             EntityType = "F"
	EntityTypeN                             EntityType = "N"
	EntityTypeS                             EntityType = "S"
	EntityTypeT                             EntityType = "T"
	EntityTypeY                             EntityType = "Y"
	EntityTypeZ                             EntityType = "Z"
	EntityTypeGovtCentral                   EntityType = "GovtCentral"
	EntityTypeGovtEdu                       EntityType = "GovtEdu"
	EntityTypeGovtLocal                     EntityType = "GovtLocal"
	EntityTypeGovtOther                     EntityType = "GovtOther"
	EntityTypeLTD                           EntityType = "LTD"
	EntityTypeULTD                          EntityType = "ULTD"
	EntityTypeCOOP                          EntityType = "COOP"
	EntityTypeASIC                          EntityType = "ASIC"
	EntityTypeNON_ASIC                      EntityType = "NON_ASIC"
)

type SearchResult struct {
	PageSize   uint             `json:"pageSize"`
	Page       uint             `json:"page"`
	TotalItems uint             `json:"totalItems"`
	SortBy     string           `json:"sortBy"`
	SortOrder  string           `json:"sortOrder"`
	Items      []BusinessEntity `json:"items"`
}

func (a *NzbnApi) Search(q SearchQuery) (*SearchResult, error) {
	var data SearchResult

	v, err := query.Values(q)
	if err != nil {
		return nil, fmt.Errorf("failed to build query string: %v", err)
	}

	err = a.Query("services/v4/nzbn/entities?"+v.Encode(), &data)
	if err != nil {
		return nil, fmt.Errorf("NZBN Api request failed: %v", err)
	}

	return &data, nil
}

func (a *NzbnApi) Lookup(nzbn string) (*BusinessEntity, error) {

	var entity BusinessEntity

	// Search by NZBN
	err := a.Query("services/v4/nzbn/entities/"+nzbn, &entity)
	if err != nil {
		return nil, fmt.Errorf("NZBN Api request failed: %v", err)
	}

	return &entity, nil
}
