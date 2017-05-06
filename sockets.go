package main

import (
	"fmt"
	"net"
)

// Linux /proc/net/tcp
//   sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode

// From: https://stackoverflow.com/questions/12142925/what-do-values-mean-in-inode-column-proc-net-tcp6
// Format: "%4d: %08X:%04X %08X:%04X %02X %08X:%08X %02X:%08lX %08X %5d %8d %lu %d %p %lu %lu %u %u %d%n",
// i, src, srcp, dest, destp, sk->sk_state,
//          tp->write_seq - tp->snd_una,
//          rx_queue,
//          timer_active,
//      jiffies_to_clock_t(timer_expires - jiffies),
//          icsk->icsk_retransmits,
//          sock_i_uid(sk),
//          icsk->icsk_probes_out,
//          sock_i_ino(sk),
//          atomic_read(&sk->sk_refcnt), sk,
//          jiffies_to_clock_t(icsk->icsk_rto),
//          jiffies_to_clock_t(icsk->icsk_ack.ato),
//          (icsk->icsk_ack.quick << 1) | icsk->icsk_ack.pingpong,
//          tp->snd_cwnd,
//          tcp_in_initial_slowstart(tp) ? -1 : tp->snd_ssthresh,
//          len);

type SocketInfo struct {
	sl, localPort, remotePort, st, tr, uid                            int32
	local, remote, txQueue, rxQueue, tmWhen, inode, retrnsmt, timeout int64
	theRest                                                           string
	remoteIP, localIP                                                 net.IP
}

// Expects line corresponding to Linux /proc/net/tcp
func NewSocketInfo(l string) *SocketInfo {
	var s SocketInfo

	fmt.Sscanf(l, "%5d: %08X:%04X %08X:%04X %02X %08X:%08X %02X:%08X %08X %5d %7d %6d %s", &s.sl, &s.local, &s.localPort, &s.remote, &s.remotePort, &s.st, &s.txQueue, &s.rxQueue, &s.tr, &s.tmWhen, &s.retrnsmt, &s.uid, &s.timeout, &s.inode, &s.theRest)

	s.remoteIP = inet_ntoa(s.remote)
	s.localIP = inet_ntoa(s.local)
	//fmt.Printf("foo    %+v\n", s)

	return &s
}

// inet_ntoa: From: https://groups.google.com/d/msg/golang-nuts/v4eJ5HK3stI/Tah15fMOC80J Author: Paul van Brouwershaven
func inet_ntoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)
	return net.IPv4(bytes[0], bytes[1], bytes[2], bytes[3])
}
