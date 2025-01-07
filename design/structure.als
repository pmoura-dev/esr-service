module esr/structure

open util/ordering[Time]

sig Time {}





// Entities
sig Entity {
	state: State,
	metrics: some Metric,
	reported_at: Time,
}


sig State {}

sig Metric {}


// Commands
sig Command {
	target: Entity,
	desired_state: State,
	status: CommandStatus,
	issued_at: Time,
	timeout_at: Time
}

enum CommandStatus {Pending, Success, Failure}

fact {
	all c: Command | c.issued_at.lt[c.timeout_at]
}

// Reports
abstract sig Report {}

sig StateReport extends Report {
	source: Entity,
	reported_state: State,
	reported_at: Time
}

sig MetricReport extends Report {
	source: Entity,
	reported_metric: Metric,
	reported_at: Time
}







run {} for 5 
