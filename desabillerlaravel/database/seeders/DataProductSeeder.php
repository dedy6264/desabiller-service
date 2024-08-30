<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;

class DataProductSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
//php artisan db:seed --class=DataProductSeeder
        
     DB::table('providers')->insert([[
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
            'product_category_name' => "PULSA",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
        ],[
            'product_category_name' => "PAKET DATA",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
        ],[
            'product_category_name' => "HP PASCABAYAR",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
        ],[
            'product_category_name' => "PLN",
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
            'product_category_name' => "MULTIFINANCE",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
        ],[
            'product_category_name' => "ASURANSI",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
        ],[
            'product_category_name' => "PDAM",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
        ],[
            'product_category_name' => "KARTU KREDIT",
            'created_by' => "sys",
            'updated_by' => "sys",
            'created_at' => now(),
            'updated_at' => now(),
        ],[
            'product_category_name' => "EWALLET",
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
        ],[
        'product_clan_name' => "XL",
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
        ],[
        'product_clan_name' => "THREE",
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
        ],[
        'product_clan_name' => "BY.U",
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
        ],[
        'product_clan_name' => "SMARTFREEN",
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
        ],[
        'product_clan_name' => "AXIS",
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
        ],[
        'product_clan_name' => "PLN",
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
     ]]);
    DB::table('product_providers')->insert([[
        'provider_id' => 1,
        'product_provider_name' => "Pulsa Telkomsel 10K",
        'product_provider_code' => "htelkomsel10000",
        'product_provider_price' => 11500,
        'product_provider_admin_fee' => 0,
        'product_provider_merchant_fee' => 0,
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
        ],[
        'provider_id' => 1,
        'product_provider_name' => "Pulsa Telkomsel 2K",
        'product_provider_code' => "htelkomsel2000",
        'product_provider_price' => 3220,
        'product_provider_admin_fee' => 0,
        'product_provider_merchant_fee' => 0,
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
        ],[
        'provider_id' => 1,
        'product_provider_name' => "PLN POSTPAID",
        'product_provider_code' => "PLNPOSTPAID",
        'product_provider_price' => 0,
        'product_provider_admin_fee' => 2750,
        'product_provider_merchant_fee' => 1800,
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
    ]]);
    DB::table('products')->insert([[
        'product_provider_id' => 2,
        'product_clan_id' => 2,
        'product_category_id' => 1,
        'product_type_id' => 2,
        'product_name' => "Pulsa Telkomsel 10000",
        'product_code' => "TSEL10",
        'product_price' => 12000,
        'product_admin_fee' => 0,
        'product_merchant_fee' => 0,
        'product_nominal' => "5000",
        'product_details' => "Pulsa Telkomsel Nominal Rp,10.000",
        'icon_url' => "",
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
        ],[
       'product_provider_id' => 3,
        'product_clan_id' => 8,
        'product_category_id' => 4,
        'product_type_id' => 1,
        'product_name' => "PLN POSTPAID",
        'product_code' => "PLNPOSTPAID",
        'product_price' => 0,
        'product_admin_fee' => 2750,
        'product_merchant_fee' => 0,
        'product_nominal' => "",
        'product_details' => "PLN TAGIHAN",
        'icon_url' => "",
        'created_by' => "sys",
        'updated_by' => "sys",
        'created_at' => now(),
        'updated_at' => now(),
     ]]);
    }
}
