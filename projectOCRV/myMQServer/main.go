package myMQServer

import pb "myMQServer/api/server/producerMQ"

type serverMQ struct {
	pb.UnimplementedProducerMQServer
}
