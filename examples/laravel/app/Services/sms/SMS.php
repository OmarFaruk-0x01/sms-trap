<?php

namespace App\Services\sms;

interface SMS
{
    public function send(array $phones, string $message, array $params): mixed;
}
