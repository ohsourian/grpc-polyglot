<?php

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use App\gRPC\HelloRequest;

class MarketController extends Controller
{

    public function __construct()
    {

    }

    public function greetUser()
    {

        return response()->json("foo!");
    }
}
