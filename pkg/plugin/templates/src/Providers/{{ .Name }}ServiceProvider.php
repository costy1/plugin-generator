<?php

namespace {{ .Name }}\Providers;

use Plenty\Plugin\ServiceProvider;

/**
 * Class {{ .Name }}ServiceProvider
 * @package {{ .Name }}\Providers
 */
class {{ .Name }}ServiceProvider extends ServiceProvider
{
    /**
    * Register the route service provider
    */
    public function register()
    {
        $this->getApplication()->register({{ .Name }}RouteServiceProvider::class);
    }

    public function boot()
    {
    }
}