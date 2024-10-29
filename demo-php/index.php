<?php

use OpenTelemetry\API\Globals;
use OpenTelemetry\Config\SDK\Configuration;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;

require __DIR__ . '/vendor/autoload.php';

// parse the config file here
$configuration = Configuration::parseFile(__DIR__ . '/../config.yaml');
$sdkBuilder = $configuration->create();
$sdkBuilder->buildAndRegisterGlobal();

$tracer = Globals::tracerProvider()->getTracer('demo');

$app = AppFactory::create();

$app->get('/rolldice', function (Request $request, Response $response) use ($tracer) {
    $span = $tracer
        ->spanBuilder('manual-span')
        ->startSpan();
    $result = random_int(1,6);
    $response->getBody()->write(strval($result));
    $span
        ->addEvent('rolled dice', ['result' => $result])
        ->end();
    return $response;
});

$app->run();
