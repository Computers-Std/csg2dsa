function insertionSort(array) {
    for (let index = 1; index < array.length; index++) {
        let key = array[index];
        let position = index - 1;

        while (position >= 0 && array[position] > key) {
            array[position + 1] = array[position];
            position = position - 1;
        }
        array[position + 1] = key;
    }
    return array;
}

let arr = [12, 11, 13, 5, 6];
console.log(insertionSort(arr));
