module esr/behavior

open esr/structure


sig Component {
	issues: set Command
}

sig Broker {
	reports: set Report
}

pred issueCommand ()


pred Stutter {
	univ' = univ
}


pred NewCommand[
	c: Component,
	e: Entity,
	s: State,
	t1, t2: Time
]{
	some cmd: Command {
		cmd.target = e
		cmd.desired_state = s
		cmd.status = Pending
		cmd.issued_at = t1
		cmd.timeout_at = t2
	}

	
}





// initial state
fact {
	no issues
	no reports

	Command.status = Pending
}

fact {
	
	always (
		Stutter or
		some e: Entity, s: State, t1, t2: Time | NewCommand[e, s, t1, t2]
	)
}



run {
	eventually (all c: Command | c.status = Pending implies c in Component.issues)
} for 3






