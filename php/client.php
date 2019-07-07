<?php

include "vendor/autoload.php";

try{
    $TcpServerAddr = "tcp://127.0.0.1:8050";
    $client = \Hprose\Socket\Client::create($TcpServerAddr, false);
    $service = $client->useService('', 'Sample');
    $rep = $service->GetUserInfo(10);
    print_r($rep);
} catch (Exception $e){
    echo $e->getMessage();
}
