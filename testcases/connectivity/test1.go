package main


import "fmt"
import "github.com/Shopify/sarama"
import "sync"
import "os"
import "os/signal"
import "log"
//import "encoding/json"

type ProducerMessage struct{
	Topic string
	Value sarama.StringEncoder
}

func main() {



	/*var devkafkainstances, UATkafkainstances, prodkafkainstances [][]string
devkafkainstances = [][]string{{"192.168.27.184:9092"},{"192.168.27.188:9092"},{"192.168.27.189:9092"}}
UATkafkainstances = [][]string{{"142.136.83.16:9092"},{"142.136.83.17:9092"},{"142.136.83.23:9092"},{"142.136.83.24:9092"}}
prodkafkainstances = [][]string{{"44.128.168.59:9092"},{"44.128.168.60:9092"},{"44.128.168.65:9092"},{"44.128.168.66:9092"}}
*/


	/*var devkafkainstances, UATkafkainstances, prodkafkainstances [][]string
devkafkainstances = [][]string{{"142.136.38.124:9092"},{"142.136.38.125:9092"},{"142.136.38.126:9092"},{"142.136.38.127:9092"},{"142.136.38.128:9092"}}
UATkafkainstances = [][]string{{"142.136.58.101:9092"},{"142.136.58.102:9092"},{"142.136.58.103:9092"},{"142.136.58.104:9092"},{"142.136.58.105:9092"},}
prodkafkainstances = [][]string{{"142.136.38.211:9092"},{"142.136.38.212:9092"},{"142.136.38.213:9092"},{"142.136.38.234:9092"},{"142.136.38.235:9092"},{"142.136.38.236:9092"},{"142.136.38.237:9092"},{"142.136.38.238:9092"},{"142.136.38.239:9092"},{"142.136.38.240:9092"},{"142.136.38.241:9092"},{"142.136.38.242:9092"},{"142.136.38.243:9092"}}
*/

	UATkafkainstances := [][]string{{"142.136.38.124:9092"},{"142.136.38.125:9092"},{"142.136.38.126:9092"},{"142.136.38.127:9092"},{"142.136.38.128:9092"}}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer := getProducer(UATkafkainstances, config)

	/*
	for _,instanceip := range devkafkainstances{
			producer, err := sarama.NewAsyncProducer(instanceip, config)
			if err != nil {
			fmt.Printf("Cannot connect to:%s", instanceip)
				panic(err)
			}
			fmt.Printf("Successfully Connected to %s, %s", instanceip, producer.Successes())

	}*/
	/*
	for _,instanceip := range UATkafkainstances{
			_, err := sarama.NewAsyncProducer(instanceip, config)
			if err != nil {
			fmt.Printf("Cannot connect to:%s", instanceip)
				panic(err)
			}
			fmt.Printf("Successfully Connected to %s", instanceip)

	}


	for _,instanceip := range prodkafkainstances{
			_, err := sarama.NewAsyncProducer(instanceip, config)
			if err != nil {
			fmt.Printf("Cannot connect to:%s", instanceip)
				panic(err)
			}
			fmt.Printf("Successfully Connected to %s", instanceip)

	}*/

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                          sync.WaitGroup
		enqueued, successes, errors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			errors++
		}
	}()

ProducerLoop:
	for {
		message := &sarama.ProducerMessage{Topic: "myTopic", Value: sarama.StringEncoder("testing 123")}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)

}

func getProducer(UATkafkainstances [][]string, config *sarama.Config) sarama.AsyncProducer {
	var producer sarama.AsyncProducer
	var err error
	for _, instance := range UATkafkainstances {
		producer, err = sarama.NewAsyncProducer(instance, config)
		if err != nil {
			fmt.Printf("Cannot connect to")
			panic(err)

		}
		break
	}
	return producer
}
