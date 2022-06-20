<?php

namespace App\Http\Controllers;
namespace App\gRPC\HelloRequest;

class MarketController extends Controller
{
    public function greetUser()
    {

        return response()->json("foo!");
    }
}
