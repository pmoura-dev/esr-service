module hauto/esr

open util/ordering[Time]

sig Time {}
sig State {}
sig Entity {}

enum CommandStatus {Pending, Success, Failure}

sig Command {
	target: Entity,
	desiredState: State,
	status: CommandStatus,
	issuedAt: Time,
	timeoutAt: Time
}



pred showCommands {

}

run showCommands
