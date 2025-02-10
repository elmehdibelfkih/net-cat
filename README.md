# TCP-Chat (NetCat Alternative)

## 📌 Project Overview
This project is a recreation of the NetCat command-line utility in a **server-client architecture**. It allows a server to run on a specified port and listen for incoming connections while multiple clients can connect to it, forming a **group chat**.

The implementation closely follows the behavior of the original `nc` (NetCat) system command, which is used for TCP, UDP, and UNIX-domain socket communications. For more details, refer to `man nc`.

---

## 🚀 Features
- **TCP Connection**: Server handles multiple clients simultaneously (1-to-many relation).
- **User Identification**: Each client must provide a name upon connection.
- **Connection Management**: Limits the number of simultaneous clients to **10**.
- **Messaging System**:
  - Clients can send messages to the chat.
  - Messages include timestamps and sender names.
  - Clients cannot send empty messages.
  - All previous messages are sent to a newly connected client.
  - Clients receive notifications when users join or leave.
- **Robustness**:
  - Clients remain connected even if another client disconnects.
  - Error handling for both server and client sides.
- **Default Port Handling**: If no port is specified, the server runs on port **8989**.
- **Go Concurrency**: Utilizes **Go-routines**, and **mutexes** for efficient concurrent handling.
<!-- - **Testing**: Recommended to include **unit tests** for both server and client connections. -->
<!-- channels -->

---

## Members
- Said Oubaaisse
- Larbi Mergaoui
- El Mehdi Belfkih

---

## 📂 Allowed Go Packages
The project uses the following standard Go packages:
- `io`
- `log`
- `os`
- `fmt`
- `net`
- `sync`
- `time`
- `strings`
---

## 🔧 Installation & Usage
### 1️⃣ Run the TCP Server
Start the server on the default port **8989**:
```sh
$ go run .
Listening on the port :8989
```
Start the server on a specific port (e.g., **2525**):
```sh
$ go run . 2525
Listening on the port :2525
```
If the user provides an incorrect number of arguments:
```sh
$ go run . 2525 localhost
[USAGE]: ./TCPChat $port
```

### 2️⃣ Connect a Client
Clients can connect to the chat using **NetCat (nc)**:
```sh
$ nc <server-ip> <port>
```
Example:
```sh
$ nc localhost 2525
```
Upon connection, the client will see a **Linux ASCII logo** and be prompted to enter their name:
```
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:
```
If the client enters an **empty name**, the connection will be rejected.

### 3️⃣ Chat Example
#### **Client 1 (Yenlik) joins and starts chatting:**
```
[2020-01-20 16:03:43][Yenlik]: hello
[2020-01-20 16:03:46][Yenlik]: How are you?
```
#### **Client 2 (Lee) joins:**
```
Lee has joined our chat...
[2020-01-20 16:04:32][Lee]: Hi everyone!
[2020-01-20 16:04:35][Lee]: How are you?
[2020-01-20 16:04:41][Yenlik]: great, and you?
[2020-01-20 16:04:44][Lee]: good!
```
#### **Client 2 (Lee) leaves:**
```
Lee has left our chat...
```

---

## 🛠️ Project Structure

```
.
├── data
│   ├── linux.logo
│   └── logs.txt
├── go.mod
├── internal
│   ├── client.go
│   ├── config.go
│   └── server.go
├── main.go
└── README.md    # Project documentation
```

---

## 📜 License
This project is open-source and licensed under the **MIT License**.

---

## 🤝 Contributing
Contributions are welcome! Feel free to submit **issues** or **pull requests**.

---

## 🏆 Acknowledgments
Inspired by the `nc` (NetCat) system command.

---

## 📞 Contact
For any questions or feedback, feel free to reach out!

