import { Injectable, OnModuleInit } from '@nestjs/common';
import { Kafka, Producer } from 'kafkajs';

@Injectable()
export class KafkaProducerService implements OnModuleInit {
  private kafka: Kafka;
  private producer: Producer;

  onModuleInit() {
    this.kafka = new Kafka({
      clientId: 'auth-service',
      brokers: ['localhost:9092'], // твой Kafka брокер
    });
    this.producer = this.kafka.producer();
    this.producer.connect();
  }

  async send(topic: string, key: string, eventType: string, payload: any) {
    await this.producer.send({
      topic,
      messages: [
        {
          key,
          value: JSON.stringify(payload),
          headers: { 'event-type': eventType }, // тип события
        },
      ],
    });
  }
}
