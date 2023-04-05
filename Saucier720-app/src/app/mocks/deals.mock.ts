import { Ingredient } from "../core/interfaces/ingredient";

    var mockDeals1: Ingredient = {
        "Name":"riley butter",
        "StoreCost":369.99,
        "OnSale":true,
        "SalePrice":0,
        "SaleDetails":"BOGO",
        "Quantity":10
    };

    var mockDeals2: Ingredient = {
        "Name":"peanut butter",
        "StoreCost":369.99,
        "OnSale":true,
        "SalePrice":0,
        "SaleDetails":"BOGO",
        "Quantity":10
    };

    var mockDeals3: Ingredient = {
        "Name":"jelly",
        "StoreCost":1,
        "OnSale":false,
        "SalePrice":0,
        "SaleDetails":"N/A",
        "Quantity":30
    };
    
    var mockDeals4: Ingredient = {
        "Name":"bread",
        "StoreCost":10.69,
        "OnSale":true,
        "SalePrice":0,
        "SaleDetails":"$2 for 2",
        "Quantity":2
    };

export const DEALS: Ingredient[] = [
    mockDeals1,
    mockDeals2,
    mockDeals3,
    mockDeals4,
];
