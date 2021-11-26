### UDP Echo Svr

A udp server for network testing

### Format

Echo server only accept and response to formatted udp packets 

|  Header   | Mode  | Data  |
|  ----  | ----  | ----  | 
|  1 byte(0xAA)  | 1 byte  |   | 

### Mode

|  Mode   | Start byte  | 2nd byte  | Usage  |
|  ----  | ----  | ----  | ----  |
|  Echo | 0xAA | 0x00 | Reply with the same data |
|  AutoIncrement | 0xAA | 0x01| Reply with a self-increasing number |

### Example

sending a echo packets with netcat

> echo -n -e "\xAA\x00hello world" | nc -4u xx.xx.xx.xx xxx
> 
> hello world
>

sending a auto-increment packets with netcat

> echo -n -e "\xAA\x011111" | nc -4u xx.xx.xx.xx xxx
> 
> 1112