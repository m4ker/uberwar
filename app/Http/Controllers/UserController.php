<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Http\Requests;
use App\Http\Controllers\Controller;
use App\Channel;
use App\User;
use Illuminate\Support\Facades\Redirect;


class UserController extends Controller
{

    public function index(Request $request)
    {
        return view('user.index');
    }

}
