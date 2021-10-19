package de.codebydenis.warehouse;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString
public class StockItem {
    private String id;
    private String name;
    private String desc;
    private int stock;
    private String imageURL;
}
