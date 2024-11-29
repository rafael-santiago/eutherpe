// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
package mdns

// INFO(Rafael): This is a well-simple MDNS facility made to be easily embed into a software.
//               The idea is being simple, not BLOATED. No frills. It only should be ran in
//               scenarios were there would not be jerk people trying to explode everything.
//               It is suitable to an embed stuff running into your friendly LAN/WLAN.
//
//               The resolutions are made based on MDNSHost struct.
//
//               This struct expects:
//
//                      - a hostname;
//                      - an addr expressed in bytes (ipv4 or ipv6);
//                      - a TTL for THIS specific resolution (YES);
//
//               Powering on MDNS stuff is fairly simple, you should use MDNSServerStart().
//
//               This function expects:
//                      - a slice of []MDNSHost (containing all possible resolution
//                                               for your LAN/WLAN);
//                      - an boolean channel (that you will be set to true when it will be time
//                                            to go);
//
//               Follows a sample that will run for 2 minutes a MDNS server capable of
//               resolving 'deepthrought.local' to '42.42.42.42' and 'hal9000.local' to
//               '9.0.0.0'.
//
//              // all import trinket and blah blah blah
//
//               func main() {
//                      MDNSHosts := make([]MDNSHost, 0)
//                      MDNSHosts = append(MDNSHosts, MDNSHost { "deepthrought.local", []byte { 42, 42, 42, 42 }, 600, })
//                      MDNSHosts = append(MDNSHosts, MDNSHost { "hal9000.local", []byte { 9, 0, 0, 0 }, 9000, })
//                      jaAcabeiJessica := make(chan bool)
//                      MDNSServerStart(MDNSHosts, jaAcabeiJessica)
//                      time.Sleep(2 * time.Minute)
//                      jaAcabeiJessica <- true
//                      os.Exit(0)
//               }
//
//               Originally, I have written this piece of code for making easy the zero conf of
//               an embed weekend project of mine and, smooth sailing, it has worked as I
//               expected. Direct and simple. Maybe it could be useful to you, too.
//
//               Feel free on copying and pasting (even blindly) and using it.
//
//               Enjoy! -- Rafael

import (
    "os"
    "net"
    "log"
    "fmt"
    "time"
)

// INFO(Rafael): Almost are useless for your minimalist goal but here we go.
const (
    MDNSQTypeA = iota + 1
    MDNSQTypeNS
    MDNSQTypeMD
    MDNSQTypeMF
    MDNSQTypeCName
    MDNSQTypeSOA
    MDNSQTypeMB
    MDNSQTypeMG
    MDNSQTypeMR
    MDNSQTypeNULL
    MDNSQTypeWKS
    MDNSQTypePTR
    MDNSQTypeHINFO
    MDNSQTypeMINFO
    MDNSQTypeMX
    MDNSQTypeTXT
    MDNSQTypeAAAA = 28
    MDNSQTypeAXFR = 252
    MDNSQTypeMAILB
    MDNSQTypeMAILA
    MDNSQTypeALL
)

// INFO(Rafael): Almost are useless for your minimalist goal but here we go.
const (
    MDNSQClassIN = iota + 1
    MDNSQClassCS
    MDNSQClassCH
    MDNSQClassHS
    MDNSQClassAny = 255
)

type MDNSQuestion struct {
    QName []byte
    QType uint16
    UnicastResp uint8
    QClass uint16
}

type MDNSResourceRecord struct {
    QName []byte
    QType uint16
    QClass uint16
    TTL uint32
    RDLength uint16
    RData []byte
}

type MDNSPacket struct {
    ID uint16
    Flags uint16
    Qdcount uint16
    Ancount uint16
    Nscount uint16
    Arcount uint16
    Questions []MDNSQuestion
    Answers []MDNSResourceRecord
    RawPkt []byte
}

type MDNSHost struct {
    Name string
    Addr []byte
    TTL uint32
}

type MDNSResolution struct {
    QName []byte
    Addr []byte
    TTL uint32
    Error error
}

func parseMDNSPacket(wireBytes []byte) (MDNSPacket, error) {
    wireBytesAmount := len(wireBytes)
    if wireBytesAmount < 12 {
        return MDNSPacket{}, fmt.Errorf("Malformed MDNS packet.")
    }
    MDNSPkt        := MDNSPacket{}
    MDNSPkt.ID      = uint16(wireBytes[ 0]) << 8 | uint16(wireBytes[ 1])
    MDNSPkt.Flags   = uint16(wireBytes[ 2]) << 8 | uint16(wireBytes[ 3])
    MDNSPkt.Qdcount = uint16(wireBytes[ 4]) << 8 | uint16(wireBytes[ 5])
    MDNSPkt.Ancount = uint16(wireBytes[ 6]) << 8 | uint16(wireBytes[ 7])
    MDNSPkt.Nscount = uint16(wireBytes[ 8]) << 8 | uint16(wireBytes[ 9])
    MDNSPkt.Arcount = uint16(wireBytes[10]) << 8 | uint16(wireBytes[11])
    if MDNSPkt.Qdcount == 0 &&
       MDNSPkt.Ancount == 0 &&
       MDNSPkt.Nscount == 0 &&
       MDNSPkt.Arcount == 0 {
        return MDNSPacket{}, fmt.Errorf("MDNS packet has no records.")
    }
    // INFO(Rafael): We are only giving real support for questions and/or answers
    if MDNSPkt.Qdcount == 0 &&
       MDNSPkt.Ancount == 0 {
        return MDNSPkt, nil
    }
    if MDNSPkt.Qdcount > 0 {
        MDNSPkt.Questions = make([]MDNSQuestion, MDNSPkt.Qdcount)
    }
    if MDNSPkt.Ancount > 0 {
        MDNSPkt.Answers = make([]MDNSResourceRecord, MDNSPkt.Ancount)
    }
    w := 12
    wireBytesAmount -= w
    for q := uint16(0); w < wireBytesAmount && q < MDNSPkt.Qdcount; q++ {
        wStart := w
        for ; w != wireBytesAmount && wireBytes[w] != 0; w++ {
        }
        if w == wireBytesAmount {
            return MDNSPacket{}, fmt.Errorf("Malformed MDNS question.")
        }
        MDNSPkt.Questions[q].QName = wireBytes[wStart:w+1]
        w += 1
        if (w + 3) >= wireBytesAmount {
            return MDNSPacket{}, fmt.Errorf("Malformed MDNS question.")
        }
        MDNSPkt.Questions[q].QType = uint16(wireBytes[w]) << 8 | uint16(wireBytes[w + 1])
        qclass := uint16(wireBytes[w + 2]) << 8 | uint16(wireBytes[w + 3])
        MDNSPkt.Questions[q].UnicastResp = uint8(qclass >> 15)
        qclass <<= 1
        MDNSPkt.Questions[q].QClass = qclass
        w += 3
    }
    // WARN(Rafael): This package only works as mDNS solver we do not give a shit
    //               for answers. We answer. Anyway the idea follows commented but
    //               you need to handle compression stuff "c0 blauZ....".
    /*
    w += 1
    for a := uint16(0) ; w < wireBytesAmount && a < MDNSPkt.Ancount; a++ {
        wStart := w
        for ; w != wireBytesAmount && wireBytes[w] != 0; w++ {
        }
        if w == wireBytesAmount {
            return MDNSPacket{}, fmt.Errorf("Malformed MDNS answer.")
        }
        MDNSPkt.Answers[a].QName = wireBytes[wStart:w+1]
        w += 1
        MDNSPkt.Answers[a].QType = uint16(wireBytes[w]) << 8 | uint16(wireBytes[w + 1])
        MDNSPkt.Answers[a].QClass = uint16(wireBytes[w + 2]) << 8 | uint16(wireBytes[w + 3])
        MDNSPkt.Answers[a].TTL = uint32(wireBytes[w + 4]) << 24 |
                                 uint32(wireBytes[w + 5]) << 16 |
                                 uint32(wireBytes[w + 6]) <<  8 |
                                 uint32(wireBytes[w + 7])
        MDNSPkt.Answers[a].RDLength = uint16(wireBytes[w + 8]) << 8 |
                                      uint16(wireBytes[w + 9])
        if (w + 10) >= wireBytesAmount {
            return MDNSPacket{}, fmt.Errorf("Malformed MDNS answer.")
        }
        w += 10
        nextOff := int(MDNSPkt.Answers[a].RDLength)
        if (w + nextOff) >= wireBytesAmount {
            return MDNSPacket{}, fmt.Errorf("Malformed MDNS answer.")
        }
        MDNSPkt.Answers[a].RData = wireBytes[w:w + nextOff]
        w += nextOff
    }
    */
    MDNSPkt.RawPkt = wireBytes
    return MDNSPkt, nil
}

func getAncountFromResolutions(resolutions []MDNSResolution) uint16 {
    if len(resolutions) == 0 {
        return 0
    }
    ancount := uint16(0)
    for _, resolution := range resolutions {
        if resolution.Error == nil {
            ancount++
        }
    }
    return ancount
}

func makeMDNSAnswer(MDNSPkt *MDNSPacket, resolutions []MDNSResolution, anCount uint16) error {
    if MDNSPkt.Qdcount == 0 || len(MDNSPkt.Questions) == 0 {
        return fmt.Errorf("MDNS packet has no questions.")
    }
    // INFO(Rafael): On Apple stuff, I observed that it includes
    //               in one packet more than one question and
    //               the requestor only accepts the response
    //               when it has the same count of answers.
    MDNSPkt.Ancount = anCount
    MDNSPkt.Qdcount = 0
    MDNSPkt.Flags = 0x8400
    MDNSPkt.Answers = make([]MDNSResourceRecord, MDNSPkt.Ancount)
    for a := uint16(0); a < MDNSPkt.Ancount; a++ {
        MDNSPkt.Answers[a].QName = resolutions[a].QName
        if len(resolutions[a].Addr) == 4 {
            MDNSPkt.Answers[a].QType = MDNSQTypeA
        } else {
            MDNSPkt.Answers[a].QType = MDNSQTypeAAAA
        }
        // WARN(Rafael): Windows boxes accepts cache-flush flag, on apple stuff it rejects
        MDNSPkt.Answers[a].QClass = /*0x80 |*/ MDNSQClassIN
        MDNSPkt.Answers[a].RDLength = uint16(len(resolutions[a].Addr))
        MDNSPkt.Answers[a].RData = resolutions[a].Addr
        MDNSPkt.Answers[a].TTL = resolutions[a].TTL
    }
    MDNSPkt.Questions = make([]MDNSQuestion, 0)
    return nil
}

func makeMDNSPacket(MDNSPkt MDNSPacket) []byte {
    var recsSize int
    for _, question := range MDNSPkt.Questions {
        recsSize += 4 + len(question.QName)
    }
    for _, answer := range MDNSPkt.Answers {
        recsSize += 10 + len(answer.QName) + int(answer.RDLength)
    }
    datagram := make([]byte, 12 + recsSize)
    datagram[ 0] = byte((MDNSPkt.ID >> 8) & 0xFF)
    datagram[ 1] = byte(MDNSPkt.ID & 0xFF)
    datagram[ 2] = byte((MDNSPkt.Flags >> 8) & 0xFF)
    datagram[ 3] = byte(MDNSPkt.Flags & 0xFF)
    datagram[ 4] = byte((MDNSPkt.Qdcount >> 8) & 0xFF)
    datagram[ 5] = byte(MDNSPkt.Qdcount & 0xFF)
    datagram[ 6] = byte((MDNSPkt.Ancount >> 8) & 0xFF)
    datagram[ 7] = byte(MDNSPkt.Ancount & 0xFF)
    datagram[ 8] = byte((MDNSPkt.Nscount >> 8) & 0xFF)
    datagram[ 9] = byte(MDNSPkt.Nscount & 0xFF)
    datagram[10] = byte((MDNSPkt.Arcount >> 8) & 0xFF)
    datagram[11] = byte(MDNSPkt.Arcount & 0xFF)
    d := 12
    for _, question := range MDNSPkt.Questions {
        copy(datagram[d:], question.QName)
        d += len(question.QName)
        datagram[d] = byte((question.QType >> 8) & 0xFF)
        datagram[d + 1] = byte(question.QType & 0xFF)
        datagram[d + 2] = byte((question.QClass >> 8) & 0xFF)
        datagram[d + 3] = byte(question.QClass & 0xFF)
        d += 4
    }
    for _, answer := range MDNSPkt.Answers {
        copy(datagram[d:], answer.QName)
        d += len(answer.QName)
        datagram[d] = byte((answer.QType >> 8) & 0xFF)
        datagram[d + 1] = byte(answer.QType & 0xFF)
        datagram[d + 2] = byte((answer.QClass >> 8) & 0xFF)
        datagram[d + 3] = byte(answer.QClass & 0xFF)
        datagram[d + 4] = byte((answer.TTL >> 24) & 0xFF)
        datagram[d + 5] = byte((answer.TTL >> 16) & 0xFF)
        datagram[d + 6] = byte((answer.TTL >>  8) & 0xFF)
        datagram[d + 7] = byte(answer.TTL & 0xFF)
        datagram[d + 8] = byte((answer.RDLength >> 8) & 0xFF)
        datagram[d + 9] = byte(answer.RDLength & 0xFF)
        d += 10
        copy(datagram[d:], answer.RData)
        d += int(answer.RDLength)
    }
    return datagram
}

func getCompressedName(offset int, rawPkt []byte) string {
    return getName(rawPkt[offset:], rawPkt)
}

func getName(qName []byte, rawPkt []byte) string {
    if len(qName) == 0 {
        return ""
    }
    if qName[0] == 0xC0 && len(qName) == 2 {
        return getCompressedName(int(qName[1]), rawPkt)
    }
    var name string
    for w := 0; w < len(qName); {
        blobSize := int(qName[w])
        if w > 0 && blobSize > 0 {
            name += "."
        }
        w += 1
        name += string(qName[w:w+blobSize])
        w += blobSize
    }
    return name
}

func getQueriedNames(MDNSPkt MDNSPacket) []string {
    names := make([]string, 0)
    for _, question := range MDNSPkt.Questions {
        name := getName(question.QName, MDNSPkt.RawPkt)
        if len(name) > 0 {
            names = append(names, name)
        }
    }
    return names
}

func resolveAddrs(MDNSPkt MDNSPacket, MDNSHosts []MDNSHost) []MDNSResolution {
    qnames := getQueriedNames(MDNSPkt)
    resolutions := make([]MDNSResolution, 0)
    for q, qname := range qnames {
        resolution := MDNSResolution { []byte{}, []byte{}, 0, fmt.Errorf("No addr resolution for '%s'.", qname) }
        for _, host := range MDNSHosts {
            if host.Name == qname {
                resolution = MDNSResolution { MDNSPkt.Questions[q].QName, host.Addr, host.TTL, nil }
                break
            }
        }
        resolutions = append(resolutions, resolution)
    }
    return resolutions
}

func doMDNSServerRunN(proto, listenAddr string,
                      MDNSHosts []MDNSHost, goinHome chan bool) error {
    addr, err := net.ResolveUDPAddr(proto, listenAddr)
    if err != nil {
        log.Fatal(err)
    }
    l, err := net.ListenMulticastUDP(proto, nil, addr)
    if err != nil {
        return err
    }
    l.SetReadBuffer(0xFFFF)
    for {
        select {
            case <-goinHome:
                break
            default:
        }
        b := make([]byte, 0xFFFF)
        l.SetReadDeadline(time.Now().Add(3 * time.Second))
        l.SetDeadline(time.Now().Add(3 * time.Second))
        bytesTotal, unicastAddr, err := l.ReadFromUDP(b)
        if bytesTotal <= 0 {
            continue
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
            continue
        }
        MDNSPkt, err := parseMDNSPacket(b)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
            continue
        }
        if MDNSPkt.Qdcount == 0 || (MDNSPkt.Flags & 0x800) != 0 {
            continue
        }

        resolutions := resolveAddrs(MDNSPkt, MDNSHosts)
        anCount := getAncountFromResolutions(resolutions)
        if anCount == 0 {
            continue
        }
        // INFO(Rafael): I think that almost all implementations of clients does not
        //               mind about avoiding multicasting abuse on their networks.
        shouldUnicast := (MDNSPkt.Questions[0].UnicastResp != 0)
        makeMDNSAnswer(&MDNSPkt, resolutions, anCount)
        MDNSReply := makeMDNSPacket(MDNSPkt)
        var addr *net.UDPAddr
        if !shouldUnicast {
            addr, err = net.ResolveUDPAddr(proto, listenAddr)
        } else {
            addr = unicastAddr
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
            continue
        }
        conn, err := net.DialUDP(proto, nil, addr)
        if err == nil {
            _, err = conn.Write(MDNSReply)
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
            continue
        }
        if conn != nil {
            conn.Close()
        }
    }

}

func doMDNSServerRun4(MDNSHosts []MDNSHost, goinHome chan bool) error {
    return doMDNSServerRunN("udp4", "224.0.0.251:5353", MDNSHosts, goinHome)
}

func doMDNSServerRun6(MDNSHosts []MDNSHost, goinHome chan bool) error {
    return doMDNSServerRunN("udp6", "[FF02::FB]:5353", MDNSHosts, goinHome)
}

func MDNSServerStart(MDNSHosts []MDNSHost, goinHome chan bool) error {
    if len(MDNSHosts) == 0 {
        return fmt.Errorf("There is no resolutions to deliver.")
    }
    mdnsHosts4 := make([]MDNSHost, 0)
    mdnsHosts6 := make([]MDNSHost, 0)
    for _, host := range MDNSHosts {
        addrLen := len(host.Addr)
        if addrLen == 4 {
            mdnsHosts4 = append(mdnsHosts4, host)
        } else if addrLen == 16 {
            mdnsHosts6 = append(mdnsHosts6, host)
        }
    }
    if len(mdnsHosts4) > 0 {
        go doMDNSServerRun4(mdnsHosts4, goinHome)
    }
    if len(mdnsHosts6) > 0 {
        go doMDNSServerRun6(mdnsHosts6, goinHome)
    }
    return nil
}
