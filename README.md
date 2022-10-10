# push-watch

## About push-watch

push-watch is the client (receiver of notifications) of pushover.  
Each time a notification is received by pushover, the specified command is executed. The title, message, etc. can be obtained from environment variables.  
Since push-watch is executed as an open client, a [desktop license](https://pushover.net/clients/desktop) is required.

## Getting Started

```bash
# Login and get secret/deviceID

$ push-watch login {username} {password}
Success!
Device ID: {deviceID}
Device Secret: {deviceSecret}

$ push-watch watch {deviceID} {deviceSecret} -- sh -c 'echo $PUSHOVER_TITLE\\n$PUSHOVER_MESSAGE'
# -> The command is executed each time a new message is retrieved
```
