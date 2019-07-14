<?php

include "vendor/autoload.php";

/**
 * @return SampleService
 * @throws Exception
 */
function getCLient()
{
    $TcpServerAddr = "tcp://127.0.0.1:8050";
    $client = \Hprose\Socket\Client::create($TcpServerAddr, false);
    $service = $client->useService('', 'Sample');
    return $service;
}

try {
    $client = getCLient();
    $rep = $client->GetUserInfo(10);
    echo $rep->errCode . PHP_EOL;
    print_r($rep);
} catch (Exception $e) {
    echo $e->getMessage();
}
