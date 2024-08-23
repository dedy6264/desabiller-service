<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateTxTable extends Migration
{
    // php artisan migrate --path=/database/migrations/2024_08_20_041414_create_tx_table.php
        // php artisan migrate:rollback --path=/database/migrations/2024_08_20_041414_create_tx_table.php
    public function up()
    {
        Schema::create('biller_trxs', function (Blueprint $table) {
            $table->id();
            $table->unsignedInteger('product_clan_id'); 
            $table->string('product_clan_name');
            $table->unsignedInteger('product_category_id');
            $table->string('product_category_name');
            $table->unsignedInteger('product_type_id');
            $table->string('product_type_name');
            
            $table->unsignedInteger('product_id');
            $table->string('product_name');
            $table->string('product_code');
            $table->double('product_price');
            $table->double('product_admin_fee');
            $table->double('product_merchant_fee');
            
            $table->unsignedInteger('product_provider_id');
            $table->string('product_provider_name');
            $table->string('product_provider_code');
            $table->double('product_provider_price');
            $table->double('product_provider_admin_fee');
            $table->double('product_provider_merchant_fee');

            $table->string('status_code');
            $table->string('status_message');
            $table->string('status_desc');
            $table->string('reference_number');
            
            $table->string('provider_status_code');
            $table->string('provider_status_message');
            $table->string('provider_status_desc');
            $table->string('provider_reference_number');
                        
            $table->unsignedInteger('client_id');
            $table->string('client_name');
            $table->unsignedInteger('group_id');
            $table->string('group_name');
            $table->unsignedInteger('merchant_id');
            $table->string('merchant_name');
            $table->unsignedInteger('merchant_outlet_id');
            $table->string('merchant_outlet_name');
            $table->string('merchant_outlet_username');
            
            $table->string('customer_id');
            $table->string('other_msg');
            
            $table->string('created_by');
            $table->string('updated_by');
            $table->timestamps();
        });
        Schema::create('trx_statuses', function (Blueprint $table) {
            $table->id();
            // $table->unsignedInteger('client_id');
            $table->string('reference_number');
            $table->string('provider_reference_number');
            $table->string('status_code');
            $table->string('status_message');
            // $table->string('created_by');
            // $table->string('updated_by');
            // $table->timestamps();
        });
        Schema::create('no_generators', function (Blueprint $table) {
            $table->id();
            $table->unsignedInteger('leadingzero');
            $table->string('data_type');
            $table->unsignedInteger('seqvalue');
            $table->string('prefix')->nullable();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('biller_trxs');
        Schema::dropIfExists('trx_statuses');
        Schema::dropIfExists('no_generators');
    }
}
