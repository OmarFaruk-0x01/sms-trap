
import requests
from django.conf import settings
from .base import SmsProvider
import logging


class SMSTrap(SmsProvider):
    def __init__(self, api_key, sender_id):
        self.base_url = settings.SMS_TRAP_BASE_URL
        # api_key and sender_id are not used in SMSTrap, but we're accepting them for demonstration
        self.api_key = api_key
        self.sender_id = sender_id

    def send(self, phones: list[str], message: str, additional: dict) -> dict:

        label = additional.get('label', 'transactional')

        params = []

        for phone in phones:
            params.append(('phones[]', phone))

        params.extend([
            ('message', message),
            ('label', label)
        ])

        response = requests.get(self.base_url, params=params)
        logging.debug(response)
        return response.json()
