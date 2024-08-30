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
        DB::table('clients')->insert([[
            'client_name' => "DYKO",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ],[
            'client_name' => "MAKARIOS",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ]]);
        DB::table('groups')->insert([[
            'client_id' => 1,
            'group_name' => "KLATEN",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ],[
            'client_id' => 2,
            'group_name' => "SEMARANG",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ]]);
        DB::table('merchants')->insert([[
            'group_id' => 1,
            'merchant_name' => "DPMALL OUTLET",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ],[
            'group_id' => 1,
            'merchant_name' => "GAPURA",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ]]);
        DB::table('merchant_outlets')->insert([[
            'merchant_id' => 1,
            'merchant_outlet_name' => "TAWCI 01",
            'merchant_outlet_username' => "abcd12345",
            'merchant_outlet_password' => "AAAAAAAAAAAAAAAAu9Hhop9zqfrnfPZAuGwMWIXt1Rh8z8++",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ],[
            'merchant_id' => 1,
            'merchant_outlet_name' => "TAWCI 02",
            'merchant_outlet_username' => "taucikuenak",
            'merchant_outlet_password' => "AAAAAAAAAAAAAAAAu9Hhop9zqfrnfPZAuGwMWIXt1Rh8z8++",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
        ]]);
        }
}
