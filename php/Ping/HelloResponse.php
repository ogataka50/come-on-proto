<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ping.proto

namespace Ping;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>ping.HelloResponse</code>
 */
class HelloResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string resMessage = 1;</code>
     */
    private $resMessage = '';

    public function __construct() {
        \GPBMetadata\Ping::initOnce();
        parent::__construct();
    }

    /**
     * Generated from protobuf field <code>string resMessage = 1;</code>
     * @return string
     */
    public function getResMessage()
    {
        return $this->resMessage;
    }

    /**
     * Generated from protobuf field <code>string resMessage = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setResMessage($var)
    {
        GPBUtil::checkString($var, True);
        $this->resMessage = $var;

        return $this;
    }

}

