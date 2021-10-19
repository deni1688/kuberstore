package de.codebydenis.warehouse;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Component;

import static de.codebydenis.warehouse.MessagingConfig.QUEUE;

@Component
public class MessageConsumer {
    @RabbitListener(queues = QUEUE)
    public void consume(StockItem item) {
        System.out.println(item);
    }
}
