let fs = require("fs");
let path = require("path");

let input = fs
  .readFileSync(path.resolve(__dirname, "./../input/day2.txt"), "utf8")
  .split("\n");

const getPosition = (arr) => {
  arr = arr.map((el) => el.split(" "));
  let [x, y] = [0, 0];
  let aim = 0;

  for (let i = 0; i < arr.length; i++) {
    const [dir, num] = arr[i];
    if (dir === "forward") {
      x += parseInt(num);
      y += parseInt(num) * aim;
    } else if (dir === "up") {
      // y += parseInt(num);
      aim -= parseInt(num);
    } else if (dir === "down") {
      // y -= parseInt(num);
      aim += parseInt(num);
    }
  }

  return x * y;
};

console.log(getPosition(input));
