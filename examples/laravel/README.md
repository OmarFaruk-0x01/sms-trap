# SMS Trap Laravel Integration

This project demonstrates the integration of SMS Trap service into a Laravel application using the Strategy pattern and Facade.

## File Structure

```
laravel-project/
├── app/
│   ├── Facades/
│   │   └── SMS.php
│   ├── Providers/
│   │   └── SMSServiceProvider.php
│   └── Services/
│       └── sms/
│           ├── SMS.php
│           └── SMSTrap.php
├── config/
│   └── sms.php
├── bootstrap/
│   └── providers.php (modified)
└── routes/
    └── web.php (modified)
```

## Components

1. **SMS Interface** (`app/Services/sms/SMS.php`): Defines the contract for SMS services.

2. **SMSTrap** (`app/Services/sms/SMSTrap.php`): Implements the SMSService interface for the SMS Trap service.

3. **SMSServiceProvider** (app/Providers/SMSServiceProvider.php): Binds the SMS service to the Laravel service container.

4. **SMS Facade** (app/Facades/SMS.php): Provides a convenient static interface to the SMS service.

## Setup

1. Register the `SMSServiceProvider` in `bootstrap/providers.php`:
```php
return [
// ...
    App\Providers\SMSServiceProvider::class,
],
```

## Usage

You can use the SMS service in your application in two ways:

1. Using dependency injection:
```php

use App\Services\sms\SMS;

public function someMethod(SMS $smsService)
{
    $result = $smsService->send([
        'phones' => ['+1234567890', '+9876543210'],
        'message' => 'Your verification code is 123456',
        'label' => 'transactional'
    ]);
}
```

2. Using the Facade:
```php
use App\Facades\SMS;

$result = SMS::send([
    'phones' => ['+1234567890', '+9876543210'],
    'message' => 'Your verification code is 123456',
    'label' => 'transactional'
]);
```

## Notes

- The SMS Trap service is configured to use `http://localhost:8080/api/v1/trap` as the base URL. Update this in `SMSTrap.php` if needed.
- The service is registered as a singleton in the service container, meaning the same instance will be used throughout the application lifecycle.
- The implementation handles the special case of formatting the `phones` parameter correctly in the query string.
