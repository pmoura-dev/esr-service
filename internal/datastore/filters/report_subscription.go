package filters

import (
	"time"

	"github.com/pmoura-dev/esr-service/internal/types"
)

type ReportSubscriptionFilter struct {
	entityID      *string
	reportType    *types.ReportType
	isActive      *bool
	updatedAfter  *time.Time
	updatedBefore *time.Time
}

func NewReportSubscriptionFilter() *ReportSubscriptionFilter {
	return &ReportSubscriptionFilter{}
}

func (f *ReportSubscriptionFilter) ByEntityID(entityID string) *ReportSubscriptionFilter {
	f.entityID = &entityID
	return f
}

func (f *ReportSubscriptionFilter) ByReportType(reportType types.ReportType) *ReportSubscriptionFilter {
	f.reportType = &reportType
	return f
}

func (f *ReportSubscriptionFilter) ByIsActive(isActive bool) *ReportSubscriptionFilter {
	f.isActive = &isActive
	return f
}

func (f *ReportSubscriptionFilter) ByTimeAfterUpdated(threshold time.Time) *ReportSubscriptionFilter {
	f.updatedAfter = &threshold
	return f
}

func (f *ReportSubscriptionFilter) ByTimeBeforeUpdated(threshold time.Time) *ReportSubscriptionFilter {
	f.updatedBefore = &threshold
	return f
}

func (f *ReportSubscriptionFilter) Check(subscription types.ReportSubscription) bool {
	if f.entityID != nil && *f.entityID != subscription.EntityID {
		return false
	}

	if f.reportType != nil && *f.reportType != subscription.ReportType {
		return false
	}

	if f.isActive != nil && *f.isActive != subscription.IsActive {
		return false
	}

	if f.updatedAfter != nil && !subscription.UpdatedAt.After(*f.updatedAfter) {
		return false
	}

	if f.updatedBefore != nil && !subscription.UpdatedAt.Before(*f.updatedBefore) {
		return false
	}

	return true
}
