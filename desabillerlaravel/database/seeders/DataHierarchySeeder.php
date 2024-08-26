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
        DB::table('providers')->insert([[
                'provider_name' => "MYBILLS",
                'created_by' => "sys",
                'updated_by' => "sys",
                'created_at' => now(),
                'updated_at' => now(),
                ],[
                'provider_name' => "IAK",
                'created_by' => "sys",
                'updated_by' => "sys",
                'created_at' => now(),
                'updated_at' => now(),
        ]]);
        DB::table('product_types')->insert([[
                'product_type_name' => "PAYMENT",
                'created_by' => "sys",
                'updated_by' => "sys",
                'created_at' => now(),
                'updated_at' => now(),
                ],[
                'product_type_name' => "PURCHASING",
                'created_by' => "sys",
                'updated_by' => "sys",
                'created_at' => now(),
                'updated_at' => now(),
        ]]);
        DB::table('product_categories')->insert([[
                'product_category_name' => "EWALET",
                'created_by' => "sys",
                'updated_by' => "sys",
                'created_at' => now(),
                'updated_at' => now(),
                ],[
                'product_category_name' => "GAME",
                'created_by' => "sys",
                'updated_by' => "sys",
                'created_at' => now(),
                'updated_at' => now(),
            ],[
                'product_category_name' => "PULSA",
                'created_by' => "sys",
                'updated_by' => "sys",
                'created_at' => now(),
                'updated_at' => now(),
        ]]);
        DB::table('product_clans')->insert([[
            'product_clan_name' => "INDOSAT",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ],[
            'product_clan_name' => "TELKOMSEL",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
         ]]);
        DB::table('product_providers')->insert([[
            'provider_id' => 1,
            'product_provider_name' => "Pulsa Telkomsel 50K",
            'product_provider_code' => "PTSEL50",
            'product_provider_price' => 49500,
            'product_provider_admin_fee' => 0,
            'product_provider_merchant_fee' => 0,
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ],[
            'provider_id' => 1,
            'product_provider_name' => "Pulsa Telkomsel 5K",
            'product_provider_code' => "PTSEL5",
            'product_provider_price' => 5250,
            'product_provider_admin_fee' => 0,
            'product_provider_merchant_fee' => 0,
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
        ]]);
        DB::table('products')->insert([[
            'product_provider_id' => 2,
            'product_clan_id' => 2,
            'product_category_id' => 3,
            'product_type_id' => 2,
            'product_name' => "Pulsa Telkomsel 5000",
            'product_code' => "TSEL5",
            'product_price' => 6000,
            'product_admin_fee' => 0,
            'product_merchant_fee' => 0,
            'product_nominal' => "5000",
            'product_details' => "Pulsa Telkomsel Nominal Rp,5.000",
            'icon_url' => "",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
            ],[
           'product_provider_id' => 1,
            'product_clan_id' => 2,
            'product_category_id' => 3,
            'product_type_id' => 2,
            'product_name' => "Pulsa Telkomsel 50000",
            'product_code' => "TSEL50",
            'product_price' => 51000,
            'product_admin_fee' => 0,
            'product_merchant_fee' => 0,
            'product_nominal' => "5000",
            'product_details' => "Pulsa Telkomsel Nominal Rp,5.000",
            'icon_url' => "",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
         ]]);
        }
}
