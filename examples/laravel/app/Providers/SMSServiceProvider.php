<?php

namespace App\Providers;

use App\Services\sms\SMS;
use App\Services\sms\SMSTrap;
use Illuminate\Support\ServiceProvider;

class SMSServiceProvider extends ServiceProvider
{
    public function register(): void
    {
        $this->app->singleton('sms', function ($app) {
            $provider = config('sms.provider');

            $apiKey = config("sms.{$provider}.api_key");
            $senderId = config("sms.{$provider}.sender_id");

            return match ($provider) {
                // SMS Trap no need any api key or any other credentials.
                // This is just a demo that you can use for other providers
                'sms-trap' => new SMSTrap($apiKey, $senderId),
            };

        });

        $this->app->singleton(SMS::class, function ($app) {
            return $app->make('sms');
        });
    }

}
