# aCdC
It is a simple software that checks if my AC cable is plugged on my pc.   
If not, it shows an alert and sends me an email with a custom body.   
If battery percentage is lower than the THRESHOLD constant, it shuts my pc down safely.   
You must to be root to shutdown your pc.

## Configuration
Edit configuration file: `src/config.go`

## Why in GoLang ?
Because I want to learn golang.

## Gopher Newbie
I'm a newbie about golang so if you have suggestions, please, let me know!

## Requirements
* acpiTool
* xmessage
* a Gmail account

## Crontab
I added the binary to my root crontab file, every 5 minutes cron checks if my AC cable is plugged on pc

```sh
*/5 * * * * DISPLAY=:0.0 /path/to/file
```
## License
MIT

## Author
* Domenico Luciani
* domenicoleoneluciani@gmail.com
* http://dlion.it

