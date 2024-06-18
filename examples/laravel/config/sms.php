<?php

return [
    'provider' => env('SMS_PROVIDER', 'sms-trap'),

    'sms-trap' => [
        'api_key' => env('SMS_TRAP_API_KEY'),
        'sender_id' => env('SMS_TRAP_SENDERID'),
    ],

    // ... other providers
];
