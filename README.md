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

## environment

| env                   | description                                                                      |
| --------------------- | -------------------------------------------------------------------------------- |
| PUSHOVER_ID           | The unique id of the message, relative to this device.                           |
| PUSHOVER_UMID         | The unique id of the message relative to all devices on the same user's account. |
| PUSHOVER_MESSAGE      | The text of the message.                                                         |
| PUSHOVER_TITLE        | The title of the message                                                         |
| PUSHOVER_PRIORITY     | The priority of the message.(-2/-1/0/1/2)                                        |
| PUSHOVER_PRIORITY_STR | The priority of the message. (lowest/low/normal/high/emergency)                  |
| PUSHOVER_APP          | The name of the application that sent the message. This may not be unique.       |
| PUSHOVER_AID          | The unique id of the application that sent the message.                          |
