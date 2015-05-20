# aCdC
It is a simple software that check if my AC cable is plugged on my pc. If not, appears an alert and send an email to me. If the percentage of the battery is lower than THRESHOLD constant, it shutdown my pc safely.

## Configuration
Edit configuration file: `src/config.go`

## Goopher Newbie
I'm a newbie about golang so if you have suggestions, please, let me know!

## Requirements
* acpiTool
* xmessage
* a Gmail account

## Crontab
I added the binary to my crontab file, every 10 minutes cron check if my AC cable is plugged on

```sh
10 * * * * DISPLAY=:0.0 /path/to/file
```
## License
MIT

## Author
* Domenico Luciani
* domenicoleoneluciani@gmail.com
* http://dlion.it

