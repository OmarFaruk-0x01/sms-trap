from django.apps import AppConfig
from . import get_sms_provider

class SmsProviderConfig(AppConfig):
    default_auto_field = 'django.db.models.BigAutoField'
    name = 'sms'

    def ready(self):
        get_sms_provider()
