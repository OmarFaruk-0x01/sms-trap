# sms_services/services.py

from django.conf import settings
from .providers.base import SmsProvider
from .providers.sms_trap import SMSTrap

class SMSProviderLocator:
    _instance = None

    @classmethod
    def get_instance(cls):
        if cls._instance is None:
            provider = settings.SMS_PROVIDER
            api_key = settings.SMS_PROVIDERS.get(provider, {}).get('api_key')
            sender_id = settings.SMS_PROVIDERS.get(provider, {}).get('sender_id')

            cls._instance = cls._get_service(provider, api_key, sender_id)

        return cls._instance

    @staticmethod
    def _get_service(provider, api_key, sender_id):
        if provider == 'sms-trap':
            # SMS Trap doesn't need api_key or sender_id, but we're passing them for demonstration
            return SMSTrap(api_key, sender_id)

        # Add other providers here as needed
        raise ValueError(f"Unsupported SMS provider: {provider}")


def get_sms_provider() -> SmsProvider:
    return SMSProviderLocator.get_instance()


def send_sms(phones: list[str], message: str,additional: dict) -> dict:
    provider = get_sms_provider()

    return provider.send(phones, message, additional)