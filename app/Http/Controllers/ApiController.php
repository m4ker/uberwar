<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Http\Requests;
use App\Http\Controllers\Controller;
use App\Channel;
use App\User;
use Illuminate\Support\Facades\Redirect;


class ApiController extends Controller
{

    public function a(Request $request)
    {
        //return view('user.index');
        echo json_encode([
            'msg' => 'ok'
        ]);
    }

    public function b(Request $request)
    {
        //return view('user.index');
        echo json_encode([
            'msg' => 'ok'
        ]);
    }

}
