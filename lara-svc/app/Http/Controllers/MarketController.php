<?php

namespace App\Http\Controllers;

class MarketController extends Controller
{
    public function greetUser()
    {
        return response()->json("foo!");
    }
}
