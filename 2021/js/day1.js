let fs = require("fs");
let path = require("path");

let input = fs
  .readFileSync(path.resolve(__dirname, "./../input/day1.txt"), "utf8")
  .split("\n");

const part1 = (arr) => {
  let max = Number(arr[0]);
  let count = 0;
  for (let i = 1; i < arr.length; i++) {
    let num = Number(arr[i]);
    if (num > max) {
      count++;
    }
    max = num;
  }
  return count;
};

const part2 = (arr) => {
  const queue = arr.slice(0, 3);
  const result = [queue.reduce((a, b) => Number(a) + Number(b))];
  let val = queue.reduce((a, b) => Number(a) + Number(b));
  for (let i = 3; i < arr.length; i++) {
    val = val - Number(arr[i - 3]) + Number(arr[i]);
    // result.push(maintainQueue(queue, arr[i]));
    result.push(val);
  }
  return part1(result);

  function maintainQueue(arr, el) {
    arr.push(el);
    arr.shift();
    return arr.reduce((a, b) => Number(a) + Number(b));
  }
};

console.log(part1(input));
console.log(part2(input));
