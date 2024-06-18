<?php

use App\Facades\SMS;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});

Route::post('/', function (Request $request) {
    $phone = $request->get('phone');
    $message = $request->get('message');
    $label = $request->get('label');

    if ($phone && $message && $label) {
        SMS::send([$phone], $message, ['label' => $label]);
    }

    return redirect('/');
});
