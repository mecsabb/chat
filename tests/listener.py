import websocket
import json
import threading
import base64

# Encoding username and password
username_password = base64.b64encode("testUser:testPass".encode()).decode()

# Headers with Basic Auth
headers = {
    "Authorization": f"Basic {username_password}"
}

def on_message(ws, message):
    print(f"Received message: {message}")
    # Assuming the message is JSON with a Base64-encoded "msg" field:
    message_dict = json.loads(message)
    msg_base64 = message_dict.get("msg")

    if msg_base64:
        # Decode the Base64 message
        msg = base64.b64decode(msg_base64).decode('utf-8')
        print(f"Decoded message: {msg}")
    else:
        print("No message or not Base64 encoded.")

def on_error(ws, error):
    print(f"Error: {error}")

def on_close(ws, close_status_code, close_msg):
    print("### Connection closed ###")
    print(f"Close status code: {close_status_code}")
    print(f"Close message: {close_msg}")

def on_open(ws):
    print("Connection is open. Listening for messages...")

def run_websocket():
    ws = websocket.WebSocketApp("ws://localhost:3333/login",
                                header=headers,
                                on_open=on_open,
                                on_message=on_message,
                                on_error=on_error,
                                on_close=on_close)

    ws.run_forever()

if __name__ == "__main__":
    websocket.enableTrace(True)
    thread = threading.Thread(target=run_websocket)
    thread.daemon = True
    thread.start()

    # Keep the main thread alive to keep listening
    input("Press enter to quit\n\n")
