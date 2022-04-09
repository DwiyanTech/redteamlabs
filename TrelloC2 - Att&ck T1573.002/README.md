# Mittre ATT&CK T1573.002
## Command And Control using Trello

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

This script utilise Trello API as Command And Control to Avoid IPS/IDS and EDR.

## Get Started
-- Python Version
- Change TRELLO_APIKEY,TRELLO_APITOKEN and IDLIST Trello with your own
```
TRELLO_APIKEY = XXX
TRELLO_APITOKEN = XXX
TRELLO_IDLIST = XXX
```
- Run Agent on Compromised Host
```
python3 agent.py
```
- Change TRELLO_APIKEY and TRELLO_APITOKEN with your own
```
TRELLO_APIKEY = XXX
TRELLO_APITOKEN = XXX
```
- Run Central Command to Control Compromised Host
```
python3 central_command.py
```

## Screenshoot
![Alt-Text](https://raw.githubusercontent.com/DwiyanTech/redteamlabs/main/TrelloC2%20-%20Att%26ck%20T1573.002/screenshoot/python-trelloc2-1.png)
![Alt-Text](https://raw.githubusercontent.com/DwiyanTech/redteamlabs/main/TrelloC2%20-%20Att%26ck%20T1573.002/screenshoot/python-trelloc2-2.png)