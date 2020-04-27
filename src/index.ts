import { Person } from "./interfaces"

console.time("Read")
let generatedData = require("../generated.json")
let people: Person[] = generatedData
console.timeEnd("Read")
console.log("People in dataset, ", people.length);

console.time("Foreach")
people.forEach((person) => {
    person.name = returner(person.name)   
})
console.timeEnd("Foreach")

console.time("iLoop")
let i = 0;
for (i = 0; i < people.length; i++) {
    // const person = people[i];
    people[i].name = returner(people[i].name)
    // person.name = returner(person.name)
}
console.timeEnd("iLoop")



const used = process.memoryUsage().heapUsed / 1024 / 1024;
console.log(`The script uses approximately ${Math.round(used * 100) / 100} MB`);

function returner<T>(input: T): T {
    return input
}
