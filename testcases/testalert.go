package main

import (
        "fmt"
//      slack "go.scope.charter.com/rabbot/slack"
        slack "rabbot/slack"
)

var TAGME = "<@Sandeep Reddy>"
var SLACK_CHANNEL = "test-alerts"
func main (){
        //fmt.Println("hello jaffa")
        slack.PostMessage(slack.Message{Channel: SLACK_CHANNEL, Message: fmt.Sprintf("TEST %s", TAGME)})
	//slack.PostMessage(slack.Message{Channel: SLACK_CHANNEL, Message: fmt.Sprintf("TEST @Sandeep Reddy")})
}

