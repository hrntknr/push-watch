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

## flags

```
> $ push-watch watch --help
Usage:
  push-watch watch [flags] device-id device-secret [...command]

Flags:
  -h, --help              help for watch
  -p, --priority string   Priority filter (default "-2,-1,0,1,2")

> $ push-watch login --help
Usage:
  push-watch login [flags] username password

Flags:
  -n, --device-name string   Device name (default "push-watch")
  -h, --help                 help for login
```

| mode  | flag | desc                                                                                                                                      |
| ----- | ---- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| watch | -p   | Specify the priority filter. If not specified, all priority events will be triggered.                                                     |
| login | -n   | You can specify a name for the device registration. To use the registered name again, you need to remove the device from pushover web ui. |

## attention

It is not possible to do more than one Listen at a time using the same device.
To listen to multiple processes at the same time, please register multiple devices.

```
> $ push-watch login -n device1 myuser mypassword
Success!
Device ID: device1_id
Device Secret: device1_secret

> $ push-watch login -n device2 myuser mypassword
Success!
Device ID: device2_id
Device Secret: device2_secret

> $ nohup push-watch watch device1_id device1_secret echo "Trigger1" &
> $ nohup push-watch watch device2_id device2_secret echo "Trigger2" &
```
