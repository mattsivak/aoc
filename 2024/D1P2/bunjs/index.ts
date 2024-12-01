const input = await Bun.file("input.txt").text()

const lines = input.split("\n")

const list1 = lines.map((line) => line.split("   ")[0]).map(Number).sort((a, b) => a - b)
const list2 = lines.map((line) => line.split("   ")[1]).map(Number).sort((a, b) => a - b)

let output = 0
for (let i = 0; i < list1.length; i++) {
  const num = list1[i]
  const sameNumbers = list2.filter((n) => n === num)
  output += num * sameNumbers.length
}

console.log(output)