import websocket
import base64

BASE_URL = "http://localhost:3333"

# Encode username and password for Basic Authentication
credentials = base64.b64encode(b"testUser:testPass").decode("utf-8")

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

    # Add any WebSocket interaction logic here

except Exception as e:
    print(f"Error establishing WebSocket connection: {e}")
finally:
    ws.close()
