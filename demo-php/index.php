<?php

use OpenTelemetry\API\Globals;
use OpenTelemetry\Config\SDK\Configuration;
use OpenTelemetry\API\Trace\Propagation\TraceContextPropagator;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;

require __DIR__ . '/vendor/autoload.php';

$env = file_get_contents(__DIR__."/.env");
$lines = explode("\n",$env);

foreach($lines as $line){
  preg_match("/([^#]+)\=(.*)/",$line,$matches);
  if(isset($matches[2])){
    $_SERVER[$matches[1]] = $matches[2];
  }
}

// parse the config file here
$configuration = Configuration::parseFile(__DIR__ . '/../config.yaml');
$sdkBuilder = $configuration->create()
    ->setAutoShutdown(true)
    ->build();

$tracer = $sdkBuilder->getTracerProvider()->getTracer('o11y-day-na-2024');

$app = AppFactory::create();

$app->get('/rolldice', function (Request $request, Response $response) use ($tracer) {
    $context = TraceContextPropagator::getInstance()->extract($request->getHeaders());
    $span = $tracer
        ->spanBuilder('roll')
        ->setParent($context)
        ->startSpan();
    $result = random_int(1,20);
    $response->getBody()->write(strval($result));
    $span
        ->addEvent('rolled dice', ['result' => $result])
        ->end();
    return $response;
});

$app->run();
