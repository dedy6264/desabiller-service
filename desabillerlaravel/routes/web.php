<?php
use App\Http\Controllers\MainController;
use App\Http\Controllers\InquiryController;

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

// Route::get('/', function () {
//     return view('welcome');
// });
Route::get('/', [MainController::class, 'index'])->name('home'); //proses logout
Route::get('/pulsa', [InquiryController::class, 'pulsaIndex'])->name('pulsa'); //proses logout
Route::get('/pulsa/checkSimProv', [InquiryController::class, 'pulsaCheckSimProvider'])->name('pulsa.checkSimProv'); //proses logout


// Route::resource('dashboard',MainController::class);
