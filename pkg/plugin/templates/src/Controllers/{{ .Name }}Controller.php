<?php

namespace {{ .Name }}\Controllers;

use Plenty\Plugin\Controller;
use Plenty\Plugin\Templates\Twig;

class {{ .Name }}Controller extends Controller
{
    /**
     * @return string
     */
    public function test():string
    {
        return "result from the test route";
    }
}