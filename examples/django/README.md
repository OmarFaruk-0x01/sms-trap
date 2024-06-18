# SMS Trap Django Integration

This project demonstrates the integration of SMS Trap service into a Laravel application using the Strategy pattern and Facade.

## File Structure

```
django-project/
├── app/
│   ├── settings.py (modified)
│   └── urls.py (modified)
├── demo/ (prepare demo ui)
└── sms/
    ├── providers/
    │   ├── base.py
    │   └── sms_trap.py
    ├── __init__.py (modified)
    └── apps.py (modified)
```

## Components

1. **SMS Base Class** (`sms/providers/base.py`):  Defines an abstract base class with a send method that accepts a dictionary of data (phones, message, additional) and returns a dictionary

2. **SMSTrap** (`sms/providers/sms_trap.py`): A concrete implementation of the SMS Base Class for the SMS Trap provider.

3. **SMSProviderLocator** (`sms/__init__.py`): A helper class that acts as a service locator for the configured SMS provider. It instantiates the appropriate provider class based on the settings and provides a singleton instance.

4. **SmsProviderConfig** (`sms/apps.py`): A Django AppConfig subclass that loads the configured SMS provider during the application's startup and makes it available through the `SMSProviderLocator`


## Setup

1. Register the `SMS Provider` in `app/settings.py`:

```py
INSTALLED_APPS = [
    # ... other apps
    'sms',
]
```

## Usage

You can use the SMS provider in your application in two ways:

1. Using serivce instance:

```py
from sms import get_sms_provider

provider = get_sms_provider()
provider.send([phone_number], message, {'label': label})

```

2. Using the helper function:

```py
from sms import send_sms

send_sms([phone_number], message, {'label': label})

```
