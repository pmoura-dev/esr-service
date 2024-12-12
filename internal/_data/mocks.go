// package that includes mock values to be used in testing
package _data

var (
	MockEntity1       = `{"id": "1", "name": "TestEntity1"}`
	MockEntity2       = `{"id": "2", "name": "TestEntity2"}`
	MockEntityInvalid = `{"id": "3", "name": "TestEnti`

	MockCommand1Pending = `{
		"id": "cmd1",
		"entity_id": "1",
		"desired_state": {
			"power": "on"
		},
		"status": "pending",
		"issued_at": "2009-11-10T23:00:00Z"
	}`
	MockCommand1Success = `{
		"id": "cmd2",
		"entity_id": "1",
		"desired_state": {
			"power": "off"
		},
		"status": "success",
		"issued_at": "2010-11-10T23:00:00Z",
		"resolved_at": "2010-11-10T23:00:10Z"
	}`
	MockCommand2Failed = `{
		"id": "cmd3",
		"entity_id": "2",
		"desired_state": {
			"power": "off"
		},
		"status": "failed",
		"issued_at": "2011-11-10T23:00:00Z",
		"resolved_at": "2011-11-10T23:00:10Z"
	}`
	MockCommandInvalid = `{
		"status": "random"
	}`

	MockReportSubscription1State = `{
		"id": 1,
		"entity_id": "entity_1",
		"report_type": "state",
		"is_active": true,
		"updated_at": "2009-11-10T23:00:00Z"
	}`
	MockReportSubscription1MetricPower = `{
		"id": 2,
		"entity_id": "entity_1",
		"report_type": "metric",
		"metric": "power",
		"is_active": false,
		"updated_at": "2009-11-10T23:00:00Z"
	}`
	MockReportSubscription2State = `{
		"id": 3,
		"entity_id": "entity_2",
		"report_type": "state",
		"is_active": true,
		"updated_at": "2011-11-10T23:00:00Z"
	}`
	MockReportSubscriptionInvalid = `{
		"report_type": "random"
	}`
)
