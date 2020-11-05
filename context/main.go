package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

/* context - нужен в двух случаях
1. Когда нужно высвободить ресурсы gorutines. Пример: Пользователь сделал запрос, после чего прервал соединение.
2. Передача каких-то значений. Пример: Есть http сервер, записываем token или id пользователя. Чтобы доставать данные token или id

В коде будет реализован сервер, который слушает http соединение и отслеживает ctr+c shutdown от системы при помощи context
ctr+c shutdown - закрывает коннект и tcp соединение(graceful shutdown proc)*/

func main() {
	//контекст - во время инициализации. возвращающий функцию, вызывая которую, мы завершаем программу и передаем gorutines event - сообщающий о завершении
	ctx, cancel := context.WithCancel(context.Background())
	//слушаем event о завершении программы в отдельной gorutine
	go handleSignals(cancel)

	if err := startServer(ctx); err != nil {
		log.Fatal(err)
	}
}

func handleSignals(cancel context.CancelFunc) {
	//создаем канал для прослушивания сигналов
	sigCh := make(chan os.Signal)
	//все уведомления о сигналах в этот канал и опишем в скобках те сигналы, что нас интеересуют
	signal.Notify(sigCh, os.Interrupt)
	//улавливаем все сигналы
	for {
		sig := <-sigCh
		// проверка, что сигнал тот
		switch sig {
		case os.Interrupt:
			cancel()
			return
		}
	}
}

func startServer(ctx context.Context) error {
	//блокирующая функция. Контекст используется, чтобы ее остановить.
	laddr, err := net.ResolveTCPAddr("tcp", ":8080")
	//если порт занят или иная ошибка, возвращаем err
	if err != nil {
		return err
	}

	// слушаем порт 8080
	l, err := net.ListenTCP("tcp", laddr)
	//если порт занят или иная ошибка, возвращаем err
	if err != nil {
		return err
	}

	//когда функция заканчивается, закрываем порт
	defer l.Close()

	// принимаем все соединения, которые приходят через порт
	for {
		//
		select {
		case <-ctx.Done():
			log.Println("server stop")
			return nil
		default:
			if err := l.SetDeadline(time.Now().Add(time.Second)); err != nil {
				return err
			}

			_, err := l.Accept()
			if err != nil {
				if os.IsTimeout(err) {
					continue
				}

				return err
			}
			log.Println("new client connected")
		}

	}
}
