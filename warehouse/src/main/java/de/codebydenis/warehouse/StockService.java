package de.codebydenis.warehouse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class StockService {
    private final StockRepository repo;

    @Autowired
    public StockService(StockRepository repo) {
        this.repo = repo;
    }

    protected void saveStockItem(StockItem item) {
        repo.save(item);
    }
}
