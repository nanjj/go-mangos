// Copyright 2018 The Mangos Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package all is used to register all transports.  This allows a program
// to support all known transports as well as supporting as yet-unknown
// transports, with a single import.
package all

import (
	"nanjj.github.io/nanomsg/go-mangos"
	"nanjj.github.io/nanomsg/go-mangos/transport/inproc"
	"nanjj.github.io/nanomsg/go-mangos/transport/ipc"
	"nanjj.github.io/nanomsg/go-mangos/transport/tcp"
	"nanjj.github.io/nanomsg/go-mangos/transport/tlstcp"
	"nanjj.github.io/nanomsg/go-mangos/transport/ws"
	"nanjj.github.io/nanomsg/go-mangos/transport/wss"
)

// AddTransports adds all known transports to the given socket.
func AddTransports(sock mangos.Socket) {
	sock.AddTransport(tcp.NewTransport())
	sock.AddTransport(inproc.NewTransport())
	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tlstcp.NewTransport())
	sock.AddTransport(ws.NewTransport())
	sock.AddTransport(wss.NewTransport())
}
