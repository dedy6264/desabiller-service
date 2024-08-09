<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;

class DataHierarchySeeder extends Seeder
{
//php artisan db:seed --class=DataHierarchySeeder
    public function run()
    {
    // merchant_outlets
    // merchants
    // groups
    // clients
        DB::table('clients')->insert([
            'client_name' => Str::random(10),
            'created_by' => Str::random(10).'@gmail.com',
            'updated_by' => Hash::make('password'),
            'created_at' => Str::random(10),
            'updated_at' => Str::random(10).'@gmail.com',
        ]);
    }
}
