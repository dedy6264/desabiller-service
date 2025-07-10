<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class AddNewHierarchyTable extends Migration
{
    public function up()
    {
        // php artisan migrate --path=/database/migrations/2024_08_08_054325_add_new_hierarchy_table.php
        // php artisan migrate:rollback --path=/database/migrations/2024_08_08_054325_add_new_hierarchy_table.php

         Schema::create('product_types', function (Blueprint $table) {            
            $table->id();
            $table->string('product_type_name')->unique();
            $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
            $table->timestamps();
        });
         Schema::create('product_categories', function (Blueprint $table) {
            $table->id();            
            $table->string('product_category_name')->unique();
            $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
            $table->timestamps();
        });
        Schema::create('product_references', function (Blueprint $table) {
            $table->id();
            $table->string('product_reference_name');//089.0888,08213
            $table->string('product_reference_code');
            $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
            $table->timestamps();
        });
         Schema::create('product_helpers', function (Blueprint $table) {
            $table->id();
            $table->string('product_prefix');//089.0888,08213
            $table->unsignedInteger('product_reference_id');
            $table->foreign('product_reference_id')->references('id')->on('product_references');
            $table->timestamps();
            $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
        });
        Schema::create('products', function (Blueprint $table) {
            $table->id();
            $table->string('product_name')->unique();
            $table->string('product_code')->unique();
            $table->double('product_price');
            $table->double('product_admin_fee');
            $table->double('product_merchant_fee');

            $table->string('product_provider_name')->unique();
            $table->string('product_provider_code')->unique();
            $table->double('product_provider_price');
            $table->double('product_provider_admin_fee');
            $table->double('product_provider_merchant_fee');
        
            $table->unsignedInteger('product_category_id');
            $table->foreign('product_category_id')->references('id')->on('product_categories');

            $table->unsignedInteger('product_type_id');
            $table->foreign('product_type_id')->references('id')->on('product_types');
            $table->string('product_type_name')->unique();

            $table->unsignedInteger('product_reference_id')->nullable() ;
            $table->string('product_reference_code')->nullable();

            $table->string('product_denom')->nullable();
            $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
            $table->timestamps();
        });
         
        Schema::create('transactions', function (Blueprint $table) {
            $table->id();
            $table->string('product_provider_name');
            $table->string('product_provider_code');
            $table->double('product_provider_price');
            $table->double('product_provider_admin_fee');
            $table->double('product_provider_merchant_fee');

            $table->unsignedInteger('product_id');
            $table->string('product_name');
            $table->string('product_code');
            $table->double('product_price');
            $table->double('product_admin_fee');
            $table->double('product_merchant_fee');

            $table->unsignedInteger('product_category_id');
            $table->string('product_category_name');

            $table->unsignedInteger('product_type_id');
            $table->string('product_type_name');

            $table->string('reference_number');
            $table->string('provider_reference_number');
            
            $table->string('status_code');
            $table->string('status_message');
            $table->string('status_desc');
           
            $table->string('status_code_detail');
            $table->string('status_message_detail');
            $table->string('status_desc_detail');
            
            $table->unsignedInteger('product_reference_id')->nullable();
            $table->string('product_reference_code')->nullable();
            
            $table->string('customer_id');
            $table->text('other_reff')->nullable();
            $table->text('other_customer_info')->nullable();
            $table->text('saving_account_name')->default('');
            $table->unsignedInteger('saving_account_id')->default(0);
            $table->text('saving_account_number')->default('');
            
            $table->double('transaction_total_amount');
                        
            $table->unsignedInteger('user_app_id');
            $table->string('username');
            
            $table->string('created_by');
            $table->string('updated_by');
            $table->timestamps();
        });
        Schema::create('no_generators', function (Blueprint $table) {
            $table->id();
            $table->unsignedInteger('leadingzero');
            $table->string('data_type');
            $table->unsignedInteger('seqvalue');
            $table->string('prefix')->default("");
        });
        Schema::create('cifs', function (Blueprint $table) {
            $table->id();
            $table->string("cif_name");
            $table->string("cif_no_id");
            $table->string("cif_type_id");
            $table->string("cif_id_index")->unique();
            $table->string("cif_phone");
            $table->string("cif_email");
            $table->string("cif_address")->nullable();
            $table->timestamps();
             $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
        });
       
        Schema::create('saving_types', function (Blueprint $table) {
            $table->id();
            $table->string("saving_type_name");
            $table->string("saving_type_desc")->nullable();
            $table->timestamps();
             $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
        });
        Schema::create('saving_segments', function (Blueprint $table) {
            $table->id();
            $table->string("saving_segment_name");
            $table->double("limit_amount");
            $table->unsignedInteger("saving_type_id");
            $table->foreign('saving_type_id')->references('id')->on('saving_types');
            $table->timestamps();
             $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
        });
        Schema::create('accounts', function (Blueprint $table) {
            $table->id();
            $table->unsignedInteger("cif_id");
            $table->string("account_number");
            $table->string("account_pin")->nullable();
            $table->double("balance");
            $table->unsignedInteger("saving_segment_id");
            $table->foreign('saving_segment_id')->references('id')->on('saving_segments');
            $table->foreign('cif_id')->references('id')->on('cifs');
            $table->timestamps();
             $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
        });
         Schema::create('user_apps', function (Blueprint $table) {
            $table->id();
            $table->string("username");
            $table->string("password");
            $table->string("name");
            $table->string("identity_type")->default('NIK') ;
            $table->string("identity_number")->nullable();
            $table->string("phone")->nullable();
            $table->string("email");
            $table->string("gender")->nullable();
            $table->string("province")->nullable();
            $table->string("city")->nullable();
            $table->string("address")->nullable();
            $table->unsignedInteger("account_id");
            $table->foreign('account_id')->references('id')->on('accounts');
            $table->string("status")->default('active');
            $table->timestamps();
             $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
        });
        Schema::create('saving_transactions', function (Blueprint $table) {
            $table->id();
            $table->string("reference_number");
            $table->string("reference_number_partner");
            $table->string("dc_type");
            $table->double("transaction_amount");
            $table->string("transaction_code");
            $table->unsignedInteger("account_id");
            $table->string("account_number");
            $table->double("last_balance");
            $table->string('created_by')->nullable();
            $table->string('updated_by')->nullable();
            $table->timestamps();
        });
    }

    public function down()
    {
        Schema::dropIfExists('saving_transactions');
        Schema::dropIfExists('user_apps');
        Schema::dropIfExists('accounts');
        Schema::dropIfExists('saving_segments');
        Schema::dropIfExists('saving_types');
        Schema::dropIfExists('cifs');
        Schema::dropIfExists('no_generators');
        Schema::dropIfExists('transactions');
        Schema::dropIfExists('products');
        Schema::dropIfExists('product_helpers');        
        Schema::dropIfExists('product_references');
        Schema::dropIfExists('product_categories');
        Schema::dropIfExists('product_types');
    }
}
