package main

import (
	"fmt"
	"log"
    "io/ioutil"
    "os"
    "bufio"
    "path/filepath"
    // "io"
    "strings"

	"golang.org/x/crypto/ssh"
)

type Connection struct {
	*ssh.Client
}

var sess *ssh.Session
func main() {
	conn, err := Connect("34.222.11.245:22", "ubuntu", "/Users/kakao_ent/Downloads/LightsailDefaultKey-us-west-2.pem")
	if err != nil {
		log.Fatal(err)
	}

    conn.SendCommands("sudo apt-get update", 1)
    conn.SendCommands(`
        sudo apt-get install  \
        bino \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg-agent \
        software-properties-common;
        `, 2)
    conn.SendCommands("sudo apt-get install bino", 2)
    // conn.SendCommands("curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -", 3)
    // fmt.Println(3)
    // conn.SendCommands(`
    //     sudo add-apt-repository \
    //     "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
    //     $(lsb_release -cs) \
    //     stable"`)
   // conn.SendCommands("sudo apt-get update")
   // conn.SendCommands("sudo apt-get install -y docker-ce docker-ce-cli containerd.io")
   // conn.SendCommands("sudo docker run hello-world")

	// fmt.Println(string(output))
}

func Connect(addr, user, password string) (*Connection, error) {
    host := "34.222.11.245"
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}
	if hostKey == nil {
		log.Fatalf("no hostkey for %s", host)
	}

    key, err := ioutil.ReadFile(password)
    if err != nil {
        log.Fatalf("unable to read private key: %v", err)
    }

    signer, err := ssh.ParsePrivateKey(key)
    if err != nil {
        log.Fatalf("unable to parse private key: %v", err)
    }

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
            ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
        fmt.Println(err)
		return nil, err
	}

	return &Connection{conn}, nil

}

func (conn *Connection) SendCommands(cmd string, i int) (error) {
    var session *ssh.Session
    if sess == nil {
        sess, err := conn.NewSession()
        if err != nil {
            log.Fatal(err)
        }
        session = sess
    }
    defer session.Close()
	// session, err := conn.NewSession()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}


    err := session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		return err
	}

	// in, err := session.StdinPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// out, err := session.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }

    // buf := make([]byte, 4096)
	// go func(out io.Reader) {
		// r := bufio.NewReader(out)
    //     for {
    //         n, err := r.Read(buf)
    //         if err == io.EOF {
    //             in.Write([]byte("n\n"))
    //             break
    //         }
    //         fmt.Print(string(buf[:n]))
    //     }
    // }(out)


    // inChan := make(chan bool)
    // in, _ := session.StdinPipe()
    // go func(in io.WriteCloser) {
    //     for {

    //     fmt.Println("GETTING USER INPUT")
    //             reader := bufio.NewReader(os.Stdin)
    //             text, _ := reader.ReadString('\n')
    //             fmt.Println("USER INPUT: ", text)
    //             in.Write([]byte(text))
    //     }
    // }(in)
    // var b bytes.Buffer
    if session.Stdout == nil {
        session.Stdout = os.Stdout
    }
    if session.Stdin == nil {
        session.Stdin = os.Stdin
    }

	err = session.Run(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute command '%s' on server: %v", cmd, err)
	}

    // session.Close()
    // session.Stdout = nil
    // session.Stdin = nil
    // if err := session.Wait(); err != nil {
    //     log.Fatal(err)
    // }
	return err
}

