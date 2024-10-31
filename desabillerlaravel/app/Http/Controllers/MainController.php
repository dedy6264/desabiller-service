<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class MainController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        $response = Http::post('http://localhost:10036/category/gets', [
            'id' => 0,
            'clientName' => "",
            'limit' => 0,
            'offset' => 0,
            'orderBy' => "",
            'startDate' => "",
            'endDate' => "",
            'username' => "",
        ])->json();

        // dd($response);
         // Mengambil data dari API
        //  dd($response);
         if (!is_array($response) || !isset($response['result']) || !is_array($response['result'])) {
            return response()->json(['error' => 'Invalid API response format or data type'], 500);
        }
        // Ambil data 'result'
        $response = $response['result'];
        $dataRes = $response['data'] ?? [];
        // dd($dataRes[1]['productCategoryName']);
        return view("dashboard.content",compact('dataRes'));
    }
    public function all(Request $request)
    {
        $response = Http::post('https://8c47-125-166-235-226.ngrok-free.app/product/gets', [
            'productCategoryId' => (int)$request->id,
        ])->json();
        // dd($response);
         // Mengambil data dari API
        //  dd($response);
         if (!is_array($response) || !isset($response['result']) || !is_array($response['result'])) {
            return response()->json(['error' => 'Invalid API response format or data type'], 500);
        }
        // Ambil data 'result'
        $response = $response['result'];
        $dataRes = $response['data'] ?? [];
        // dd($dataRes);
        return view("dashboard.content",compact('dataRes'));
  
    }
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        //
    }
}
