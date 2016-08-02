/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
	"log"
	"os"
	"time"

	pb "github.com/HuKeping/examples/grpc_health_check/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	out, err := healthCheck(1*time.Second, conn, "")
	if err != nil {
		log.Fatalf("Health/Check error %v, want <nil>", err)
	}
	if out.Status != healthpb.HealthCheckResponse_SERVING {
		log.Fatalf("Got the serving status %v, want SERVING", out.Status)
	}
	log.Println("For \"\" Got the out.Status is:", out.Status)

	wantErr := grpc.Errorf(codes.NotFound, "unknown service")
	if _, err := healthCheck(1*time.Second, conn, "David"); !equalErrors(err, wantErr) {
		log.Fatalf("Health/Check error %v, want error code %d", err, codes.NotFound)
	}
	log.Println("For \"unknown service \"Got the out.Status is:", out.Status)

	// Expected to be SERVING
	out, err = healthCheck(1*time.Second, conn, "grpc.health.v1.Health")
	if err != nil {
		log.Fatalf("Health/Check(_, _) = _, %v, want _, <nil>", err)
	}
	if out.Status != healthpb.HealthCheckResponse_SERVING {
		log.Fatalf("Got the serving status %v, want SERVING", out.Status)
	}
	log.Println("For \"grpc.health.v1.Health\"Got the out.Status is:", out.Status)

	// Expected to be NOT_SERVING
	out, err = healthCheck(1*time.Second, conn, "grpc.health.v1.NotHealth")
	if err != nil {
		log.Fatalf("Health/Check(_, _) = _, %v, want _, <nil>", err)
	}
	if out.Status != healthpb.HealthCheckResponse_NOT_SERVING {
		log.Fatalf("Got the serving status %v, want NOT_SERVING", out.Status)
	}
	log.Println("For \"grpc.health.v1.NotHealth\" Got the out.Status is:", out.Status)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func healthCheck(d time.Duration, cc *grpc.ClientConn, serviceName string) (*healthpb.HealthCheckResponse, error) {
	ctx, _ := context.WithTimeout(context.Background(), d)
	hc := healthpb.NewHealthClient(cc)
	req := &healthpb.HealthCheckRequest{
		Service: serviceName,
	}
	return hc.Check(ctx, req)
}

func equalErrors(l, r error) bool {
	return grpc.Code(l) == grpc.Code(r) && grpc.ErrorDesc(l) == grpc.ErrorDesc(r)
}
