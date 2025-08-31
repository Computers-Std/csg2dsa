function twoNumberProducts(array) {
    let products = [];
    for (let i = 0; i <= array.length; i++) {
        for (let j = i + 1; j < array.length; j++) {
            let product = array[i] * array[j];
            products.push(product);
        }
    }
    return products;
}

let arr = [1, 2, 3, 4, 5];
console.log(twoNumberProducts(arr));
