# mocacinno.com
go version of mocacinno.com

Currently in pre-alpha stage... Defenately NOT production ready!!!

You're free to patch anything you'd like and create pull requests in case you want added functionality at an increased pace.

The code is free, if you have your own node running and want a quick and easy gui, don't hesistate to use the code... It should be OK to run out of the box (after updating the config.ini), but maybe it's a good idear to modify the layout.html :)

## How to use/install
* Make sure you have the latest version of go (i think you need go > 1.11, but i'm using 1.15 to test everything)
* clone this repo (duh)
* edit config.ini, and in main.go change the startTLS line to use YOUR certificates instead of mine
* All the templates and static files can be found in ./views and ./staticfiles. Everything should work out of the box, but adding a personal touch might be nice
* to start the server, run "go run main.go"
* you'll probably be missing some dependencys... When running the code these dependencys will be shown in STDOUT, just do: "go get github.com/dependency"
* once you've confirmed everything works, you can compile the code: "go build main.go"
* you can create a systemd service to restart the server on crash... If you do, don't forget to submit the code to me ;). IF nobody made a service file BEFORE i reach the actual beta-stage, i'll write one myself

## Credits and tipjar
BTW: i'd still like to be mentioned in the credits (maybe a link to mocacinno.com) in case you run my code... And if you run it for profit, a tip to my tipjar would certainly be appreciated...

Tipjar: BTC/BCH: 1MocACiWLM8bYn8pCrYjy6uHq4U3CkxLaa  ETH: 0xD65400079AdAba8cCcaD4f3D20D21f13B6240514
