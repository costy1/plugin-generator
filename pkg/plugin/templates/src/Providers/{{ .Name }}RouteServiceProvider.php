<?php

namespace {{ .Name }}\Providers;

use Plenty\Plugin\RouteServiceProvider;
use Plenty\Plugin\Routing\ApiRouter;

/**
 * Class {{ .Name }}RouteServiceProvider
 * @package {{ .Name }}\Providers
 */
class {{ .Name }}RouteServiceProvider extends RouteServiceProvider
{
    /**
     * @param ApiRouter $apiRouter
     */
    public function map(ApiRouter $apiRouter)
    {
        $apiRouter->version(['v1'], ['namespace' => '{{ .Name }}\Controllers', 'middleware' => ['oauth']], function ($router) {
            $router->get('test', '{{ .Name }}Controller@test');
        });
    }
}