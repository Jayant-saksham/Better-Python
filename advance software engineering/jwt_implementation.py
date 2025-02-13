import json
import hmac
import hashlib
import base64
from typing import Dict, Tuple

class Base64Url:
    @staticmethod
    def encode(data: bytes) -> str:
        return base64.b64encode(data).decode().rstrip("=").replace("+", "-").replace("/", "_")

    @staticmethod
    def decode(data: str) -> bytes:
        data += "=" * (-len(data) % 4)  # Pad to make length a multiple of 4
        return base64.b64decode(data.replace("-", "+").replace("_", "/"))

class JWT:
    def __init__(self, secret: str):
        self.secret = secret

    def create(self, payload: Dict) -> str:
        header = {"alg": "HS256", "typ": "JWT"}

        header_json = json.dumps(header, separators=(",", ":")).encode()
        payload_json = json.dumps(payload, separators=(",", ":")).encode()

        header_b64 = Base64Url.encode(header_json)
        payload_b64 = Base64Url.encode(payload_json)

        signature = self._sign(f"{header_b64}.{payload_b64}")
        signature_b64 = Base64Url.encode(signature)

        return f"{header_b64}.{payload_b64}.{signature_b64}"

    def verify(self, token: str) -> bool:
        try:
            header_b64, payload_b64, signature_b64 = self._split_token(token)
            expected_signature = self._sign(f"{header_b64}.{payload_b64}")
            expected_signature_b64 = Base64Url.encode(expected_signature)
            return hmac.compare_digest(signature_b64, expected_signature_b64)
        except Exception:
            return False

    def decode(self, token: str) -> Dict:
        try:
            _, payload_b64, _ = self._split_token(token)
            payload_json = Base64Url.decode(payload_b64)
            return json.loads(payload_json)
        except Exception:
            return {}

    def _sign(self, data: str) -> bytes:
        return hmac.new(self.secret.encode(), data.encode(), hashlib.sha256).digest()

    def _split_token(self, token: str) -> Tuple[str, str, str]:
        return token.split(".")

# Example Usage
secret_key = "my_secret"
jwt_handler = JWT(secret_key)

payload = {"user_id": 123, "role": "admin"}

jwt_token = jwt_handler.create(payload)
print("JWT Token:", jwt_token)

is_valid = jwt_handler.verify(jwt_token)
print("Is JWT valid?", is_valid)

decoded_payload = jwt_handler.decode(jwt_token)
print("Decoded Payload:", decoded_payload)
