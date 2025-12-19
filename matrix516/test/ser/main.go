package main

import (
	"fmt"
	"io"
	"log"
	"time"
	
	// serial library
	"go.bug.st/serial"
)

const (
	BaudRate   = 19200
	DataFormat = 128    // Payload + Header
	Header1    = 0x5A
	Header2    = 0xA5
)

type RS485Packet struct {
	Header   [2]byte // 5A A5
	Payload  []byte  // 126 bytes
}

func receiveData(port serial.Port) {
	log.Println("Start listing...")

	// Create a sufficiently large buffer to read data that may overflow
	readBuffer := make([]byte, DataFormat*2)
	
	// Data buffer
	dataBuffer := make([]byte, 0, DataFormat*4)

	for {
		// Read data from the serial port
		n, err := port.Read(readBuffer)
		if err != nil {
			if err == io.EOF || err.Error() == "The operation timed out" { 
				continue
			}
			log.Printf("COM Port Error: %v\n", err)
			return
		}

		if n > 0 {
			// Append the read data to the data parsing buffer
			dataBuffer = append(dataBuffer, readBuffer[:n]...)
			// fmt.Printf("DEBUG: Read %d bytes, Buffer total length: %d\n", n, len(dataBuffer))

			// Attempt to parse complete data packets from the buffer
			for {
				packet, consumedBytes := parsePacket(dataBuffer)

				if packet != nil {
                    now := time.Now()
                    msec := now.UnixMilli()
					// Successfully parsed a data packet
					fmt.Printf("\n%d Successfully received a complete data packet. (%d bytes):\n", msec, consumedBytes)
					fmt.Printf("   Header: 0x%X 0x%X\n", packet.Header[0], packet.Header[1])
					fmt.Printf("   Payload (first 10 bytes): %X...\n", packet.Payload[:10]) 
					
					// Clear the dataBuffer
					dataBuffer = dataBuffer[consumedBytes:]
					
				} else if consumedBytes > 0 {
					// Header found, but data is incomplete (awaiting next read to complete)
					break
				} else {
					// Header not found or insufficient data
					break
				}
			}
		}

		// Sleep 1ms for prevent excessive CPU usage when no data is available
		time.Sleep(1 * time.Millisecond)
	}
}

// --------------------------------------------------------------------------------
// Data Analysis
// --------------------------------------------------------------------------------

// parsePacket Parse a data packet from the buffer
// Return: (*RS485Packet, consumedBytes)
// - packet: Successfully parsed data package (nil if not found or incomplete)
// - consumedBytes: Number of bytes consumed from the buffer (used for trimming dataBuffer)
func parsePacket(dataBuffer []byte) (*RS485Packet, int) {
	bufferLen := len(dataBuffer)
	
	// Locate Header (5A A5)
	for i := 0; i < bufferLen-1; i++ {
		if dataBuffer[i] == Header1 && dataBuffer[i+1] == Header2 {
			
			// Find Header
			headerIndex := i
			remainingDataLen := bufferLen - headerIndex
			
			// Verify data length completeness (Header 2 bytes + Payload 126 bytes = 128 bytes)
			if remainingDataLen >= DataFormat {
				// Data length is complete; proceed to parse the data packet
				
				// Total data packet length
				packetLen := DataFormat 
				
				packet := &RS485Packet{
					Header:   [2]byte{dataBuffer[headerIndex], dataBuffer[headerIndex+1]},
					Payload:  dataBuffer[headerIndex+2 : headerIndex+packetLen], // 126 bytes
				}
				
				// Returns the parsed data packet and the number of bytes consumed 
                // The byte count includes the entire packet length from the header to the end
				return packet, headerIndex + packetLen
				
			} else {
				// Found Header (5A A5) but subsequent data length is insufficient  
                // At this point, move the Header and following data to the buffer's start and wait for more data  
                // Return nil, indicating that i bytes of data before the Header have been skipped
				return nil, i
			}
		}
	}
	
	// Header not found
	return nil, 0
}


func main() {
	// - Windows: "COM3", "COM4", ...
	// - Linux: "/dev/ttyUSB0", "/dev/ttyS0", ...
	portName := "/dev/ttyS1" 
	
	mode := &serial.Mode{
		BaudRate: BaudRate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Fatalf("Error Unable to open the port %s: %v", portName, err)
	}
	defer port.Close()

	log.Printf("Successfully open the port %s, Baud：%d\n", portName, BaudRate)

	if err := port.SetReadTimeout(50 * time.Millisecond); err != nil {
		log.Fatalf("Unable to set read timeout: %v", err)
	}

	// Goroutine for data receive
	go receiveData(port)
	
	// keep the main thread running to wait for incoming data
	select {} 
}