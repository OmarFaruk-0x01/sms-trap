<?php

namespace App\Services\sms;

use Illuminate\Support\Facades\Http;

class SMSTrap implements SMS
{
    protected string $baseUrl = 'http://localhost:8080/api/v1/trap';

    public function send(array $phones, string $message, array $params): mixed
    {
        $query = [
            'message' => $message,
            'label' => $data['label'] ?? 'transactional',

        ];

        foreach ($phones as $phone) {
            $query['phones[]'] = $phone;
        }

        $response = Http::withOptions(['debug' => true])->get($this->baseUrl, $query);

        return $response->json();
    }
}
