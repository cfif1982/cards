package kafkacons

import (
	"encoding/json"
	"log/slog"

	"github.com/IBM/sarama"

	"github.com/cfif1982/cards/internal/controller"
	swgModels "github.com/cfif1982/cards/internal/infrastructure/swagger/models"
)

const messageReceived = "Получено сообщение: %s, partition: %d, offset: %d\n"
const topicKeyAddUser = "add_user"
const topicKeyAddBank = "add_bank"

type KafkaConsumer struct {
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
	log               *slog.Logger
	controller        controller.Controller
}

func NewKafkaConsumer(
	log *slog.Logger,
	controller controller.Controller,
	topicName,
	host string,
) *KafkaConsumer {
	// Настройка консюмера
	config := sarama.NewConfig()
	// config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest // Читаем только новые сообщения

	// Создание консюмера
	consumer, err := sarama.NewConsumer([]string{host}, config)
	if err != nil {
		log.Info(err.Error())
	}

	// Подписка на топик. Читаем только новые сообщения
	partitionConsumer, err := consumer.ConsumePartition(topicName, 0, sarama.OffsetNewest)
	if err != nil {
		log.Info(err.Error())
	}

	return &KafkaConsumer{
		consumer:          consumer,
		partitionConsumer: partitionConsumer,
	}
}

func (k *KafkaConsumer) Close() {
	k.consumer.Close()
	k.partitionConsumer.Close()
}

func (k *KafkaConsumer) Run() {
	// Чтение сообщений
	go func() {
		k.log.Info("Ожидание сообщений...")
		for message := range k.partitionConsumer.Messages() {
			k.log.Info(messageReceived, string(message.Value), message.Partition, message.Offset)
			k.checkMessage(message)
		}
	}()
}

// проверяем полученное сообщение
func (k *KafkaConsumer) checkMessage(message *sarama.ConsumerMessage) {
	key := string(message.Key)

	switch key {
	// получили сообщение от кафки о созданни банка
	case topicKeyAddBank:
		// Десериализация JSON в структуру
		var msg *swgModels.NewBank
		err := json.Unmarshal(message.Value, msg)
		if err != nil {
			k.log.Info("Ошибка при десериализации сообщения: %v", err)
			return
		}

		// выполняем вставку
		_, err = k.controller.AddBank(msg)
		if err != nil {
			k.log.Info("Ошибка при добавлении банка: %v", err)
		}

	// получили сообщение от кафки о созданни юзера
	case topicKeyAddUser:
		// Десериализация JSON в структуру
		var msg *swgModels.NewUser
		err := json.Unmarshal(message.Value, msg)
		if err != nil {
			k.log.Info("Ошибка при десериализации сообщения: %v", err)
		}

		// выполняем вставку
		_, err = k.controller.AddUser(msg)
		if err != nil {
			k.log.Info("Ошибка при добавлении пользователя: %v", err)
		}
	}
}
