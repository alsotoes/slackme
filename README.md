# slackme

A little web app that helps to deploy a tiny web portal so users can auto-invite themselves to Slack workspaces.

# Motivation

I work with several open-source communities, and many of them started using IRC. It's also impossible to work against people using one app or another; people use whatever they want; Slack, IRC, Discord, Telegram, etc., and many communities out there don't have the money to spend on Slack. Slack It's fantastic, but expensive for local communities.

Also, several well-known apps doing the same started getting deprecated, and after analyzing the needed fixes I preferred to start from scratch following the up-to-date Slack API documentation.
- [slackin](https://github.com/rauchg/slackin) 
- [slack-invite-automation](https://github.com/outsideris/slack-invite-automation)

# Disclaimer

This project is heavily based on the [slackin](https://github.com/rauchg/slackin) project, but reimplemented using Go and focused on long-term support, from user libraries to testing pipelines.
- [gin-gonic](https://github.com/gin-gonic/gin)
- [slack-go](https://github.com/slack-go/slack)

## How to use

# Parameters

All parameters can alternatively be configured through environment variables.

| Flag | Short | Environment variable | Default | Description |
| --- | --- | --- | --- | --- |
| `SLACK_API_TOKEN` | **Required** | [API token](https://api.slack.com/tutorials/tracks/getting-a-token) |
| `GOOGLE_CAPTCHA_SECRET` | `''` | reCAPTCHA secret |
| `GOOGLE_CAPTCHA_SITEKEY` | `''` | reCAPTCHA sitekey |
| `SLACKME_COC` | `''` | Full URL to a CoC that needs to be agreed to |
| `PORT` | `3000` | Port to listen on, 3000 by default|
| `DEBUG` | `true` | Enable debug log output, false by default|

Alternatively, you can specify the configuration parameters in a `.env` file in the root directory of your project and add environment-specific variables on new lines in the form of NAME=VALUE. For example:

```
GOOGLE_CAPTCHA_SECRET="your_secret"
GOOGLE_CAPTCHA_SITEKEY="your_site_key"
SLACK_API_TOKEN="xoxb-not-a-real-token-this-will-not-work"
SLACKME_COC="url_pointing_to_your_code_of_condut"
PORT=8080
DEBUG=true
```

#### Deploy on any server
I'll add the steps later on

#### Deploy on a containers
I'll add the steps later on

#### Deploy on Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/alsotoes/slackme/tree/master)

  

## Credits

 - The [slackin](https://github.com/alsotoes/slackin) project as I used the look and feel.  
 - Inhered from the slackin project:  
	 - The SVG badge generation was taken from the excellent [shields](https://github.com/badges/shields) project.  
	 - The button CSS is based on [github-buttons](https://github.com/mdo/github-buttons)


## Licenses (both included in the repo)

- slackme is under GPL-3
- slackin is under MIT
