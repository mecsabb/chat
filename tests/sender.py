import websocket
import json
import base64

BASE_URL = "http://localhost:3333"

# Encode username and password for Basic Authentication
credentials = base64.b64encode(b"sender:testPass").decode("utf-8")

headers = {
    "Connection": "Upgrade",
    "Upgrade": "websocket",
    "Sec-WebSocket-Version": "13",
    "Sec-WebSocket-Key": "x3JJHMbDL1EzLkh9GBhXDw==",
    # Add the encoded credentials to the header
    "Authorization": f"Basic {credentials}"
}

try:
    # Try to initiate a WebSocket connection
    ws = websocket.create_connection(
        f"ws://localhost:3333/login", header=headers)
    print("WebSocket connection established!")
    while(1):
        input("Press Enter to send a message...")
        message_json = json.dumps({
            "from": "sender",
            "to": "testUser",
            "msg": base64.b64encode("Hello from the sender script!".encode()).decode()
        })
        ws.send(message_json)

    # Add any WebSocket interaction logic here

except Exception as e:
    print(f"Error establishing WebSocket connection: {e}")
finally:
    ws.close()
