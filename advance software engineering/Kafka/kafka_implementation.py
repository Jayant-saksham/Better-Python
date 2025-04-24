import queue
import threading
import time

event_queue = queue.Queue()

def sign_in(username):
    print(f"[Auth] {username} signed in.")
    event = {
        'type': 'SEND_WELCOME_EMAIL',
        'username': username
    }
    event_queue.put(event)
    print(f"[Producer] Event queued for {username}.")


def email_service():
    while True:
        if not event_queue.empty():
            event = event_queue.get()
            if event['type'] == 'SEND_WELCOME_EMAIL':
                username = event['username']
                print(f"[EmailService] Sending welcome email to {username}...")
                time.sleep(1)
                print(f"[EmailService] Welcome email sent to {username}.")
        time.sleep(0.5)


email_thread = threading.Thread(target=email_service, daemon=True)
email_thread.start()


sign_in("alice")
time.sleep(0.5)
sign_in("bob")
time.sleep(2)

print("[Main] Done.")
