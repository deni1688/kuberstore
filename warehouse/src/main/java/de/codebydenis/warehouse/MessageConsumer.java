package de.codebydenis.warehouse;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import static de.codebydenis.warehouse.MessagingConfig.QUEUE;

@Component
public class MessageConsumer {
    private final StockService service;

    @Autowired
    public MessageConsumer(StockService service) {
        this.service = service;
    }

    @RabbitListener(queues = QUEUE)
    public void consume(StockItem item) {
        System.out.println("received message with: " + item);
        service.saveStockItem(item);
    }
}
