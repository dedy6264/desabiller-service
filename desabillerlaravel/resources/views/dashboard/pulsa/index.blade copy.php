@extends('dashboard.app')
@section('activeMenu')
@php
    $activeMenu="pulsa";
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

{{-- @section('script')
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
@endsection --}}

@section('content')
    <!-- DataTales Example -->
    <div class="card shadow mb-4">
        <div class="card-header py-3">
            {{-- <h6 class="m-0 font-weight-bold text-primary ">DataTables Example</h6> --}}
            <a href="#" class="d-none d-sm-inline-block btn btn-sm btn-primary shadow-sm"><i
                class="fas fa-plus fa-sm text-white-50"></i> Check Sim Provider</a>
        </div>
        @if (Session::get('fail'))
        {{Session::get('fail')}}
    @endif
        <div class="card-body">
            <div class="table-responsive">
                <form action="{{route('pulsa.checkSimProv')}}" method="post"  class="d-none d-sm-inline-block form-inline mr-auto ml-md-3 my-2 my-md-0 mw-100 navbar-search">
                    @csrf
                    <div class="row">
                        <div class="input-group col-md-6 mb-4">
                            <label for="id_customer">no HP</label>
                            <input type="text" name="id_customer" class="form-control bg-light border-0 small" placeholder="Id Customer">
                            {{-- <div class="input-group-append">
                                <button class="btn btn-primary" type="button">
                                    <i class="fas fa-search fa-sm"></i>
                                </button>
                            </div> --}}
                        </div>
                        <div class="input-group col-md-6 mb-4 align-right">
                            {{-- <div class="input-group-append"> --}}
                                <button class="btn btn-primary" type="submit" value="submit">
                                    <i class="fas fa-plus fa-sm"></i> Save
                                </button>
                            {{-- </div> --}}
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>

@endsection

