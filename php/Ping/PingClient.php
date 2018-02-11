<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Ping;

/**
 */
class PingClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Ping\HelloReqest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Hello(\Ping\HelloReqest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/ping.Ping/Hello',
        $argument,
        ['\Ping\HelloResponse', 'decode'],
        $metadata, $options);
    }

}
