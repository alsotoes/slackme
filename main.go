package main

import (
  "os"
  "fmt"
  "strconv"
  "strings"
  "net/http"
  "html/template"

  "github.com/gin-gonic/gin"
  "github.com/slack-go/slack"
)

type ServerParameters struct {
  SLACKME_COC string
  SLACKME_PORT string
  SLACKME_DEBUG string
  SLACK_API_TOKEN string
  GOOGLE_CAPTCHA_SECRET string
  GOOGLE_CAPTCHA_SITEKEY string
}

var defaultParameters = ServerParameters {
  SLACKME_PORT: "3000",
  SLACKME_DEBUG: "false",
}

func DefaultIfNotSet(key string, defaultValue string) string {
  var parameter string
  value, set := os.LookupEnv(key)
  if !set {
    parameter = defaultValue
  } else {
    parameter = value
  }
  return parameter
}

func main() {
  var slackme = ServerParameters{
    SLACKME_COC: DefaultIfNotSet("SLACKME_COC",defaultParameters.SLACKME_COC),
    SLACKME_PORT: DefaultIfNotSet("PORT",defaultParameters.SLACKME_PORT),
    SLACKME_DEBUG: DefaultIfNotSet("DEBUG",defaultParameters.SLACKME_DEBUG),
    SLACK_API_TOKEN: DefaultIfNotSet("SLACK_API_TOKEN",defaultParameters.SLACK_API_TOKEN),
    GOOGLE_CAPTCHA_SECRET: DefaultIfNotSet("GOOGLE_CAPTCHA_SECRET",defaultParameters.GOOGLE_CAPTCHA_SECRET),
    GOOGLE_CAPTCHA_SITEKEY: DefaultIfNotSet("GOOGLE_CAPTCHA_SITEKEY",defaultParameters.GOOGLE_CAPTCHA_SITEKEY),
  }

  defer func() {
    fmt.Printf("%+v\n",slackme)
  }()

  debug, err := strconv.ParseBool(slackme.SLACKME_DEBUG)
  if err != nil || debug {
    gin.SetMode(gin.DebugMode)
  } else {
    gin.SetMode(gin.ReleaseMode)
  }

  router := gin.New()
  router.Use(gin.Logger())
  router.Use(gin.Recovery())
  router.SetFuncMap(template.FuncMap{
    "upper": strings.ToUpper,
  })

  // Create a new Slack API client
  api := slack.New(slackme.SLACK_API_TOKEN)

  // Call the users.list method to retrieve a list of all users in the workspace
  userList, err := api.GetUsers()
  if err != nil {
    panic(err)
  }

  // Iterate over the list of users but only count
  // active and real users
  var users int
  for _, user := range userList {
    if !user.Deleted && !user.IsBot {  
      users++
    }
  }

  // Call the team.teamInfo method to retrieve the team's information
  teamInfo, err := api.GetTeamInfo()
  if err != nil {
    panic(err)
  }

  router.Static("/assets", "./assets")
  router.LoadHTMLGlob("templates/*.tmpl")

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "team_members": users-1, // -1 as slackbot is not considered a bot
      "team_name": teamInfo.Name,
      "team_domain": teamInfo.Domain,
      "team_coc": slackme.SLACKME_COC,
      "team_logo": teamInfo.Icon["image_132"],
      "reCAPTCHA_site_key": slackme.GOOGLE_CAPTCHA_SITEKEY,
    })
  })

  router.Run(":" + slackme.SLACKME_PORT)
}
