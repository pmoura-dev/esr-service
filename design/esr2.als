module esr2

open util/ordering[Time]

sig User {
	var issues: set Command
}



sig Time {}



enum CommandStatus {Pending, Success, Failure}

sig Command {
	target: one Entity,
	desired_state: one State,

	var status: one CommandStatus,
	
	issued_at: one Time,
	timeout_at: one Time
}

sig Entity {}

sig State {}

sig StateReport {
	source: one Entity,
	reported_state: one State,
	reported_at: one Time
}







// operations
pred stutter {
	issues' = issues

	target' = target
	desired_state' = desired_state

	status' = status

	issued_at' = issued_at
	timeout_at' = timeout_at
}


pred new_command {
	some u: User, c: Command, e: Entity, t: Time | 
		issue_command[u, c, e, t]
}

pred issue_command[u: User, c: Command, e: Entity, t: Time] {
	c not in User.issues
	c.target = e
	c.status = Pending
	c.issued_at = t

	u.issues' = u.issues + c

	(User - u).issues' = (User - u).issues
	target' = target
	desired_state' = desired_state
	status' = status
	issued_at' = issued_at
	timeout_at' = timeout_at
}

pred state_report {
	
}




// initial state
fact {
	no issues
	all c: Command | c.status = Pending
}

fact UserCantIssueTwoCommandsAtTheSameTime {
	all u: User {
		all t1, t2: u.issues.issued_at | t1 != t2
	}
}




fact {
	always (
		stutter or
		new_command
	)
}

assert Prop1 {
	always (all c: User.issues | c.status = Pending)
}


run {
	#status.Pending = 3
	eventually (all c: Command | c.status = Pending implies c in User.issues)
} for 5 but exactly 1 User






