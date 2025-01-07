module esr_service

open util/ordering[Time]

sig Time {}

sig User {
	issues: set Command
}

abstract sig CommandStatus {}
one sig Pending, Success, Failed extends CommandStatus {}

sig Command {
	target: one Entity,
	desired_state: one State,
	status: one CommandStatus,
	issued_at: one Time,
	timeout_at: one Time
}

sig Entity {
	current_state: one State
}


sig State {}

sig StateUpdate {
	source: one Entity,
	reported_state: one State,
	updated_at: one Time
}



pred in_time[t: Time, lower: Time, upper: Time] {
	lte[lower, t] and lte[t, upper]
}

// assertions
pred CommandFailsAfterTimeout {
	all c: Command | no u: StateUpdate |
		c.target = u.source and c.desired_state = u.reported_state and
		in_time[u.updated_at, c.issued_at, c.timeout_at]
		implies c.status = Success
}


fact {
	// all commands are issued by an user
	Command in User.issues
}

fact {
	// initially, there are no commands nor state updates
	no Command
	no StateUpdate
}

run {} for 3 but exactly 1 User, exactly 1 Entity
