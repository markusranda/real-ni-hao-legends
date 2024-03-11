interface Command {
    // multiplayer commands to be sent to the server
    // these are the commands that the server will understand
    // and execute
    // using these commands, the server with mutate the backend state and send an updated state to the clients
    type: string
}
