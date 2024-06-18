<?php

namespace App\Facades;

use Illuminate\Support\Facades\Facade;

class SMS extends Facade
{
    protected static function getFacadeAccessor(): string
    {
        return 'sms';
    }
}
