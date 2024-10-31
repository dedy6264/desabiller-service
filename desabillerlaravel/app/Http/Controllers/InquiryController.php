<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class InquiryController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(){
        // $response = Http::post('https://8c47-125-166-235-226.ngrok-free.app/client/gets', [
        //     'id' => 0,
        //     'clientName' => "",
        //     'limit' => 0,
        //     'offset' => 0,
        //     'orderBy' => "",
        //     'startDate' => "",
        //     'endDate' => "",
        //     'username' => "",
        // ])->json();
        // // dd($response);
        //  // Mengambil data dari API
        // //  dd($response);
        //  if (!is_array($response) || !isset($response['result']) || !is_array($response['result'])) {
        //     return response()->json(['error' => 'Invalid API response format or data type'], 500);
        // }
        // // Ambil data 'result'
        // $response = $response['result'];
        // $dataRes = $response['data'] ?? [];
        // return view('dashboard.product.pulsa',compact('datares'));
        return view('dashboard.product.pulsa');
    }
    public function create(){
        //
    }
    public function store(Request $request){
        //
    }
    public function show($id){
        //
    }
    public function edit($id){
        //
    }
    public function update(Request $request, $id){
        //
    }
    public function destroy($id){
        //
    }
    //Pulsa
    public function pulsaIndex(){//menampilkan home input id customer
    return view("dashboard.pulsa.index");
    }
    public function pulsaCheckSimProvider(Request $request){//get sim provider & get product by sim provider
       dd("Uiuhhiu");
        $request->validate([
            'id_customer' => 'required|numeric'
        ]);
        $response = Http::post('https://localhost:10036/category/gets', [
            'id' => 0,
            'clientName' => "",
            'limit' => 0,
            'offset' => 0,
            'orderBy' => "",
            'startDate' => "",
            'endDate' => "",
            'username' => "",
        ])->json();
        dd($response);
    }
    public function pulsaInquiry(){//inquiry

    }
    public function pulsaConfirm(){//confirm, return print struk page

    }
}
