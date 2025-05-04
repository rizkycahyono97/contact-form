<?php

use App\Http\Controllers\ContactController;
use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});

// Define a resource route for the ContactController
// This automatically sets up standard routes for index, create, store, show, edit, update, and destroy actions
Route::resource('contact', ContactController::class);