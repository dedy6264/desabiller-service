<ul class="navbar-nav bg-gradient-primary sidebar sidebar-dark accordion" id="accordionSidebar">

    <!-- Sidebar - Brand -->
    <a class="sidebar-brand d-flex align-items-center justify-content-center" href="{{route('home')}}">
        <div class="sidebar-brand-icon rotate-n-15">
            <i class="fas fa-laugh-wink"></i>
        </div>
        <div class="sidebar-brand-text mx-3">SB Admin <sup>2</sup></div>
    </a>

    <!-- Divider -->
    <hr class="sidebar-divider my-0">

    <!-- Nav Item - Dashboard -->
    <li class="nav-item active">
        <a class="nav-link" href="#">
        {{-- <a class="nav-link" href="{{route('dashboard')}}"> --}}
            <i class="fas fa-fw fa-tachometer-alt"></i>
            <span>Dashboard</span></a>
    </li>

    {{-- <!-- Divider -->
    <hr class="sidebar-divider">

    <!-- Heading -->
    <div class="sidebar-heading">
        User
    </div>

    <!-- Nav Item - Pages Collapse Menu -->
    <li class="nav-item ">
        <a class="nav-link collapsed" href="#" data-toggle="collapse" data-target="#collapseUser"
            aria-expanded="true" aria-controls="collapseUser">
            <i class="fas fa-fw fa-cog"></i>
            <span>User</span>
        </a>
        <div id="collapseUser" class="collapse" aria-labelledby="headingTwo" data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Custom Components:</h6>
                <a class="collapse-item" href="buttons.html">User</a>
                <a class="collapse-item" href="cards.html">Cards</a>
            </div>
        </div>
    </li> --}}

    <!-- Divider -->
    <hr class="sidebar-divider">

    <!-- Heading -->
    <div class="sidebar-heading">
        User
    </div>

    <!-- Nav Item - Pages Collapse Menu -->
    <li class="nav-item {{in_Array($activeMenu,['user','otorisasi']) ? 'active':''}}">
        {{-- collapsed --}}
        <!-- {{dump($activeMenu)}} -->
        <a class="nav-link {{in_Array($activeMenu,['user','otorisasi']) ? '':'collapsed'}}" href="#" data-toggle="collapse" data-target="#collapseUser"
            aria-expanded="{{in_Array($activeMenu,['user','otorisasi']) ? 'true':''}}" aria-controls="collapseUser">
            <i class="fas fa-fw fa-cog"></i>
            <span>User</span>
        </a>
        <div id="collapseUser" class="collapse {{in_Array($activeMenu,['user','otorisasi']) ? 'show':''}}" aria-labelledby="headingTwo" data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Custom Components:</h6>
                <a class="collapse-item {{in_Array($activeMenu,['user']) ? 'active':''}}" href="#">User</a>
                {{-- <a class="collapse-item {{in_Array($activeMenu,['user']) ? 'active':''}}" href="{{route('user.index')}}">User</a> --}}
                {{-- <a class="collapse-item {{in_Array($activeMenu,['otorisasi']) ? 'active':''}}" href="{{route('transaction.index')}}">Otorisasi</a> --}}
                <a class="collapse-item {{in_Array($activeMenu,['otorisasi']) ? 'active':''}}" href="#">Otorisasi</a>
            </div>
        </div>
    </li>

    <!-- Nav Item - Pages Collapse Menu -->
    <li class="nav-item {{in_Array($activeMenu,['product','transaction','sales']) ? 'active':''}}">
        {{-- collapsed --}}
        <a class="nav-link {{in_Array($activeMenu,['product','transaction','sales']) ? '':'collapsed'}}" href="#" data-toggle="collapse" data-target="#collapseProduct"
            aria-expanded="{{in_Array($activeMenu,['product','transaction','sales']) ? 'true':''}}" aria-controls="collapseProduct">
            <i class="fas fa-fw fa-cog"></i>
            <span>Pembelian</span>
        </a>
        <div id="collapseProduct" class="collapse {{in_Array($activeMenu,['product','transaction','sales']) ? 'show':''}}" aria-labelledby="headingTwo" data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Custom Components:</h6>
                <a class="collapse-item {{in_Array($activeMenu,['product']) ? 'active':''}}" href="#">Pulsa</a>
                {{-- <a class="collapse-item {{in_Array($activeMenu,['product']) ? 'active':''}}" href="{{route('product.index')}}">Produk</a> --}}
                <a class="collapse-item {{in_Array($activeMenu,['transaction']) ? 'active':''}}" href="#">Voucer Game</a>
                {{-- <a class="collapse-item {{in_Array($activeMenu,['transaction']) ? 'active':''}}" href="{{route('transaction.index')}}">Transaksi</a> --}}
                <a class="collapse-item {{in_Array($activeMenu,['sales']) ? 'active':''}}" href="#">PLN Token</a>
                {{-- <a class="collapse-item {{in_Array($activeMenu,['sales']) ? 'active':''}}" href="{{route('sales.index')}}">Penjualan</a> --}}
            </div>
        </div>
    </li>
<!-- Divider -->
<hr class="sidebar-divider">

<!-- Heading -->
<div class="sidebar-heading">
    Front End
</div>
    <!-- Nav Item - Pages Collapse Menu -->
     <li class="nav-item">
        <a class="nav-link collapsed" href="#" data-toggle="collapse" data-target="#collapseTwo"
            aria-expanded="true" aria-controls="collapseTwo">
            <i class="fas fa-fw fa-cog"></i>
            <span>Components</span>
        </a>
        <div id="collapseTwo" class="collapse" aria-labelledby="headingTwo" data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Custom FrontEnd:</h6>
                <a class="collapse-item" href="buttons.html">Landing Page</a>
                <a class="collapse-item" href="cards.html">Slide Project</a>
                <a class="collapse-item" href="cards.html">Step</a>
                <a class="collapse-item" href="cards.html">Motto</a>
                <a class="collapse-item" href="cards.html">Gallery</a>
            </div>
        </div>
    </li>
    {{--
    <!-- Nav Item - Utilities Collapse Menu -->
    <li class="nav-item">
        <a class="nav-link collapsed" href="#" data-toggle="collapse" data-target="#collapseUtilities"
            aria-expanded="true" aria-controls="collapseUtilities">
            <i class="fas fa-fw fa-wrench"></i>
            <span>Utilities</span>
        </a>
        <div id="collapseUtilities" class="collapse" aria-labelledby="headingUtilities"
            data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Custom Utilities:</h6>
                <a class="collapse-item" href="utilities-color.html">Colors</a>
                <a class="collapse-item" href="utilities-border.html">Borders</a>
                <a class="collapse-item" href="utilities-animation.html">Animations</a>
                <a class="collapse-item" href="utilities-other.html">Other</a>
            </div>
        </div>
    </li>

    <!-- Divider -->
    <hr class="sidebar-divider">

    <!-- Heading -->
    <div class="sidebar-heading">
        Addons
    </div>

    <!-- Nav Item - Pages Collapse Menu -->
    <li class="nav-item">
        <a class="nav-link collapsed" href="#" data-toggle="collapse" data-target="#collapsePages"
            aria-expanded="true" aria-controls="collapsePages">
            <i class="fas fa-fw fa-folder"></i>
            <span>Pages</span>
        </a>
        <div id="collapsePages" class="collapse" aria-labelledby="headingPages" data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Login Screens:</h6>
                <a class="collapse-item" href="login.html">Login</a>
                <a class="collapse-item" href="register.html">Register</a>
                <a class="collapse-item" href="forgot-password.html">Forgot Password</a>
                <div class="collapse-divider"></div>
                <h6 class="collapse-header">Other Pages:</h6>
                <a class="collapse-item" href="404.html">404 Page</a>
                <a class="collapse-item" href="blank.html">Blank Page</a>
            </div>
        </div>
    </li>

    

    <!-- Nav Item - Tables -->
    <li class="nav-item">
        <a class="nav-link" href="tables.html">
            <i class="fas fa-fw fa-table"></i>
            <span>Tables</span></a>
    </li>--}}
    <!-- Nav Item - Logout -->
    <li class="nav-item">
        {{-- <a class="nav-link" href="{{route('logout')}}"> --}}
        <a class="nav-link" href="#">
            <i class="fas fa-fw fa-sign-out-alt"></i>
            <span>Logout</span></a>
    </li>
    <!-- Divider -->
    <hr class="sidebar-divider d-none d-md-block">

    <!-- Sidebar Toggler (Sidebar) -->
    <div class="text-center d-none d-md-inline">
        <button class="rounded-circle border-0" id="sidebarToggle"></button>
    </div> 

    <!-- Sidebar Message -->
    {{-- <div class="sidebar-card d-none d-lg-flex">
        <img class="sidebar-card-illustration mb-2" src="img/undraw_rocket.svg" alt="...">
        <p class="text-center mb-2"><strong>SB Admin Pro</strong> is packed with premium features, components, and more!</p>
        <a class="btn btn-success btn-sm" href="https://startbootstrap.com/theme/sb-admin-pro">Upgrade to Pro!</a>
    </div> --}}

</ul>