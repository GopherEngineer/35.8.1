package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const (
	// Служба будет слушать запросы на всех IP-адресах
	// компьютера на порту 8080.
	addr = "0.0.0.0:8080"
	// Протокол сетевой службы.
	proto = "tcp4"
)

func main() {

	// Читаем файл с поговорками.
	data, err := os.ReadFile("proverbs.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Преобразуем прочитанное и делим на отдельные поговорки.
	proverbs := strings.Split(string(data), "\n")

	// Запускаем сетевую службу.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	// Принимаем новые подключения и работаем
	// с ними в отдельных потоках для обеспечения
	// множественных одновременных дальнейших подключений.
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handler(conn, proverbs)
	}

}

// Обработчик подключения записывает случайную поговорку
// раз в 3 секунды. В случае завершения подключения со стороны
// клиента обрабатывам ошибку записи и завершаем работу обработчика.
func handler(conn net.Conn, proverbs []string) {
	defer conn.Close()

	for {
		time.Sleep(time.Second * 3)

		proverb := proverbs[rand.Intn(len(proverbs))]

		if _, err := conn.Write([]byte(proverb + "\n")); err != nil {
			break
		}
	}
}
