@extends('dashboard.app')
@section('activeMenu')
@php
    $activeMenu="product";
@endphp
@endsection
@section('customLink')
{{-- <link href="{{url('admin/vendor/datatables/dataTables.bootstrap4.min.css')}}" rel="stylesheet"> --}}
<link href="https://cdn.datatables.net/1.10.19/css/jquery.dataTables.min.css" rel="stylesheet">
<link href="https://cdn.datatables.net/responsive/2.2.3/css/responsive.dataTables.min.css" rel="stylesheet">
<style>
    /*Form fields*/
    .dataTables_wrapper select,
    .dataTables_wrapper .dataTables_filter input {
        color: #4a5568; 			/*text-gray-700*/
        padding-left: 1rem; 		/*pl-4*/
        padding-right: 1rem; 		/*pl-4*/
        padding-top: .5rem; 		/*pl-2*/
        padding-bottom: .5rem; 		/*pl-2*/
        line-height: 1.25; 			/*leading-tight*/
        border-width: 2px; 			/*border-2*/
        border-radius: .25rem; 		
        border-color: #edf2f7; 		/*border-gray-200*/
        background-color: #edf2f7; 	/*bg-gray-200*/
    }

    /*Row Hover*/
    table.dataTable.hover tbody tr:hover, table.dataTable.display tbody tr:hover {
        background-color: #ebf4ff;	/*bg-indigo-100*/
    }
    
    /*Pagination Buttons*/
    .dataTables_wrapper .dataTables_paginate .paginate_button		{
        font-weight: 700;				/*font-bold*/
        border-radius: .25rem;			/*rounded*/
        border: 1px solid transparent;	/*border border-transparent*/
    }
    
    /*Pagination Buttons - Current selected */
    .dataTables_wrapper .dataTables_paginate .paginate_button.current	{
        color: #fff !important;				/*text-white*/
        box-shadow: 0 1px 3px 0 rgba(0,0,0,.1), 0 1px 2px 0 rgba(0,0,0,.06); 	/*shadow*/
        font-weight: 700;					/*font-bold*/
        border-radius: .25rem;				/*rounded*/
        background: #667eea !important;		/*bg-indigo-500*/
        border: 1px solid transparent;		/*border border-transparent*/
    }

    /*Pagination Buttons - Hover */
    .dataTables_wrapper .dataTables_paginate .paginate_button:hover		{
        color: #fff !important;				/*text-white*/
        box-shadow: 0 1px 3px 0 rgba(0,0,0,.1), 0 1px 2px 0 rgba(0,0,0,.06);	 /*shadow*/
        font-weight: 700;					/*font-bold*/
        border-radius: .25rem;				/*rounded*/
        background: #667eea !important;		/*bg-indigo-500*/
        border: 1px solid transparent;		/*border border-transparent*/
    }
    
    /*Add padding to bottom border */
    table.dataTable.no-footer {
        border-bottom: 1px solid #e2e8f0;	/*border-b-1 border-gray-300*/
        margin-top: 0.75em;
        margin-bottom: 0.75em;
    }
    
    /*Change colour of responsive icon*/
    table.dataTable.dtr-inline.collapsed>tbody>tr>td:first-child:before, table.dataTable.dtr-inline.collapsed>tbody>tr>th:first-child:before {
        background-color: #667eea !important; /*bg-indigo-500*/
    }

</style>
<script type="text/javascript" src="https://code.jquery.com/jquery-3.4.1.min.js"></script>
<script src="https://cdn.datatables.net/1.10.19/js/jquery.dataTables.min.js"></script>
<script src="https://cdn.datatables.net/responsive/2.2.3/js/dataTables.responsive.min.js"></script>
@endsection

@section('pageheading')
<div class="d-sm-flex align-items-center justify-content-between mb-4">
    <h1 class="h3 mb-0 text-gray-800">Pulsa</h1>
    <a href="#" class="d-none d-sm-inline-block btn btn-sm btn-primary shadow-sm"><i
            class="fas fa-download fa-sm text-white-50"></i> Generate Report</a>
</div>
@endsection

@section('script')
<script>
    var datatable=$('#dataTable').DataTable( {
        ajax:{
            url:'{!!url()->current()!!}',
        },
        columnDefs: [{
                            targets: [5],
                            render: $.fn.dataTable.render.number( '.', ',', 2)
                        }],
        columns:[
            { data: 'DT_RowIndex', orderable: false, searchable: false },
            {data:'action', name:'action', orderable:false, searchable:false},
            {data:'product_name', name:'product_name'},
            {data:'product_code', name:'product_code'},
            {data:'product_desc', name:'product_desc'},
            {data:'product_price', name:'product_price'},
            {data:'created_by', name:'created_by'},
        ],
    } );
</script>
@endsection

@section('content')
    <!-- DataTales Example -->
    <div class="card shadow mb-4">
        <div class="card-header py-3">
            {{-- <h6 class="m-0 font-weight-bold text-primary ">DataTables Example</h6> --}}
            <a href="{{route('product.create')}}" class="d-none d-sm-inline-block btn btn-sm btn-primary shadow-sm"><i
                class="fas fa-plus fa-sm text-white-50"></i> Add Product</a>
        </div>
        @if (Session::get('fail'))
        {{Session::get('fail')}}
    @endif
        <div class="card-body">
            <div class="table-responsive">
                <table class="table table-bordered" id="dataTable" width="100%" cellspacing="0">
                    <thead>
                        <tr>
                            <th>#</th>
                            <th>Action</th>
                            <th>Product Name</th>
                            <th>Product Code</th>
                            <th>Product Desc</th>
                            <th>Price</th>
                            <th>Created By</th>
                        </tr>
                    </thead>
                    {{-- <tfoot>
                        <tr>
                            <th>Name</th>
                            <th>Position</th>
                            <th>Office</th>
                            <th>Age</th>
                            <th>Start date</th>
                            <th>Salary</th>
                        </tr>
                    </tfoot> --}}
                    <tbody>
                        {{-- <tr>
                            <td>Tiger Nixon</td>
                            <td>System Architect</td>
                            <td>Edinburgh</td>
                            <td>61</td>
                            <td>2011/04/25</td>
                            <td>$320,800</td>
                        </tr>
                        <tr>
                            <td>Garrett Winters</td>
                            <td>Accountant</td>
                            <td>Tokyo</td>
                            <td>63</td>
                            <td>2011/07/25</td>
                            <td>$170,750</td>
                        </tr> --}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
    <div class="container-fluid">
        <div class="col-sm-6">
            <h2 id="nav-tabs">Tabs</h2>
            <div class="bs-component">
            <ul class="nav nav-tabs">
                <li class="nav-item">
                <a class="nav-link active" data-toggle="tab" href="#home">Home</a>
                </li>
                <li class="nav-item">
                <a class="nav-link" data-toggle="tab" href="#profile">Profile</a>
                </li>
                <li class="nav-item">
                <a class="nav-link disabled" href="#">Disabled</a>
                </li>
                <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle" data-toggle="dropdown" href="#" role="button" aria-haspopup="true" aria-expanded="false">Dropdown</a>
                <div class="dropdown-menu" style="will-change: transform;">
                    <a class="dropdown-item" href="#">Action</a>
                    <a class="dropdown-item" href="#">Another action</a>
                    <a class="dropdown-item" href="#">Something else here</a>
                    <div class="dropdown-divider"></div>
                    <a class="dropdown-item" href="#">Separated link</a>
                </div>
                </li>
            </ul>
            <div id="myTabContent" class="tab-content">
                <div class="tab-pane fade active show" id="home">
                <p>Raw denim you probably haven't heard of them jean shorts Austin. Nesciunt tofu stumptown aliqua, retro
                    synth master cleanse. Mustache cliche tempor, williamsburg carles vegan helvetica. Reprehenderit
                    butcher retro keffiyeh dreamcatcher synth. Cosby sweater eu banh mi, qui irure terry richardson ex
                    squid. Aliquip placeat salvia cillum iphone. Seitan aliquip quis cardigan american apparel, butcher
                    voluptate nisi qui.</p>
                </div>
                <div class="tab-pane fade" id="profile">
                <p>Food truck fixie locavore, accusamus mcsweeney's marfa nulla single-origin coffee squid. Exercitation
                    +1 labore velit, blog sartorial PBR leggings next level wes anderson artisan four loko farm-to-table
                    craft beer twee. Qui photo booth letterpress, commodo enim craft beer mlkshk aliquip jean shorts
                    ullamco ad vinyl cillum PBR. Homo nostrud organic, assumenda labore aesthetic magna delectus mollit.</p>
                </div>
                <div class="tab-pane fade" id="dropdown1">
                <p>Etsy mixtape wayfarers, ethical wes anderson tofu before they sold out mcsweeney's organic lomo retro
                    fanny pack lo-fi farm-to-table readymade. Messenger bag gentrify pitchfork tattooed craft beer, iphone
                    skateboard locavore carles etsy salvia banksy hoodie helvetica. DIY synth PBR banksy irony. Leggings
                    gentrify squid 8-bit cred pitchfork.</p>
                </div>
                <div class="tab-pane fade" id="dropdown2">
                <p>Trust fund seitan letterpress, keytar raw denim keffiyeh etsy art party before they sold out master
                    cleanse gluten-free squid scenester freegan cosby sweater. Fanny pack portland seitan DIY, art party
                    locavore wolf cliche high life echo park Austin. Cred vinyl keffiyeh DIY salvia PBR, banh mi before
                    they sold out farm-to-table VHS viral locavore cosby sweater.</p>
                </div>
            </div>
            </div>
        </div>
    </div>

@endsection

