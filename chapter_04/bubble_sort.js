function bubbleSort(array) {
  let unsortedIndex = array.length - 1;
  let sorted = false;

  while (!sorted) {
    sorted = true;
    for (let i = 0; i < unsortedIndex; i++) {
      if (array[i] > array[i + 1]) {
        let temp = array[i];
        array[i] = array[i + 1];
        array[i + 1] = temp;
        sorted = false;
      }
    }
    unsortedIndex -= 1;
  }
  return array;
}

let arr = [65, 55, 45, 35, 25, 15, 10];
console.log(bubbleSort(arr));
